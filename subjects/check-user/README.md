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

- Wrong number of arguments: `"Error: expect 2 arguments"`, exit with `1`.
- First argument different from `-e` or `-i`: `"Error: unknown flag"`, exit with `1`.

### Hints

- `getent` is a command to get entries from a database. `passwd` is the database where information about users is stored.
- `getent passwd` will give you the list of all users.
- `getent passwd <username>` will give you information about a specific user.
- If the user doesn't exists `getent` returns an empty string, use this at your advantage for `-e` flag.

> `man getent` will provide extensive documentation about this command.

### Resources

> [List Linux users](https://linuxize.com/post/how-to-list-users-in-linux/)
