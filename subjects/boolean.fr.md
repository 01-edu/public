## boolean

### Instructions

Créer un fichier `.go`.

- Le code ci-dessous doit être copié dans ce fichier.

- Les changements nécéssaires doivent être appliqués pour que le programme fonctionne.

- Le programme doit être rendu dans un dossier nommé `boolean`.

### Code à copier

```go
func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
```

### Utilisation

```console
student@ubuntu:~/[[ROOT]]/boolean$ go build
student@ubuntu:~/[[ROOT]]/boolean$ ./boolean "not" "odd"
I have an even number of arguments
student@ubuntu:~/[[ROOT]]/boolean$ ./boolean "not even"
I have an odd number of arguments
```
