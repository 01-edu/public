## OSINTMaster

<center>
<img src="./resources/osint-meme.png?raw=true" style = "width: 673px !important; height: 439px !important;"/>
</center>

### Introduction:
Open-source intelligence (OSINT) is a key component of cybersecurity, providing valuable insights into potential vulnerabilities and security risks. This project involves creating a tool that performs comprehensive passive reconnaissance using publicly available data.

### Objective:
The goal is to build a multi-functional tool capable of retrieving detailed information based on user inputs such as `Full names`, `IP addresses`, `usernames`, and `domains`. This project will enhance your skills in data analysis, ethical considerations, and the use of various cybersecurity tools and APIs.

By completing this project, You will:
- Develop an understanding of OSINT techniques and their applications.
- Gain practical experience in programming, API integration, and data handling.
- Learn to identify and mitigate security risks, including subdomain takeovers.
- Understand the ethical and legal implications of cybersecurity practices.

### Resources
Some useful resources:
- [Open-source intelligence](https://en.wikipedia.org/wiki/Open-source_intelligence)
- [Doxing](https://en.wikipedia.org/wiki/Doxing)
- [Kali Tools - Recon](https://en.kali.tools/all/?category=recon)
- [OSINT Tools on GitHub](https://github.com/topics/osint-tools)

Before asking for help, ask yourself if you have really thought about all the possibilities.

### Role play
To enhance the learning experience and assess your knowledge, a role play question session will be included as part of this project.
This section will involve answering a series of questions in a simulated real-world scenario where you assume the role of a Cyber Security Expert explaining how to protect information from OSINT techniques to a team or stakeholder.

The goal of the role play question session is to:
- Assess your understanding of OSINT risks and mitigation strategies.
- Test your ability to communicate effectively and explain security measures related to this project.
- Challenge you to think critically about the importance of information security and consider alternative approaches.
- Explain what subdomain takeovers are.

Prepare for a role play question session in the audit.

### Project Requirements
#### Input Handling:
The tool should accept the following inputs: `Full Name`, `IP Address`, `Username`, and `Domain`.

#### Information Retrieval:
- Full Name:
Parse the input to extract "First Name" and "Last Name".
Look up associated information such as phone numbers, addresses, and social media profiles using directory APIs or web scraping.

- IP Address:
Retrieve geolocation data, ISP details, and check for any historical data associated with the IP (e.g., from abuse databases).

- Username:
Check for the presence of the username on at least five known social networks and public repositories.
Retrieve public profile information, such as profile bio, activity status, and follower count.

- Domain and Subdomain Enumeration:
Enumerate subdomains and gather information including IP addresses, SSL certificate details, and potential vulnerabilities.
Identify potential subdomain takeover risks by analyzing DNS records and associated resources.

#### Subdomain Takeover Detection:
Detect and report any subdomains pointing to potentially unclaimed or deprecated resources, indicating a risk of takeover.

#### Output Management:
Store the results in a well-organized file format.

### Usage Examples

#### Command Line Interface:
```sh
$> osintmaster --help

Welcome to osintmaster multi-function Tool

OPTIONS:
    -n  "Full Name"        Search information by full name
    -i  "IP Address"       Search information by IP address
    -u  "Username"         Search information by username
    -d  "Domain"           Enumerate subdomains and check for takeover risks
    -o  "FileName"         File name to save output
```

#### Example Outputs:
```sh
$> osintmaster -n "FNAME LNAME" -o result1.txt
First name: FNAME
Last name: LNAME
Phone Number: +1234567890
Address: Address123, CITY, COUNTRY-CODE
LinkedIn: linkedin.com/in/XX.XX
Facebook: facebook.com/XX.XX
Data Saved in result1.txt
```

#### IP Address:
```sh
$> osintmaster -i 8.8.8.8 -o result2.txt
ISP: Google LLC
City: Mountain View
Country: COUNTRY
ASN: 15169
Known Issues: No reported abuse
Data Saved in result2.txt
```

#### Username:
```sh
$> osintmaster -u "@username" -o result3.txt
Facebook: Found
Twitter: Found
LinkedIn: Found
Instagram: Not Found
GitHub: Found
Recent Activity: Active on GitHub, last post 1 days ago
Data Saved in result3.txt
```

#### Domain and Subdomain Enumeration:
```sh
$> osintmaster -d "example.com" -o result4.txt
Main Domain: example.com

Subdomains found: 3
  - www.example.com (IP: 123.123.123.123)
    SSL Certificate: Valid until 2030-03-01
  - mail.example.com (IP: 123.123.123.123)
    SSL Certificate: Valid until 2030-03-01
  - test.example.com (IP: 123.123.123.123)
    SSL Certificate: Not found

Potential Subdomain Takeover Risks:
  - Subdomain: test.example.com
    CNAME record points to a non-existent AWS S3 bucket
    Recommended Action: Remove or update the DNS record to prevent potential misuse

Data saved in result4.txt
```

### Bonus
If you complete the mandatory part successfully, and you still have free time, you can implement anything that you feel deserves to be a bonus, for example:

- User Interface: Develop a graphical user interface (GUI) for better user accessibility.
- PDF Generation: Add a feature to generate your OSINT result as PDF files.

Challenge yourself!

### Documentation
Create a `README.md` file that provides comprehensive documentation for your tool (prerequisites, setup, configuration, usage, ...). This file must be submitted as part of the solution for the project.

### Ethical and Legal Considerations
- Get Permission: Always obtain explicit permission before gathering information.
- Respect Privacy: Collect only necessary data and store it securely.
- Follow Laws: Adhere to relevant laws such as GDPR and CFAA.
- Report Responsibly: Privately notify affected parties of any vulnerabilities.
- Educational Use Only: Use this tool and techniques solely for learning and improving security.

> ⚠️ Disclaimer: This project is for educational purposes only. Ensure all activities comply with legal and ethical standards. The institution is not responsible for misuse of the techniques and tools demonstrated.

### Submission and audit
Upon completing this project, you should submit the following:

- Your documentation in the `README.md` file.
- The Source code of your tool.
- Any required files to run your tool.