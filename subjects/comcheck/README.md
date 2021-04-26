## comcheck

### Instructions

Write a program `comcheck` that displays on the standard output `Alert!!!` followed by newline (`'\n'`) if at least one of the arguments passed in parameter matches the `string`:

- `01`, `galaxy` or `galaxy 01`.

- If none of the parameters match, the program displays a nothing.

### Usage

```console
student@ubuntu:~/comcheck$ go build
student@ubuntu:~/comcheck$ ./comcheck "I" "Will" "Enter" "the" "galaxy"
Alert!!!
student@ubuntu:~/comcheck$ ./comcheck "galaxy 01" "do" "you" "hear" "me"
Alert!!!
student@ubuntu:~/comcheck$
```
