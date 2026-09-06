# go-log-parser

A minimal Go project with a hello-world example.

## Requirements

- Go 1.22 or newer

## Run

```sh
go run .
```

Expected output:

```text
Hello, world!
```

## Test

```sh
go test ./...
```

`TestGreeting` is the positive test: it verifies that `greeting()` returns exactly `Hello, world!`.

`TestGreetingDoesNotReturnIncorrectMessage` is the negative test: it verifies that `greeting()` does not return the incorrect message `Goodbye, world!`.
