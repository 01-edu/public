## scan

![scan1](scan1.png)

> You're going to do like Trinity in the movie _The Matrix Reloaded_. Except that you're not going to turn off the electricity in a city, and you probably don't wear as much leather, but that's okay.

---

By design and by default, most network entities have a high discoverability.

Whatever your motivations for exploring a network (hacking, curiosity...), you generally proceed in this order :

1. Scan a network to find hosts
2. Scan a host to get information (hardware, OS & exposed services)

`1.` can be done by analyzing the link layer of Internet protocols, in particular `ARP` in order to gather MAC addresses (the identifier of a network interface, such as Wi-Fi or Ethernet cards).

`2.` is usually performed by analyzing the network and transport layers, in particular the `IP` & `TCP` protocols with the `nmap` tool.

For the context, when you open a website, here are the protocols involved, from the highest to lowest level (the protocols at the top depend on those at the bottom):

- `HTTP`: application layer (website data) - messages
- `TLS`: application layer ("lock" sign) - secure connection
- `TCP`: transport layer (port number) - reliable connection
- `IP`: network layer (IP address) - global communication
- `ARP`: link layer (MAC address) - local communication

For this challenge, you will focus on `IP`, `ARP` & `TCP`, and therefore on IP/MAC addresses and ports.

You will need to add these 2 VMs:

### For VirtualBox

- [01_scan_RRF-CONTROL](https://assets.01-edu.org/sys/01_scan_RRF-CONTROL.tar.gz)
- [01_scan_laptop](https://assets.01-edu.org/sys/01_scan_laptop.tar.gz)

### For UTM

- [01_scan_RRF-CONTROL](https://assets.01-edu.org/sys/01_scan_RRF-CONTROL.utm.zip)
- [01_scan_laptop](https://assets.01-edu.org/sys/01_scan_laptop.utm.zip)

To perform your tasks you only need to use the two VMs.

You will only have control over "01_scan_laptop". A port forwarding is set on 10122 so you can connect through SSH, the password is a single space.

Your mission, should you choose to accept it, is to scan the network interface (`enp0s8` for VirtualBox,  `enp0s1` for UTM) and find a way in the server "01_scan_RRF-CONTROL", you will know you have succeeded when you see :

```
RRF-CONTROL> █
```

Because the virtualized Network is very slow (10 Mbps, Ethernet is usually 1000 Mbps), expect long scan times :

- ARP scanning takes up to 5 minutes
- port scanning takes up to half an one hour with the option `-T4` (even more without).

May the Fourth be with you.

---

---

![scan2](scan2.png)

> Another depiction of the `nmap` tool: in _Ocean's 8_, Rihanna uses [Kali Linux](https://www.kali.org) to steal a valuable diamond.
>
> Same strategy, different styles
