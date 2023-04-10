## Env-format

> 🕵️ **Operation Data Centre**
>
> [Confidential]
>
> You are on a secret mission inside an underground data centre. You are alone, there is no internet access and there is no mobile phone signal.
>
> You cannot interact with anyone.
>
> You cannot not access the internet - although you can read the mission below.
>
> You may access man pages via the command line. This is your only resource.
>
> Good luck agent, we're counting on your knowledge.

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
