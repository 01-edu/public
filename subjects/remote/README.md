## remote

### Connect

To type commands in a distant shell, you can use your peripherals (keyboard, monitor) or [SSH](<https://en.wikipedia.org/wiki/SSH_(Secure_Shell)>).

It is more comfortable to use SSH because you can use your usual terminal, with the right keymap, theme, etc.

To do this exercise you will need to add this VM :

### For VirtualBox 

- [01_remote](https://assets.01-edu.org/sys/01_remote.tar.gz)

### For UTM

- [01_remote](https://assets.01-edu.org/sys/01_remote.utm.zip)

Because the VM is behind the virtualization software router you can't access it directly. Unless you add a port forwarding rule in the VM settings that maps a host port to a guest port. Host refers to your machine and guest to the VM.

Host & guest IP addresses don't need to be specified, guest port needs to be 22 because that's SSH default listening port.

Then connect to it via SSH (this is the only way since consoles are disabled) :

```
ssh -p HOST_PORT root@localhost
```

### Configure

It is recommended to change the default SSH port (22) to prevent bots from trying to connect to it.

Since we are pretending that the guest VM is a server, change the SSH service port and make sure the port forwarding still works !

In addition, you will need to allow the new port in the firewall `ufw`.
