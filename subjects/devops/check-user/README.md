## Check User

### Instructions

In this exercise you will make a script `check-user.sh` that will take 2 arguments and return information about the selected user, always ended by a new line.

The first argument will be a flag defining the behavior of the script:
- `-e`: check if the user exists, returns `yes` or `no` appropriately.
- `-i`: returns information about the user.

The second argument will be the name of the checked user.

> The information about the user will be formatted in the same way it appears in `/etc/passwd`.

### Usage

```console
$ ./check-user.sh -e root
yes
$ ./check-user.sh -i root
root:x:0:0:root:/root:/bin/bash
$ ./check-user.sh -e unknown
no
$ ./check-user.sh -i unknown
$
```

> Your results may appear slightly different.

### Error handling

All errors will print a specific message on **stderr** (ending with a newline) and returns a specific non-zero value:
- Wrong number of arguments: `"Error: expect 2 arguments"`, returns `1`.
- First argument different from `-e` or `-i`: `"Error: unknown flag"`, exit with `1`.

### Hints

> `man getent` will be a great resource to explore  
> [List Linux users](https://linuxize.com/post/how-to-list-users-in-linux/)  
