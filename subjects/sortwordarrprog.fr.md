## sortwordarrprog

##**AVERTISSEMENT! TRÈS IMPORTANT!**

Pour cet exercice une fonction sera testée **avec le main de l'examen**. Cependant l'étudiant **doit quand même** rendre un programme structuré:

Cela signifie que:

- Le package doit être nommé `package main`.
- Le code rendu doit avoir une fonction main déclarée(```func main()```) même si elle est vide.
- La fonction main déclarée doit **aussi passer** le `Restrictions Checker`(le testeur de fonctions illégales). Il est conseillé à l'étudiant de rendre une fonction main vide après ses tests finis.
- Toutes les autres régles sont les mêmes que pour un `programme`.

### Instructions

Écrire une fonction `SortWordArr` qui trie par ordre `ascii` (dans l'ordre croissant) un tableau de `string`.

### Fonction attendue

```go
func SortWordArr(array []string) {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
	"fmt"
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}
```

Et son résultat :

```console
student@ubuntu:~/test$ go build
student@ubuntu:~/test$ ./test
[1 2 3 A B C a b c]
student@ubuntu:~/test$
```
