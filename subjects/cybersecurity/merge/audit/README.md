#### General

##### Check the Repo Content

Files that must be present in the repository:

- `README.md` containing your explanation, walkthrough, and ethical considerations.
- The source code of your binder program.
- The test cases to demonstrate the merging of two simple programs.

###### Are all the required files present?

##### Play the Role of a Stakeholder

Conduct a simulated scenario where the student plays the role of a **Binary Analyst** presenting their findings to a team of stakeholders (auditors). Evaluate their understanding, communication skills, and depth of knowledge. Suggested questions include:

- Can you explain the structure of the executable file formats you worked with?
- How does your binder program work, and what are its potential applications?
- What are the ethical considerations and legal implications of binary modification?
- How would you recommend securing executable files against tampering?

###### Did the student demonstrate a thorough understanding of the project and concepts?

###### Did the student discuss the potential real-world applications and implications of their program?

###### Was the student able to justify the recommendations for securing binaries?

##### Review the Student Documentation

Verify that the `README.md` file contains:

- **Explanation of the Binder**: Provide a detailed explanation of how your program works.
- **Walkthrough**: Describe the step-by-step process of merging binaries.
- **Binary File Structure**: Include your analysis of the target binary file format.
- **Usage Instructions**: Provide clear instructions on how to use your program.
- **Ethical and Legal Report**: Discuss the ethical responsibilities and legal considerations of binary modification.

###### Does the `README.md` file include a complete and clear explanation of the binder?

###### Is the walkthrough detailed and easy to follow?

###### Does the documentation provide an analysis of the binary file structure?

###### Are ethical and legal considerations discussed comprehensively?

##### Test the Challenge

Verify that the student can:

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

1. **Run the Binder Program**:

- Ask the student to merge two binaries into one using their binder program.
- Verify that the program executes without errors and generates the merged binary.

2. **Execute the Merged Binary**:

- Confirm that the merged binary runs both programs seamlessly.
- Verify the execution flow preserves the functionality of both binaries.

3. **Analyze Binary File Structure**:

- Ask the student to explain the binary file formats they worked with and how the binder modifies the file headers and entry points.

###### Did the student successfully run the binder program and generate the merged binary?

###### Does the merged binary execute both programs as expected?

###### Was the student able to analyze and explain the structure of the binary file formats?

##### Bonus

###### + Did the student extend the binder to support multiple executable formats (e.g., ELF, PE, Mach-O)?

###### + Did the student implement advanced execution flow control (e.g., user-defined order or conditions for execution)?

###### + Did the student create a graphical interface for the binder program?

###### + Is this project an outstanding project that exceeds the basic requirements?
