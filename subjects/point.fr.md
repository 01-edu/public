## point

### Instructions

Créer un fichier `.go`.

-   Le code ci-dessous doit être copié dans ce fichier.

-   Les changements nécéssaires doivent être appliquer et the code below into that file
    and do the necessary changes so that the program works.

-   Le programme doit être rendu dans un dossier nommé `boolean`.

### Code à copier

```go
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}
```

### Usage

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
x = 42, y = 21
student@ubuntu:~/piscine-go/test$
```
