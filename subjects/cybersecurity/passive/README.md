# Passive

### Introduction

Information gathering is one of the most important steps during a pentest, it can be considered the longest step.

### Objective

The goal of this project is for you to become more comfortable with open source investigative methods

### Advice

Before asking help, ask yourself if you have really thought about all the possibilities.  
https://en.kali.tools/all/?category=recon  
https://github.com/topics/osint-tools  
https://en.wikipedia.org/wiki/Open-source_intelligence  
https://en.wikipedia.org/wiki/Doxing

### Guidelines

You are going here to create your first passive recognition tool, you have the choice of language, however your program will have to recognize the information entered (FULL NAME, IP, @login).

For the case of the full name, it will have to recognize the entry:
Last name First Name,
then look in the directories for the telephone number and the address.

If it is the ip address, your tool should display at least the city and the name of the internet service provider.

If it is an username, your tool will have to check if this username is used in at least 5 known social networks.

The result should be stored in a result.txt file (result2.txt if the file already exists)

### Bonus

You can add more API features

### Usage

```
$> passive --help

Welcome to passive v1.0.0

OPTIONS:
    -fn         Search with full-name
    -ip         Search with ip address
    -u          Search with username

$>  passive -fn "Jean Dupont"
First name: John
Name: Dupont
Address: 7 rue du Progrès
75016 Paris
Number: +33601010101
Saved in result.txt

$>  passive -ip 127.0.0.1
ISP: FSociety, S.A.
City Lat/Lon:	(13.731) / (-1.1373)
Saved in result2.txt

$>  passive -u "@user01"
Facebook : yes
Twitter : yes
Linkedin : yes
Instagram : no
Skype : yes
Saved in result3.txt
```

### Submission and audit

Files that must be inside your repository:

- Your program source code.
- a README.md file, Which clearly explains how to use the program.

Don’t hesitate to double check the names of your folders and files to ensure they are correct!
