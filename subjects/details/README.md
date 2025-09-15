## details

### Instructions

Create a script `details.sh` that does the following:

- Change the size of the `file1.txt` to "1000".
- Change the permissions of the `file1.txt`:

```console
$ ls -l   # from this:
-rw-rw-r-- 1 user user     0 dez 27 12:25  file1.txt
$ ls -l   # to this:
-rw------- 1 user user     0 dez 27 12:25  file1.txt
$
```

- Update both `Modification` and `Access time` of the `file1.txt` to `2022-01-01`.

### Usage

```console
$ stat file1.txt
  File: file1.txt
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: ff01h/64769e    Inode: 1839372     Links: 1
Access: (0664/-rw-rw-r--)  Uid: ( 1000/  user)   Gid: ( 1000/  user)
Access: 2022-12-27 12:44:56.397966181 +0000
Modify: 2022-12-27 12:44:56.397966181 +0000
Change: 2022-12-27 12:44:56.397966181 +0000
 Birth: -
$ ./details.sh
$ stat file1.txt
  File: file1.txt
  Size: 1000      	Blocks: 0          IO Block: 4096   regular file
Device: ff01h/64769e	Inode: 1839472     Links: 1
Access: (0600/-rw-------)  Uid: ( 1000/  user)   Gid: ( 1000/  user)
Access: 2022-01-01 00:00:00.000000000 +0000
Modify: 2022-01-01 00:00:00.000000000 +0000
Change: 2022-12-27 14:53:43.886486188 +0000
 Birth: -
$
```

### Hints

Use the `stat` command to view the current details of the file, including the last modification time and size.

```console
$ stat file.txt
File: file.txt
  Size: 4030      	Blocks: 8          IO Block: 4096   regular file
Device: 801h/2049d	Inode: 13633379    Links: 1
Access: (0644/-rw-r--r--)  Uid: ( 1000/   linuxize)   Gid: ( 1000/   linuxize)
Access: 2019-11-06 09:52:17.991979701 +0100
Modify: 2019-11-06 09:52:17.971979713 +0100
Change: 2019-11-06 09:52:17.971979713 +0100
 Birth: -
```

The `touch` command's primary function is to modify a timestamp. Commonly, the utility is used for file creation, although this is not its primary function. The terminal program can change the modification and access time for any given file.

The fundamental syntax for the touch command is:

`touch <options> <file or directory name>`

Some of the touch Command Options:

- `-a` Changes the access time.
- `-d=<string>` Changes a timestamp using a date string.
- `-m` Changes the modification time.

```console
$ touch test.txt
$ ls -lu
-rw-rw-r-- 1 user user     0 dez 27 12:13  test.txt
$ touch -a test.txt
$ ls -lu
-rw-rw-r-- 1 user user     0 dez 27 12:20  test.txt
$ touch -d tomorrow test.txt
$ ls -l
-rw-rw-r-- 1 user user     0 dez 28  2022  test.txt
$ touch -m test.txt
$ ls -l
-rw-rw-r-- 1 user user     0 dez 27 12:25  testing.txt
```

Sometimes we need to remove the content of a file without deleting the file. For that Linux operating system offers a command called `truncate`. It is used to extend or reduce the file size. Truncating a file is much quicker and simpler without modifying the permissions and ownership of the file.

```console
$ touch test.txt
$ ls -l
-rw-rw-r-- 1 user user     0 dez 27 12:13  test.txt
$ truncate -s 100 test.txt
$ ls -l
-rw-rw-r-- 1 user user   100 dez 27 12:15  test.txt
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
