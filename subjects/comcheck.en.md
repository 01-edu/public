## comcheck

### Instructions

Write a program `comcheck` that displays on the standard output `Alert!!!` followed by newline(`'\n'`) if at least one of the arguments passed in parameter matches the `strings`:

- `01`, `galaxy` or `galaxy 01`.

- If none of the parameters match, the program displays a newline(`'\n'`).

### Usage

```console
student@ubuntu:~/piscine/comcheck$ go build
student@ubuntu:~/piscine/comcheck$ ./comcheck "I" "Will" "Enter" "the" "galaxy"
Alert!!!
student@ubuntu:~/piscine/comcheck$ ./comcheck "galaxy 01" "do" "you" "hear" "me"
Alert!!!
student@ubuntu:~/piscine/comcheck$
```
