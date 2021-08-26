## ascii-art-color

### Objectives

You must follow the same [instructions](../README.md) as in the first subject but this time with colors.

The output should manipulate colors using the **flag** `--color=<color>`, in which `--color` is the flag and `<color>` is the color desired by the user. These colors can be achieved using different notations (color code systems, like `RGB`, `hsl`, `ANSI`...), it is up to you to choose which one you want to use.

- You should be able to specify a single or a set of letters that you want to color (use your imagination for this one).
- If the letter is not specified, the whole `string` should be colored.
- The flag must have exactly the same format as above, any other formats must return the following usage message:

```console
Usage: go run . [STRING] [BANNER] [OPTION]

EX: go run . something standard --color=<color>
```

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](../../good-practices/README.md).
- It is recommended that the code should present a **test file**.

### Allowed packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed

This project will help you learn about :

- The Go file system(**fs**) API
- Color converters
- Data manipulation
- Terminal display
