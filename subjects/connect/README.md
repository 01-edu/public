## connect

To communicate over a network, a computer must have an IP address.

The computer can choose its own IP address (static) or can ask a DHCP server to assign one (dynamic).

Generally, clients (smartphones, laptops, etc...) rely on DHCP servers to have a dynamic IP address and servers have a static IP address.

For this project you will need to add these 3 VM :

### For VirtualBox 

- [01_connect_box](https://assets.01-edu.org/sys/01_connect_box.tar.gz)
- [01_connect_machine1](https://assets.01-edu.org/sys/01_connect_machine1.tar.gz)
- [01_connect_machine2](https://assets.01-edu.org/sys/01_connect_machine2.tar.gz)

### For UTM

- [01_connect_box](https://assets.01-edu.org/sys/01_connect_box.utm.zip)
- [01_connect_machine1](https://assets.01-edu.org/sys/01_connect_machine1.utm.zip)
- [01_connect_machine2](https://assets.01-edu.org/sys/01_connect_machine2.utm.zip)

The VM are configured like this :

```
        N E T W O R K S                C O M P U T E R S
_______________________________     ________________________

.-----------------------------.
|          Internet           |
'-----------------------------'
               ^
               |
               v
.-----------------------------.
|              NAT            |
|                             |
|         DHCP server         |     .----------------------.
|          DNS server         |     |         box          |
|                             |     |                      |
|              (10.0.2.2) NIC |<--->|  INT_2 (10.0.2.15)   |
|                             |     |   ^                  |
'-----------------------------'     |   |                  |
.-----------------------------.     |   |                  |
|      Internal Network       |     |   | DHCP server      |
|                             |     |   v                  |
|                           |<----->|  INT_1 (192.168.0.1) |
|                           | |     |                      |
|                           | |     '----------------------'
|                           | |     .----------------------.
|                           | |     |       machine1       |
|                           | |     |                      |
|                           |<----->|  INT_1 (192.168.0.2) |
|                           | |     |                      |
|                           | |     '----------------------'
|                           | |     .----------------------.
|                           | |     |       machine2       |
|                           | |     |                      |
|                           |<----->|  INT_1 (192.168.0.2) |
|                             |     |                      |
'-----------------------------'     '----------------------'
```

You will only have control over "machine2". This computer have Internet access through the "box".

Start the 3 VM and test on machine2 the connectivity quality with this command :

```
timeout --signal SIGINT 1m ping google.com
```

After one minute the result shows the percentage of lost packets. It should be quite high (above 10%).
machine1 and machine2 have the same IP address, which leads to connectivity problems.

Find how to :

- change the IP address to avoid the conflict
- make the IP address dynamic (attributed by the box DHCP server)
