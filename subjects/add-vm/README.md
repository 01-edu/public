## add-vm

### Add a virtual machine

We provide virtual machines for both, X86-64 chips and Apple Silicon chips bases systems, you need to download & add to your virtualization software.

For X86-64 chips, you have to use [VirtualBox](https://www.virtualbox.org/wiki/Downloads).

For Apple Silicon chips, you have to use [UTM](https://mac.getutm.app/).

Each VM will contain a system in a certain state that you will need to use in order to crack the sysadmin exercises. Their names start with "01_" in order to avoid conflict with your VM names.

You can download these VM here :

https://assets.01-edu.org/sys

### VirtualBox instructions:

For this exercise you need to download the archive : [01_add-vm.tar.gz](https://assets.01-edu.org/sys/01_add-vm.tar.gz).

Extract it in the folder `VirtualBox VMs` which is located in your [home directory](https://en.wikipedia.org/wiki/Home_directory).

In VirtualBox :

- Select on the menu bar : <u>M</u>achine → <u>A</u>dd
- Open the file `01_add-vm.vbox`

The VM should appear on the list under the name "01_add-vm".

### UTM instructions:

For this exercise you need to download the archive : [01_add-vm.utm.zip](https://assets.01-edu.org/sys/01_add_vm.utm.zip).

Create UTM_VMs folder in your Desktop and extract **01_add-vm.utm** to it

In UTM:

- Select on the menu bar : Files→ open
- Navigate to`01_add-vm.utm` and select it

The VM should appear on the list under the name "01_add-vm".

(Note: the **.utm** are not files, they are directories.)

### Snapshots

A snapshot is a state of a system at a given time, just like a version or a copy.

Let's imagine that you want to modify a text or a drawing, that you make a copy before, keeping the original... This is a snapshot.

Software that makes extensive use of this feature :

- Git (a commit is a snapshot)
- Windows System Restore (a restore point is a snapshot)
- macOS Time Machine (each backup is a snapshot)
- Linux LVM (Logical Volume Manager)
- Filesystems : Btrfs, ZFS
- Hypervisors like the one you are currently using : VirtualBox or UTM

**We recommend that this be the first thing you do after importing a VM.**
This way you can break it, go in the wrong direction or want to try another way and at any time you can go back to the original state.

If you forgot to do so and want to make a fresh start, you can always delete the VM (all files) and add it again in VirtualBox.

### Check

Since all VMs are based on this one, check that it is working properly by starting and stopping it gracefully.

The username is `root` and the password is a single space.
