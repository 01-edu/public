## ascii-art-fs

### Objectives

Ascii-art-fs consists on receiving two `strings` has arguments. The first string being a random phrase and the second string being the template for the graphic representation in ascii of the first argument.

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
- You should build your own files with the templates you desire.

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard
  _                _    _           
 | |              | |  | |          
 | |__      ___   | |  | |    ___   
 |  _ \    / _ \  | |  | |   / _ \  
 | | | |  |  __/  | |  | |  | (_) | 
 |_| |_|   \___|  |_|  |_|   \___/  
                                    
                                    
student@ubuntu:~/ascii-art$ ./ascii-art "hello" shadow
 oooo                    oooo   oooo             
 `888                    `888   `888             
  888 .oo.     .ooooo.    888    888    .ooooo.  
  888P"Y88b   d88' `88b   888    888   d88' `88b 
  888   888   888ooo888   888    888   888   888 
  888   888   888    .o   888    888   888   888 
 o888o o888o  `Y8bod8P'  o888o  o888o  `Y8bod8P' 
                                                 
student@ubuntu:~/ascii-art$
student@ubuntu:~/ascii-art$
```
