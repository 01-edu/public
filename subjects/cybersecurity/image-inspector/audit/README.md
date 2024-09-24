#### General

##### Check the Repo Content

Files that must be inside the repository:

- Detailed documentation in the `README.md` file.
- Source code for the Image Inspector tool.
- Any required configuration files and scripts for running the tool.

###### Are all the required files present?

##### Play the Role of a Stakeholder

Organize a simulated scenario where the student takes on the role of a Digital Forensics Expert and explains their solution and knowledge to a team or stakeholder. Evaluate their grasp of the concepts and technologies used in the project, their communication efficacy, and their critical thinking about their solution and the knowledge behind this project.

Suggested role play questions include:

- What is metadata in the context of digital images, and why is it important?
- How does steganography work, and what are its potential uses and risks?
- What challenges did you face while developing the Image Inspector tool, and how did you address them?
- How can this tool be used in real-life digital forensics or cybersecurity scenarios?
- What ethical considerations should be taken into account when analyzing images for hidden data?

###### Were the students able to answer all the questions?

###### Did the students demonstrate a thorough understanding of the concepts and technologies used in the project?

###### Were the students able to communicate effectively and justify their decisions and explain the knowledge behind this project?

###### Were the students able to evaluate the value of this project in real-life scenarios?

###### Did the students demonstrate an understanding of ethical and legal considerations related to digital forensics?

##### Check the Student Documentation in the `README.md` File

###### Does the `README.md` file contain all the necessary information about the tool (prerequisites, setup, configuration, usage, ...)?

###### Does the `README.md` file contain clear guidelines and warnings about the ethical and legal use of the tool?

##### Review the Tool's Design and Implementation

1. **Help Command:**

```sh
$> image-inspector --help
```

###### Does the output include an explanation of how to use the tool, with all options clearly described?

2. **Metadata Extraction Option:**

```sh
$> image-inspector -m -o metadata.txt image-example1.jpeg
```

###### Does the output correctly extract and display metadata such as geolocation, device information, and date/time?

###### Is the output stored in the file specified in the output parameter?

3. **Steganography Detection Option:**

```sh
$> image-inspector -s -o hidden_data.txt image-example1.jpeg
```

###### Does the output correctly detect and extract any hidden PGP keys or other concealed information within the image?

###### Is the output stored in the file specified in the output parameter?

##### Testing with Images

**You will be provided with an example image to test the students tool. Feel free to test with other images.**
The example image attached: 
[image-example1.jpeg](resources/image-example1.jpeg)
[image-example2.jpeg](resources/image-example2.jpeg)
[image-example3.jpeg](resources/image-example3.jpeg)
[image-example4.jpeg](resources/image-example4.jpeg)

###### Test the tool with the provided example image and at least one other image to ensure the tool's robustness.

###### Does the tool produce accurate and expected results for different images?

##### Ensure that the student submission meets the project requirements:

1. **Functionality:** Does the tool perform its intended functions accurately (metadata extraction and steganography detection)?
2. **Data Accuracy:** Is the retrieved information accurate and relevant?
3. **Ethical Considerations:** Are there clear guidelines and warnings about the ethical and legal use of the tool?
4. **Usability:** Is the tool user-friendly and well-documented?

###### Did the tool design and implementation align with all the project requirements above?

###### Were the students able to implement a functional and reliable tool that meets the project requirements?

#### Bonus

###### + Did the student implement additional valuable features (e.g., additional steganography methods, GUI, facial recognition)?

###### + Is this project an outstanding project that exceeds the basic requirements?
