# demo-go

A minimal Go demo project demonstrating a small package, a main program, and a unit test.
- **Main package**: `greet` with a `Hello(name)` function
- **CLI application**: Prints "Hello, World!" to the console
- **Tests**: Unit tests verify that `Hello()` works correctly

## Project Structure

```
demo-go/
├── go.mod           # Module file
├── main.go          # Main CLI application
├── demo-go          # Do not commit this since it is a binary
└── greet/
    ├── greet.go     # Greet package
    └── greet_test.go # Unit tests
```

## Prerequisites

- **Go:** Install Go (recommended 1.20+). On macOS with Homebrew:

```zsh
brew install go
```

Tested with:
- Go: `go1.24.0`

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