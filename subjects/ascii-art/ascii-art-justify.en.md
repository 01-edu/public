## ascii-art-justify

### Objectives

You must follow the same [instructions](https://github.com/01-edu/public/blob/master/subjects/ascii-art/ascii-art.en.md) as in the first subject but the representation should be formatted using a **flag** `--align=<type>`, in which `type` can be :

- center
- left
- right
- justify

- You must adapt your representation to the terminal size. If you reduce the terminal window the graphical representation should be adapted to the terminal size.

This project will help you learn about :

- Client utilities.
- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Manipulation of structures.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://github.com/01-edu/public/blob/master/subjects/good-practices.en.md).
- It is recommended that the code should present a **test file**.

### Usage

```console
|student@ubuntu:~/ascii-art$ go build                                                                                       |
|student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard --align=center                                                    |
|                                             _                _    _                                                       |
|                                            | |              | |  | |                                                      |
|                                            | |__      ___   | |  | |    ___                                               |
|                                            |  _ \    / _ \  | |  | |   / _ \                                              |
|                                            | | | |  |  __/  | |  | |  | (_) |                                             |
|                                            |_| |_|   \___|  |_|  |_|   \___/                                              |
|                                                                                                                           |
|                                                                                                                           |
|student@ubuntu:~/ascii-art$ ./ascii-art "Hello There" standard --align=left                                                |
| _    _           _    _                 _______   _                                                                       |
|| |  | |         | |  | |               |__   __| | |                                                                      |
|| |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___                                           |
||  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \                                          |
|| |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/                                          |
||_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___|                                          |
|                                                                                                                           |
|                                                                                                                           |
|student@ubuntu:~/ascii-art$ ./ascii-art "hello" shadow --align=right                                                       |
|                                                                                                                           |
|                                                                                          _|                _| _|          |
|                                                                                          _|_|_|     _|_|   _| _|   _|_|   |
|                                                                                          _|    _| _|_|_|_| _| _| _|    _| |
|                                                                                          _|    _| _|       _| _| _|    _| |
|                                                                                          _|    _|   _|_|_| _| _|   _|_|   |
|                                                                                                                           |
|                                                                                                                           |
|student@ubuntu:~/ascii-art$ ./ascii-art "how are you" shadow --align=justify                                               |
|                                                                                                                           |
|_|                                                                                                                         |
|_|_|_|     _|_|   _|      _|      _|                  _|_|_| _|  _|_|   _|_|                    _|    _|   _|_|   _|    _| |
|_|    _| _|    _| _|      _|      _|                _|    _| _|_|     _|_|_|_|                  _|    _| _|    _| _|    _| |
|_|    _| _|    _|   _|  _|  _|  _|                  _|    _| _|       _|                        _|    _| _|    _| _|    _| |
|_|    _|   _|_|       _|      _|                      _|_|_| _|         _|_|_|                    _|_|_|   _|_|     _|_|_| |
|                                                                                                      _|                   |
|                                                                                                  _|_|                     |
|student@ubuntu:~/ascii-art$                                                                                                |
```
