## easy-perm

### Instructions

Create a file `easy-perm.sh`, which will change the default permissions for the `example.txt` and `example2.txt` files inside the folder `easy-perm`, to the ones below:

Expected Output:

```console
$ ls -l easy-perm
total 8
-rwxr--rw- 1 user user 689 dez 13 16:14 example2.txt
-rw-r--r-- 1 user user 348 dez 13 16:14 example.txt
$
```

### Hints

-`chmod` The chmod, or change mode, command allows an administrator to set or modify a file’s permissions. Every UNIX/Linux file has an owner user and an owner group attached to it, and every file has permissions associated with it. The permissions are as follows: read, write, or execute.

This is what the default permissions looks like when you create a file.

```console
$ touch example.txt
$ ls -l example.txt
-rw-rw-r-- 1 user user 348 dez 13 15:31 example.txt
$
```

This is what it looks like if you want to give permissions to read, write and execute to every group.

```console
$ chmod 777 example.txt
$ ls -l example.txt
-rwxrwxrwx 1 user user 348 dez 13 15:31 example.txt
$
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [Chmod](https://www.linode.com/docs/guides/modify-file-permissions-with-chmod/#modify-file-permissions-with-chmod)
