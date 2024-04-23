#### General

##### Check the Repo content:

- A `README.md` file, Which explains all the steps to bypass all exercises.
- All used tools and scripts.

###### Are all the required files present?

##### Set up the virtual machine:

1. Download the virtual machine image [hole-in-bin.ova](https://assets.01-edu.org/cybersecurity/hole-in-bin/hole-in-bin.ova).
   For machine using Apple Silicon or equivalent get [hole-in-bin.utm.zip](https://assets.01-edu.org/cybersecurity/hole-in-bin/hole-in-bin.utm.zip).
   This image contains all the binaries you will need for the audit.

SHA1 for `hole-in-bin.ova`: 7db09b7a8fdfe25c286561dfa7ca5b50718bd60c
SHA1 for `hole-in-bin.utm.zip`: fc93533b2054d10d03b09d53c223e57bf7ac7b62

> If it's already downloaded in the student machine, please check the SHA1 running the following command

```console
$ sha1sum <filename>
<SHA1>
```

2. Load the virtual machine image into your virtualization software of choice (e.g., VirtualBox, VMWare).

3. Login using the provided credentials (username: user, password: user).

##### Ask the student to disassemble and explain the binaries:

> Using a decompiler is forbidden, use a disassembler instead of it!

- The compiler is used to convert high-level programming language code into machine language code.
- The assembler converts assembly-level language code into machine language code.

###### Was the student capable to disassemble the binaries?

###### Was the student capable to explain the functionality of all the binaries?

###### Has the student shown the ability to understand and analyze binary structures and operations?

###### Did the student showcase an understanding of reverse engineering concepts?

##### Ask the student to exploit the binaries:

> It's forbidden to use external scripts!

###### Have all binaries been exploited successfully?

###### Did the student demonstrate an understanding of various binary exploitation techniques?

##### Check the student Documentation:

###### Is the documentation clear and complete, including well-structured explanations and thorough descriptions?

###### Did the student explain their thought process and approach to each challenge?

###### Have the student’s notes clearly described the tools and techniques used during the exercise?
