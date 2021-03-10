## add-vm

### Add a virtual machine

We provide virtual machines you need to download & add to VirtualBox.

Each VM will contain a system in a certain state that you will need to use in order to crack the sysadmin exercises. Their names start with "01\_" in order to avoid conflict with your VM names.

You can download these VM here :

https://assets.01-edu.org/sys

For this exercise you need to download the archive : [01_add-vm.tar.gz](https://assets.01-edu.org/sys/01_add-vm.tar.gz).

Extract it in the folder `VirtualBox VMs` which is located in your [home directory](https://en.wikipedia.org/wiki/Home_directory).

In VirtualBox :

- Select on the menu bar : <u>M</u>achine → <u>A</u>dd
- Open the file `01_add-vm.vbox`

The VM should appear on the list under the name "01_add-vm".

### Snapshots

A snapshot is a state of a system at a given time, just like a version or a copy.

Let's imagine that you want to modify a text or a drawing, that you make a copy before, keeping the original... This is a snapshot.

Software that makes extensive use of this feature :

- Git (a commit is a snapshot)
- Windows System Restore (a restore point is a snapshot)
- macOS Time Machine (each backup is a snapshot)
- Linux LVM (Logical Volume Manager)
- Filesystems : Btrfs, ZFS
- Hypervisors like the one you are currently using : VirtualBox

**We recommend that this be the first thing you do after importing a VM.**
This way you can break it, go in the wrong direction or want to try another way and at any time you can go back to the original state.

If you forgot to do so and want to make a fresh start, you can always delete the VM (all files) and add it again in VirtualBox.

### Check

Since all VMs are based on this one, check that it is working properly by starting and stopping it gracefully (`poweroff` or ACPI shutdown).

The password is a single space.
