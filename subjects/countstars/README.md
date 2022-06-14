## CountStars

### Instructions
Write a function named `CountStar` that makes u fall asleep by counting stars, it takes an integer in the parameter
- If the number is negative or equal to 0, return "`No star`"
- No need to manage overflow.

```go
func CountStars(num int) string {

}
```
```go
import (
	"fmt"
	"strconv"   
)

func main() {
	fmt.Println(CountStars(5))
	fmt.Println(CountStars(4))
	fmt.Println(CountStars(-1))
}
```

```console
$ go run . | cat -e
1 star...2 star...3 star...4 star...5$
1 star...2 star...3 star...4$
No star$
```
