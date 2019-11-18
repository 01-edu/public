## ascii-art-reverse

### Objectives

Ascii-art-reverse consists on reversing the process, convertin the graphic representation into a text.

- You will have to create a `txt` file with a graphic representation of a random phrase;
- The program will have **one** argument being the name of the file created;
- The program must print this phrase in **normal text**.

This project will help you learn about :

- Client utilities;
- The go file system(**fs**) API;
- Ways to receive data;
- Ways to output data;
- Manipulation of strings;
- Manipulation of structures.

### Instructions

- Your project must be written in **GO**;
- The code must respect the **good practices**;
- It is recommended that the code should present a **test file**;

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ cat file.txt
  _                _    _           
 | |              | |  | |          
 | |__      ___   | |  | |    ___   
 |  _ \    / _ \  | |  | |   / _ \  
 | | | |  |  __/  | |  | |  | (_) | 
 |_| |_|   \___|  |_|  |_|   \___/  
                                    
                                    
student@ubuntu:~/ascii-art$ ./ascii-art file.txt
hello
student@ubuntu:~/ascii-art$
```
