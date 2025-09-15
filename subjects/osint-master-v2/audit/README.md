#### General

##### Check the Repo content.

Files that must be inside the repository:

- Comprehensive documentation in the `README.md` file with all required sections.
- Source code for the OSINT-Master tool (`osintmaster.py`).
- Dependencies list (`requirements.txt`).
- Sample output files demonstrating functionality.

###### Are all the required files present and properly structured?

##### Play the role of a stakeholder

Organize a simulated scenario where the student take on the role of Cyber Security Experts and explain their solution and knowledge to a team or stakeholder. Evaluate their grasp of the concepts and technologies used in the project, their communication efficacy, and their critical thinking about their solution and knowledge behind this project.

Suggested role play questions include:

- What is OSINT and why is it significant in cybersecurity?
- What types of information can be gathered using OSINT techniques?
- How does the OSINT-Master tool help identify publicly available information?
- What challenges did you face while developing the OSINT-Master tool and how did you address them?
- How can we protect our critical information from OSINT techniques?
- How can this tool help in a defensive approach?
- What are the ethical and legal considerations when using OSINT tools?
- How does your tool handle API failures and network errors?

###### Were the student able to answer all the questions?

###### Did the student demonstrate a thorough understanding of the concepts and technologies used in the project?

###### Were the students able to communicate effectively and justify their decisions and explain the knowledge behind this project?

###### Was the student able to evaluate the value of this project in real-life scenarios?

###### Did the students demonstrate an understanding of ethical and legal considerations related to OSINT?

##### Check the Student Documentation in the `README.md` File

###### Does the `README.md` file contain all the necessary information about the tool (prerequisites, setup, configuration, usage, ...)?

###### Does the `README.md` file contain installation instructions with pip requirements?

###### Does the `README.md` file contain usage examples with command-line demonstrations?

###### Does the `README.md` file contain API configuration instructions?

###### Does the `README.md` file contain troubleshooting section for common issues?

###### Does the `README.md` file contain clear guidelines and warnings about the ethical and legal use of the tool?

##### Review the Tool's Design and Implementation

1. **Help Command:**

```sh
$> python osintmaster.py --help
```

###### Does the output include clear explanation of how to use the tool with all available options?

2. **Full Name Option:**

```sh
$> python osintmaster.py -n "John Doe" -o result1.txt
```

###### Does the output correctly parse the name into first and last name components?

###### Does the output include social media profile checking results?

###### Does the output get stored to the file specified in the output parameter?

3. **IP Address Option:**

```sh
$> python osintmaster.py -i "8.8.8.8" -o result2.txt
```

###### Does the output include basic geolocation data (city, country) and ISP information?

###### Does the output get stored to the file specified in the output parameter?

4. **Username Option:**

```sh
$> python osintmaster.py -u "testuser" -o result3.txt
```

###### Does the output check the presence of the username on at least three social platforms?

###### Does the output return "Found" or "Not Found" status for each platform?

###### Does the output get stored to the file specified in the output parameter?

5. **Domain Option:**

```sh
$> python osintmaster.py -d "example.com" -o result4.txt
```

###### Does the output check domain existence and resolve the main IP address?

###### Does the output perform basic domain validation?

###### Does the output get stored to the file specified in the output parameter?

##### Ensure that the student submission meets the project requirements:

1. **Functionality:** Does the tool retrieve information based on the given inputs (Full Name, IP Address, Username, and Domain) as specified in the mandatory requirements?
2. **Data Accuracy:** Is the retrieved information accurate and relevant to the search type?
3. **Error Handling:** Does the tool handle network timeouts, invalid domains, and API failures gracefully?
4. **File Output:** Are results properly saved in JSON format with timestamps?
5. **Ethical Considerations:** Are there clear guidelines and warnings about the ethical and legal use of the tool?
6. **Usability:** Is the tool user-friendly with proper command-line interface and well-documented?
7. **Dependencies:** Are all required libraries properly listed in requirements.txt?

###### Did the tool design and implementation align with all the mandatory project requirements?

###### Were the students able to implement a functional and reliable tool that meets the simplified project scope?

###### Does the tool demonstrate proper understanding of API integration and basic OSINT concepts?

#### Bonus

###### +Did the student implement additional valuable features such as:

- Enhanced social media analysis beyond basic checking?
- Multiple API sources for IP geolocation cross-validation?
- Advanced username analysis with profile information extraction?
- Subdomain discovery functionality with common wordlists?
- Server technology detection and HTTP header analysis?

###### +Is this project an outstanding project that exceeds the basic requirements while maintaining ethical standards?

###### +Did the student demonstrate creativity and technical depth in their bonus implementations?
