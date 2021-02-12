## Linux

### Introduction

Linux is the most widely used kernel (the core of an operating system).

Linux can be found everywhere there is a computer : data centers, phones, gaming consoles, cars, planes, submarines, space stations...

Operating systems using the Linux kernel are called "Linux distributions".

![Tux mascot](tux.png)

> "But this wasn't to be just any penguin. Above all, Linus wanted one that looked happy, as if it had just polished off a pitcher of beer and then had the best sex of its life."
>
> _Just for Fun: The Story of an Accidental Revolutionary_ (Linus Torvalds, David Diamond)

---

This series of projects will introduce you to a Linux distribution called Debian and its use in servers (always-on computers providing the required services for Internet and applications).

You'll become a better programmer by knowing what's underneath your programs and what's behind Cloud services.

It will also allow you to learn some tricks to better control your machine and servers.

### Virtualization

To practice in a standardized way and without the risk of altering your own operating system, install the hypervisor [VirtualBox](https://www.virtualbox.org/wiki/Downloads).

This will allow you to run virtual machines (virtual computers accessible through a window).

This way, whatever your operating system, you can, for example, run Linux, Windows & macOS in 3 separate windows at the same time.

### Installation

Operating systems are usually installed with a "disk image" that can be written to a USB stick or CD/DVD for installation on a computer, or kept as is for installation in a virtual machine (which is the case here).

Install Debian (ISO installer "netinst" is recommended) in a VM (Virtual Machine). The steps are :

- Create the VM.
- Start the VM and attach the ISO file as an optical drive (CD/DVD).
- Follow the steps of the installer, it will reboot with a Debian system ready to use.

### Shut down

Gracefully close the VM using "ACPI Shutdown". It's like pressing the power button of a real computer.
