## OSINT Master

![OSINT Meme](./resources/osint-meme.png)

## Introduction:

Open-source intelligence (OSINT) is a key component of cybersecurity, providing valuable insights into potential vulnerabilities and security risks. This project involves creating a tool that performs comprehensive passive reconnaissance using publicly available data.

### Objective:

The goal is to build a multi-functional tool using a programming language of your choice, The tool is capable of retrieving detailed information based on user inputs such as Full names, IP addresses, usernames, and domains. This project will enhance your skills in data analysis, ethical considerations, and the use of various cybersecurity tools and APIs.

By completing this project, You will:

- Develop an understanding of OSINT techniques and their applications.
- Gain practical experience in programming, API integration, and data handling.
- Learn to identify basic security information gathering techniques.
- Understand the ethical and legal implications of cybersecurity practices.

### Resources

Some useful resources:

- [Open-source intelligence](https://en.wikipedia.org/wiki/Open-source_intelligence)
- [Doxing](https://en.wikipedia.org/wiki/Doxing)
- [Python Requests Library](https://docs.python-requests.org/)
- [Free IP API - ipapi.co](https://ipapi.co/)
- [OSINT Tools on GitHub](https://github.com/jivoi/awesome-osint)
- [Python String Methods](https://docs.python.org/3/library/stdtypes.html#string-methods)
- [HTTP Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
- [Python Socket Programming](https://docs.python.org/3/library/socket.html)
- [DNS Resolution Examples](https://docs.python.org/3/library/socket.html#socket.gethostbyname)
- [GitHub API Documentation](https://docs.github.com/en/rest)
- [Social Media URL Patterns](https://github.com/sherlock-project/sherlock)

### Hints

Before starting this project, you should know the following:

**Basic Programming Concepts:**

- Command-line argument parsing using `argparse` library
- HTTP requests and JSON data handling with `requests` library
- File I/O operations for saving results
- Error handling and exception management

**Networking Fundamentals:**

- Understanding of IP addresses and domain names
- Basic DNS resolution using Python's `socket` library
- HTTP status codes and their meanings

**Recommended Python Libraries:**

```bash
pip install requests argparse
```

**API Integration Tips:**

- Always check API rate limits and terms of service
- Handle API failures gracefully with try-catch blocks
- Use free tiers: ipapi.co offers 1000 free requests per month
- For social media checks, simple HTTP GET requests to profile URLs work

**Implementation Approach:**

- Start with a simple CLI structure using argparse
- Build one function at a time (name → IP → username → domain)
- Test each function individually before integration
- Use JSON format for structured data storage

Before asking for help, ask yourself if you have really thought about all the possibilities.

### Role play

To enhance the learning experience and assess your knowledge, a role play question session will be included as part of this project. This section will involve answering a series of questions in a simulated real-world scenario where you assume the role of a Cyber Security Expert explaining how to protect information from OSINT techniques to a team or stakeholder.

The goal of the role play question session is to:

- Assess your understanding of OSINT risks and basic mitigation strategies.
- Test your ability to communicate effectively and explain security measures related to this project.
- Challenge you to think critically about the importance of information security and consider alternative approaches.
- Explain basic information gathering techniques.
- Prepare for a role play question session in the audit.

### Project Requirements

#### Input Handling:

The tool should accept the following inputs: Full Name, IP Address, Username, and Domain.

#### Information Retrieval:

**Full Name**: Parse the input to extract "First Name" and "Last Name". Basic name validation and formatting. Add social media profile checking for name searches and implement basic information extraction from found profiles.

**IP Address**: Retrieve basic geolocation data (city, country) and ISP information using one free API such as ipapi.co.

**Username**: Check for the presence of the username on three social platforms using simple HTTP requests. Return only "Found" or "Not Found" status based on HTTP response codes.

**Domain**: Check if the domain exists and resolve its main IP address. Basic domain validation and DNS resolution.

**Documentation**: Create a README.md file that provides comprehensive documentation for your tool (prerequisites, setup, configuration, usage, ...). This file must be submitted as part of the solution for the project. Add clear guidelines and warnings about the ethical and legal use of the tool to your documentation.

**Documentation should include:**

- Installation instructions with pip requirements
- Usage examples with command-line demonstrations
- API configuration instructions
- Troubleshooting common issues
- Ethical use guidelines and legal disclaimers

You are responsible for choosing the way you want to find the data and for using free APIs only. Be aware of API Terms of Use!

#### Output Management:

Store the results in a well-organized file format.

**Implementation Hints:**

- Use JSON format for structured data storage
- Include timestamps in your output files
- Create separate functions for file operations
- Handle file writing errors gracefully

#### Usage Examples

##### Command Line Interface:

```
$> osintmaster --help

Welcome to osintmaster multi-function Tool

OPTIONS:
    -n  "Full Name"        Search information by full name
    -i  "IP Address"       Search information by IP address
    -u  "Username"         Search information by username
    -d  "Domain"           Check domain existence and resolve IP address
    -o  "FileName"         File name to save output
```

#### Example Outputs:

```
$> osintmaster -n "John Doe" -o result1.txt
First name: John
Last name: Doe
Social media check: Completed
LinkedIn: Found
Facebook: Not Found
Data Saved in result1.txt
```

**IP Address:**

```
$> osintmaster -i 8.8.8.8 -o result2.txt
IP: 8.8.8.8
ISP: Google LLC
City: Mountain View
Country: United States
Data Saved in result2.txt
```

**Username:**

```
$> osintmaster -u "@username" -o result3.txt
**GitHub: Found** - Public profile with repositories
**Twitter: Found** - Active account verified
**Instagram: Found** - Profile accessible
**LinkedIn: Found** - Professional profile exists
Data Saved in result3.txt
```

**Domain:**

```
$> osintmaster -d "example.com" -o result4.txt
Domain: example.com
IP Address: 93.184.216.34
Status: Active
Data saved in result4.txt
```

### Common Challenges & Solutions:

- **Network Timeouts:** Set reasonable timeout values in requests (5-10 seconds)
- **Invalid Domains:** Use try-catch blocks for DNS resolution and handle errors gracefully
- **API Failures:** Handle cases where APIs are unavailable with appropriate error messages
- **File Operations:** Ensure proper file handling with error checking

### Additional Resources:

- [Python Socket Programming](https://docs.python.org/3/library/socket.html)
- [Argparse Tutorial](https://docs.python.org/3/tutorial/stdlib.html#command-line-arguments)
- [Requests Quickstart](https://docs.python-requests.org/en/latest/user/quickstart/)
- [JSON Handling in Python](https://docs.python.org/3/library/json.html)

### Bonus

If you complete the mandatory part successfully, and you still have free time, you can implement anything that you feel deserves to be a bonus, for example:

**Advanced Username Analysis:**

- Add more social media platforms (Reddit, YouTube, LinkedIn)
- Retrieve basic profile information like repository counts or account status

**IP Address Enhancements:**

- Use multiple IP geolocation APIs for cross-validation
- Add reverse DNS lookup to find associated domain names
- Integrate basic IP reputation checking against threat databases

**Domain Analysis:**

- Implement subdomain discovery with common wordlists (www, mail, ftp, admin, test)
- Add HTTP header analysis and server technology detection
- Check for basic security misconfigurations

Challenge yourself!

### Documentation

Create a README.md file that provides comprehensive documentation for your tool (prerequisites, setup, configuration, usage, ...). This file must be submitted as part of the solution for the project. Add clear guidelines and warnings about the ethical and legal use of the tool to your documentation.

**Documentation should include:**

- Installation instructions with pip requirements
- Usage examples with screenshots
- API configuration instructions
- Troubleshooting common issues
- Ethical use guidelines

### Ethical and Legal Considerations

**Get Permission**: Always obtain explicit permission before gathering information.
**Respect Privacy**: Collect only necessary data and store it securely.
**Follow Laws**: Adhere to relevant laws such as GDPR and CFAA.
**Report Responsibly**: Use information gathering techniques ethically.
**Educational Use Only**: Use this tool and techniques solely for learning and improving security.

**Rate Limiting**: Respect API rate limits to avoid getting blocked.
**Terms of Service**: Always read and comply with API terms of service.
**Data Storage**: Don't store sensitive information unnecessarily.

⚠️ **Disclaimer**: This project is for educational purposes only. Ensure all activities comply with legal and ethical standards. The institution is not responsible for misuse of the techniques and tools demonstrated.

### Submission and audit

Upon completing this project, you should submit the following with proper repository structure:

**Repository Structure:**

```
osint-master/
├── README.md                 # Comprehensive documentation
├── requirements.txt          # Python dependencies
├── osintmaster.py           # Main application code
├── examples/                # Sample output files
│   ├── result1.txt
│   ├── result2.txt
│   ├── result3.txt
│   └── result4.txt
└── docs/                    # Additional documentation (optional)
    └── ethical-guidelines.md
```

**Required Files:**

- **README.md**: Complete documentation with installation, usage, and ethical guidelines
- **osintmaster.py**: Your main source code file
- **requirements.txt**: List of required Python packages
- **Sample outputs**: Example result files demonstrating functionality

**Documentation Requirements:**
Your README.md must include:

- Project description and objectives
- Installation instructions (`pip install -r requirements.txt`)
- Usage examples with command-line options
- Sample outputs for each search type
- API configuration instructions where applicable
- Troubleshooting section for common issues
- Comprehensive ethical use guidelines and legal disclaimers

### Ethical and Legal Considerations

**Get Permission**: Always obtain explicit permission before gathering information.
**Respect Privacy**: Collect only necessary data and store it securely.
**Follow Laws**: Adhere to relevant laws such as GDPR and CFAA.
**Report Responsibly**: Use information gathering techniques ethically.
**Educational Use Only**: Use this tool and techniques solely for learning and improving security.

**Rate Limiting**: Respect API rate limits to avoid getting blocked.
**Terms of Service**: Always read and comply with API terms of service.
**Data Storage**: Don't store sensitive information unnecessarily.

**Disclaimer**: This project is for educational purposes only. Ensure all activities comply with legal and ethical standards. The institution is not responsible for misuse of the techniques and tools demonstrated.
