## deep-in-system

![sysadmin](pictures/sysadmin.jpeg)

In this project you will learn how to administer a Linux server, You will set up security and network for a ubuntu server, and will install some popular services.

### Objectives

Implement some learned skills from the scripting pool in a real-life project.

Having a first experience in a ubuntu server setup.

Discovering some network and security implementations in Linux.

Discovering some popular services in Linux.

### Advice

Read the entire project before starting implementation!

Try to understand all the commands that you will use in your setup.

Save any command you entered, it will be useful if you want to reinstall the server or do some debugging.

Create a backup file for each config file you will modify, this will be useful if you broke the config.

> In this project we have put some passwords and private keys exposed, It is not recommended to do this in any way!
> And don't use these passwords and private keys outside this learning project!

### Instructions

#### The Virtual Machine Part:

Install a ubuntu server's latest LTS as a virtual machine.

- The VM disk size must be 30GB.

- You must divide your VM disk into these partitions:
`swap:` 4G
`/`: 15G
`/home`: 5G
`/backup`: 6G

- Your username must be your login name.

- You have to set your hostname with the format of `{username}-host`, if your login is `potato`, then your hostname must be `potato-host`.

#### The Network Part:

Set a static private IP address, you are free to choose which netmask to use.

You must be able to connect to the Internet!, you can test with:

```console
$> ping -c 5 google.com
```

> You should not have any internet interface with dynamic ip assignment.

#### The Security Part:

> You do not have to use the root user in your setup process!
> You won't need it when you have `sudo`.
> Sudo provides fine-grained access control. It grants elevated permissions to only a particular program that requires it. You know which program is running with elevated privileges, rather than working with a root shell (running every command with root privileges).

- You have to disable remote root login via ssh.

- Change the ssh port to: `2222`.

- Configure the Firewall, and close all incoming ports, only used ports must be opened.
  > All open ports must be justified in the audit!

#### User Management Part:

You have to create 2 users in your server as follows:

##### 1- _luffy_:

- SSH authentication Method: Public key-based authentication
- Home directory: `/home/luffy`
- Sudoer: yes
- Public key:

```
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC9NYZT5ueK+2JupNLOAg3xTSd117NwrPgVN15S2gXfijJlpmO1dBgR+ro6N2yngoLTLi2QGUU53xRj0p/SJf+9GOdlkt55ePxs2GNKxACJcrZLPiyOBbb5VRyyRn3ie84qdw7EUSnWROTBWVIqkcxE+YFP9e06gQCuUxm7FyjwUfEMSXEaWCLMC2qREz8H92ZtEcXqQNKotG0CtIuuFsVX1CQdEh86v+SVN6pVbVaXLWFKkpZSubAvGe5g4ffiLjTSMfzmZ+Ayley0DmX+7nsV0OXgIpixMmW1KV8NNo5oxTQFPG3z5v7AgCUM8Hc1R2dj2AjbmDRlh9amTjQd1dPR99TJ84Nu1fIwsar5eG5u/oIA3cUTT028gcAL85GLy7YERUyXpbbaap1QgsJGViCYETflUcvwfdxrDetLBbnQ2aKqo/KxyXFDXt7uR618p2hrotWE9nWZnIQ90FRFUhEIcoq1Gg1on/0G+4M9WIqlChh6qUAq/Gi3IHURXlTBP9M= luffy@luffy
```

- Private key:

```
-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----
```

##### 2- _zoro_:

- SSH authentication Method: Password authentication
- Home directory: `/home/zoro`
- password: `^wb@92Sq&ls644@5*Je0`
- sudoer: no

#### Services Part:

- Install an FTP server and create a user `nami`.
  The user `nami` can access via FTP only to `/backup` with read-only access.
  `nami` user password: `mYdb6HA^5W4o`

> Don't enable anonymous access!
> This will be risky!

#### The Database Part:

- You have to install MySQL Server

- Disable the remote connection to the root user.

- Do not allow connection to MySQL from outside the server!
  To improve the security of your website, you should keep your MySQL server accessible only by applications in the server, As long as it does not affect your solution.

- You must create a MySQL user, which has the only required access to the WordPress database.
  > Don't use the root user in your WordPress website!

#### WordPress Part:

- You have to install WordPress

- WordPress must be in your web server root directory: `http://{host}/`

- Your WordPress must work in a normal way, try to post something or create another user, any way you are free to do anything.

> The configuration file must not be public accessible!, try `http://{host}/wp-config.php`

#### Backup Part:

Backups protect against human errors, hardware failure, virus attacks, power failures, and natural disasters. Backups can help save time and money if these failures occur.

In this exercise, you will set up a simple backup method by using cron jobs.

- Set up a cron job that starts every Day At 00:00, it must create a tar file of the WordPress database.

- The backup files must be created in the `/backup` folder.

- The backup file name must contain the creation date.

- You must add a line to a log file (`/var/log/backup.log`) This line should contain a message informing you that the backup was successful and the timing of the backup.

- Your backup files must be downloadable from the `nami` FTP user.

### Bonus Part

If you complete the mandatory part perfectly, you can move to this part.
You can add anything you feel deserves to be a bonus, some of the suggested ideas:

- Install the Minecraft server, It must be always running even if you reboot your server.

- Automate all instructions with ansible, so you can remake this long process in more than 1 server in a short time.

- Set up the SSL in the web server and FTP server, you can use self-signed SSL.

_Challenge yourself!_

### Submission and audit

You must export your VM to a safe place, you will need it in the audit.
You will use your exported VM to run a new VM for each audit.
Push the shasum of your exported VM, you can get it this way:

```console
user:~$ sha1sum deep-in-system.ova > deep-in-system.sha1
user:~$ cat deep-in-system.sha1 | cat -e
<...>255bfef9560<...>  deep-in-system.ova$
user:~$
```

Files that must be inside your repository:

- deep-in-system.sha1

> It's Forbidden to use external scripts!
> Any use of external scripts or use of commands without understanding their jobs is considered cheating!
