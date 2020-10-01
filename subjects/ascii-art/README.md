## ascii-art

### Objectives

Ascii-art consists on receiving a `string` as an argument and outputting the `string` in a graphic representation of ASCII.

- This project should handle numbers, letters, spaces, special characters and `\n`.
- Take a look at the ASCII manual.

This project will help you learn about :

- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Manipulation of structures.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices/).
- It is recommended that the code present a **test file**.

- It will be given some [**banner**](https://github.com/01-edu/public/blob/master/subjects/ascii-art) files with a specific graphical template representation of ASCII. The files are formatted in a way that it is not necessary to change them.

### Banner Format

- Each character has an height of 8 lines.
- Characters are separated by a new line `\n`.
- Here is an example of ' ', '!' and '"'(one dot represents one space) :

```console

......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

._._..
(.|.).
.V.V..
......
......
......
......
......

etc
```

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello"
  _              _   _
 | |            | | | |
 | |__     ___  | | | |   ___
 |  _ \   / _ \ | | | |  / _ \
 | | | | |  __/ | | | | | (_) |
 |_| |_|  \___| |_| |_|  \___/


student@ubuntu:~/ascii-art$ ./ascii-art "HeLlO"
  _    _          _        _    ____
 | |  | |        | |      | |  / __ \
 | |__| |   ___  | |      | | | |  | |
 |  __  |  / _ \ | |      | | | |  | |
 | |  | | |  __/ | |____  | | | |__| |
 |_|  |_|  \___| |______| |_|  \____/


student@ubuntu:~/ascii-art$ ./ascii-art "Hello There"
  _    _           _    _                 _______   _
 | |  | |         | |  | |               |__   __| | |
 | |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___
 |  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \
 | |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/
 |_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___|


student@ubuntu:~/ascii-art$ ./ascii-art "1Hello 2There"
     _    _           _    _                       _______   _
 _  | |  | |         | |  | |               ____  |__   __| | |
/ | | |__| |   ___   | |  | |    ___       |___ \    | |    | |__      ___    _ __     ___
| | |  __  |  / _ \  | |  | |   / _ \        __) |   | |    |  _ \    / _ \  | '__|   / _ \
| | | |  | | |  __/  | |  | |  | (_) |      / __/    | |    | | | |  |  __/  | |     |  __/
|_| |_|  |_|  \___|  |_|  |_|   \___/      |_____|   |_|    |_| |_|   \___|  |_|      \___|


student@ubuntu:~/ascii-art$ ./ascii-art "{Hello There}"
   __   _    _           _    _                 _______   _                              __
  / /  | |  | |         | |  | |               |__   __| | |                             \ \
 | |   | |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___   | |
/ /    |  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \   \ \
\ \    | |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/   / /
 | |   |_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___|  | |
  \_\                                                                                    /_/

student@ubuntu:~/ascii-art$
```
