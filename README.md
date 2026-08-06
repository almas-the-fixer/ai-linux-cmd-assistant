# ai-linux-cmd-assistant

A terminal-based AI Linux Command Assistant built with Go and Ollama.

The assistant uses a local Large Language Model (LLM) to answer Linux and Bash-related questions. Before generating a response, it performs **intent detection** to classify the user's query (e.g., command explanation, how-to, troubleshooting, off-topic, or malicious prompt injection) and builds an appropriate prompt for the model.

Future versions will integrate **Retrieval-Augmented Generation (RAG)** using Linux documentation to provide more accurate and grounded responses.

---

## Features

* Terminal-based interactive CLI
* Runs completely with a local LLM using Ollama
* Intent Detection

  * Explain Linux commands
  * How-to questions
  * Linux troubleshooting
  * Off-topic detection
  * Prompt injection detection
* Intent-specific prompt building
* Modular Go project structure

---

## Requirements

### 1. Ollama

Install Ollama:

```bash
Linux:
curl -fsSL https://ollama.com/install.sh | sh

macOS:
curl -fsSL https://ollama.com/install.sh | sh

Windows:
Download the installer from:
https://ollama.com/download/windows
```

---

### 2. Go

Install the latest stable version of Go.

---

### 3. Pull a Local Model

This project currently uses:

```text
qwen2.5:3b
```

Pull it with:

```bash
ollama pull qwen2.5:3b
```

> **Note**
>
> If you use another model, update the `model` constant inside the Ollama client accordingly.
>
> In a future version, this configuration will be moved to a `.env` file.

---

## Running the Project

Clone the repository:

```bash
git clone <repository-url>
cd ai-linux-cmd-assistant
```

Start Ollama (if it is not already running).

Run the application:

```bash
go run cmd/main.go
```

You should see:

```text
AI Linux Assistant
Type 'exit' to quit.

>
```

You can now ask Linux-related questions such as:

```text
What is grep?

How do I recursively copy a directory?

Permission denied while running chmod
```

Type `exit` to quit.

---

## Current Architecture

```text
User
 │
 ▼
CLI
 │
 ▼
Intent Detection
 │
 ▼
Prompt Builder
 │
 ▼
Ollama (Local LLM)
 │
 ▼
Response
```

---

## Roadmap

### Phase 1 ✅

* Interactive CLI
* Ollama integration
* Local LLM support
* Intent Detection
* Prompt Builder

### Phase 2

* Streaming responses
* Better prompt engineering
* Refactor duplicated HTTP request logic

### Phase 3

* Retrieval-Augmented Generation (RAG)
* Linux documentation indexing
* Context retrieval
* Better troubleshooting responses

### Phase 4

* Shell command generation
* Command safety checks
* Command explanations with documentation references

---

This project is being built primarily as a learning project to explore Go, Generative AI, Retrieval-Augmented Generation (RAG), prompt engineering, and backend system design.