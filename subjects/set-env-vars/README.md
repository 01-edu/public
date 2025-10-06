## set-env-vars

### Instructions

Create a script `set-env-vars.sh`, which will allow you to set the following variables as environment variables and print only the ones you created:

- `MY_MESSAGE` which contains the string `"Hello World"`.
- `MY_NUM` which contains the number `100`.
- `MY_PI` which contains the number `3.142`.

Expected output:

```console
$ printenv  # The env variables present are just an example, yours will be different.
SHELL=/bin/bash
TERM=xterm
USER=demouser
MAIL=/var/mail/demouser
PATH=/usr/local/bin:/usr/bin:/bin:/usr/local/games:/usr/games
PWD=/home/demouser
SHLVL=1
HOME=/home/demouser
LOGNAME=demouser
_=/usr/bin/printenv
$ ./set-env-vars.sh
MY_NUM=100
MY_PI=3.142
MY_MESSAGE=Hello World
$
```

### Hints

Setting values to environment variables:

In order to set a value to an existing environment variable, we use an assignment expression. For instance, to set the value of the "LANG" variable to "he_IL.UTF-8", we use the following command:

```console
$ LANG=he_IL.UTF-8
```

If we use an assignment expression for a variable that doesn't exist, the shell will create a shell variable, which is similar to an environment variable but does not influence the behavior of other applications.

A shell variable can be exported to become an environment variable with the export command. To create the "EDITOR" environment variable and assign the value "nano" to it, you can do:

```console
$ EDITOR=nano
$export EDITOR
```

`echo` is not allowed in the exercise! Try to use the command `printenv` and filter with `grep` to get only the variables that you created.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
