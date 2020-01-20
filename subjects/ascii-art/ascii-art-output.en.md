## ascii-art-output

### Objectives

- You must follow the same [instructions](https://public.01-edu.org/subjects/ascii-art/ascii-art.en) as in the first subject but writing the result into a file.

- The file must be named by using the flag `--output=<fileName.txt>`, in which `--output` is the flag and `<fileName.txt>` is the file name.

This project will help you learn about :

- Client utilities.
- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Learning about the choice of outputs.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommended that the code should present a **test file**.

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard --output=banner.txt
student@ubuntu:~/ascii-art$ cat banner.txt
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
student@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" shadow --output=banner.txt
student@ubuntu:~/ascii-art$ cat banner.txt
                                                                                         
_|    _|          _| _|                _|_|_|_|_| _|                                  _| 
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| 
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| 
                                                                                         
                                                                                         
student@ubuntu:~/ascii-art$
```
