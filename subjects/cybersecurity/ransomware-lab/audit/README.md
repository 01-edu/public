#### General

##### Check the Repo Content

Files that must be present in the repository:

- Source code for both the ransomware and decryption programs.
- Detailed documentation in the `README.md` file containing program overviews, technical explanations, testing instructions, and ethical considerations.

###### Are all the required files present?

##### Play the Role of a Stakeholder

Conduct a simulated scenario where the student acts as a **Malware Developer** presenting their ransomware simulation to a hypothetical team of stakeholders. Evaluate their understanding, communication skills, and depth of knowledge. Suggested questions include:

- What is ransomware, and how does it operate on a system?
- How does your ransomware encrypt files and generate the unique encryption code?
- How does your decryption program restore the encrypted files?
- What techniques did you use to bypass antivirus detections?
- What ethical considerations should be addressed when working with ransomware simulations?

###### Did the student demonstrate a thorough understanding of ransomware mechanics and encryption techniques?

###### Was the student able to communicate effectively and explain their programs?

###### Did the student discuss the ethical and real-world implications of ransomware?

###### Was the student able to justify their program’s design and implementation?

##### Review the Student Documentation

Verify that the `README.md` file contains:

- **Ransomware Program Overview**: Explanation of how the ransomware operates.
- **Decryption Program Overview**: Description of how the decryption program restores encrypted files.
- **Technical Explanation**: Details about encryption, code generation, and decryption logic.
- **Testing and Usage Instructions**: Clear steps for safely testing the ransomware and decryption tools in a secure environment.
- **Ethical and Legal Report**: Discussion of ethical responsibilities and legal considerations.

###### Does the documentation provide a clear and complete overview of the ransomware and decryption programs?

###### Are the technical details of encryption and decryption well-explained?

###### Does the documentation highlight the importance of ethical and legal practices?

##### Check the detection avoidance

Scan the ransomware program in **Windows Defender** and **VirusTotal**.

###### Does the ransomware bypass detection in Windows Defender?

###### Does the ransomware bypass detection in more than 80% of security vendors on VirusTotal?

##### Check the file encryption

Create files with random data in the virtual machine and run the ransomware program.

###### Does the ransomware encrypt all files in the specified directory or with the targeted extension?

###### Does the ransomware create a text file on the desktop with the encryption code and contact information?

##### Check the file decryption

- Ask the student to provide the decryption program for testing.
- Run the decryption program using the encryption code.

###### Does the decryption program successfully restore all encrypted files?

##### Manual Verification:

Ask the student to explain how the encryption code is generated and used.
Verify that the encryption code is random and unique for each affected system.

###### Is the encryption code unique for each affected system?

###### Is the encryption code generation method random and effective?

#### Bonus

###### + Did the student extend the ransomware functionality to support multiple file types?

###### + Did the student implement advanced stealth techniques?

###### + Did the student develop a custom encryption algorithm?

###### + Is this project an outstanding submission that exceeds the basic requirements?
