## comcheck

### Instructions

écrire un programme `comcheck` qui affiche sur la sortie standard `Alert!!!` suivi d'un newline(`'\n'`) si au moins un des arguments passé ne paramètre correspond aux `strings`:

- `01`, `galaxy` ou `galaxy 01`.

- si aucun des paramètres correspeond, le programme affiche un newline(`'\n'`).

### Usage

```console
student@ubuntu:~/piscine/comcheck$ go build
student@ubuntu:~/piscine/comcheck$ ./comcheck "I" "Will" "Enter" "the" "galaxy"
Alert!!!
student@ubuntu:~/piscine/comcheck$ ./comcheck "galaxy 01" "do" "you" "hear" "me"
Alert!!!
student@ubuntu:~/piscine/comcheck$
```
