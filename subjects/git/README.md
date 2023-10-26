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

#### Changed your mind?
- Make the following changes to the latest version of the file, then revert it before staging (Do not use CTRL+Z)
```
# This is a bad comment.  We want to revert it.
name = ARGV.first || "World"

puts "Hello, #{name}!"
```
- Modify the file to have a bad comment, stage it, then clean the staging area.

```
# This is an unwanted but staged comment
name = ARGV.first || "World"

puts "Hello, #{name}!"
```
- Add the following changes again, stage the file and commit the changes, then revert it.
```
# This is an unwanted but committed change
name = ARGV.first || "World"

puts "Hello, #{name}!"
````
- Mark the latest commit with `oops` tag, then remove commits coming after `v1` version (HEAD should be in `v1`)
- Show the logs with the deleted commits displayed (eg: log for commit tagged `oops` shoud be displayed).
- Now make sure the unreferenced commits are deleted from the history (no logs for deleted commits).
- Add an author comment to the file and commit the changes
```
# Default is World
# Author: Jim Weirich
name = ARGV.first || "World"

puts "Hello, #{name}!"
```
- Oops we forgot the author email, Update the file to include the email. don't make a new commit, but include the change in the last commit. 

#### Move it
- Move the program `hello.rb` into a `lib/` directory (use git commnad) and commit the move.
- Add a Rakefile to the root of the repository and commit it.
```
#!/usr/bin/ruby -wKU

task :default => :run

task :run do
  require './lib/hello'
end
```
#### blobs, trees and commits
- Explore .git/ directory and explain its contents: `objects/`, `config`, `refs` and `HEAD`.
- Grab the latest object hash within ``.git/object`` directory print its type and its content using `git` command.
- Dump the directory tree referenced in the commit
- Dump the `lib` directory, then `hello.rb` file.