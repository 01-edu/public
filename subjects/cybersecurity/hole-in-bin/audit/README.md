#### General

##### Check the Repo content:

- A `README.md` file that explains all the steps taken to bypass the challenges in the project.
- All tools and scripts used, including their purpose and implementation.

###### Are all the required files present in the repository?

##### Play the Role of a Stakeholder

Conduct a simulated scenario where the student plays the role of a **Security Analyst** presenting their findings to a team of stakeholders (auditors). Evaluate their understanding, communication skills, and depth of knowledge. Suggested questions include:

- How did you analyze and exploit the binaries?
- What vulnerabilities did you identify, and what is their impact?
- What tools and techniques did you use for exploitation?
- How would you recommend mitigating these vulnerabilities?
- How did you ensure adherence to ethical guidelines?

###### Did the student demonstrate a thorough understanding of the project and concepts?

###### Was the student able to communicate effectively and explain their findings?

###### Did the student discuss the real-world implications and propose practical solutions?

##### Review the Student's `README.md` File:

Verify that the documentation contains:

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

###### Is the documentation clear, well-structured, and complete?

###### Does the documentation reflect the student's thought process and understanding?

##### Set up the virtual machine:

1. Download the virtual machine image:

   - [hole-in-bin.ova](https://assets.01-edu.org/cybersecurity/hole-in-bin/hole-in-bin.ova)
   - For machines using Apple Silicon or ARM architecture, use [hole-in-bin.utm.zip](https://assets.01-edu.org/cybersecurity/hole-in-bin/hole-in-bin.utm.zip)

2. Ensure the SHA1 checksums match:
   - SHA1 for `hole-in-bin.ova`: `7db09b7a8fdfe25c286561dfa7ca5b50718bd60c`
   - SHA1 for `hole-in-bin.utm.zip`: `fc93533b2054d10d03b09d53c223e57bf7ac7b62`

```console
$ sha1sum <filename>
<SHA1>
```

2. Load the VM into the virtualization software (e.g., VirtualBox, UTM).

3. Confirm successful login with the provided credentials:
   - **Username**: `user`
   - **Password**: `user`

###### Is the VM properly configured and running?

##### Disassemble and Explain the Binaries:

- The disassembler is a computer program that translates machine language into assembly language.
- The assembler converts assembly-level language code into machine language code.

> **Note:** Decompilers are not allowed. Students must use disassemblers to analyze the binaries.

###### Has the student successfully disassembled all the binaries?

###### Can the student explain the purpose and functionality of the binaries?

###### Did the student demonstrate an understanding of reverse engineering principles and binary mechanics?

##### Exploit the Binaries:

1. Analyze and exploit each binary according to the guidelines.
2. Students must not use external scripts or tools for automated exploitation.

###### Did the student successfully exploit all the binaries?

###### Did the student demonstrate a clear understanding of binary exploitation techniques?

#### Bonus:

###### + Did the student explore alternative exploitation paths and document them?

###### + Is this project an outstanding project that exceeds the basic requirements?
