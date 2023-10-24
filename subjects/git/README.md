# Git Ready
## Introduction
The Git project is designed to introduce you to the world of version control and collaboration with Git. 
Git is a powerful and widely used tool for tracking changes in your projects, collaborating with others, 
and ensuring the integrity of your code.

To aid your learning journey, we provide you with a resource called "git-ready.zip," which you can find in the "resources" folder.

This Git project is structured as a series of exercises. You will progressively build your Git skills,
starting from the basics and gradually moving towards more advanced topics.

Let's Git ready for it! 

#### Setting Up Git and Python for Work
- Install Git on your local machine by following the instructions for your operating system on the official Git website.
- Configure Git with your name and email address.

#### Git commits to commit
- In `work` directory, create a folder named `hello` then a file `hello.rb` with the content below:
  ```
  puts "Hello, World"
  ```
- Initialize the git repository in `hello` directory, then add `hello.rb` to the repository.
- Check the status and act accordingly with the output of the executed command.
- Change the “Hello, World” program. Change the file to be
```
puts "Hello, #{ARGV.first}!"
```
- Stage the changed file and commit the changes, the working tree should be clean.
- Change again `hello.rb` file and add the file to staging area, make two commits one for the comment in line1, the other for lines 3 and 4.
```
# Default is "World"
name = ARGV.first || "World"

puts "Hello, #{name}!"
```

#### History
- Show the history of the working directory.
- Show One line history; only the hashes and commit messages are displayed.
- Show logs with control over entries displayed: 2 last entries, since 5 minutes ago, until 5 minutes ago.
- Show logs in a personalized format like: `* e4e3645 2023-06-10 | Added a comment (HEAD -> main) [John Doe]`

#### Check it out
- Restore the first snapshot of your working tree, then print the content of `hello.rb` file.
- Restore the second recent snapshot and print the content of `hello.rb` file.
- Return the latest version in the main branch (do not use commit hash)

#### TAG me
- Refer the current version of the repository as `v1`.
- Tag the version immediately prior to the current version as `v1-beta`, do not use commit hashes to move through the history.
- Go back and forth between the two tagged versions.
- Show the list of tags




