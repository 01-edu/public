# Hole-In-Bin

<center>
<img src="./pictures/meme.png?raw=true" style = "width: 600px !important; height: 737px !important;"/>
</center>

### Objective

This exercise is designed to test your skills and understanding of binary exploitation and reverse engineering. You will need to work through a series of binary exploitation challenges using a provided virtual machine.

### Setup

1. Download the virtual machine image [hole-in-bin.ova](https://assets.01-edu.org/cybersecurity/hole-in-bin/hole-in-bin.ova). For machine running with Apple Silicon CPU or ARM architecture (e.g. M1, M2 ...) use the following [hole-in-bin.utm.zip](https://assets.01-edu.org/cybersecurity/hole-in-bin/hole-in-bin.utm.zip). This image contains all the binaries you will need to exploit.

SHA1: 7db09b7a8fdfe25c286561dfa7ca5b50718bd60c

2. Load the virtual machine image into your virtualization software of choice (e.g., VirtualBox, VMWare).

3. Login using the provided credentials (username: user, password: user).

> You will set up a new VM in the audit session!

### The Challenges

Inside the `/opt/hole-in-bin` directory, you will find a set of binaries (ex00 to ex11) that you need to exploit. Each binary presents a unique challenge and will test different aspects of your knowledge about binary exploitation and reverse engineering.
You can find more details in a `README.txt` file inside each exercise folder.

Your task is to exploit these binaries, following ethical hacking guidelines.

> Using a decompiler is forbidden, use a disassembler instead of it!

- The compiler is used to convert high-level programming language code into machine language code.
- The assembler converts assembly-level language code into machine language code.

> ⚠️ These methods and tools are for educational purposes only, so that you have a better understanding of how to protect against similar vulnerabilities. You must ensure that you do not attempt any exploit-type activity without the explicit permission of the owner of the machine, system or application. Failure to obtain permission risks breaking the law.

### Submission and audit

Files that must be inside your repository:

- A `README.md` file, Which explains all the steps you went through to bypass each exercise.
- All tools you use, and any scripts you write.

> It's forbidden to use external scripts, in the audit you will be asked different questions about the concepts and the practices of this project, prepare yourself!
