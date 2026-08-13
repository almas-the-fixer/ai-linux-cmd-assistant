# AI Linux Command Assistant

A terminal-based Linux assistant powered by a **local Large Language Model (LLM)**. The application understands Linux-related requests, identifies the user's intent, generates Linux/Bash commands when appropriate, validates generated commands against a safety layer, and requires explicit user confirmation before executing commands.

The project is designed as a practical demonstration of **Generative AI + backend engineering + Linux command execution** without relying on cloud AI APIs.

## Features

* Local LLM inference using Ollama
* Qwen2.5 3B as the language model
* Linux/Bash-focused conversational assistant
* Intent classification
* Linux command generation
* Command explanation
* Command parsing
* Basic command safety validation
* Explicit `y/N` confirmation before execution
* Actual Linux command execution
* Command output displayed in the terminal
* Colored and bordered terminal UI
* Handles general Linux questions, how-to questions, and troubleshooting

## Architecture

```text
User
 │
 ▼
CLI
 │
 ▼
Intent Detector
 │
 ├── EXPLAIN_COMMAND
 ├── HOW_TO
 ├── TROUBLESHOOT
 ├── GENERATE_COMMAND
 ├── OFF_TOPIC
 └── MALICIOUS
 │
 ▼
Prompt Builder
 │
 ▼
Ollama
 │
 ▼
Qwen2.5:3b
 │
 ▼
Response
 │
 ├── General Response ──────────────► UI
 │
 └── Generated Command
          │
          ▼
     Response Parser
          │
          ▼
     Security Validator
          │
          ▼
     User Confirmation
          │
       ┌──┴──┐
       │     │
      NO    YES
       │     │
       ▼     ▼
     Stop   Executor
               │
               ▼
          Command Output
```

## How It Works

### 1. User Input

The user enters a Linux-related request through the terminal.

Example:

```text
> give me a command to find all .log files
```

### 2. Intent Detection

The application sends the user's request to the local LLM and classifies it into one of the supported intents.

Examples:

```text
EXPLAIN_COMMAND
HOW_TO
TROUBLESHOOT
GENERATE_COMMAND
OFF_TOPIC
MALICIOUS
```

### 3. Response Generation

The detected intent is used to construct an appropriate prompt.

For command-generation requests, the model is instructed to return a structured response:

```text
COMMAND: find . -type f -name "*.log"
EXPLANATION: Searches the current directory recursively for files ending in .log.
```

### 4. Command Parsing

The generated response is parsed into a Go structure:

```go
type Command struct {
    Command     string
    Explanation string
}
```

This separates the executable command from its explanation.

### 5. Security Validation

Before execution, the generated command passes through a deterministic validation layer.

The validator checks for potentially dangerous patterns such as:

* Destructive filesystem operations
* Disk formatting
* Filesystem wiping
* Dangerous block-device operations
* Fork bombs
* Security-disabling commands
* Remote code piped directly into a shell

The LLM's safety instructions are **not treated as the security boundary**. The Go validator provides an additional application-level safety check.

### 6. User Confirmation

Even if the command passes validation, the application does not execute it automatically.

The user must explicitly approve it:

```text
Execute this command? [y/N]:
```

Only `y` or `yes` results in execution.

Anything else is treated as rejection.

### 7. Command Execution

Approved commands are executed using Go's `os/exec` package.

The application uses:

```go
exec.Command(binary, args...)
```

rather than passing the generated command to a shell using `sh -c`.

The command's combined standard output and error output are captured and displayed to the user.

## Technology Stack

| Technology | Purpose                            |
| ---------- | ---------------------------------- |
| Go         | Application/backend implementation |
| Ollama     | Local LLM inference                |
| Qwen2.5 3B | Local language model               |
| bufio      | Terminal input handling            |
| os/exec    | Linux command execution            |
| Lipgloss   | Terminal UI styling                |
| Git        | Version control                    |

## Project Structure

```text
ai-linux-cmd-assistant/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── cli/
│   │   └── cli.go
│   │
│   ├── command/
│   │   └── command.go
│   │
│   ├── executor/
│   │   └── executor.go
│   │
│   ├── intent/
│   │   └── intent.go
│   │
│   ├── ollama/
│   │   ├── ollama.go
│   │   └── types.go
│   │
│   ├── prompt/
│   │   └── prompt.go
│   │
│   ├── security/
│   │   └── security.go
│   │
│   └── ui/
│       └── ui.go
│
├── go.mod
├── go.sum
└── README.md
```

> Note: The Project still contains knowledge/ and knowledgebase/
> These were meant to add RAG in future versions so these are dormant for now...

## Requirements

* Linux
* Go
* Ollama
* Qwen2.5 3B model

Install Ollama and pull the model:

```bash
ollama pull qwen2.5:3b
```

Start Ollama if it is not already running:

```bash
ollama serve
```

## Running the Application

Clone the repository and enter the project directory:

```bash
git clone <repository-url>
cd ai-linux-cmd-assistant
```

Install dependencies:

```bash
go mod download
```

Run the application:

```bash
go run cmd/main.go
```

The application starts with:

```text
AI Linux Assistant
Type 'exit' to quit.

>
```

## Example

```text
> can you give me a command to list all hidden files?

User's Intent: GENERATE_COMMAND

Command
╭──────────────────────────────╮
│ ls -a                        │
╰──────────────────────────────╯

Explanation
╭──────────────────────────────╮
│ Lists all files including    │
│ hidden files and directories.│
╰──────────────────────────────╯

Execute this command? [y/N]: yes

Answer
╭──────────────────────────────╮
│ .                            │
│ ..                           │
│ .git                         │
│ README.md                    │
│ ...                          │
╰──────────────────────────────╯
```

## Security Model

The project uses multiple layers of protection:

```text
LLM Safety Instructions
        ↓
Generated Command
        ↓
Go Command Parser
        ↓
Deterministic Denylist
        ↓
Explicit User Confirmation
        ↓
os/exec
```

The application does **not** consider the LLM's safety instructions sufficient by themselves.

The generated command is checked by application code before it can reach the executor.

### Important Limitation

The validator is a **simple denylist**, not a complete security sandbox.

Linux commands and shell environments are extremely powerful, and there are many ways to express similar operations. Therefore, this project should not be treated as a production-grade secure command execution environment.

The user confirmation step is intentionally retained as an additional safety layer.

## Design Decisions

### Why a Local Model?

The project uses Ollama and Qwen2.5 3B so that inference can happen locally without sending user queries to an external AI API.

This also makes the project suitable for experimentation without requiring an API key.

### Why Intent Detection?

Intent detection separates different types of requests before generating the final response.

For example:

```text
"What does grep do?"
        ↓
EXPLAIN_COMMAND
```

while:

```text
"give me a command to find large files"
        ↓
GENERATE_COMMAND
```

This allows different prompt instructions to be used for different tasks.

### Why Explicit Confirmation?

The model can generate commands, but the application should not blindly execute AI-generated instructions.

Therefore:

```text
Generated ≠ Executed
```

The user must explicitly authorize execution.

## Limitations

* The application currently targets Linux environments.
* The LLM can occasionally misunderstand user requests.
* Intent classification can be incorrect.
* Generated commands may not always be technically optimal.
* The command safety layer uses a denylist and cannot guarantee complete protection.
* Command execution requires careful user judgment.
* The application currently has limited conversation/context memory.
* The local model's capabilities are constrained by the relatively small 3B parameter size.

## Future Improvements

Potential future improvements include:

* Better command validation
* More robust command parsing
* Improved conversational context
* Streaming model responses
* Better terminal syntax highlighting
* Command history
* Configuration options
* More detailed execution feedback
* Sandboxed command execution
* Automated tests
* Support for additional local models

## Project Goal

The primary goal of this project is to demonstrate how a **local Generative AI model can be integrated with a traditional Go application** to create a practical Linux assistant.

Rather than allowing the AI model to directly control the system, the application places deterministic application logic and explicit user authorization between the model and command execution.

This makes the project a combination of:

**Generative AI + Go + Linux + CLI development + basic AI safety engineering.**