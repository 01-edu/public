## Ransomware-Lab

<center>
<img src="pictures/ransomware-lab.jpg" style="width: 560px; height: auto;">
</center>

### Introduction

To effectively defend against ransomware attacks, it’s essential to understand their mechanics. This project challenges you to think like a black hat and develop a controlled ransomware simulation. By creating both encryption and decryption programs, you’ll gain insight into how ransomware like WannaCry operates.

This exercise is conducted in a secure virtual machine environment to ensure safety and controlled experimentation.

### Objective

The goal of this project is to develop a ransomware simulation and a corresponding decryption tool to understand ransomware mechanics. By completing this project, you will:

- Learn the core principles of file encryption and decryption.
- Understand how ransomware manipulates files and affects systems.
- Gain hands-on experience in developing secure cryptographic programs.
- Enhance your understanding of how to defend against such attacks in real-world scenarios.

### Role Play

As part of the project, you will participate in a role-play session where you act as a **Malware Developer** presenting your ransomware simulation to a team of stakeholders. Be prepared to:

- Explain the design and functionality of your ransomware and decryption tool.
- Discuss how ransomware can bypass antivirus detection and the challenges involved.
- Analyze the ethical considerations and potential risks of working with ransomware simulations.
- Provide recommendations for improving ransomware detection and system security.

### Project Requirements

#### Setup and Environment

- Set up a Windows-based virtual machine for development and testing.

> Ensure that the environment is isolated to prevent accidental damage or spread of the ransomware!

#### The Challenge

1. **Ransomware Development**:

   - Develop a ransomware program that encrypts all files.
   - Place a text file on the desktop with the following message:
     ```
     All of your files have been encrypted.
     To unlock them, contact me with your encryption code at email@email.com.
     Your encryption code is: {randomly_generated_code}
     ```
   - Generate a unique random encryption code for each affected system.

2. **Decryption Program**:

   - Develop a decryption program that uses the encryption code to restore the encrypted files.
   - Ensure the decryption program works reliably for each affected system based on their encryption code.

3. **Detection Avoidance**:
   - Your ransomware program must use advanced techniques to avoid detection by antivirus software.
   - Ensure the ransomware bypasses detection in **Windows Defender** and more than 80% of security vendors on [VirusTotal](https://www.virustotal.com/).

### Documentation

Create a `README.md` file that includes:

- **Ransomware Program Overview**: Explain how your program works and its intended functionality.
- **Decryption Program Overview**: Describe how the decryption program works and how it interacts with the encryption tool.
- **Technical Explanation**: Provide details on the encryption algorithm used, how the encryption code is generated, and how files are decrypted.
- **Testing and Usage Instructions**: Include clear instructions for testing the ransomware and decryption tools in the virtual machine.
- **Ethical and Legal Report**: Discuss the ethical responsibilities of developing ransomware simulations and the importance of using such knowledge to improve security defenses.

### Bonus

If you complete the mandatory part successfully, and you still have free time, you can implement additional features, such as:

- **Multiple File Type Support**: Expand the ransomware to encrypt multiple file types (e.g., images, videos).
- **Stealth Techniques**: Implement advanced techniques to improve ransomware stealth.
- **Custom Encryption Algorithms**: Develop your encryption algorithm for enhanced security.

Challenge yourself!

### Ethical and Legal Considerations

You are responsible for ensuring all ransomware testing is conducted within a secure, isolated environment. Do not use or share ransomware outside of this project. Misuse of these techniques is strictly prohibited and may violate local laws.

> ⚠️ Disclaimer: This project is for educational purposes only. Unauthorized use of these techniques is prohibited and may be illegal.

### Submission and Audit

Submit the following:

- Source code for both the ransomware and decryption programs.
- `README.md` file with detailed documentation.

Ensure VirtualBox or equivalent software is installed for the audit.

### Resources

Some useful resources:

- [Windows Cryptographic Functions](https://docs.microsoft.com/en-us/windows/win32/api/bcrypt/): Learn about cryptographic APIs in Windows.
- [File Management Functions](https://docs.microsoft.com/en-us/windows/win32/fileio/file-management-functions): Understand file manipulation in Windows.
- [VirusTotal](https://www.virustotal.com/): Check your ransomware against antivirus detections.
- [Microsoft Security Basics](https://docs.microsoft.com/en-us/windows/security/): Learn about Windows security features.
