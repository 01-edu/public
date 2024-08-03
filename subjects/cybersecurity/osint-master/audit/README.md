#### General

##### Check the Repo content.
Files that must be inside the repository:
- Detailed documentation in the `README.md` file.
- Source code for the OSINT-Master tool.
- Any required configuration files and scripts for running the tool.
###### Are all the required files present?

##### Play the role of a stakeholder
Organize a simulated scenario where the student take on the role of Cyber Security Experts and explain their solution and knowledge to a team or stakeholder. Evaluate their grasp of the concepts and technologies used in the project, their communication efficacy, and their critical thinking about their solution and knowledge behind this project.
Suggested role play questions include:

- What is OSINT and why is it significant in cybersecurity?
- What types of information can be gathered using OSINT techniques?
- Explain what subdomain takeovers are, and how to protect against it?
- How does the OSINT-Master tool help in identifying sensitive information?
- What challenges did you face while developing the OSINT-Master tool and how did you address them?
- How we can protect our cretical information from OSINT techniques?
- How can this tool help in a defensive approach? 
###### Were the student able to answer all the questions?
###### Did the student demonstrate a thorough understanding of the concepts and technologies used in the project?
###### Were the student able to communicate effectively and justify their decisions and explain the knowledge behind this project?
###### Did the student able to evalute the value of this project in the real life scenarios?
###### Did the students demonstrate an understanding of ethical and legal considerations related to OSINT?

##### Check the Student Documentation in the `README.md` File
###### Does the `README.md` file contain all the necessary information about the tool (prerequisites, setup, configuration, usage, ...)?
###### Does the `README.md` file contain clear guidelines and warnings about the ethical and legal use of the tool?

##### Review the Tool's Design and Implementation
1. **Help Command:** 
```sh
$> osintmaster --help
```
###### Does the output include explanation how to use the tool?

2. **Full Name Option:**
```sh
$> osintmaster -n "Full Name" -o filename
```
###### Does the output include accurate details such as phone numbers, addresses, and social media profiles?
###### Does the output stored to the file specified in the output parameter?

3. **IP Adress Option:**
```sh
$> osintmaster -i "IP Address" -o filename
```
###### Does the output include geolocation data, ISP details, and historical data?
###### Does the output stored to the file specified in the output parameter?

4. **Username Option:**
```sh
$> osintmaster -u "Username" -o filename
```
###### Does the output check the presence of the username on multiple social networks and public repositories?
###### Does the output stored to the file specified in the output parameter?

5. **Domain Option:**
```sh
$> osintmaster -d "Domain" -o filename
```
###### Does the output enumerate subdomains, gather relevant information, and identify potential subdomain takeover risks?
###### Does the output stored to the file specified in the output parameter?

##### Ensure that the student submission meets the project requirements:
1. **Functionality:** Does the tool retrieve detailed information based on the given inputs (Full Name, IP Address, Username, and Domain)?
2. **Data Accuracy:** Is the retrieved information accurate and relevant?
3. **Ethical Considerations:** Are there clear guidelines and warnings about the ethical and legal use of the tool?
4. **Usability:** Is the tool user-friendly and well-documented?
###### Did the tool design and implementation align with all the project requirements above?
###### Were the students able to implement a functional and reliable tool that meets the project requirements?

#### Bonus
###### + Did the student implement additional valuable features?
###### + Is this project an outstanding project that exceeds the basic requirements?