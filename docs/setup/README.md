# Setup

## Introduction

Our programming exercises require a Unix-like OS (operating system), in particular GNU/Linux.

There are several ways to get a working Linux environment :

- Buy a computer which comes with Linux pre-installed
- Install Linux yourself
  - Natively
    - In dual-boot, which allows you to choose between Windows and Linux when your computer starts
    - In replacement of any existing OS
  - Virtualized
    - Using a third-party hypervisor (VirtualBox)
    - Using the Windows Subsystem for Linux

This document focuses on the latest method : "Windows Subsystem for Linux".

## Install Windows 10

Skip this part if Windows 10 is up-to-date.

### Download

- Visit https://www.microsoft.com/en-us/software-download/windows10ISO (if you already are on Windows, you might need to modify the user-agent)
- Select "Windows 10" edition
- Confirm
- Select "English" product language
- Confirm
- Click on "64-bit Download" button

### Install

- Burn on flash drive (using Rufus https://rufus.ie is advised)
- Boot on the flash drive
- If you see "Press any key to boot from CD or DVD..." do it immediately
- This tutorial is using Windows 10 Home
- The system will reboot several times

#### Wait for Windows 10 to be idle

- Open the Task Manager (right-click in Taskbar or use the shortcut <kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>Esc</kbd>)
- Click on "More details"
- Select "Performance" tab
- Wait for all the graphs to be completely steady (especially CPU, Wi-Fi, Ethernet)
- Reboot
- Repeat [Wait for Windows 10 to be idle](#wait-for-windows-10-to-be-idle) until the system is idle shortly after startup

#### Install Updates

- Click on :
  - Start
  - Settings
  - Update & Security
  - Check for updates (even if it says "No updates are available")
- Wait until every component Status is "Pending restart"
- Click on "Restart now"
- Repeat [Install Updates](#install-updates) until you see "You're up to date"

## Install Linux

Skip this part if you already have WSL with Debian (or similar system e.g. Ubuntu).

### Install Windows Subsystem for Linux (WSL2)

Follow the guide : https://docs.microsoft.com/en-us/windows/wsl/install-win10 \
To "Open PowerShell as Administrator", right-click on the Start menu.

### Install Debian

- Visit https://www.microsoft.com/en-us/p/debian/9msvkqc78pk6
- Click on "Get" button
- Launch
- Create user as requested (the password doesn't need to be secure)
- Leave (by using the shortcut <kbd>Ctrl</kbd>+<kbd>D</kbd>, closing the window or typing the `exit` command)

## Install VSCode

### Install remote extension

## Configure Linux

Connect to remote WSL

You can use right-click to copy/paste.

Open a terminal then type the following commands :

```
sudo apt update
sudo apt -y upgrade
sudo apt -y install curl
curl ADDR | bash
```

- Then go to Gitea (https://git.DOMAIN/user/settings/keys) in Settings => SSH / GPG Keys
- In "Manage SSH Keys" section, click on "Add Key" button
- Paste the public key in the "Content" text area (the "Key Name" will be automatically set)
- Click on "Add Key"
