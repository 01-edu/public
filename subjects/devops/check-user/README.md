## Check User

### Instructions

In this exercise, you will create a script called `check-user.sh` that accepts two arguments and returns information about a specified user. Each output should end with a newline.

The **first argument** is a flag that defines the script’s behavior:

- `-e`: Check if the user exists, returning `yes` or `no` accordingly.
- `-i`: Display information about the user.

The **second argument** is the username to be checked.

The user information must be displayed in the same format as it appears in `/etc/passwd`.

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
