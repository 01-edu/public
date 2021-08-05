# Setup

Table of contents :

- [Introduction](#introduction)
- [Install Windows 10](#install-windows-10)
  - [Prepare](#prepare)
  - [Install](#install)
    - [Wait for Windows 10 to be idle](#wait-for-windows-10-to-be-idle)
    - [Install updates](#install-updates)
- [Install a web browser](#install-a-web-browser)
- [Install Linux](#install-linux)
  - [Install Windows Subsystem for Linux (WSL2)](<#install-windows-subsystem-for-linux-(wsl2)>)
  - [Install Debian](#install-debian)
- [Install VSCode](#install-vscode)
  - [Install remote extension](#install-remote-extension)
- [Configure Linux](#configure-linux)
  - [Connect to remote WSL](#connect-to-remote-wsl)
  - [Configure tools](#configure-tools)
    - [Configure Gitea](#configure-gitea)
    - [Configure Go extension](#configure-go-extension)

## Introduction

Our programming exercises require a Unix-like OS (operating system), in particular GNU/Linux.

There are several ways to get a working Linux environment :

- Buy a computer which comes with Linux pre-installed
- Install Linux yourself :
  - Natively
    - In dual-boot, which allows you to choose between Windows and Linux when your computer starts
    - In replacement of any existing OS
  - Virtualized
    - Using a third-party hypervisor (VirtualBox)
    - Using the Windows Subsystem for Linux

This document focuses on the latest method : "Windows Subsystem for Linux".
If you want to do it differently, ensure you still do the [Configure tools](#configure-tools) part.

## Install Windows 10

[Skip this part](#install-a-web-browser) if Windows 10 is up-to-date.

### Prepare

You can check the [video](https://www.youtube.com/watch?v=9mnsKrGsOOo).

- Visit [Download Windows 10](https://www.microsoft.com/en-us/software-download/windows10)
- Click on "Download tool now"
- Run the downloaded tool
- Accept the license
- Select "Create installation media (USB flash drive..."
- Select "English (United States)" product language
- Select "Windows 10" edition
- Select "64-bit (x64)" architecture
- Click on "Next"
- Select "USB flash drive"
- Click on "Next"
- If you see the message "We can't find a USB flash drive" make sure you connected a USB flash drive that is formatted
  - Click on "Refresh drive list" to see it
- Click on "next"
- Click on "Finish"

### Install

You can check the [video](https://www.youtube.com/watch?v=4D6SwICBkFA).

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

#### Install updates

- Click on :
  - Start
  - Settings
  - Update & Security
  - Check for updates (even if it says "No updates are available")
- Wait until every component Status is "Pending restart"
- Click on "Restart now"
- Repeat [Install updates](#install-updates) until you see "You're up to date"

## Install a web browser

You can check the [video](https://www.youtube.com/watch?v=8QkXFo5kKg4).

[Skip this part](#install-linux) if you have a modern & full-featured web browser (e.g. Chrome, Firefox, etc).

We recommend to use Firefox with the extension "uBlock Origin" and the additional filters lists :

- Annoyances
  - EasyList Cookie
  - uBlock filters - Annoyances

Make sure to apply changes.

## Install Linux

You can check the [video](https://www.youtube.com/watch?v=u0vM78D2MLo).

[Skip this part](#install-vscode) if you already have Debian (or similar system e.g. Ubuntu).

### Install Windows Subsystem for Linux (WSL2)

Follow the guide : [Windows Subsystem for Linux Installation Guide for Windows 10](https://docs.microsoft.com/en-us/windows/wsl/install-win10)

To "Open PowerShell as Administrator", right-click on the Start menu.

### Install Debian

- Visit [Debian](https://www.microsoft.com/en-us/p/debian/9msvkqc78pk6)
- Click on "Get" button
- Launch
- Create user as requested (the password doesn't need to be secure and you have to remember it)
- Close (by using the shortcut <kbd>Ctrl</kbd>+<kbd>D</kbd>, closing the window or typing the `exit` command)

## Install VSCode

You can check the [video](https://www.youtube.com/watch?v=TWw2Xc7A3YI).

If you run Linux natively, go to [Configure tools](#configure-tools).

Download it from the [official website](https://code.visualstudio.com).
Run the installer, check all "Additional Tasks".

### Install remote extension

Launch VSCode, once started it should detect WSL and propose you to install the "Remote - WSL" extension, which is needed for the next steps.

## Configure Linux

You can check the [video](https://www.youtube.com/watch?v=c8HoeMDoOXk).

### Connect to remote WSL

Click on the bottom-left green button (overlay says "Open a remote window"), select "Remote-WSL: New Window".

It will open a new VSCode window then install and run a server program in Linux to help VSCode running in its context.

The bottom-left green button (Remote Host) in the VSCode status bar should now indicate "WSL: Debian".

You can close the previous VSCode window.

### Configure tools

In the Menu Bar of VSCode, click on "Terminal", then "New Terminal".

A new panel should appear looking like :

```
user@DESKTOP-XXXXXXX:~$ █
```

This is a command-line interface, you can type commands and execute them pressing <kbd>Enter</kbd>. You can also complete your command using <kbd>Tab ↹</kbd>. For instance type "`ec`" :

```
user@DESKTOP-XXXXXXX:~$ ec█
```

Then press <kbd>Tab ↹</kbd>. The command will be completed to "`echo`" and a space is added, waiting for arguments to be provided, type "`Hello World`" then <kbd>Enter</kbd> :

```
user@DESKTOP-XXXXXXX:~$ echo Hello World
Hello World
user@DESKTOP-XXXXXXX:~$ █
```

`echo` is a program that displays text. Now you will execute commands that will install all the necessary programs you will need, this is the first one :

```
sudo apt update
```

Type the password you entered during [Linux installation](#install-debian).
When it's over (the command prompt appears again) run the 3 next commands (we hide their output):

```
sudo apt -y upgrade
sudo apt -y install curl
curl https://raw.githubusercontent.com/01-edu/public/master/docs/setup/configure.sh | bash
```

Now you are able to push & pull code to Gitea using `git`.

Close the terminal (by using the shortcut <kbd>Ctrl</kbd>+<kbd>D</kbd> or typing the `exit` command)

#### Configure Go extension

- In the Menu Bar of VSCode, click on "View", then "Extensions"
- In the search field type "go"
- Click on "Install" for the first extension in the list : "Go"
- Click on "Reload Required"
- Open the "Command Palette" (<kbd>F1</kbd> or <kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>P</kbd>)
- Type "go install"
- Click on "Go: Install/Update Tools"
- Check the first box in order to check all of them
- Click on "OK"
