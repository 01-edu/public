## ascii-art-fs

### Objectives

You must follow the same [instructions](https://github.com/01-edu/public/ascii-art.en.md) as the first subject but the second argument must be the name of the template.

This project will help you learn about :

- Client utilities.
- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Manipulation of structures.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://github.com/01-edu/public/good-practices.en.md).
- It is recommended that the code should present a **test file**.

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
