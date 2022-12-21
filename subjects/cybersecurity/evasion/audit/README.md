#### General

###### Is the student able to explain how the Anti-Viruses detect the viruses?

###### Is the student able to explain clearly how he can bypass the Anti-Viruses?

###### Is the student able to explain clearly how his program works?

##### Check the Repo content

Files that must be inside your repository:

- Your program source code.

- a README.md file, Which clearly explains how to use the program.

###### Are the required files present?

##### Evaluate the student's submission

Open the student program in a Windows system and add as argument a simple program that you can find on your Windows (calc.exe, ...)

Compare the hash with a checker, before and after binary encryption

###### Has the signature of the binary argument been modified by the student's program?

##### Open the program with a hex editor or disassembler and compare binary argument before and after binary encryption

###### Has the form of the program been modified?

##### Launch the program that has just been encrypted

###### Does the program run normally after 101 seconds?

#### Bonus

Add [mal-track.exe](</res/mal-track(Fynloski%20sample%2C%20ON%20VM%20ONLY).zip>) as an argument to the student project without running it.

> It is a program that is currently detected by 61/68 antivirus.
> https://www.virustotal.com/gui/file/a164abbb6778e1378af208b4a3d4833c2b226c68452d2151fb14e2e01a578fdd?nocache=1

Add mal-track.exe as an argument to the student program, and upload the new encrypted version of mal-track.exe to an online Virus Scanner.

The student can refuse to have his program uploaded to VirusTotal and therefore choose another scanner that does not send samples to preserve his algorithm.

###### +Does the new encrypted version of the binary upgrade to at least 40/68?
