## brackets

### Instructions

Écrire un programme qui prend un nombre indéfini de `strings` en arguments. Pour chaque
argument, si l'expression est correctement "entre parenthèses" (bracketed), le programme affiche sur la sortie standard `OK` suivi d'un newline(`'\n'`), autrement il affiche `Error` suivi d'un newline.

Les symboles considérés comme des `parenthèses` sont les parenthèses `(` et `)`, les crochets `[`
et `]` et les accolades `{` et `}`. Tout autre symbole est simplement ignoré.

Une parenthèse ouvrante doit toujours être fermée par la parenthèse correspondante dans l'ordre correct. Une `string` qui ne contient aucune parenthèse est considérée comme une `string` correctement "entre parenthèses".

Si il n'y a pas d'argument, le programme affiche seulement un newline.

### Utilisation

```console
student@ubuntu:~/student/brackets$ go build
student@ubuntu:~/student/brackets$ ./brackets '(johndoe)' | cat -e
OK$
student@ubuntu:~/student/brackets$ ./brackets '([)]' | cat -e
Error$
student@ubuntu:~/student/brackets$ ./brackets '' '{[(0 + 0)(1 + 1)](3*(-1)){()}}' | cat -e
OK$
OK$
student@ubuntu:~/student/brackets$ ./brackets | cat -e
$
student@ubuntu:~/student/brackets$

```
