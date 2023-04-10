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
