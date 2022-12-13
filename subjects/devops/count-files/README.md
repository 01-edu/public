## Count files

### Instructions

Create a file `count-files.sh`, which will print the number **(and only the number)** of regular files and folders contained in the current directory and its sub-folders (the current directory must be included in the count).

### Usage

```console
$ ./count-files.sh | cat -e
12$
$
```

### Hints

Here are some Commands that can help you:

- `find`. Find files or directories under the given directory tree, recursively.

Find all php files in a directory:

```console
$ find . -type f -name "*.php"

./tec.php
./login.php
./index.php
```

- `wc`. Count lines, words, and bytes.

```console
$ cat README.md
this is a simple example
with two lines
$ wc -lwc example.md
 2  8 40 example.md
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [Find](https://manned.org/find).
- [wc](https://www.gnu.org/software/coreutils/wc).
