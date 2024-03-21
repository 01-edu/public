## Git Ready

### Introduction

This project is designed to introduce you to the world of version control and collaboration using **Git**. Git is a powerful and widely used tool for tracking changes in your projects, collaborating with others, and ensuring the integrity of your code.

Throughout this project, you will embark on a journey of progressively building your Git skills. Starting from the basics, you'll gradually explore more advanced topics, equipping yourself with the essential knowledge and practices for effective version control and collaboration.

Let's Git ready for it!

### Instructions

To begin, create a `work` directory and organize all your tasks within it. Each exercise should be encapsulated in its own file, named after the corresponding task for clarity and ease of reference.

Accompanying your work, provide documentation or a report detailing the process followed for each exercise. This documentation should include any challenges faced, solutions implemented, and lessons learned. It could be in the form of a README file or a separate document. Make sure to show it to the auditor during evaluation.

> ⚠️ Your completion of tasks will be evaluated based on the commit history reflecting the changes made throughout the exercises and the presence of accompanying documentation detailing the process followed.

Here is an example of a file that you can deliver to your auditor to help with the review process:

```md
#### Conflicts, merging and rebasing

# Merge Main into Greet Branch

<Write the command here>

# Switch to main branch and make changes to hello.sh file

<Write the command here>

# Merging Main into Greet Branch (Conflict)

<Write the command here>

# Resolve the conflict (manually or using merge tools)

<Write the command here>

# After resolving, stage the changes and commit

<Write the command here>

# Rebasing Greet Branch

<Write the command here>

# Merging Greet into Main

<Write the command here>
```

#### Setting Up Git

- Install Git on your local machine by following the instructions for your operating system on the official [Git website](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).
- Configure Git with your username and email address.

#### Git commits to commit

- Within the `work` directory, establish a subdirectory named `hello`. Inside this directory, generate a file titled `hello.sh` and input the following content:

```sh
echo "Hello, World"
```

- Initialize the git repository in the `hello` directory.
- Check the status and act accordingly with the output of the executed command.
- Change the `hello.sh` content to the following:

```sh
#!/bin/bash

echo "Hello, $1"
```

- Stage the changed file and commit the changes, the working tree should be clean.
- Modify the `hello.sh` file to include comments and stage it.

```sh
#!/bin/bash

# Default is "World"
name=${1:-"World"}
echo "Hello, $name"
```

- Make two separate commits:
  - The first commit should be for the comment in line 1.
  - The second commit should include changes made to lines 3 and 4.

#### History

- Show the history of the working directory.
- Show One-Line History for a condensed view showing only commit hashes and messages.
- **Controlled Entries**:
  - You need to customize the log output by specifying the number of entries or a time range. Customize it to display the last `2 entries` and to view the `commits made within the last 5 minutes`.
- **Personalized Format**:
  - Show logs in a personalized format, including the commit hash, date, message, branch information, and author name, resembling `* e4e3645 2023-06-10 | Added a comment (HEAD -> main) [John Doe]`

#### Check it out

- **Restore First Snapshot**:
  - Revert the working tree to its initial state, as captured in the first snapshot, and then print the content of the `hello.sh` file.
- **Restore Second Recent Snapshot**:
  - Revert the working tree to the second most recent snapshot and print the content of the `hello.sh` file.
- **Return to Latest Version**:
  - Ensure that the working directory reflects the latest version of the `hello.sh` file present in the main branch, without referring to specific commit hashes.

#### TAG me

- **Referencing Current Version**:
  - Tag the current version of the repository as `v1`.
- **Tagging Previous Version**:
  - Tag the version immediately prior to the current version as `v1-beta`, without relying on commit hashes to navigate through the history.
- **Navigating Tagged Versions**:
  - Move back and forth between the two tagged versions, `v1` and `v1-beta`.
- **Listing Tags**:
  - Display a list of all tags present in the repository to verify successful tagging.

#### Changed your mind?

- **Reverting Changes**:
  - Modify the latest version of the file with unwanted comments, then revert it back to its original state before staging using a `Git` command.

```sh
#!/bin/bash

# This is a bad comment. We want to revert it.
name=${1:-"World"}

echo "Hello, $name"
```

- **Staging and Cleaning**:
  - Introduce unwanted changes to the file, stage them, then clean the staging area to discard the changes.

```sh
#!/bin/bash

# This is an unwanted but staged comment
name=${1:-"World"}

echo "Hello, $name"
```

- **Committing and Reverting**:
  - Add the following unwanted changes again, stage the file, commit the changes, then revert them back to their original state.

```sh
#!/bin/bash

# This is an unwanted but committed change
name=${1:-"World"}

echo "Hello, $name"
```

- **Tagging and Removing Commits**:
  - Tag the latest commit with `oops`, then remove commits made after the `v1` version. Ensure that the `HEAD` points to `v1`.
- **Displaying Logs with Deleted Commits**:
  - Show the logs with the deleted commits displayed, particularly focusing on the commit tagged `oops`.
- **Cleaning Unreferenced Commits**:
  - Ensure that unreferenced commits are deleted from the history, meaning there should be no logs for these deleted commits.
- **Author Information**:
  - Add an author comment to the file and commit the changes.

```sh
#!/bin/bash

# Default is World
# Author: Jim Weirich
name=${1:-"World"}

echo "Hello, $name"
```

- Oops the author email was forgotten, update the file to include the email without making a new commit, but include the change in the last commit.

#### Move it

- **Moving hello.sh**:
  - Using Git commands, move the program `hello.sh` into a `lib/` directory, and then commit the move.
  - Create a `Makefile` in the root directory of the repository with the provided content and commit it to the repository.

```sh
TARGET="lib/hello.sh"

run:
	bash ${TARGET}
```

#### blobs, trees and commits

- **Exploring `.git/` Directory**:
  - Navigate to the `.git/` directory in your project and examine its contents.You will have to explain the purpose of each subdirectory, including `objects/`, `config`, `refs`, and `HEAD` in the audit.
- **Latest Object Hash**:
  - Find the latest object hash within the `.git/objects/` directory using Git commands and print the type and content of this object using Git commands.
- **Dumping Directory Tree**:
  - Use Git commands to dump the directory tree referenced by this commit.
  - Dump the contents of the `lib/` directory and the `hello.sh` file using Git commands.

#### Branching

It’s time to do a major rewrite of the hello world functionality. Since this might take a while, you’ll want to put these changes into a separate branch to isolate them from changes in the main branch.

- **Create and Switch to New Branch**:
  - Create a local branch named `greet` and switch to it.
  - In the `lib` directory, create a new file named `greeter.sh` and add the provided code to it. Commit these changes.

```sh
#!/bin/bash

Greeter() {
    who="$1"
    echo "Hello, $who"
}

```

- Update the `lib/hello.sh` file by adding the content below, stage and commit the changes.

```sh
#!/bin/bash

source lib/greeter.sh

name="$1"
if [ -z "$name" ]; then
    name="World"
fi

Greeter "$name"
```

- Update the `Makefile` with the following comment and commit the changes.

```sh
# Ensure it runs the updated lib/hello.sh file
TARGET="lib/hello.sh"

run:
	bash ${TARGET}
```

- Switch back to the `main` branch, compare and show the differences between the `main` and `greet` branches for `Makefile`, `hello.sh`, and `greeter.sh` files.
- Generate a `README.md` file for the project with the provided content. Commit this file.

```console
This is the Hello World example from the git project.
```

- Draw a commit tree diagram illustrating the diverging changes between all branches to demonstrate the branch history.

#### Conflicts, merging and rebasing

- **Merge Main into Greet Branch**:
  - Start by merging the changes from the `main` branch into the `greet` branch.
  - Switch to `main` branch and make the changes below to the `hello.sh` file, save and commit the changes.

```sh
#!/bin/bash

echo "What's your name"
read my_name

echo "Hello, $my_name"
```

- **Merging Main into Greet Branch (Conflict)**:
  - Attempt to merge the `main` branch into `greet`. Bingooo! There you have it, a `conflict`.
  - Resolve the conflict (manually or using graphical merge tools), accept changes from `main` branch, then commit the conflict resolution.
- **Rebasing Greet Branch**:
  - Go back to the point before the initial merge between `main` and `greet`.
  - Rebase the `greet` branch on top of the latest changes in the `main` branch.
- **Merging Greet into Main**:
  - Merge the changes from the `greet` branch into the `main` branch.
- **Understanding Fast-Forwarding and Differences**:
  - Explain fast-forwarding and the difference between merging and rebasing.

#### Local and remote repositories

- In the `work/` directory, make a clone of the repository `hello` as `cloned_hello`. (Do not use `copy` command)

- Show the logs for the cloned repository.

- Display the name of the remote repository and provide more information about it.

- List all remote and local branches in the `cloned_hello` repository.

- Make changes to the original repository, update the `README.md` file with the provided content, and commit the changes.

```
This is the Hello World example from the git project.
(changed in the original)
```

- Inside the cloned repository (`cloned_hello`), fetch the changes from the remote repository and display the logs. Ensure that commits from the `hello` repository are included in the logs.

- Merge the changes from the remote `main` branch into the local `main` branch.

- Add a local branch named `greet` tracking the remote `origin/greet` branch.

- Add a `remote` to your Git repository and push the `main` and `greet` branches to the `remote`.

- Be ready for this question in the audit!

**"What is the single git command equivalent to what you did before to bring changes from remote to local `main` branch?"**

#### Bare repositories

- What is a bare repository and why is it needed?
- Create a bare repository named `hello.git` from the existing `hello` repository.
- Add the bare `hello.git` repository as a remote to the original repository `hello`.
- Change the `README.md` file in the original repository with the provided content:

```
This is the Hello World example from the git project.
(Changed in the original and pushed to shared)
```

- Commit the changes and push them to the shared repository.

- Switch to the cloned repository `cloned_hello` and pull down the changes just pushed to the shared repository.

### Submission and Evaluation

Your work must be submitted at the `gitea` link provided. The evaluation will be carried out based on your submission and according to the following criteria:

- Correctness of the git commands you are using.
- Clear understanding of the git commands and concepts.

### Notions

- [Git Branching](https://learngitbranching.js.org/)
- [Git Gud](https://github.com/benthayer/git-gud)
