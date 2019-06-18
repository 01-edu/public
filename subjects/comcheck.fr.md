## ROT 14

### Instructions

Write a function `rot14` that returns the string within the parameter but transformed into a rot14 string.

- If you not certain what we are talking about, there is a rot13 already.

### Expected function

```go
func rot14(str string) string {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	result := rot14("Hello How are You")
	arrayRune := []rune(result)

	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}

```

And its output :

```console
student@ubuntu:~/piscine/comcheck$ go build
student@ubuntu:~/piscine/comcheck$ ./comcheck "I" "Will" "Enter" "the" "galaxy"
Alert!!!
student@ubuntu:~/piscine/comcheck$ ./comcheck "galaxy 01" "do" "you" "hear" "me"
Alert!!!
student@ubuntu:~/piscine/comcheck$
```
