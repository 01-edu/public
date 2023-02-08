## largest

### Instructions

Crate a script `largest.sh` that will look for, and display on the console, the seven biggest files in the current directory (including subdirectories) in a specific format.

The script will also have to find and take into account hidden files.

### Usage

Here is an example of how to use the script:

```console
$  tree -ha
[4.0K]  .
├── [4.0K]  d1
│   ├── [4.0K]  d1-1
│   ├── [   2]  f1-1
│   ├── [  27]  f1-2
│   ├── [  24]  f1-3
│   ├── [  24]  f1-4
│   ├── [   3]  f2-1
│   ├── [  34]  f2-2
│   ├── [  38]  f2-3
│   └── [  87]  .hf1-1
├── [   5]  f1
├── [  18]  f2
└── [  34]  f3

2 directories, 11 files
$ bash largest.sh | cat -e
   87 | ./d1/.hf1-1$
    5 | ./f1$
    3 | ./d1/f2-1$
   38 | ./d1/f2-3$
   34 | ./d1/f2-2$
   34 | ./f3$
    2 | ./d1/f1-1$
$
```

> The output needs to be formatted as the example shown above.

### Hints

- `awk` might be handy.
