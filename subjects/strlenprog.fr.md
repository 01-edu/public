## strlenprog

##**AVERTISSEMENT! TRÈS IMPORTANT!**

Pour cet exercice une fonction sera testée **avec le main de l'examen**. Cependant l'étudiant **doit quand même** rendre un programme structuré:

Cela signifie que:

- Le package doit être nommé `package main`.
- Le code rendu doit avoir une fonction main déclarée(```func main()```) même si elle est vide.
- La fonction main déclarée doit **aussi passer** le `Restrictions Checker`(le testeur de fonctions illégales). Il est conseillé à l'étudiant de rendre une fonction main vide après ses tests finis.
- Toutes les autres régles sont les mêmes que pour un `programme`.

### Instructions

- Écrire une fonction qui compte le nombre de `runes` d'une `string` et qui retourne le nombre trouvé.

### Fonction attendue

```go
func StrLen(str string) int {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import (
	"fmt"
)

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
12
student@ubuntu:~/[[ROOT]]/test$
```
