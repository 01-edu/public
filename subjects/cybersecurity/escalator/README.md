## Escalator

<center>
<img src="resources/escalator.jpg" style="width: 702px; height: 395px;">
</center>

### Introduction

Privilege escalation is a fundamental concept in cybersecurity, allowing attackers to elevate their privileges on a system to gain higher levels of access. In this project, you will explore and understand privilege escalation techniques using a virtual machine (VM) designed for this challenge. The goal is to gain root access to the system and retrieve a hidden flag.

### Objective

The goal of this project is to help you gain hands-on experience with privilege escalation. You will install a provided VM locally, identify the IP address, enumerate the system for potential vulnerabilities, and ultimately escalate your privileges from a regular user to root.

By completing this project, you will:

- Develop a practical understanding of privilege escalation in Linux systems.
- Learn to identify and exploit vulnerabilities that can lead to unauthorized access.
- Gain experience in ethical hacking and penetration testing methodologies.
- Understand the importance of securing systems against privilege escalation attacks.

### Role Play

As part of the project, you will participate in a role-play session where you will act as a **Penetration Tester** presenting your findings to a hypothetical team of stakeholders. Prepare to discuss:

- How you identified and leveraged the vulnerabilities.
- The impact these vulnerabilities could have in real-world scenarios.
- Recommendations for securing the system against privilege escalation.
- The importance of responsible disclosure and ethical hacking practices.

### Project Requirements

#### Setup and Installation

Download the provided VM image and set it up in VirtualBox:

- **Download Links**:
  - [VM Image - OVA Format](https://assets.01-edu.org/cybersecurity/local/01-Local.ova)
  - For Apple Silicon or ARM architecture (e.g., M1, M2), use [VM Image - UTM Format](https://assets.01-edu.org/cybersecurity/local/01-Local1.utm.zip)
- **SHA1 Checksums**:
  - `01-Local1.ova`: `f3422f3364fd38e8183740f8f57fa951d3f6e0bf`
  - `01-Local1.utm.zip`: `d4a40ca50044778ddc01a57ac16382e4140000e0`

Ensure VirtualBox is installed and properly configured for the audit.

#### The Challenge

- **Achieve Root Access**: Your main objective is to escalate privileges and gain root access.
- **Find the Flag**: Once root access is achieved, locate and display the flag in `/root/root.txt`.
- **Document Your Process**: Keep thorough notes on your methodology for reporting purposes.

#### Guidances

- **IP Discovery**: Use network scanning tools to identify the IP address of the VM.
- **Enumeration**: Leverage system enumeration tools to gather information on running processes, services, and configurations.
- **Privilege Escalation**: Investigate user permissions, files, and services that could be used for privilege escalation.
- **Document Your Process**: Keep thorough notes on your methodology for reporting purposes.

#### Vulnerability Report Email

Draft an email as if reporting the vulnerability through a bug bounty program. The report should show proof of root access and include the flag without disclosing full reproduction steps.

**Your report should include:**

- **Subject Line**: "Security Vulnerability Report: Privilege Escalation in [System/Environment Name]"
- **Summary**: A concise description of the vulnerability.
- **Impact**: The potential risk and implications of the vulnerability.
- **Proof of Root Access**: Indicate root access and display the flag.

**Example Email**:

---

**To**: security@[organization].com  
**Subject**: Security Vulnerability Report: Privilege Escalation in [System/Environment Name]

**Dear Security Team,**

I am writing to report a potential security vulnerability identified during an educational penetration testing exercise. The vulnerability involves privilege escalation within [System/Environment Name]. Below are the details:

**Summary**:  
The vulnerability allows an authenticated user to escalate their privileges to root due to [brief reason, e.g., misconfigured file permissions or an exploitable SUID binary].

**Impact**:  
If exploited, this could lead to unauthorized root access, compromising the system's integrity and exposing sensitive data.

**Proof of Root Access**:  
I successfully gained root access and retrieved the flag: {flag}
<screenshot>

Please feel free to reach out for further details or clarification.

**Best regards,**  
[Your Name]  
[Your Contact Information]

### Documentation

Create a `README.md` file that contains the following:

- **Walkthrough**: Describe the step-by-step process of how you exploited the vulnerability.
- **Remediation**: Suggest ways to fix or mitigate the vulnerability.
- **Vulnerability Report Email**: Include your drafted report.
- **Ethical Hacking Report**: Discusses the ethical responsibilities when performing security testing. This report should cover the following points:

1. The importance of obtaining proper authorization before testing.
2. The legal and ethical boundaries of vulnerability testing.
3. How to report vulnerabilities responsibly and avoid causing harm.

### Bonus

If you complete the mandatory part successfully, and you still have free time, you can implement anything that you feel deserves to be a bonus, for example:

- **Exploring Alternative Methods**: Document different privilege escalation paths.
- **Post-Exploitation Analysis**: Explain potential next steps after gaining root access.

Challenge yourself!

### Ethical and Legal Considerations

You are responsible for following ethical hacking guidelines and only performing security testing in the provided VM environment. Do not use these techniques on unauthorized systems.

> ⚠️ Disclaimer: This project is for educational purposes only. All testing must be done ethically and following legal standards. Unauthorized use of these techniques is prohibited and may be illegal.

### Submission and Audit

Submit the following:

- `README.md` with your walkthrough and vulnerability report email.
- Any scripts or files used during the project.

Ensure VirtualBox is installed for the audit.

### Resources

Some useful resources:

- [Privilege_escalation](https://en.wikipedia.org/wiki/Privilege_escalation): Privilege escalation - Wikipedia.
- [Nmap](https://nmap.org/): A powerful network scanning tool for discovering hosts and services.
- [Dirsearch](https://github.com/maurosoria/dirsearch): A web path scanner useful for directory enumeration.
