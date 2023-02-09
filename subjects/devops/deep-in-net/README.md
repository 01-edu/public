## deep-in-net

![serverRoomMeme](https://assets.01-edu.org/devops-branch/DeepInNet/serverRoomMeme.jpg)

In this project, you will discover the Cisco Packet tracer which is a useful tool for networking students, And you will understand some important networking concepts as a cloud & DevOps student.

### Objectives

Discover and deal with some important networking devices.

Understand and usage of some important Services and protocols.

Understand OSI Model.

Discover some important networking commands in Linux.

### Advice

Avoid using any of the IP Subnet Calculator tools.

Use CLI in the cisco packet tracer to configure the devices.

Use commands to test the connectivity for different protocols.

Debug communication using commands instead of a GUI.

> Networking is really important in different IT specialties and specialties for cloud and DevOps engineering.
> Be curious and never stop searching!

### Guideline

> You must respect the defined netmask and IP address for each network!

### Instructions

#### Cisco Packet Tracer:

In your machine or your virtual machine, you have to install Cisco Packet Tracer:

![PacketTracer](https://assets.01-edu.org/devops-branch/DeepInNet/PacketTracer.jpg)

Take your time to discover it!

> You will need it in the audit sessions.

#### Exercise 1:

In your Cisco PacketTracer create this network:

![ex01](https://assets.01-edu.org/devops-branch/DeepInNet/ex01.jpg)

- PC0 can communication with PC1.
- PC2 can communication with PC3.
- PC4 can communication with PC5.

![ex01-scenario](https://assets.01-edu.org/devops-branch/DeepInNet/ex01-scenario.jpg)

- Knowledge: What is RJ-45 cable?
- Knowledge: Understand what is difference between straight-through and crossover RJ-45 cables.

#### Exercise 2:

In your Cisco PacketTracer create this network:

![ex02](https://assets.01-edu.org/devops-branch/DeepInNet/ex02.jpg)

- All computers connected to the Switch must be connected.
- All computers connected to the Hub must be connected.

![ex02-scenario](https://assets.01-edu.org/devops-branch/DeepInNet/ex02-scenario.jpg)

- Knowledge: What is the Switch and what is its role?
- Knowledge: What is the Hub and what is its role?
- Knowledge: What is the difference between the Switch and the Hub?
- Knowledge: At what layer of the OSI model do the switch and the hub operate?

#### Exercise 3:

In your Cisco PacketTracer create this network:

![ex03](https://assets.01-edu.org/devops-branch/DeepInNet/ex03.jpg)

- All Servers must have static IP addresses.
- Your servers should only provide the service specified for them!
- All PC's IP addresses must be assigned by the DHCP server.
- Your HTTPS Server must show a hello message and HTTP must be disabled!

![ex03-https](https://assets.01-edu.org/devops-branch/DeepInNet/ex03-https.jpg)

- You must have a user "deepinnet" with RWDNL access in your FTP Server.

![ex03-ftp](https://assets.01-edu.org/devops-branch/DeepInNet/ex03-ftp.jpg)

- In your DNS Server you must add these records:
  deep-in-net.local > 192.168.1.99
  deep-in-net.com > deep-in-net.local

- "https://deep-in-net.com" must redirect to your HTTPS server.

![ex03-dns](https://assets.01-edu.org/devops-branch/DeepInNet/ex03-dns.jpg)

- Knowledge: What is a Server and what is its role?
- Knowledge: What is DHCP and what is its role?
- Knowledge: What is DNS and what is its role?
- Knowledge: What is HTTP and what is its role?
- Knowledge: What is HTTPS and what is its role?
- Knowledge: What is FTP and what is its role?
- Knowledge: What is TCP and UDP communication and difference between them?
- Knowledge: What layer are TCP and UDP?
- Knowledge: What is the port in networking?
- Knowledge: What is the port and OSI Model layer of each used protocol?
- Knowledge: You must understand the DNS Records types!

#### Exercise 4:

In your Cisco PacketTracer create this network:

![ex04](https://assets.01-edu.org/devops-branch/DeepInNet/ex04.jpg)

- The 2 PCs must communicate with each other.

![ex04-scenario](https://assets.01-edu.org/devops-branch/DeepInNet/ex04-scenario.jpg)

- Knowledge: What is the Router and what is its role?
- Knowledge: What is the difference between the Switch and the Router?
- Knowledge: At what layer of the OSI model does a network Router?
- Knowledge: What is Default Gateway?

#### Exercise 5:

In your Cisco PacketTracer create this network:

![ex05](https://assets.01-edu.org/devops-branch/DeepInNet/ex05.jpg)

- All devices connected to the same switch must be able to communicate with each other.
- All devices in subnet 1 can communicate with all devices in subnet 2.
- All devices in subnet 2 can communicate with all devices in subnet 1.

![ex05-scenario](https://assets.01-edu.org/devops-branch/DeepInNet/ex05-scenario.jpg)

#### Exercise 6:

In your Cisco PacketTracer create this network:

![ex06](https://assets.01-edu.org/devops-branch/DeepInNet/ex06.jpg)

- The PC in subnet 1 can communicate with The PC in subnet 2.
- The PC in subnet 2 can communicate with The PC in subnet 1.

![ex06-scenario](https://assets.01-edu.org/devops-branch/DeepInNet/ex06-scenario.jpg)

- Knowledge: What is a routing table and what is its role?

#### Exercise 7:

In your Cisco PacketTracer create this network:

![ex07](https://assets.01-edu.org/devops-branch/DeepInNet/ex07.jpg)

- All devices connected to the same switch must be able to communicate with each other.
- All devices in subnet 1 can communicate with all devices in subnet 2.
- All devices in subnet 2 can communicate with all devices in subnet 1.

![ex07-scenario](https://assets.01-edu.org/devops-branch/DeepInNet/ex07-scenario.jpg)

#### Exercise 8:

In your Cisco PacketTracer create this network:

![ex08](https://assets.01-edu.org/devops-branch/DeepInNet/ex08.jpg)

- All devices connected to the same switch must be able to communicate with each other.
- All devices in subnet 1 can communicate with all devices in subnet 2.
- All devices in subnet 1 can communicate with all devices in subnet 3.
- All devices in subnet 2 can communicate with all devices in subnet 1.
- All devices in subnet 2 can communicate with all devices in subnet 3.
- All devices in subnet 3 can communicate with all devices in subnet 1.
- All devices in subnet 3 can communicate with all devices in subnet 2.

![ex08-scenario](https://assets.01-edu.org/devops-branch/DeepInNet/ex08-scenario.jpg)

### Bonus

If you complete the mandatory part perfectly, you can move to this part. You can add anything you feel deserves to be a bonus.

Challenge yourself!

### Submission and audit

You must save your Exercises solutions to "ptk" files, and then push them to your repository.

Your repository must look like this:

```console
user:~/deep-in-net$ ls
ex01.pkt
ex02.pkt
ex03.pkt
ex04.pkt
ex05.pkt
ex06.pkt
ex07.pkt
ex08.pkt
bonus.pkt
user:~/deep-in-net$
```

> In the audit you must answer a group of questions, and recreate a network in the cisco packet tracer.
> You must also explain any calculations that you have made without the use of tools!
> If you fail to do any of this, then your project will be considered a failed project! so take your time to learn and practice.
