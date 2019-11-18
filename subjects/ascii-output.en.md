## ascii-art-output

### Objectives

Ascii-art-output consists on receiving two strings. The first argument will be converted into a graphic representation of ASCII letters and numbers and written into a file named by using the second argument.

- In case the second argument is not present it should print the graphic representation.

This project will help you learn about :

- Client utilities.
- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Learning about the choice of outputs.
- Manipulation of structures.

### Instructions

- Your project must be written in **Go**.
- The code must respect the **good practices**.
- It is recommended that the code should present a **test file**.
- You may use the same `banner` file.

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello" banner.txt
student@ubuntu:~/ascii-art$ cat banner.txt
  _                _    _           
 | |              | |  | |          
 | |__      ___   | |  | |    ___   
 |  _ \    / _ \  | |  | |   / _ \  
 | | | |  |  __/  | |  | |  | (_) | 
 |_| |_|   \___|  |_|  |_|   \___/  
                                    
                                    
student@ubuntu:~/ascii-art$ ./ascii-art "Hello There" banner.txt
student@ubuntu:~/ascii-art$ cat banner.txt
  _    _           _    _                 _______   _                              
 | |  | |         | |  | |               |__   __| | |                             
 | |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___  
 |  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \ 
 | |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/ 
 |_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___| 
                                                                                   
                                                                                   
student@ubuntu:~/ascii-art$
```
