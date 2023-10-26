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

#### Branching
It’s time to do a major rewrite of the hello world functionality. Since this might take awhile, you’ll want to put these changes into a separate branch to isolate them from changes in main.
- Create a local branch named `greet` and switch to it.
- Create `greeter.rb` file in `lib` directory, add the following content to it and commit the changes
```
class Greeter
  def initialize(who)
    @who = who
  end
  def greet
    "Hello, #{@who}"
  end
end
```
- Update `lib/hello.rb` file by adding the content below, stage and commit the changes.
```
require 'greeter'

# Default is World
name = ARGV.first || "World"

greeter = Greeter.new(name)
puts greeter.greet
```
- Update `lib/Rakefile` too and commit the changes
```
#!/usr/bin/ruby -wKU

task :default => :run

task :run do
  ruby '-Ilib', 'lib/hello.rb'
end
```
- Switch to `main` branch, show the difference between the versions in `main` and `greeter` branches for the these files: `Rakefile`, `hello.rb` and `greeter.rb`
- Create `README.md` file with the content below and commit the changes.
```
This is the Hello World example from the git project.
```
- Draw the commit tree for all the branches to show the diverging changes.

#### Conflicts, merging and rebasing
- Merge `main` branch into `greeter` branch.
- Switch to `main` branch and make the changes above to `hello.rb` save and commit.
```
puts "What's your name"
my_name = gets.strip

puts "Hello, #{my_name}!"
```
- Now switch to `greeter` branch and try to merge `main` into it (Bingooo! there you have a conflict!).
- Resolve the conflict (manually or using graphical merge tools), accept changes from `main` branch, then commit the conflict resolution.
- Go back in time before the very first merge. now rebase the branch `greeter` on top of `main` branch.
- Now merge changes from `greeter` into `main` branch.
- Explain fast-forwarding and the difference between merging and rebasing.

#### Local and remote repositories
- In `work/` directory make a clone of the repository `hello` as `cloned_hello` (do not use `copy` command).
- Show the logs for the cloned repository, what are `origin/main`, `origin/greet` and `origin/HEAD` ?.
- Display the name of the remote repository, show more information about it.
- List all the remote and local branches.
- Make changes to the original repository, Update `README.md` file and commit the changes.
```
This is the Hello World example from the git tutorial.
(changed in original)
```
- Inside the cloned copy fetch the changes from remote and display the logs (commits from `hello` repository should be included in the logs).
- Merge remote `main` branch into local `main` branch.
- What is the single git command equivalent to what you did before to bring changes from remote to local `main` branch?
- Add a `greet` local branch tracking the remote `origin/greet` branch.

#### Bare repositories
- What is a bare repository and what is it needed for?
- Create a bare repository from `hello` repository and name it `hello.git`.
- Add the bare `hello.git` repository as a remote to our original repository `hello`.
- Change README.md file, commit and push the change to the shared repository.
```
This is the Hello World example from the git tutorial.
(Changed in the original and pushed to shared)
```
- Quick hop over to the clone repository `cloned_hello` and pull down the changes just pushed to the shared repository.