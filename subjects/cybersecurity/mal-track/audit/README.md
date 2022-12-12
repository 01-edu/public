#### General

Knowledge-Check:

###### Is the student able to explain clearly how we can manage the startup programs in windows?

###### Is the student able to explain clearly how he get the ip of the attacker from the malware?

###### Is the student able to explain clearly how his program works?

Create an offical Windows virtual machine.

##### Check the Repo content

Files that must be inside your repository:

- Your program source code.
- a README.md file, Which clearly explains how to use the program.

###### Does the required files present?

##### Evaluate the student's submission

Create a Windows virtual machine

Add [mal-track.exe](<../resources/mal-track(Fynloski%20sample%2C%20ON%20VM%20ONLY).zip>) to anti-virus exceptions

Launch the mal-track.exe program check if it's exist in the task manager, then launch the student program

We are going to verify that the program developed by the student kills the malware process.
Open task manager CTRL+ALT+DEL or Windows key+R -> msconfig.

###### Has mal-track.exe been killed from task manager?

We are going to verify that the program developed by the student removes our malware from the start of the machine
Open Registry Editor Windows key+R -> regedit

###### maltrack not located in HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run?

###### maltrack not located in HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\RunOnce?

###### maltrack not located in HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run?

###### maltrack not located in HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce?

The attacker's IP address is "127.0.0.1"

###### Does the student program display the attacker's IP address?

###### Can the student trace the IP address of the attacker manually with an hexadecimal editor?
