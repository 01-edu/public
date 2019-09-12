# Ubuntu

## OS Installation

Download and boot the [last Ubuntu release](http://releases.ubuntu.com/19.04/ubuntu-19.04-desktop-amd64.iso).

Follow the steps :

![img1](img/ubuntu-installation/1.png)
![img2](img/ubuntu-installation/2.png)
![img3](img/ubuntu-installation/3.png)

The partitioning is :

1. 256 MB : EFI partition
2. 20 GB : system partition

![img4](img/ubuntu-installation/4.png)
![img5](img/ubuntu-installation/5.png)
![img6](img/ubuntu-installation/6.png)

Remove the installation disk and then reboot.

Skip the welcoming window.

Don't install updates if Ubuntu asks to. The scripts will.

## OS configuration

Run a terminal and type :

```console
student@ubuntu:~$ unset HISTFILE
student@ubuntu:~$ sudo apt-get -y install curl
student@ubuntu:~$ bash <(curl -sSL raw.githubusercontent.com/01-edu/public/master/scripts/kickstart.sh)
[...]
Ask for student user password (will be removed later)
[...]
Long installation/configuration process then reboots
```

The system is now read-only, every data is written to a temporary partition.

The session is password-less.

To gain a superuser terminal, use SSH :

```console
user@remote:~$ ssh -p521 root@IP_ADDRESS
```

To gain access with read/write access to the filesystem, use this command :

```console
root@ubuntu:~# overlayroot-chroot
INFO: Chrooting into [/media/root-ro]
root@ubuntu:/#
```
