## CountStars

### Instructions
Write a function named `CountStar` that makes u fall asleep by counting stars, it takes an integer in the parameter
- If the number is negative or equal to 0, return "`No star`"
- No need to manage overflow.

```go
func CountStar(num int) string {

}
```
```go
import (
	"fmt"
	"strconv"   
)

func main() {
	fmt.Println(CountStar(5))
	fmt.Println(CountStar(4))
	fmt.Println(CountStar(-1))
}
```

```console
$ go run . | cat -e
1 star...2 star...3 star...4 star...5$
1 star...2 star...3 star...4$
No star$
```
