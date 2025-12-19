# demo-go

A minimal Go demo project demonstrating a small package, a main program, and a unit test.

**Prerequisites:**
- **Go:** Install Go (recommended 1.20+). On macOS with Homebrew:

```zsh
brew install go
```

**Project layout:**
- `go.mod` — module file
- `main.go` — small CLI that prints a greeting
- `greet/` — package with `greet.go` and `greet_test.go`

**Quick start (from workspace root):**

```zsh
cd gitRepos/demo-go

# Download dependencies (if any)
go mod download

# Build the binary
go build -o demo-go

# Run tests
go test ./...

# Run directly
go run .

# Or run the built binary
./demo-go
```

If you plan to use a different module path, update `module` in `go.mod` accordingly.

That's it — the program prints a simple greeting and the unit test demonstrates package testing.