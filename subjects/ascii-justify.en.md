## ascii-art-justify

### Objectives

You must follow the same [instructions](https://github.com/01-edu/public/ascii-art.en.md) as the first subject but the representation should be formatted using the second argument, that can be:

- center
- justify
- align-left
- align-right

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
- You should build your one files with the templates you desire.

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard center
                                                         _                _    _           
                                                        | |              | |  | |          
                                                        | |__      ___   | |  | |    ___   
                                                        |  _ \    / _ \  | |  | |   / _ \  
                                                        | | | |  |  __/  | |  | |  | (_) | 
                                                        |_| |_|   \___|  |_|  |_|   \___/  
                                                                                           
                                                                                           
student@ubuntu:~/ascii-art$ ./ascii-art "Hello There" standard align-right
                                                          _    _           _    _                 _______   _                              
                                                         | |  | |         | |  | |               |__   __| | |                             
                                                         | |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___  
                                                         |  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \ 
                                                         | |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/ 
                                                         |_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___| 
                                                                                                                                           
                                                                                                                                           
student@ubuntu:~/ascii-art$ ./ascii-art "hello" shadow center
                                                   oooo                    oooo   oooo             
                                                   `888                    `888   `888             
                                                    888 .oo.     .ooooo.    888    888    .ooooo.  
                                                    888P"Y88b   d88' `88b   888    888   d88' `88b 
                                                    888   888   888ooo888   888    888   888   888 
                                                    888   888   888    .o   888    888   888   888 
                                                   o888o o888o  `Y8bod8P'  o888o  o888o  `Y8bod8P' 
                                                                                                   
student@ubuntu:~/ascii-art$
```
