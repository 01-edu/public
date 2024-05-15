## quadchecker

### Instructions

This raid is based on the `quad` functions.

Create a program `quadchecker` that takes a `string` as an argument and displays the name of the matching `quad` and its dimensions.

- If the argument is not a `quad` the program should print `Not a quad function`.

- All answers must end with a newline (`'\n'`).

- If there is more than one `quad` matches, the program must display them all alphabetically ordered and separated by a `||`.

The solution should be submitted as follow:

```console
$ tree .
quadchecker
├── go.mod
└── main.go
```

### Usage

- If it's `quadA`

```console
$ ls -l
-rw-r--r-- 1 student student  nov 23 14:30 main.go
-rwxr-xr-x 1 student student  nov 23 19:18 quadchecker
-rwxr-xr-x 1 student student  nov 23 19:50 quadA
-rwxr-xr-x 1 student student  nov 23 19:50 quadB
-rwxr-xr-x 1 student student  nov 23 19:50 quadC
-rwxr-xr-x 1 student student  nov 23 19:50 quadD
-rwxr-xr-x 1 student student  nov 23 19:50 quadE
$ ./quadA 3 3 | go run .
[quadA] [3] [3]
$
```

- If it's `quadC 1 1` :

```console
$ ./quadC 1 1
A
$ ./quadD 1 1
A
$ ./quadE 1 1
A
$ ./quadC 1 1 | go run .
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
$
```

- If it's `quadC 1 2` :

```console
$ ./quadE 1 2
A
C
$ ./quadC 1 2
A
C
$ ./quadE 1 2 | go run .
[quadC] [1] [2] || [quadE] [1] [2]
$
```

- If it's not a quad function:

```console
$ echo 0 0 | go run .
Not a quad function
$ echo -n "o--o"$'\n'"|"$'\n'"o"
o--o
|
o$ echo -n "o--o"$'\n'"|"$'\n'"o" | go run .
Not a quad function
```
