## CountStars

### Instructions
Write a function named `CountStar` that makes you fall asleep by counting stars, it takes an integer in the parameter
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
$ go run . 
1 star...2 stars...3 stars...4 stars...5 stars...
1 star...2 stars...3 stars...4 stars...
No star
```
