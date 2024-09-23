## Image-Inspector

<center>
<img src="./resources/steganography-meme.png?raw=true" style="width: 600px !important; height: 516px !important;"/>
</center>

### Introduction

Images can contain more than just visual information, they often carry hidden data such as metadata or even concealed messages through techniques like steganography. This project is designed to help you explore and analyze these hidden aspects of images.

### Objective

The goal is to develop a tool using a programming language of your choice (Python is recommended) that can analyze images to extract hidden information. Specifically, your tool should:

1. **Extract Metadata**: Identify and display metadata from images, such as geolocation (latitude and longitude) where the photo was taken, the device used, and other relevant information.
2. **Detect Steganography**: Discover and extract any hidden PGP keys or other data concealed within the image using steganography techniques.

By completing this project, you will:

- Develop an understanding of image recognition techniques.
- Gain practical experience with steganography and metadata extraction.
- Learn to identify hidden information in images, which can be crucial in cybersecurity investigations.

### Resources

Some useful resources:

- [Steganography](https://en.wikipedia.org/wiki/Steganography)
- [Exif Metadata](https://en.wikipedia.org/wiki/Exif)
- [Python Imaging Library (PIL)](https://pillow.readthedocs.io/en/stable/)

Before asking for help, explore these resources to fully understand the concepts behind image analysis and steganography.

### Role Play

To enhance your learning experience and assess your knowledge, a role play question session will be included as part of this project. This session will involve answering a series of questions in a simulated real-world scenario where you assume the role of a digital forensics expert explaining how your tool identifies hidden information in images.

The goal of the role play question session is to:

- Assess your understanding of image analysis and steganography.
- Test your ability to communicate effectively and explain the techniques used in your project.
- Challenge you to think critically about the implications of hidden data in images.

Prepare for a role play question session in the audit.

### Project Requirements

#### Metadata Extraction:

Your tool should be able to extract and display key metadata from an image file, including but not limited to:

- Geolocation (Latitude and Longitude)
- Device information (camera make/model)
- Date and time the photo was taken

#### Steganography Detection:

- Detect hidden PGP keys or other concealed information in the image using steganography techniques.
- Display the extracted hidden data clearly.

#### Input Handling:

The tool should accept an image file as input and provide options for the user to specify which analysis to perform (metadata extraction or steganography detection).

#### Output Management:

Store the results in a well-organized file format, with clear and concise reporting of both metadata and any hidden data found.

### Usage Examples

#### Command Line Interface:

```sh
$> image-inspector --help

Welcome to Image Inspector

OPTIONS:
    -m  Metadata          Extract metadata from the image (e.g., geolocation, device info)
    -s  Steganography     Detect and extract hidden data from the image using steganography techniques
    -o  "FileName"        Specify the file name to save output
```

#### Example Outputs:

```sh
$> image-inspector -m -o metadata.txt image-example1.jpeg
Lat/Lon: (13.731) / (-1.1373)
Device: Canon EOS 5D Mark III
Date: 2023-07-20 14:32:10
Data saved in metadata.txt
```

```sh
$> image-inspector -s -o hidden_data.txt image-example1.jpeg
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: 01
...
-----END PGP PUBLIC KEY BLOCK-----
Data saved in hidden_data.txt
```

**You will be provided with an example image to test your tool. Make sure your tool can successfully analyze this image and produce the expected outputs.**
The example image attached: 
[image-example1.jpeg](resources/image-example1.jpeg)
[image-example2.jpeg](resources/image-example2.jpeg)
[image-example3.jpeg](resources/image-example3.jpeg)
[image-example4.jpeg](resources/image-example4.jpeg)

#### Documentation

Create a `README.md` file that provides comprehensive documentation for your tool (prerequisites, setup, configuration, usage, ...). Ensure the documentation includes clear guidelines on the ethical use of the tool and warnings about the legal implications of analyzing images without permission.

### Bonus

If you complete the mandatory part successfully and still have time, consider adding the following features:

- **Additional Steganography Methods**: Implement detection for other steganography techniques beyond PGP keys.
- **Graphical User Interface (GUI)**: Create a user-friendly GUI using libraries like Tkinter or PyQT.

Challenge yourself!

### Ethical and Legal Considerations

- **Get Permission**: Always obtain explicit permission before analyzing any image.
- **Respect Privacy**: Be aware of the sensitive nature of metadata and hidden data, and handle it responsibly.
- **Follow Laws**: Adhere to relevant laws regarding data privacy and the analysis of digital media.

> ⚠️ Disclaimer: This project is for educational purposes only. Ensure all activities comply with legal and ethical standards. The institution is not responsible for misuse of the techniques and tools demonstrated.

### Submission and Audit

Upon completing this project, you should submit the following:

- Your documentation in the `README.md` file.
- The source code for your tool.
- Any required files to run your tool.
