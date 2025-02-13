## Merge

<center>
<img src="pictures/merge-meme.jpg" style="width: 414px; height: auto;">
</center>

### Introduction

Executable files play a vital role in software development and system operations. By analyzing binary headers and modifying executable behavior, we can gain a deeper understanding of how operating systems handle executables. In this project, you will develop a **binder**, a tool that merges two executable programs into one, providing insight into executable file structures and entry points.

### Objective

The goal of this project is to develop a binary binder that merges two executable programs into a single executable. This will help you:

- Understand executable file formats and their structures (e.g., ELF, PE, Mach-O).
- Learn about entry points and program execution flow.
- Gain experience in low-level programming and file manipulation.
- Enhance debugging and reverse engineering skills.

### Role Play

As part of the project, you will participate in a role-play session where you will act as a **Binary Analyst** presenting your findings to a hypothetical team of stakeholders. Be prepared to:

- Explain the structure of the executable file formats you used.
- Demonstrate how your binder program works and its potential applications.
- Discuss the ethical considerations and legal implications of binary modification.
- Provide recommendations for securing executable files against tampering.

### Project Requirements

#### Setup and Environment

Choose a Linux operating system. Ensure you have the necessary tools and libraries for binary analysis and manipulation.

#### The Challenge

Using a programming language from your choice create a program that merges two executable files into one.
Ensure that the output executable runs both programs seamlessly and manages the execution flow to preserve the functionality of both programs.

##### Usage Examples

```sh
$> ./bin1
Message from bin1

$> ./bin2
Message from bin2

$> ./merge
Welcome to the merge program.
Usage: ./merge source-binary1 source-binary2 -o output-binary

$> ./merge bin1 bin2 -o bin3
bin1 and bin2 merged into bin3 successfully!

$> ./bin3
Message from bin1
Message from bin2
```

### Documentation

Create a `README.md` file that includes:

- **Explanation of the Binder**: Provide a detailed explanation of how your program works.
- **Walkthrough**: Describe the step-by-step process of merging binaries.
- **Binary File Structure**: Include your analysis of the target binary file format.
- **Usage Instructions**: Provide clear instructions on how to use your program.
- **Ethical and Legal Report**: Discuss the ethical responsibilities and legal considerations of binary modification.

### Bonus

If you complete the mandatory part successfully, and you still have free time, you can implement anything that you feel deserves to be a bonus, for example:

- **Support for Multiple Formats**: Extend your binder to support multiple executable formats (e.g., ELF, PE, Mach-O).
- **Advanced Execution Flow**: Let the user define the order and conditions for executing the merged programs.
- **Graphical Interface**: Develop a simple GUI for merging executables.

Challenge yourself!

### Ethical and Legal Considerations

You are responsible for following ethical guidelines when analyzing and modifying binaries. Do not use these techniques on unauthorized systems or software. Misuse of these techniques is strictly prohibited and may violate local laws.

> ⚠️ Disclaimer: This project is for educational purposes only. Unauthorized use of these techniques is prohibited and may be illegal.

### Submission and Audit

Submit the following:

- The source code of your binder program.
- `README.md` containing your explanation, walkthrough, and ethical considerations.
- The test cases to demonstrate the merging of two simple programs.

Ensure the necessary tools and environment are installed for the audit.

### Resources

Some useful resources:

- [File Binder](https://en.wikipedia.org/wiki/File_binder): Introduction to file binders.
- [Executable File Formats](https://en.wikipedia.org/wiki/Executable_and_Linkable_Format): Learn about ELF, PE, and Mach-O file formats.
- [Entry Point](https://en.wikipedia.org/wiki/Entry_point): Understand how programs start execution.
- [Hex Editors and Debuggers](https://hexed.it/): A tool for analyzing and editing binary files.
