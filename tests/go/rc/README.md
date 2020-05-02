## rc (restrictions checker)

This program analyses a go source file and displays in standard output the imports, functions, array types and loops used without authorization.

### By default:

- NO imports and NO built-in functions are allowed.
- NO casting is allowed either.
- Only functions declared inside the source file are allowed.
- All array types are allowed
- Loops are allowed

### Flags

- Two flags are defined:
  - `--cast` allows casting to every built-in type.
  - `--no-for` prohibits the use of `for` loops in the program or function.
  - `--no-array`:
    - Prohibits all array types if no types are specified after the flag.
      Ex.
      ```console
	_$ ./rc main.go fmt.* github.com/01-edu/z01.PrintRune len --no-array
      ```
      All array type in main.go will cause an error message.
    - Prohibits only the types specified after the flag
      Ex.
      ```console
	_$ ./rc main.go fmt.* github.com/01-edu/z01.PrintRune len --no-array rune string
      ```
      Only array from the type rune and string are prohibit. All other array from built-in types are allowed

### Arguments:

- First Argument:

The program must be executed passing the go source file to be analyze as the first argument

- The remaining argument (from 2 to ...):

  Can be (without any particular order):

  - Allowed imports and functions from a package
    - `<package>.*` for full imports (all functions from that package are allowed)
    - `<package>`.`<function>` for partial imports (only the function is allowed)
    - `<package>`.`<function>#amout` for certain amounts (only certain amount os a function is allowed)
    - Ex: `fmt.*` (all functions from `fmt` are allowed), `github.com/01-edu/z01.PrintRune` (only `z01.PrintRune` is allowed), `append#2` (the only amount of `append`'s allowed is 2)
  - Allowed built-in functions

    - Use the name of the built-in function
    - Ex: `make`, `append`, `len`.
    - Allowed casting
    - by using the type of casting, ex: for allowing `string` casting, use `string`
    - Or use the flag `--cast`, to allow every type of casting

  - Import relative packages

    - Use the relative path
    - Ex: `../piscine`, `..`, `.`

  - Unallow for loops
    - Use the flags `--no-for`.
    - Note: remember to use it before the `--no-array` flag.
      - ex:
      ```console
      _$ ./rc main.go fmt.* github.com/01-edu/z01.PrintRune len --no-array <...> --no-for
      ```
      the last line produces undesired behaviors.
  - Unallow literals
    - Use the flag `--no-lit="{PATTERN}"`
    - Note: `"{PATTERN}"` must be a valid RegExp.
      - ex:
      ```console
      _$ ./rc main.go fmt.* github.com/01-edu/z01.PrintRune len --no-array --no-lit=[b-yB-Y]
      ```

- Optional lasts arguments
  - The flag `--no-array` must be given as the last argument or to signal that all the arguments after are unallowed array types

### Usage:

- To allow the import of the whole `fmt` package, `z01.PrintRune` and the built-in functions len in the file `main.go`

  The imports must be writen exactly the way are writen inside the source code, example:

```console
   _$ ./rc main.go fmt.* github.com/01-edu/z01.PrintRune len
```

- More examples:

- import "fmt" is allowed by executing

  ```console
  _$ ./rc sourcefile.go fmt.*
  ```

  - import "go/parser" is allowed by executing

  ```console
  _$ ./rc sourcefile.go go/parser.*
  ```

  - import "github.com/01-edu/z01" is allowed by executing

  ```console
  ./rc sourcefile.go github.com/01-edu/z01.*
  ```

  - import "../../../all/tests/go/solutions" is allowed by executing

  ```console
  _$ ./rc sourcefile.go ../../../all/tests/go/solutions
  ```

  (no `.*` is needed, all the functions from this relative package are allowed)

- allow all type of casting

  ```console
  _$ ./rc sourcefile.go ../../../all/tests/go/solutions/ztail/ztail.go fmt.* github.com/01-edu/z01 os.* strconv.* make len append --cast
  ```

  - this will allow all type of casting in the file ztail.go

- to allow just one type of casting

  ```console
  _$ ./rc sourcefile.go ../../../all/tests/go/solutions/ztail/ztail.go fmt.* github.com/01-edu/z01 os.* strconv.* make len append rune
  ```

  - this will allow `rune`, but not `int8`, ..., `string`, `float32`, ...
