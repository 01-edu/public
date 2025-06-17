#### Git

> ⚠️ The student must provide you with a file containing the solutions for each task. Furthermore, they should showcase their commit history on GitHub, facilitating your review of the evolution of their work and the strategies employed to complete each task. This commit history is crucial to the evaluation process. Please ensure that the submission includes both the solution file and the link to the GitHub repository containing the commit history. In the absence of the link, kindly request the student to provide it.

#### Setup and Installation

###### Did the student successfully install Git on their local machine?

###### Did the student configure Git with a valid username and email address?

#### Git commits to commit

###### Did the student navigate to the `work` directory and create a subdirectory named `hello`?

###### Did the student generate a file named `hello.sh` with the content `echo "Hello, World"` inside the `hello` directory?

###### Did the student initialize a Git repository in the `hello` directory?

###### Did the student use the `git status` command to check the status of the repository?

###### Did the student modify the `hello.sh` file content with the provided `echo "Hello, $1"`?

###### Did the student stage the modified `hello.sh` file, commit the changes to the repository, and ensure that the working tree is clean afterward?

###### Did the student further modify the `hello.sh` file to include comments, and then make two separate commits as instructed?

###### Did the student make two separate commits, with the first commit for the comment in line 1 and the second commit for the changes made to lines 3 and 4, as instructed?

#### History

###### Did the student display the Git history of the working directory with the `git log` command?

###### Did the student successfully display a condensed view of the Git history, showing only commit hashes and messages using the "One-Line History" format?

###### Was the student able to customize the log output to display the last 2 entries?

###### Did the student successfully demonstrate viewing commits made within the last 5 minutes?

###### Did the student successfully customize the format of Git logs and display them according to this example `* e4e3645 2023-06-10 | Added a comment (HEAD -> main) [John Doe]`?

#### Check it out

###### Did the student successfully restore the first snapshot of the working tree and print the content of `hello.sh`?

###### Did the student successfully restore the second recent snapshot and print the content of `hello.sh`?

###### Did the student ensure that the working directory reflects the latest version of `hello.sh` from the main branch without using commit hashes?

#### TAG me

###### Did the student successfully tag the current version of the repository as `v1`?

###### Did the student successfully tag the version immediately prior to the current version as `v1-beta`, without relying on commit hashes?

###### Did the student navigate back and forth between the two tagged versions, `v1` and `v1-beta`?

###### Did the student display a list of all tags present in the repository to verify successful tagging?

#### Changed your mind?

###### Did the student successfully revert the modifications made to the latest version of the file, restoring it to its original state before staging using a `Git` command?

###### Did the student introduce unwanted changes to the file, stage them, and then successfully clean the staging area to discard the changes?

###### Did the student add unwanted changes again, stage the file, commit the changes, and then revert them back to their original state?

###### Did the student tag the latest commit with oops and remove commits made after the v1 version, ensuring that the HEAD points to v1?

###### Did the student display the logs with the deleted commits, particularly focusing on the commit tagged `oops`?

###### Did the student ensure that unreferenced commits were deleted from the history, with no logs remaining for these deleted commits?

###### Did the student add author information to the file and commit the changes?

###### Did the student update the file to include the author email without making a new commit, but included the change in the last commit?

#### Move it

###### Did the student successfully move the `hello.sh` program into a `lib/` directory using Git commands?

###### Did the student commit the move of `hello.sh`?

###### Did the student create and commit a `Makefile` in the root directory of the repository with the provided content?

#### blobs, trees and commits

##### Ask the student to navigate to the `.git/` directory and explain to you the purpose of each subdirectory, including `objects/`, `config`, `refs`, and `HEAD`.

###### Was the student able to explain the purpose of each subdirectory, including `objects/`, `config`, `refs`, and `HEAD`?

###### Did the student successfully find the latest object hash within the `.git/objects/` directory using Git commands?

###### Was the student able to print the type and content of this object using Git commands?

###### Did the student use Git commands to dump the directory tree referenced by a specific commit?

###### Were they able to dump the contents of the `lib/` directory and the `hello.sh` file using Git commands?

#### Branching, Merging & Rebasing

###### Did the student successfully create and switch to a new branch named `greet`?

###### Did the student create and commited a new file named `greeter.sh` in the `lib` directory with the provided code in it?

###### Did the student update the `lib/hello.sh` file with the provided content, stage, and commit the changes?

###### Did the student update the `Makefile` with the comment, stage, and commit the changes?

###### Was the student able to compare and show the differences between the `main` and `greet` branches for the `Makefile`, `hello.sh`, and `greeter.sh` files?

###### Did the student generate a `README.md` file with the provided content and commit it?

###### Did the student draw a commit tree diagram illustrating the diverging changes between all branches to demonstrate the branch history?

#### Conflicts, merging and rebasing

###### Did the student successfully merge the changes from the `main` branch into the `greet` branch?

###### Did the student make the specified changes to the `hello.sh` file in the `main` branch and commit them?

###### Did the student attempt to merge the `main` branch into the `greet` branch creating a conflict during the merge?

###### Did the student successfully resolve the conflict, accepting changes from the `main` branch?

###### Did the student commit the conflict resolution changes?

###### Did the student return to the point before the initial merge between `main` and `greet`?

###### Did the student rebase the `greet` branch on top of the latest changes in the `main` branch?

###### Did the student successfully merge the changes from the `greet` branch into the `main` branch?

##### Ask the student to explain the difference between merging and rebasing and if he understand Fast-Forwarding.

###### Did the student demonstrate an understanding of fast-forwarding?

###### was the student able to explain the difference between merging and rebasing?

#### Local & Remote Repositories

###### Did the student complete the cloning process of the `hello` repository to `cloned_hello`?

###### Did the student fetch and merge changes from the remote repository into the `main` branch?

###### Did the student list both remote and local branches, make changes to the original repository, and synchronize the cloned repository with remote changes?

###### Did the student successfully clone the `hello` repository into the `work/` directory as `cloned_hello`, without using the `copy` command?

###### Did the student show the logs for the `cloned_hello` repository?

###### Did the student display the name of the remote repository (`origin`) and provide more information about it?

###### Did the student list all remote and local branches in the `cloned_hello` repository?

###### Did the student make changes to the original repository, update the `README.md` file with the provided content, and commit the changes?

###### Inside the cloned repository (`cloned_hello`), did the student fetch the changes from the remote repository and display the logs, ensuring commits from the `hello` repository are included?

###### Did the student merge the changes from the remote `main` branch into the local `main` branch?

###### Did the student add a local branch named `greet` tracking the remote `origin/greet` branch?

###### Did the student add a `remote` reference to their Git repository?

###### Did the student push the `main` and `greet` branches to the `remote` repository?

##### Ask the following question to the student:

**What is the single git command equivalent to what you did before to bring changes from remote to local `main` branch?**

###### Did the student provide an accurate response?

#### Bare Repositories

##### Ask the following question to the student:

**What is a bare repository and why is it needed?**

###### Did the student correctly explain what a bare repository is and why it is needed?

###### Did the student successfully create a bare repository named `hello.git` from the existing `hello` repository?

###### Did the student add the bare `hello.git` repository as a remote to the original repository `hello`?

###### Did the student change the `README.md` file in the original repository, commit the change, and push it to the shared repository?

###### Did the student switch to the cloned repository `cloned_hello` and successfully pull down the changes just pushed to the shared repository?
