## splitprog

## **AVERTISSEMENT! TRÈS IMPORTANT!**

Pour cet exercice une fonction sera testée **avec le main de l'examen**. Cependant l'étudiant **doit quand même** rendre un programme structuré:

Cela signifie que:

- Le package doit être nommé `package main`.
- Le code rendu doit avoir une fonction main déclarée(```func main()```) même si elle est vide.
- La fonction main déclarée doit **aussi passer** le `Restrictions Checker`(le testeur de fonctions illégales). Il est conseillé à l'étudiant de rendre une fonction main vide après ses tests finis.
- Toutes les autres régles sont les mêmes que pour un `programme`.

### Instructions

Écrire une fonction qui sépare les mots d'une `string`, qui les met dans un tableau de `string`.

Les séparateurs sont les caractères de la `charset string` donnée en paramètre.

### Fonction attendue

```go
func Split(str, charset string) []string {

}
```

### Utilsation

Voici un programme possible pour tester votre fonction :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(piscine.Split(str, "HA"))
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[Hello how are you?]
student@ubuntu:~/[[ROOT]]/test$
```
