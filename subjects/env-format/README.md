## Env-format

### Instructions

Create a script `env-format.sh`, which will print the environment variables that have the following formats:

- Print the `PWD` value.
- All the environment variables that have a letter `H` in the name, without printing the value of those variables.

```console
$ printenv
SHELL=/bin/bash
QT_ACCESSIBILITY=1
NVM_RC_VERSION=
COLORTERM=truecolor
HOMEBREW_PREFIX=/home/linuxbrew/.linuxbrew
DESKTOP_SESSION=i3
...
PWD=/home/user
LOGNAME=user
MANPATH=/home/linuxbrew/.linuxbrew/share/man:
XDG_SESSION_DESKTOP=i3
XAUTHORITY=/run/user/1000/gdm/Xauthority
WINDOWPATH=2
SHELL=/bin/bash
SSH_AGENT_PID=2900
HOME=/home/user
HOMEBREW_SHELLENV_PREFIX=/home/linuxbrew/.linuxbrew
USERNAME=user
$
```

Expected output:

```console
$ ./env-format.sh
/home/user/Documents/public/sh/tests/student
SHELL
SSH_AGENT_PID
HOMEBREW_SHELLENV_PREFIX
MANPATH
XAUTHORITY
WINDOWPATH
HOME
$
```

### Hints

Environment variables are variables that contain values necessary to set up a shell environment. Contrary to shell variables, environment variables persist in the shell’s child processes.

```console
VARIABLE_NAME=value
```

Most Common Environment Variables:

- PWD – Current working directory.
- HOME – The user’s home directory location.
- SHELL – Current shell (bash, zsh, etc.).
- LOGNAME – Name of the user.
- UID – User’s unique identifier.
- HOSTNAME – Device’s hostname on the network.
- MAIL – User’s mail directory.
- EDITOR – The Linux system default text editor.
- TEMP – Directory location for temporary files.

How to Check Environment Variables:

Structurally, environment and shell variables are the same – both are a key-value pair, separated by an equal sign.

- `printenv`. Prints the values of all or some environment variables.

```console
$ printenv HOME
/home/user
```

Search Specific Environment Variables:

To find all the variables containing a certain character or pattern, use the `grep` command:

```console
$ printenv | grep "W"
HOMEBREW_PREFIX=/home/linuxbrew/.linuxbrew
HOMEBREW_SHELLENV_PREFIX=/home/linuxbrew/.linuxbrew
PWD=/home/user/Documents/
WINDOWPATH=2
OLDPWD=/home/user/Public
```

- awk. The awk command is a Linux tool and programming language that allows users to process and manipulate data and produce formatted reports.

This flag will come in handy:

`-F` [separator] Used to specify a file separator. The default separator is a blank space.

```console
$ cat example.txt
Hello:my:name:is:John
$ awk -F ":" '{print $5}' example.txt
John
$
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [printenv](https://www.gnu.org/software/coreutils/manual/html_node/printenv-invocation.html#printenv-invocation).
- [grep](https://www.gnu.org/software/grep/manual/grep.html).
- [awk](https://www.ibm.com/docs/en/zos/2.4.0?topic=descriptions-awk-process-programs-written-in-awk-language)
