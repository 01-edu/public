## reversestrcap

### Instructions

Écrire un programme avec une ou plusieurs `string` comme arguments et qui, **pour chaque argument**:
-place le dernier caractère de chaque mot (si c'est une lettre) en majuscule et le reste en minuscule.
-affiche ensuite le résultat suivi d'un retour à la ligne (`'\n'`).

Un mot est une suite de caractères alphanumériques.

Si il n'y a pas de paramètre, le programme affiche un retour à la ligne.

### Utilisation

```console
student@ubuntu:~/[[ROOT]]/reversestrcap$ go build
student@ubuntu:~/[[ROOT]]/reversestrcap$ ./reversestrcap "First SMALL TesT" | cat -e
firsT smalL tesT$
student@ubuntu:~/[[ROOT]]/reversestrcap$ ./reversestrcap go run reversestrcap.go "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
seconD tesT iS A littlE easieR$
bewarE it'S noT harD wheN $
 gO A dernieR 0123456789 foR thE roaD E$
student@ubuntu:~/[[ROOT]]/reversestrcap$ ./reversestrcap | cat -e
$
student@ubuntu:~/[[ROOT]]/reversestrcap$
```
