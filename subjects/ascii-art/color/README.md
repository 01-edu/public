## ascii-art-color

### Objectives

You must follow the same [instructions](../README.md) as in the first subject but this time with colors.

The output should manipulate colors using the **flag** `--color=<color> <substring to be colored>`, in which `--color` is the flag and `<color>` is the color desired by the user and `<substring to be colored>` is the substring that you can chose to be colored. These colors can be achieved using different notations (color code systems, like `RGB`, `hsl`, `ANSI`...), it is up to you to choose which one you want to use.

- The substring to be colored can be a single letter or more
- If the substring is not specified, the whole `string` should be colored.
- The flag must have exactly the same format as above, any other formats must return the following usage message:

```console
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
```

### Usage

```shell
$ go run . --color=red kit "a king kitten have kit"
```

For the example above, the substring `kit` in the word `kitten` and the word `kit` at the end should be colored.

If there are other `ascii-art` optional projects implemented, the program should accept other correctly formatted `[OPTION]` and/or `[BANNER]`.
Additionally, the program must still be able to run with a single `[STRING]` argument.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](../../good-practices/README.md).
- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

### Allowed packages

- Only the [standard Go](https://golang.org/pkg) packages are allowed

This project will help you learn about :

- The Go file system(**fs**) API
- Color converters
- Data manipulation
- Terminal display
