# Active

### Objective

In this project you will have to make a simple port scanner, which will tell you if the port is open or closed.
You must create the project from scratch.

### Advice

https://en.wikipedia.org/wiki/Nmap

### Bonus

Show the name of the service that uses the port

### Usage

```console
$>  tinyscanner --help
Usage: tinyscanner [OPTIONS] [HOST] [PORT]
Options:
  -p               Range of ports to scan
  -u               UDP scan
  -t               TCP scan
  --help           Show this message and exit.
$>  tinyscanner -u 20.78.06.364 -p 80
Port 80 is open
$> tinyscanner -t 127.0.0.1 -p 1604
Port 1604 is closed
$> tinyscanner -t 10.53.224.5 -p 80-83
Port 80 is open
Port 81 is open
Port 82 is close
Port 83 is open
```

### Submission and audit

Files that must be inside your repository:

- Your program source code.
- a README.md file, Which clearly explains how to use the program.

Don’t hesitate to double check the names of your folders and files to ensure they are correct!
