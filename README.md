# ai-linux-cmd-assistant
A Generative AI-based Linux Command Assistant that uses Retrieval-Augmented Generation (RAG) to answer Linux-related queries, explain commands, generate shell scripts, and troubleshoot common errors using Linux documentation. The system supports both local and API-based language models.

### Requirements

- Ollama

Install Ollama:

```
Linux:   curl -fsSL https://ollama.com/install.sh | sh
Mac:     curl -fsSL https://ollama.com/install.sh | sh
Windows: Goto: https://ollama.com/download/windows and Download .exe or paste command in powershell
```

- Golang

- Pull any model for this project we used qwen2.5:3b

> Note: If you are pulling other model make sure to change const model = qwen2.5:3b to whatever model you are pulling, It will be moved into .env in future...

```
ollama pull qwen2.5:3b
```

### Steps

1. Clone Repo
2. Setup Ollama
3. Pull Model
4. Run ```go run cmd/main.go```

You will get response to the prompt passed in main.go

That's it for Phase 1 of this Project...