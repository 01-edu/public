#### General

##### Check the Repo Content:

Files that must be present in the repository:

- Source code for both the Evasion and Polymorphic programs.
- `README.md` file with detailed documentation.

###### Are all the required files present?

##### Play the Role of a Stakeholder

Conduct a simulated scenario where the student plays the role of a **Binary Analyst** presenting their programs to a team of stakeholders (auditors). Evaluate their understanding, communication skills, and depth of knowledge. Suggested questions include:

- How does your evasion program encrypt and modify binaries to bypass detection?
- What techniques did you use to evade antivirus detection, and how do they work?
- Can you explain how your polymorphic program generates self-modifying binaries?
- What are the ethical and legal implications of developing such programs?
- How would you recommend detecting and defending against similar techniques?

###### Did the student demonstrate a thorough understanding of the project and concepts?

###### Was the student able to communicate effectively and explain their findings?

###### Did the student discuss the potential real-world impact of these techniques?

###### Was the student able to justify the security measures and recommendations they provided?

##### Review the Student Documentation

Verify that the README.md file contains the following:

- **Evasion Program Explanation**: Clearly describe how the program encrypts and modifies binaries.
- **Polymorphic Program Explanation**: Detailed explanation of how the self-modifying binaries are generated.
- **Walkthroughs**: Step-by-step usage instructions for both programs.
- **Technical Insights**: Analysis of binary structure, encryption methods, stealth techniques, and reverse shell payload integration (e.g., how the payload connects to a designated IP and port, and steps to ensure secure testing in an isolated environment).
- **Ethical and Legal Report**: Discussion of the ethical responsibilities and legal considerations.

###### Does the README.md file include a complete and clear walkthrough for both programs?

###### Are technical explanations of the methods and techniques provided and well-documented?

###### Does the documentation discuss the ethical implications and responsibilities of working with these techniques?

##### Test the Challenge

###### Evasion Program:

Prepare a simple binary file (e.g., simple_program.exe) that prints "Hello, World!".
Run the evasion program to encrypt the binary:

```sh
$> ./evasion --encrypt simple_program.exe --output evaded_program.exe --add-size 101 --delay 101
```

**Verify the following functionalities:**

- File size increased by 101 MB.
- Execution delay works as expected (e.g., a delay of 101 seconds).
- The encrypted binary bypasses Windows Defender.
- The encrypted binary bypasses detection by at least 60% of vendors on VirusTotal.

###### Does the program increase the file size by the specified amount?

###### Does the program enforce the execution delay correctly?

###### Does the program bypass Windows Defender?

###### Does the program bypass at least 60% of vendors on VirusTotal?

###### Polymorphic Program:

Prepare the reverse shell payload. A simple reverse shell payload example can be found in the subject. Ensure that the payload connects to a designated IP and port in a controlled environment.

Run the polymorphic program to generate a self-modifying binary:

```sh
$> ./polymorphic --generate polymorphic.exe --payload reverse_shell_payload
```

**Verify the following functionalities:**

- The generated binary maintains its core functionality after modification.
- The reverse shell payload initializes correctly and connects to the designated IP and port upon execution.

###### Does the polymorphic binary maintain functionality after modification?

###### Does the reverse shell payload execute and connect to the expected IP and port?

#### Bonus

###### + Did the student implement advanced obfuscation techniques, such as anti-debugging or stealth memory allocation?

###### + Did the student use a custom encryption algorithm for enhanced obfuscation?

###### + Did the student create a graphical interface for ease of use?

###### + Is this project an outstanding project that exceeds the basic requirements?
