## set-internal-vars

### Instructions

Create a script `set-internal-vars.sh`, which will allow you to create and print the following variables.

- `MY_MESSAGE` which contains the string `"Hello World"`.
- `MY_NUM` which contains the number `100`.
- `MY_PI` which contains the number `3.142`.
- `MY_ARR` which contains `(one two three four five)`

Expected output:

```console
$ ./set-internal-vars.sh
Hello World
100
3.142
one, two, three, four, five
$
```

### Hints

Creating variables:

Variables can be created either at the shell or in shell-scripts. Any variable created within a shell script is lost when the script stops executing. A variable created at the prompt, however, will remain in existence until the shell is terminated. The syntax for creating a variable is :

`<variable name>=<value>`

If we type the `set` command without any additional parameters, we will get a list of all shell variables, environmental variables, local variables, and shell functions.

Use the `unset` command to remove a variable from your shell environment.

```console
$ MyVar=8472
$ echo $MyVar
8472
$ unset MyVar
$ echo $MyVar

$
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
