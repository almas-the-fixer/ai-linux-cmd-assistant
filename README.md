# AI Linux Command Assistant

A terminal-based Linux assistant built in Go. It uses a local **Qwen2.5:3B** model through **Ollama** for intent detection and response generation, and retrieval-augmented generation (RAG) to ground Linux answers in a small Markdown knowledge base.

For semantic retrieval, the application creates embeddings remotely with the **Jina Embeddings API** and stores/searches those vectors in **Weaviate Cloud**. Qwen inference remains local; embedding requests and the corresponding text sent to Jina use an external API.

## Features

- Local LLM inference with Ollama and Qwen2.5:3B
- Intent classification for Linux-related requests
- Linux command explanations, how-to guidance, and troubleshooting
- Linux/Bash command generation with a structured response format
- Markdown knowledge base split into chunks for retrieval
- Jina text embeddings for knowledge chunks and user queries
- Weaviate vector storage and semantic search
- Retrieved documentation included in prompts sent to Qwen
- Basic application-level command validation
- Explicit confirmation before attempting to execute generated commands
- Colored, bordered terminal output

> **Security notice:** The command validator is a basic denylist, not a security sandbox. It cannot guarantee that generated commands are safe. Review every command before approving it, and do not run this project with elevated privileges or against important data unless you understand the risks.

## Architecture

### Answering a Linux question

```text
User query
   |
   v
Local Qwen intent detection (Ollama)
   |
   +---- OFF_TOPIC / MALICIOUS / unknown --> response path without retrieval
   |
   v
Jina: embed query (retrieval.query)
   |
   v
Weaviate: near-vector search
   |
   v
Top relevant knowledge chunks
   |
   v
Prompt builder adds retrieved context
   |
   v
Local Qwen2.5:3B generates the answer
   |
   v
Terminal UI
```

### Indexing the knowledge base

```text
Markdown files
   |
   v
Load and split into chunks
   |
   v
Jina: embed each chunk (retrieval.passage)
   |
   v
Store text, source filename, and vector in Weaviate
```

The collection uses **bring-your-own vectors**: Jina generates the vectors, and Weaviate stores and searches them. The application uses 768-dimensional vectors with the `jina-embeddings-v5-text-nano` model.

### Generated-command path

```text
User request
   |
   v
Intent detection --> GENERATE_COMMAND
   |
   v
Retrieval + prompt construction
   |
   v
Qwen generates COMMAND and EXPLANATION fields
   |
   v
Go parser --> basic safety validator
   |
   v
Display command and request explicit confirmation
   |
   v
Executor --> display command output
```

The model's prompt instructions are not treated as a security boundary. The application validator is only a basic additional check.

## Technology stack

| Technology | Purpose |
|---|---|
| Go | CLI and application logic |
| Ollama | Local LLM runtime |
| Qwen2.5:3B | Intent detection and response generation |
| Jina Embeddings API | Remote text embeddings for documents and queries |
| Weaviate Cloud | Vector storage and semantic retrieval |
| Lip Gloss | Terminal styling |
| Go `os/exec` | Command execution |
| Go `net/http` / JSON | API requests and responses |

## Requirements

- Linux
- Go installed
- Ollama installed and running
- Qwen2.5:3B available in Ollama
- A Jina API key
- A Weaviate Cloud cluster and API key
- Internet access for Jina embeddings and Weaviate Cloud

Pull the model:

```bash
ollama pull qwen2.5:3b
```

If Ollama is not already running as a service, start it in a separate terminal:

```bash
ollama serve
```

## Configuration

Create a `.env` file in the repository root. Do not commit this file or share its credentials.

```dotenv
JINA_API_KEY=your_jina_api_key
WEAVIATE_HOST=your-cluster-hostname
WEAVIATE_API_KEY=your_weaviate_api_key
```

Use the **bare Weaviate hostname** for `WEAVIATE_HOST`; do not include `https://`. The application uses HTTPS separately.

Keep `.env` out of version control. If a real API key is accidentally committed, revoke it and create a replacement.

## Setup and run

Clone the repository and enter it:

```bash
git clone https://github.com/almas-the-fixer/ai-linux-cmd-assistant.git
cd ai-linux-cmd-assistant
```

Download Go dependencies:

```bash
go mod download
```

### 1. Index the knowledge base

Run this after configuring `.env`:

```bash
go run cmd/index/main.go
```

The indexer loads the Markdown files, creates chunks, obtains a Jina embedding for each chunk, and inserts the objects into the `LinuxKnowledge` collection in Weaviate.

The current knowledge base contains eight Markdown files and approximately 174 chunks. The indexer checks whether the collection already contains objects to avoid importing the same set repeatedly. This is a simple one-time indexing guard, not an incremental update system.

**If you edit the knowledge files and need to rebuild the index**, delete the `LinuxKnowledge` collection in Weaviate Cloud and run the indexer again. This removes the existing indexed objects, so do it only when you intend to rebuild the knowledge base.

### 2. Run the assistant

```bash
go run cmd/main.go
```

The normal application connects to Ollama and Weaviate at startup. Ensure the `.env` variables are present and both services are reachable.

### Optional: test retrieval

The repository may include a small search diagnostic under `cmd/search`:

```bash
go run cmd/search/main.go
```

This is a development/testing utility, not required for normal use. It may use a hard-coded sample query.

## Knowledge base

The current Markdown knowledge base covers topics such as:

- Archives
- File management
- File search
- Networking
- Package management
- Permissions and ownership
- Process management
- Text processing

The indexer splits the documents into chunks and stores each chunk's content, source filename, and Jina-generated vector in Weaviate.

## Example interaction

```text
AI Linux Assistant
Type 'exit' to quit.

> what does chmod do?
User's Intent: EXPLAIN_COMMAND

Answer
...a concise explanation grounded in retrieved Linux documentation...

> what does pgrep do?
User's Intent: EXPLAIN_COMMAND

Answer
...an explanation using relevant process-management documentation...
```

For command-generation requests, the model is asked to return a structure like:

```text
COMMAND: find . -type f -name "*.log"
EXPLANATION: Finds files ending in .log under the current directory.
```

The application parses the response, validates the extracted command, displays it, and asks for explicit confirmation before execution. The generated command may still be wrong or unsafe, so inspect it carefully.

## Safety and limitations

- The validator uses a limited denylist and is not a comprehensive shell-security mechanism.
- A command passing validation does not mean it is safe.
- The command parser is intentionally simple and may not support all shell quoting, substitutions, pipelines, or compound-command syntax correctly.
- Never approve a command you do not understand.
- Do not run the application as root or with `sudo`.
- Qwen2.5:3B may misclassify requests, produce incorrect answers, or fail to follow the requested output format.
- Retrieval can return related but imperfect chunks; answers are not guaranteed to be correct just because context was retrieved.
- If the knowledge base does not cover a topic, retrieval may not provide useful evidence.
- Intent detection and response generation are separate local model calls; the application does not currently provide full conversational memory.
- Query text and knowledge chunks are sent to the Jina Embeddings API to create vectors. Do not send secrets or private documents through this embedding path.
- Weaviate Cloud and Jina API availability, quotas, and terms are controlled by their providers and may change.

## Project structure

```text
ai-linux-cmd-assistant/
├── cmd/
│   ├── main.go          # Run the assistant
│   ├── index/main.go    # Build the vector knowledge index
│   └── search/main.go   # Optional retrieval diagnostic
├── internal/
│   ├── cli/             # Input loop and application flow
│   ├── command/         # Parse structured command responses
│   ├── embedding/       # Jina embedding API client
│   ├── executor/        # Execute approved commands
│   ├── intent/          # Intent classification
│   ├── knowledge/       # Source Markdown files
│   ├── knowledgebase/   # Load documents and create chunks
│   ├── ollama/          # Local Ollama client
│   ├── prompt/          # Intent-specific prompts and RAG context
│   ├── security/        # Basic command validation
│   ├── ui/              # Colored terminal presentation
│   └── weaviate/        # Connection, collection, indexing and search
├── .env                 # Local secrets; do not commit
├── go.mod
├── go.sum
└── README.md
```

## Future improvements

- Automated tests for parsing, retrieval, and validation
- Incremental re-indexing when Markdown files change
- Better chunk boundaries and retrieval relevance filtering
- Stronger command validation and isolated/sandboxed execution
- More robust shell-argument parsing
- More visible retrieval diagnostics and source citations
- Improved conversation context handling
- Streaming responses and richer terminal UX

## Goal

This project demonstrates how a locally hosted Generative AI model can be combined with Go application logic, remote embeddings, vector search, and user-controlled command execution to build a practical Linux assistant. The LLM generates suggestions; deterministic application code and explicit user confirmation remain between generated commands and execution.
