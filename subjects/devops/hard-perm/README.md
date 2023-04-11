## hard-perm

### Instructions

Create a file `hard-perm.sh`, which will change the default permissions for the files and folders inside the folder `hard-perm`, to the ones below:

Right now the folder looks like this:

```console
$ ls -l hard-perm
total 12
drwxrwxr-x 2 user user 4096 dec 13 18:10 0
-rw-rw-r-- 1 user user    0 dec 13 18:11 1
-rw-rw-r-- 1 user user    0 dec 13 18:11 2
drwxrwxr-x 2 user user 4096 dec 13 18:10 3
-rw-rw-r-- 1 user user    0 dec 13 18:11 4
-rw-rw-r-- 1 user user    0 dec 13 18:11 5
-rw-rw-r-- 1 user user    0 dec 13 18:11 6
-rw-rw-r-- 1 user user    0 dec 13 18:11 7
-rw-rw-r-- 1 user user    0 dec 13 18:11 8
-rw-rw-r-- 1 user user    0 dec 13 18:11 9
drwxrwxr-x 2 user user 4096 dec 13 18:10 A

```

Expected output:

```console
$ ls -l hard-perm
total 12
dr-------x 2 user user 4096 dec 13 17:50 0
-r------w- 1 user user    0 dec 13 17:51 1
-rw----r-- 1 user user    0 dec 13 17:51 2
drwxrwxrwx 2 user user 4096 dec 13 17:51 3
-r-x--x--- 1 user user    0 dec 13 17:51 4
-r--rw---- 1 user user    0 dec 13 17:51 5
-r--rw---- 1 user user    0 dec 13 17:51 6
-r-x--x--- 1 user user    0 dec 13 17:51 7
-rw----r-- 1 user user    0 dec 13 17:51 8
-r------w- 1 user user    0 dec 13 17:51 9
dr-------x 2 user user 4096 dec 13 17:50 A
```

### Hints

- The `chmod`, or change mode, command allows an administrator to set or modify a file’s permissions. Every UNIX/Linux file has an owner user and an owner group attached to it, and every file has permissions associated with it. The permissions are as follows: read, write, or execute.

This is what the default permissions looks like when you create a file.

```console
$ touch example.txt
$ ls -l example.txt
-rw-rw-r-- 1 user user 348 dec 13 15:31 example.txt
$
```

This is what it looks like if you want to give permissions to read, write and execute to every group.

```console
$ chmod 777 example.txt
$ ls -l example.txt
-rwxrwxrwx 1 user user 348 dec 13 15:31 example.txt
$
```

You can also achieve the same result using symbolic notation.

```console
$ chmod a+rwx example.txt
$ ls -l example.txt
-rwxrwxrwx 1 user user 348 dec 13 15:31 example.txt
$
```

In this example we use:

- `"a"` which is a shorthand for user `"u"`, group `"g"`, and others `"o"`.

- The `"+"` sign which specifies that permissions will be added.

- `"rwx"` which is a shorthand for read `"r"`, write `"w"`, and execute `"x"`.

**Symbolic links**, also known as symlinks, are files that act as pointers or aliases to other files or directories on a file system.

To modify the permissions of a `symbolic link`, you would use the same `chmod` command as you would for a regular file or directory. However, the permissions that you set will apply to the `link` itself, not the file or directory that it points to.

```console
$ ls -l my_link
lrwxrwxrwx 1 user user  11 Apr  3 17:35 my_link -> target_file
$
```

The `l` at the beginning of the string indicates that this is a symbolic link, and the next nine characters `rwxrwxrwx` indicate the permissions of the link. The last column `my_link -> target_file` indicates the name of the link followed by an arrow `->` and the name of the target file.

Here are some examples of using chmod with `symbolic links`:

```console
$ ls -l my_link
lrwxrwxrwx 1 user user  11 Apr  3 17:35 my_link -> target_file
$ chmod 600 my_link
$ ls -l my_link
lrw------- 1 user user  11 Apr  3 17:35 my_link -> target_file
$
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [Chmod](https://www.linode.com/docs/guides/modify-file-permissions-with-chmod/#modify-file-permissions-with-chmod)
- [Symbolic notation](https://www.linode.com/docs/guides/modify-file-permissions-with-chmod/#using-symbolic-notation-syntax-with-chmod)
