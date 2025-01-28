## Hole-In-Bin

<center>
<img src="./pictures/meme.png?raw=true" style="width: 808px !important; height: auto !important;"/>
</center>

### Introduction

**Hole-In-Bin** is a comprehensive learning platform designed to teach participants the fundamentals of reverse engineering and binary exploitation. By analyzing and exploiting vulnerabilities in binaries, you will strengthen your understanding of low-level system mechanics and learn essential techniques for identifying and mitigating security risks.

### Objective

The primary objective of this project is to develop your skills in binary exploitation and reverse engineering. By analyzing and exploiting the provided binaries, you will gain practical experience with concepts such as buffer overflows and memory vulnerabilities.

By completing this project, you will:

- Analyze and exploit binaries to uncover security vulnerabilities.
- Understand assembly code and memory structures.
- Gain hands-on experience with debugging and disassembly tools.
- Enhance your ability to communicate complex technical findings effectively.

### Role Play

As part of the project, you will participate in a role-play session where you act as a **Security Analyst**. Be prepared to:

- Explain your approach to analyzing and exploiting binaries.
- Discuss the real-world implications of the vulnerabilities you identified.
- Propose practical solutions and prevention methods.

The role-play session will test your ability to communicate complex technical concepts effectively and ethically.

### Project Requirements

#### Setup and Installation

Download the provided VM image and set it up in VirtualBox or UTM:

- **Download Links**:

  - [hole-in-bin.ova](https://assets.01-edu.org/cybersecurity/hole-in-bin/hole-in-bin.ova)
  - For Apple Silicon or ARM architecture (e.g. M1, M2 ...), use [VM Image - UTM Format](https://assets.01-edu.org/cybersecurity/hole-in-bin/hole-in-bin.utm.zip).

- **SHA1 Checksums**:

  - SHA1 for `hole-in-bin.ova`: 7db09b7a8fdfe25c286561dfa7ca5b50718bd60c
  - SHA1 for `hole-in-bin.utm.zip`: fc93533b2054d10d03b09d53c223e57bf7ac7b62

This VM contains all the binaries you will need to exploit.

> Ensure the VM is installed and properly configured for the audit.

#### Access:

Log in using the following credentials:

- **Username**: `user`
- **Password**: `user`

#### The Challenges

Navigate to `/opt/hole-in-bin` and review the binaries. Each folder contains:

- A binary file for exploitation.
- A `README.txt` file explaining the exercise requirements and providing hints.

Your task is to exploit these binaries, following ethical hacking guidelines.

### Guidelines

- **Allowed Tools**:

  - Debuggers/Disassembler: Ghidra, GDB, PEDA
  - Scripting: Python, Bash

- **Prohibited**:
  - Automated external scripts or tools for exploitation.
  - Decompiler.

> Using a Decompiler is forbidden, use a Debuggers/Disassembler instead of it!

- The disassembler is a computer program that translates machine language into assembly language.
- The assembler converts assembly-level language code into machine language code.

### Documentation

Create a `README.md` file that includes:

1. **Challenge Walkthroughs**:

   - Step-by-step explanation of how you exploited each binary.
   - Tools and commands used.
   - Key takeaways for each challenge.

2. **Remediation Suggestions**:

   - Practical steps to fix or mitigate the identified vulnerabilities.

3. **Ethical Hacking Report**:
   - Importance of proper authorization.
   - Legal and ethical boundaries.
   - Responsible disclosure practices.

### Bonus

If you complete the mandatory part successfully, and you still have free time, you can implement anything that you feel deserves to be a bonus, for example:

- **Exploring Alternative Exploitation Paths**: Document different approaches to solving the challenges.

Challenge yourself!

### Ethical and Legal Considerations

This project is strictly for educational purposes. All testing must be conducted in the provided VM environment. Unauthorized attempts to exploit vulnerabilities on live systems or networks are illegal and unethical.

> ⚠️ **Disclaimer**: This project is for learning purposes only. Adhere to ethical hacking practices and legal standards. Misuse of these techniques is prohibited.

### Submission and Audit

Submit the following:

- `README.md` with your walkthrough and vulnerability report email.
- Any scripts or files used during the project.

Ensure VirtualBox is installed for the audit.

> It's forbidden to use external scripts, in the audit you will be asked different questions about the concepts and the practices of this project, prepare yourself!

### Resources

Some useful resources:

- [HackTricks - Binary Exploitation](https://book.hacktricks.xyz/binary-exploitation): A detailed guide to binary exploitation techniques.
- [Radare2](https://radare.org/n/radare2.html): An open-source framework for reverse engineering and analyzing binaries.
- [Ghidra](https://ghidra-sre.org/): A software reverse engineering suite developed by the NSA.
- [GDB Documentation](https://www.gnu.org/software/gdb/documentation/): Official GNU Debugger documentation.
