## hard-perm

### Instructions

Create a file `hard-perm.sh`, which will change the default permissions for the files and folders inside the folder `hard-perm`, to the ones bellow:

Right now the folder looks like this:

```console
$ ls -l hard-perm
total 12
drwxrwxr-x 2 user user 4096 dez 13 18:10 0
-rw-rw-r-- 1 user user    0 dez 13 18:11 1
-rw-rw-r-- 1 user user    0 dez 13 18:11 2
drwxrwxr-x 2 user user 4096 dez 13 18:10 3
-rw-rw-r-- 1 user user    0 dez 13 18:11 4
-rw-rw-r-- 1 user user    0 dez 13 18:11 5
-rw-rw-r-- 1 user user    0 dez 13 18:11 6
-rw-rw-r-- 1 user user    0 dez 13 18:11 7
-rw-rw-r-- 1 user user    0 dez 13 18:11 8
-rw-rw-r-- 1 user user    0 dez 13 18:11 9
drwxrwxr-x 2 user user 4096 dez 13 18:10 A

```

Expected output:

```console
$ ls -l hard-perm
total 12
dr-------x 2 user user 4096 dez 13 17:50 0
-r------w- 1 user user    0 dez 13 17:51 1
-rw----r-- 1 user user    0 dez 13 17:51 2
drwxrwxrwx 2 user user 4096 dez 13 17:51 3
-r-x--x--- 1 user user    0 dez 13 17:51 4
-r--rw---- 1 user user    0 dez 13 17:51 5
-r--rw---- 1 user user    0 dez 13 17:51 6
-r-x--x--- 1 user user    0 dez 13 17:51 7
-rw----r-- 1 user user    0 dez 13 17:51 8
-r------w- 1 user user    0 dez 13 17:51 9
dr-------x 2 user user 4096 dez 13 17:50 A
```

### Hints

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [Chmod](https://www.linode.com/docs/guides/modify-file-permissions-with-chmod/#modify-file-permissions-with-chmod)
