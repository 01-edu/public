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

###### Does the output correctly detect and extract any hidden PGP keys within the image?

###### Is the output stored in the file specified in the output parameter?

##### Testing with Images

**You will be provided with an example image to test the students tool. Feel free to test with other images.**
The example image attached:
[image-example1.jpeg](resources/image-example1.jpeg)
[image-example2.jpeg](resources/image-example2.jpeg)
[image-example3.jpeg](resources/image-example3.jpeg)
[image-example4.jpeg](resources/image-example4.jpeg)

###### Test the tool with the provided example image and at least one other image to ensure the tool's robustness.

```sh
$> image-inspector -s -o hidden_data1.txt image-example1.jpeg
Enter   -----BEGIN PGP PUBLIC KEY BLOCK-----
Version: 01

mQENBGIwpy4BC<...>
=N8hc
-----END PGP PUBLIC KEY BLOCK-----
$>
```

```sh
$> image-inspector -s -o hidden_data2.txt image-example2.jpeg
-----BEGIN PGP PUBLIC KEY BLOCK-----
xo0EZuV9/AEEANRklh3<...>
=BhsK
-----END PGP PUBLIC KEY BLOCK-----
$>
```

```sh
$> image-inspector -s -o hidden_data3.txt image-example3.jpeg
-----BEGIN PGP PUBLIC KEY BLOCK-----
xo0EZuWBWQEEALrj<...>
=SXXm
-----END PGP PUBLIC KEY BLOCK-----
$>
```

```sh
$> image-inspector -s -o hidden_data4.txt image-example4.jpeg
-----BEGIN PGP PUBLIC KEY BLOCK-----
xo0EZuWA9AE<...>
=/1dE
-----END PGP PUBLIC KEY BLOCK-----
$>
```

> The results are cut. you can compare the beginning and the end of the results!

###### Does the tool produce accurate and expected results for different images?

#### Bonus

###### + Did the student implement additional valuable features (e.g., additional steganography methods, GUI, facial recognition)?

###### + Is this project an outstanding project that exceeds the basic requirements?
