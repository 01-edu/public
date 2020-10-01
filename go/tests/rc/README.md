### RC (Restrictions Checker)

This program analyzes a go source file and displays in standard output the imports, functions, slice types and loops used without authorization.

### By default:

- Allowed

  - All functions declared inside the source file are allowed
  - Slices of all types are allowed
  - Loops are allowed
  - Relative imports are allowed

- Disallowed

  - NO absolute imports are allowed
  - NO built-in functions are allowed.
  - NO casting is allowed

### Flags

- `-h` for help
- `-cast` allows casting to every built-in type.
- `-no-for` prohibits the use of `for` loops in the program or function.
- `-allow-builtin` allows all builtin functions and casting to builtin types
- `-no-slices` disallows the use of all slices types
- `-no-these-slices=type1,type2`: disallows the slices of type1 and type2
- `-no-relative-imports`: disallows the use of relative imports
- `--no-lit="{PATTERN}"`: disallows character and string literals that match the pattern `PATTERN` which represent a Regular Expression

### Arguments:

- Flags can be passed at any point (in the beginning, middle or end) of the argument list

- The remaining arguments represent the allowed functions
  - Allowed imports and functions from a package
    - `<package>.*` for full imports (all functions from that package are allowed)
    - `<package>`.`<function>` for partial imports (only the function is allowed)
    - `<package>`.`<function>#amount` the function is only allowed to be used `amount` number of times
    - Ex: `fmt.*` (all functions from `fmt` are allowed), `github.com/01-edu/z01.PrintRune` (only `z01.PrintRune` is allowed), `fmt.Println#2` (fmt.Println can only be used 2 times or less)
  - Allowed built-in functions
    - Use the name of the built-in function
    - It is possible to limit the number of calls of a functions like with the imports using the '#' character
    - Ex: `make`, `append`, `len`, `print#2`.

### Example:

- To allow the import of the whole `fmt` package, `z01.PrintRune` and the built-in functions `len` for the file `main.go`
  Note: The imports must be written exactly the way they are written inside the source code, example:

  ```console
  _$ rc main.go fmt.* github.com/01-edu/z01.PrintRune len
  ```

- Import "fmt" is allowed by executing

  ```console
  _$ rc sourcefile.go fmt.*
  ```

- Import "go/parser" is allowed by executing

  ```console
  _$ rc sourcefile.go go/parser.*
  ```

- Import "github.com/01-edu/z01" is allowed by executing

  ```console
  _$ rc sourcefile.go github.com/01-edu/z01.*
  ```

- Disallow literals

  - Use the flag `--no-lit="{PATTERN}"`
    ```console
    _$ rc -no-slices --no-lit=[b-yB-Y] main.go fmt.* github.com/01-edu/z01.PrintRune len
    ```

- Allow all type of casting

  ```console
  _$ rc -cast sourcefile.go fmt.* github.com/01-edu/z01 os.* strconv.* make len append
  ```

  - this will allow all type of casting in the file sourcefile.go

- Disallow the use of the slices of type `string` and `int`

  ```console
  _$ rc -no-these-slices=string,int sourcefile.go
  ```

- To allow just one type of casting

  ```console
  _$ rc sourcefile.go fmt.* github.com/01-edu/z01 os.* strconv.* make len append rune
  ```

  - this will allow the casting to `rune`, but not `int8`, ..., `string`, `float32`, ...

### How to read the error message

Let us look to an example snipped of code, let us imagine this code in a file called `main.go`:

```go
package main

import "fmt"

func main() {
	for _, v := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Println(v)
	}
	fmt.Println()
}
```

Now let us run the `rc` and understand the message

```console
_$ rc main.go github.com/01-edu/z01.PrintRune
Parsing:
	Ok
Cheating:
	TYPE:             	NAME:      	LOCATION:
	illegal-import    	fmt        	main.go:3:8
	illegal-access    	fmt.Println	main.go:7:3
	illegal-access    	fmt.Println	main.go:10:2
	illegal-definition	main       	main.go:5:1
```

The important part is printed after the `Cheating` tag:

- The import of of the package `fmt` is not allowed
- In go the dot (.) is also known as the access operator for that reason the use of fmt.Println is shown as an illegal-access
- Finally the main function is shown as illegal-definition because the function is using disallowed functions that does not mean that the function can not be defined it just mean that the definition of the function must be changed to not use disallowed functions.
- Notice that the third column of the output with the tag "LOCATION:" show the location in the following way filepath:line:column
  This mean that you have to substitute the illegal function for ones that are allowed or write your own function with allowed functions
