# secepp

A CLI secrets manager for storing and retrieving secrets (similar to AWS Secrets Manager).

Setup

- Prerequisites: install Go (1.18+) and Git.
- Clone the repository:

```bash
git clone https://github.com/skarekroe666/secepp.git
cd secepp
```

Alternatively, you can run the app via Docker without cloning or building locally (see the Dockerfile in the repository).

Build

```bash
go build -o secepp .
```

Run

- Run the compiled binary:

```bash
./secepp <command>
```

- Or run directly with `go` during development:

```bash
go run main.go <command>
```

Execute (examples)

- Create (example):

```bash
./secepp create files/employee.json
```

- List all secrets:

```bash
./secepp list
```

- Get a secret by id (example):

```bash
./secepp get <filename>
```

- Delete a secret by id (example):

```bash
./secepp delete <hash>
```

Replace `<command>` and flags with the command you want to run. See CLI help for available commands:

```bash
./secepp --help
```
