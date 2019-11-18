## ascii-art

### Objectives

Ascii-art consists on receiving a `string` has an argument and outputing the `string` in a graphic representation of ascii letters and numbers.

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
- It will be given a `banner` file with a specific graphical template representation of ascii letters and numbers.

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello"
  _                _    _           
 | |              | |  | |          
 | |__      ___   | |  | |    ___   
 |  _ \    / _ \  | |  | |   / _ \  
 | | | |  |  __/  | |  | |  | (_) | 
 |_| |_|   \___|  |_|  |_|   \___/  
                                    
                                    
student@ubuntu:~/ascii-art$ ./ascii-art "HeLlO"
  _    _           _        _     ____   
 | |  | |         | |      | |   / __ \  
 | |__| |   ___   | |      | |  | |  | | 
 |  __  |  / _ \  | |      | |  | |  | | 
 | |  | | |  __/  | |____  | |  | |__| | 
 |_|  |_|  \___|  |______| |_|   \____/  
                                         
                                         
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
                                                                                            
                                                                                            
student@ubuntu:~/ascii-art$
```
