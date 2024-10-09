# Introduction

A place to store all scripts that learning go from https://www.go.dev/

# Getting started

After installing g as per instructed, before starting the project,
there are some steps to learn:

### Dependencies management: create `go.mod` file

For dependencies tracking at the project root, following [managing-dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module).
Let init a `go.mod` file for our first module.

```shell
go mod init hello
```

In the same directory level, create a file called `hello.go`.
The relevent code can be seen in `./hello.go`.

To run the code, use `go run path/to/go/file` command.

> [!info] Keeping track of your imports:
> In case we would like to use an external package, we can add to `go.mod` file.
> To add to the `go.mod` file, we can import the package in our code.
> Then run `go mod tidy` command, the command will automatically detect imports,
> and update our `go.mod` file. Simultaneously, the command also creates a `go.sum` file.


