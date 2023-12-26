#### General

###### Is the student able to explain clearly what port means?

###### Is the student able to explain clearly what ports scanning means?

###### Is the student able to explain clearly why the ports scanning is important in the pentesting?

###### Is the student able to explain clearly how his program works?

##### Check the Repo content

Files that must be inside your repository:

- Your program source code.

- a README.md file, Which clearly explains how to use the program.

###### Are the required files present?

##### Evaluate the student's submission

The student must launch his program by passing the IP address of a website as the first argument and as the second argument "80" which corresponds to port 80, open to all IP addresses using the HTTP protocol. Port 80 is therefore open on all HTTP sites.

21 FTP  
22 SSH  
23 Telnet  
25 SMTP  
53 DNS  
80 HTTP  
110 POP3  
115 SFTP  
135 PRC  
139 NetBIOS  
143 IMAP  
194 CRI  
443 SSL  
445 SMB  
1433 MSSQL  
3306 mysql  
3389 Remote Desktop  
5632 PCAnywhere  
5900 VNC  
25565 Minecraft

##### Run `tinyscanner -t 127.0.0.1 -p 80`

###### Does port 80 show as open?

##### Run a local server using udp protocol with the port 8080 and run `tinyscanner -u 127.0.0.1 -p 80`

###### Does port 80 show as open?

##### Run `tinyscanner --help`

```console
$>  tinyscanner --help
Usage: tinyscanner [OPTIONS] [HOST] [PORT]
Options:
  -p               Range of ports to scan
  -u               UDP scan
  -t               TCP scan
  --help           Show this message and exit.
```

###### Is the program display an output similar to that?

#### Bonus

###### +Is the service name displayed?
