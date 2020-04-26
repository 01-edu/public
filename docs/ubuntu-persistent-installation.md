# Ubuntu

## OS Installation

Install the latest version of [Virtual Box](https://www.virtualbox.org/wiki/Downloads).

> The text instructions are more important than the screenshots

Screenshots of the installation of Ubuntu in Virtual Box are [here](https://github.com/xpetit/vbox-ubuntu-install/blob/master/README.md).

Download and boot the [last Ubuntu release](http://releases.ubuntu.com/19.10/ubuntu-19.10-desktop-amd64.iso).

- Create a new virtual machine named "Ubuntu" with at least 4096 MB of RAM
- Use the fixed size storage allocation (to have more performance)
- In the settings of the VM
  - System -> Motherboard : check "Enable EFI"
  - System -> Processor : Select at least 2 processors
  - Display -> Screen : Put "Video Memory" to the maximum
    - Enable 3D acceleration
  - Storage
    - Remove IDE controller
    - Add Optical Drive to the SATA controller
      - Choose your Ubuntu ISO image
  - Close the settings (click OK)
- Run the VM

Follow the screenshots (some settings can be personalized, such as keyboard layout, location, password, login automatically, but **do not change the username**)

Skip the welcoming window.

Don't install updates if Ubuntu asks to. The scripts will.

## OS customization

You can overwrite the files of the folder `system` by setting an environment variable named `OVERWRITE` with the format : `Destination folder;Git URL`.

For example to write the content of the repository [github.com/xpetit/custom](https://github.com/xpetit/custom) in the system folder :

```shell
export OVERWRITE='.;https://github.com/xpetit/custom.git'
```

## OS configuration

Run a terminal and type these commands :

```shell
unset HISTFILE
sudo apt-get -y install git
export PERSISTENT=
git clone https://github.com/01-edu/public.git
public/scripts/setup.sh
```

After reboot you should install Virtual Box additions (and reboot again) :

```shell
sudo apt -y install virtualbox-guest-x11
```

Then it is advised to use the virtual machine in full screen mode (Host key - F)
