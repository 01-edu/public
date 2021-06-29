## ascii-art

### Objectives

Ascii-art is a program which consists in receiving a `string` as an argument and outputting the `string` in a graphic representation using ASCII. Time to write big.

What we mean by a graphic representation using ASCII, is to write the `string` received using ASCII characters, as you can see in the example below:

```````````console
@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
""WW^^@@@@^^``````^^BB@@^^``^^&&``
^^&&^^@@````^^``&&``@@````^^^^&&``
``WW&&^^""``^^WW&&&&""``^^^^&&88``
^^8888&&&&&&WW88&&88WW&&&&88&&WW``
@@``&&88888888WW&&WW88&&88WW88^^$$
@@""88&&&&&&&&888888&&``^^&&88``$$
@@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
@@@@@@^^888888&&88&&&&MM88^^BB$$$$
@@@@@@BB````&&&&&&&&88""``BB@@BB$$
$$@@$$$$$$$$``````````@@$$@@$$$$$$
```````````

- This project should handle an input with numbers, letters, spaces, special characters and `\n`.
- Take a look at the ASCII manual.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices/).
- It is recommended that the code present a **test file**.

- Some [**banner**](https://github.com/01-edu/public/blob/master/subjects/ascii-art) files (.txt files) with a specific graphical template representation using ASCII will be given. The files are formatted in a way that is not necessary to change them.

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

```

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Usage

```console
$ go run . "hello"
  _              _   _
 | |            | | | |
 | |__     ___  | | | |   ___
 |  _ \   / _ \ | | | |  / _ \
 | | | | |  __/ | | | | | (_) |
 |_| |_|  \___| |_| |_|  \___/


$ go run . "HeLlO"
  _    _          _        _    ____
 | |  | |        | |      | |  / __ \
 | |__| |   ___  | |      | | | |  | |
 |  __  |  / _ \ | |      | | | |  | |
 | |  | | |  __/ | |____  | | | |__| |
 |_|  |_|  \___| |______| |_|  \____/


$ go run . "Hello There"
  _    _           _    _                 _______   _
 | |  | |         | |  | |               |__   __| | |
 | |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___
 |  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \
 | |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/
 |_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___|


$ go run . "1Hello 2There"
     _    _           _    _                       _______   _
 _  | |  | |         | |  | |               ____  |__   __| | |
/ | | |__| |   ___   | |  | |    ___       |___ \    | |    | |__      ___    _ __     ___
| | |  __  |  / _ \  | |  | |   / _ \        __) |   | |    |  _ \    / _ \  | '__|   / _ \
| | | |  | | |  __/  | |  | |  | (_) |      / __/    | |    | | | |  |  __/  | |     |  __/
|_| |_|  |_|  \___|  |_|  |_|   \___/      |_____|   |_|    |_| |_|   \___|  |_|      \___|


$ go run . "{Hello There}"
   __   _    _           _    _                 _______   _                              __
  / /  | |  | |         | |  | |               |__   __| | |                             \ \
 | |   | |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___   | |
/ /    |  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \   \ \
\ \    | |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/   / /
 | |   |_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___|  | |
  \_\                                                                                    /_/

$
```

This project will help you learn about :

- The Go file system(**fs**) API
- Data manipulation
