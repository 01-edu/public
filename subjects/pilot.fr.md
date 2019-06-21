## pilot

### Instructions

Écrire un fichier `go` de manière à ce que le programme ci-dessous compile.

### Utilisation

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	var donnie student.Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = student.AIRCRAFT1

	fmt.Println(donnie)
}
```
