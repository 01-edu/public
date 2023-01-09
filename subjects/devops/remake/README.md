## remake

### Instructions

Create a file `remake.sh`, which will take one argument, the relative path of a directory, and will create new files and directories in it.

If the number of given arguments is not one, your script should print `Error` and exit with the status code 1.

Below the expected behavior of your script: 

```console
$  bash remake.sh given-path
$  ls -ltr given-path 
total 8
-r--r---w- 1 nprimo nprimo    0 Jan  1 00:01 ciao 
drwxrwxrwx 2 nprimo nprimo 4096 Jan  2 00:01 mamma
-r-------- 1 nprimo nprimo    0 Jan  3 00:01 guarda
-rw-r---w- 1 nprimo nprimo    0 Jan  4 00:01 come 
dr--r-x-w- 2 nprimo nprimo 4096 Jan  5 00:01 mi
-r---w---x 1 nprimo nprimo    0 Jan  6 00:01 diverto
$ 
```

### Hints

- `mkdir <relative-path>` command is used to create a new directory in the specified `<relative-path>`. For example:

```console
$  ls -l
total 0
-rw-rw-r-- 1 nprimo nprimo 0 Jan 12 14:26 a
-rw-rw-r-- 1 nprimo nprimo 0 Jan 12 14:26 b
-rw-rw-r-- 1 nprimo nprimo 0 Jan 12 14:26 c
$  mkdir d
$  ls -l 
total 4
-rw-rw-r-- 1 nprimo nprimo    0 Jan 12 14:26 a
-rw-rw-r-- 1 nprimo nprimo    0 Jan 12 14:26 b
-rw-rw-r-- 1 nprimo nprimo    0 Jan 12 14:26 c
drwxrwxr-x 2 nprimo nprimo 4096 Jan 12 14:26 d
$
```

- `touch <file-path>` command is used to change the modification and/or access time of the specified `<file-path>` to the current time. If the file does not exist yet, a new empty file is created at the specified `<file-path>`.

The flag `-t` allow to specify the time in the format `[[CC]YY]MMDDhhmm[.ss]` instead of the current time.

- `chmod` The chmod, or change mode, command allows an administrator to set or modify a file’s permissions. Every UNIX/Linux file has an owner user and an owner group attached to it, and every file has permissions associated with it. The permissions are as follows: read, write, or execute.

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
