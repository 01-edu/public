## remake

### Instructions

Create a file `remake.sh`, which will take one argument, the relative path of a directory, and will create new files and directories in it. The new files and directories created must have the same name, permission and modification dates as shown in the example below.

If the number of given arguments is not one or if the directory given as argument does not exist, your script should print `Error` and exit with the status code 1.

Below the expected behavior of your script:

```console
$ bash remake.sh given-path
$ ls -ltr given-path
total 8
-r--r---w- 1 user user    0 Jan  1 00:01 ciao
drwxrwxrwx 2 user user 4096 Jan  2 00:01 mamma
-r-------- 1 user user    0 Jan  3 00:01 guarda
-rw-r---w- 1 user user    0 Jan  4 00:01 come
dr--r-x-w- 2 user user 4096 Jan  5 00:01 mi
-r---w---x 1 user user    0 Jan  6 00:01 diverto
$
```

### Hints

- `touch -t <file-path>` allows to specify the modification time in the format `[[CC]YY]MMDDhhmm[.ss]` instead of the current time for the file at `<file-path>`.

- it is possible to check whether a file or a directory exists with the following commands:

```bash
if [[ -d $DIRPATH ]]; then # for a directory
    echo "the $DIRPATH exists"
fi
```

```bash
if [[ -f $FILEPATH ]]; then # for a file
    echo "the $FILEPATH exists"
fi
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
