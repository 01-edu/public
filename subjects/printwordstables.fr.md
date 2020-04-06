## printwordstables

### Instructions

Écrire une fonction qui affiche les mots d'un tableau de `string` qui sera le resultat d'une fonction `SplitWhiteSpaces`. (les tests seront effectués avec la notre)

Chaque mot est sur une seule ligne.

Chaque mot fini avec un `\n`.

### Fonction attendue

```go
func PrintWordsTables(table []string) {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import piscine ".."

func main() {

	str := "Hello how are you?"
	table := piscine.SplitWhiteSpaces(str)
	piscine.PrintWordsTables(table)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build main.go
student@ubuntu:~/[[ROOT]]/test$ ./main
Hello
how
are
you?
student@ubuntu:~/[[ROOT]]/test$
```
