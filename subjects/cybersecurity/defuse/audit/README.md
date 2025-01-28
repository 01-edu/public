#### General

##### Check the Repo Content:

- The student's malware removal program source code.
- Detailed documentation in the README.md file explains how to use the program and includes the malware mitigation report email.

###### Are all the required files present?

##### Play the Role of a Stakeholder

Conduct a simulated scenario where the student plays the role of a **Malware Analyst** presenting their findings to a team of stakeholders. Evaluate their understanding, communication skills, and depth of knowledge. Suggested questions include:

- How did you analyze the malware and identify its behavior?
- Can you explain how your program works and the steps it takes to neutralize the malware?
- What impact could this malware have if executed in an uncontrolled environment?
- What measures would you recommend to prevent similar malware infections in the future?
- How did you ensure that the malware was analyzed and mitigated ethically within a secure environment?

###### Did the student demonstrate a thorough understanding of the project and its concepts?

###### Was the student able to communicate effectively and explain their findings?

###### Did the student discuss the real-world impact of malware and their recommendations for mitigation?

##### Review the Student Documentation

Verify that the `README.md` file contains:

- **Program Explanation**: Explain the functionality of your program and how it neutralizes the malware.
- **Walkthrough**: Describe the step-by-step process of how you analyzed and eradicated the malware.
- **Remediation**: Suggest ways to prevent similar malware infections.
- **Malware Mitigation Report Email**: Include your drafted report.
- **Ethical Hacking Report**: Discuss the ethical responsibilities when performing malware analysis.

###### Does the README file clearly explain the program's functionality?

###### Does the README include a walkthrough of the analysis and removal process and the remediation?

###### Does the README file include the malware mitigation report email?

###### Does the README file include the ethical hacking report?

##### Set Up the Virtual Machine:

1. Create a Windows virtual machine.
2. Download the provided malware sample ([Malware Sample](<./resources/Fynloski(ON VM ONLY).zip>)).
3. Add the malware sample to antivirus exceptions.
4. Launch the malware executable.

> Ensure the malware is executed within a secure, isolated environment to prevent accidental spread or damage.

###### Does the malware executable appear in the Task Manager?

##### Launch the Student's Program:

Run the student's malware removal program.

###### Does the program developed by the student terminate the malware process?

###### Is the malware process removed from the Task Manager?

###### Does the program remove the malware's persistence mechanisms from the system?

##### Check Startup Persistence:

Open Task Manager using **CTRL+ALT+DEL** or **Windows key+R -> msconfig**.

###### Is the malware removed from the startup programs?

Open the Registry Editor using **Windows key+R -> regedit**. Verify the following registry keys:

- **HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run**
- **HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\RunOnce**
- **HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run**
- **HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce**

###### Is the malware absent from all the above registry locations?

##### Attacker's IP Address:

Confirm that the program extracts and displays the attacker's IP address.

The attacker's IP address is `127.0.0.1`

###### Does the student's program display the attacker's IP address correctly?

##### Review the Student's Malware Mitigation Report Email

Ask the student to present their **Mitigation Report Email**. Ensure that it includes:

- **Subject Line**: "Malware Analysis Report: Mitigation of [Malware Name]"
- **Summary**: A concise description of the malware's behavior and impact.
- **Proof of Mitigation**: Evidence that the malware process was terminated and persistence mechanisms were removed.
- **Attacker Information**: Display the attacker's IP address.

###### Does the email contain a concise summary of the malware's behavior and impact?

###### Does the email include proof of mitigation (e.g., terminated process, removed persistence mechanisms)?

###### Does the email display the attacker's IP address?

#### Bonus

###### + Did the student implement additional features, such as dynamic analysis automation or threat intelligence integration?

###### + Is this project an outstanding project that exceeds the basic requirements?
