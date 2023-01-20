#### General

##### Check the Repo content

Files that must be inside the repository:

- DeepInSystem.sha1

###### Are the required files present?

##### Verify that the virtual machine that will be audited is the one that is submitted:

```console
user:~$ sha1sum deep-in-system.ova > deep-in-system-toaudit.sha1
user:~$ diff deep-in-system.sha1  deep-in-system-toaudit.sha1 ; echo $?
0
user:~$
```

###### Is the SHA1 of the provided machine the same as the machine being audited?

##### Check the Virtual machine aliases

###### The virtual machine is clean of any alias that may affect the results of the audit commands?

#### The Virtual Machine Part:

##### Check the Linux distribution

To get information about the OS release:

```console
user:~$ cat /etc/os-release
PRETTY_NAME="Ubuntu <...> LTS"
NAME="Ubuntu"
VERSION_ID="<...>"
VERSION="<...> LTS <...>"
VERSION_CODENAME=<...>
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=<...>
user:~$
```

##### Check if ubuntu is a server and not a desktop:

```console
user:~$ dpkg -l ubuntu-desktop
dpkg-query: no packages found matching ubuntu-desktop
user:~$
```

You can check the versions of the ubuntu server from here: [Get Ubuntu Server](https://ubuntu.com/download/server)

###### Is the installed Linux distribution is Ubuntu server's latest LTS?

##### Check the VM disk and partitions

Check the VM disk and partitions with this command:

```console
user:~$ lsblk -o NAME,FSTYPE,SIZE,MOUNTPOINT /dev/sda
NAME   FSTYPE SIZE MOUNTPOINT
sda            30G
├─sda<...>          1M
├─sda<...> swap     4G [SWAP]
├─sda<...> ext4    15G /
├─sda<...> ext4     5G /home
└─sda<...> ext4     6G /backup
user:~$
```

- The VM disk size must be 30GB.

- VM disk must be divided into these partitions:
"swap:" 4G
"/": 15G
"/home": 5G
"/backup": 6G

###### Is the VM Disk size correct?

> There is no problem if the size of the divisions is not very accurate (Authorized error rate: <= 0.5G)!

###### Are the VM disk partitions correct?

##### Check the hostname and user name

To check the hostname:

```console
user:~$ hostname
<username>-host
user:~$
```

To check the user name and groups:

```console
user:~$ id
uid=<...>({username}) gid=<...>({username}) groups=<...>({username}),<...>(sudo),<...>
user:~$
```

###### Does the hostname in the format of "{username}-host"?

###### Does the student use a user different from the "root" user?

###### Does the username contain the student login?

###### Does the user in the sudo group?

###### Does the student can explain what is sudo group in Linux?

#### The Network & Security Part:

##### Check the VM IP address

The student must show the file that was modified to set a static IP address.

###### Does the student can explain the configuration?

###### Does the student What is a netmask?

##### Check if the IP address is static with this command:

```console
user:~$ ip a | grep dynamic
user:~$
```

###### There is no internet interface with dynamic IP assignment?

##### Check if the internet works fine with the static IP address:

```console
user:~$ ping -c 5 google.com
```

###### Can connect to the internet properly?

###### Can The student explain why a static IP address is important for a web server?

##### Check the sshd configuration

The student must show the file that was modified to secure the ssh server.

###### Does the student can explain the configuration?

###### Is the root access disabled in the sshd config (PermitRootLogin: no)?

###### Is the port of the sshd "2222"?

##### Try to connect from outside the VM

```console
outsideTheVM:~$ ssh {username}@{machine-ip} -p 2222
{username}@{machine-ip}'s password:
Welcome to Ubuntu <......>
InsideTheVM:~$ hostname
{username}-host
InsideTheVM:~$
```

###### Can connect to the ssh properly?

###### Does the student can explain what is ssh server and what the role of it?

##### Check the firewall

If the student uses ufw you can check it with this command:

```console
user:~$ sudo ufw status
Status: active

To                         Action      From
--                         ------      ----
2222/tcp                   ALLOW       Anywhere
<...>
Apache                     ALLOW       Anywhere
<...>
user:~$
```

Otherwise, the student must show what firewall is used.

###### Is the firewall activated?

##### Ask the student to justify why each open port is open

###### Are all open ports justified?

###### Is the MySQL port not open in the firewall?

###### Does the student can explain what is firewall and what the role of it in a server?

#### User Management Part:

##### Check luffy user

The student should connect to the machine with the "luffy" user by using this private key:

- Private key:

```
-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----
```

###### Is the student able to connect to the machine with the "luffy" user by using the private key and without using any password?

##### Check the groups of luffy user:

```console
luffy:~$ groups luffy
luffy : luffy sudo
luffy:~$
```

##### Check the home directory of luffy user:

```console
luffy:~$ echo ~
/home/luffy
luffy:~$ echo $HOME
/home/luffy
luffy:~$
```

###### Is the "luffy" user assigned to the sudo group?

###### Is the home directory of "lufy" user: /home/luffy?

##### Check zoro user

The student should connect to the machine with the "zoro" user by using this password: `^wb@92Sq&ls644@5*Je0`

###### Is the student able to connect to the machine with the "zoro" user by the mentioned password?

##### Try to execute a command with sudo:

```console
zoro:$ sudo cat /etc/shadow
zoro is not in the sudoers file.  This incident will be reported.
zoro:~$
```

##### Check the groups of zoro user:

```console
zoro:~$ groups zoro
zoro : zoro
zoro:~$
```

##### Check the home directory of zoro user:

```console
zoro:~$ echo ~
/home/zoro
zoro:~$ echo $HOME
/home/zoro
zoro:~$
```

###### Is the "zoro" user can't perform a command with sudo?

###### Is the "zoro" user not assigned to the sudo group?

###### Is the home directory of "zoro" user: /home/zoro?

##### Ask the student to:

In less than 10 minutes the student must create a user called "kratos" this user must be a sudoer and must be able to connect with a private key.
The private ssh key must be created by the student during this exam.
After the student finishes creating and setting up the user, the student must show that the user can be connected with the private key and can perform a sudo command.

> If the student can't solve this exam, he must directly fail in this project.
> If did not pass this exam and was able to succeed in this project, a temporal crater will open and the world will be destroyed!

###### Does the student can create a private key?

###### Does the student can create the user?

###### Does the student assign the public key to the user?

###### Does the student add the user to the sudo group?

###### Is user "kratos" can connect with the private key?

###### Is user "kratos" can perform a sudo command?

#### Services Part:

##### Check nami user:

##### By using SSH create a file inside /backup:

```console
$ sudo touch /backup/audit-check
```

##### Try to connect to the "nami" user via FTP:

```console
user:~$ ftp {vm-ip}
Connected to {vm-ip}.
<...>
Name ({vm-ip}:{username}): nami
331 Please specify the password.
Password:
230 Login successful.
Remote system type is UNIX.
Using binary mode to transfer files.
ftp> ls
<...>
<...> audit-check
<...>
226 Directory send OK.
ftp> get audit-check
<...>
226 Transfer complete.
ftp>
```

- "nami" user password: `mYdb6HA^5W4o`

###### Can connect with user "nami" and mentioned password to the FTP Server properly?

###### Is the created file exist in the FTP Server?

###### Can get the audit-check file from the FTP Server?

##### Check anonymous user:

##### Try to connect with an anonymous user and a blank password:

```console
user:~$ ftp {vm-ip}
Connected to {vm-ip}.
<...>
Name ({vm-ip}:{username}): anonymous
331 Please specify the password.
Password:
530 Login incorrect.
ftp: Login failed
ftp>
```

###### Can't connect to FTP Server with an anonymous user and blank password?

###### Does the student can explain what is FTP Server and what the role of it?

#### WordPress Part:

##### From your browser, enter "http://{vm-ip}/"

> it can be https instead of http if the student installs an SSL certificate!

##### Ask the student to log in with the admin user.

WordPress must be installed.

##### Try to post something, any way you are free to do anything.

###### Is WordPress installed and working properly?

##### Try to access to "http://{vm-ip}/wp-config.php"

###### The WordPress config file content is not displayed?

#### Backup Part:

##### Check the cronjob:

The student must show created cronjob.

###### Is they are a cron job that starts every Day At 00:00 (0 0 \* \* \*)?

###### Is the cronjob command creating a tar file of the WordPress database in /backup?

##### Check the FTP system functionality:

> Before starting this test you have to remove all WordPress backup files in "/backup" and delete the logs file "/var/log/backup.log".

##### In the crontab, you have to change the scheduling to :

`* * * * *`

##### After 1 minute, check the FTP Server with the "nami" user:

```console
user:~$ ftp {vm-ip}
Connected to {vm-ip}.
<...>
Name ({vm-ip}:{username}): nami
331 Please specify the password.
Password:
230 Login successful.
Remote system type is UNIX.
Using binary mode to transfer files.
ftp> ls
<...>
<...> {wordpress-backupfile}
<...>
226 Directory send OK.
ftp> get audit-check
<...>
226 Transfer complete.
ftp>
```

###### Does a WordPress database backup file with the date of today exists in the FTP Server?

##### Check the backup logs file:

```console
user:~$ cat /var/log/backup.log
<...>wordpress backup created!, date: <...>
user:~$
```

###### Is the backup logs file existing and contains a message informing you that the backup was successful and the timing of the backup?

###### Does the student can explain what is cronjob and what the role of it?

###### Does the student can explain why backup is important?

#### Bonus

###### + Did the student pass the account creation exam without error and in a short time?

###### + Did the student add any optional bonus?

###### + Is the student a genius of the system administration?
