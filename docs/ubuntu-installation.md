# Ubuntu

## OS Installation

Download and boot the [last Ubuntu release](http://releases.ubuntu.com/19.04/ubuntu-19.04-desktop-amd64.iso).

Follow the steps :

![img1](https://user-images.githubusercontent.com/32063953/56804679-85867580-681e-11e9-8965-e87c6a89fac0.png)
![img2](https://user-images.githubusercontent.com/32063953/56963599-3eb3bb00-6b51-11e9-9778-4f3bb9993c74.png)
![img3](https://user-images.githubusercontent.com/32063953/56963600-3eb3bb00-6b51-11e9-94cc-279406f37def.png)

The partitioning is :

- 256 MB : EFI partition
- 20 GB : system partition
- 32 GB : unused partition (will be used later)
- rest : unused partition (will be used later)

![img4](https://user-images.githubusercontent.com/32063953/56963602-3eb3bb00-6b51-11e9-8977-38e4e67d6ce1.png)
![img5](https://user-images.githubusercontent.com/32063953/56963603-3f4c5180-6b51-11e9-9349-46ab90287691.png)
![img6](https://user-images.githubusercontent.com/32063953/56963604-3f4c5180-6b51-11e9-8df2-5016771e6e07.png
)

Remove the installation disk and then reboot.

Skip the welcoming window.

Don't install updates if Ubuntu asks to. The scripts will.

![img8](https://user-images.githubusercontent.com/32063953/56804701-8d461a00-681e-11e9-8825-dfc69f8268bf.png)
![img9](https://user-images.githubusercontent.com/32063953/56804703-8d461a00-681e-11e9-840c-498ccab7d911.png)
![img10](https://user-images.githubusercontent.com/32063953/56804704-8ddeb080-681e-11e9-96ff-6c8783c5aacc.png)
![img11](https://user-images.githubusercontent.com/32063953/56804706-8ddeb080-681e-11e9-85e1-20c5b6956a36.png)

## OS configuration

```shell
student@tmp-hostname:~$ wget github.com/01-edu/public/archive/master.zip
student@tmp-hostname:~$ unzip master.zip
student@tmp-hostname:~$ cd public-master/scripts
student@tmp-hostname:~$ sudo ./install_client.sh
[...]
Ask for student user password (will be removed later)
[...]
Ask to set the root password
[...]
Long installation/configuration process
[...]
student@tmp-hostname:~$ cat dconfig.txt | dconf load /
student@tmp-hostname:~$ reboot
```

The system is now read-only, every data is written to a temporary partition.

The session is password-less.

To gain a superuser terminal with read/write access to the filesystem, type these commands:

```shell
student@tmp-hostname:~$ su -
Password:
root@tmp-hostname:~# overlayroot-chroot
```
