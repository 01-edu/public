## hello_rust

### Introduction

Hello and welcome to rust.
The goal of this first exercise is to configure your repository properly and to give you a set of rules for the whole piscine rust.
Make sure those rules are followed consistenly during the whole piscine.

### Instructions

#### 1- get-ready

Create in your [Gitea](<https://git.((DOMAIN))>) account the repository named `((ROOT))`.

This repository will be the folder where all the exercices must be uploaded.

Once created, clone that repository on your desktop.

First, tell Git to remember your password (like a web browser would):

```
git config --global credential.helper store
```

If your username was `choumi` this is the command that will need to be used:

```
git clone https://git.((DOMAIN))/choumi/((ROOT)).git
```

This command needs to be adapted with **your own username**.

#### gitignore file

Once you repository is cloned, create and edit the .gitignore file in your repository and add this line:

```console
**/target/*
```

The goal of this setup is to avoid any binary files to be pushed in your Gitea accidentally.

Do not forget to push it to your repository.

1. `git add .gitignore`
2. `git commit -m "My very first gitignore commit"`
3. `git push`

#### 2- get-ready

Below are the commands that you must use during this piscine for initiating programs and functions. This command that you are going to depend in what is asked in the subject.

#### Commands for a program

```console
cargo new --vcs=none name-of-exercise
```

#### Commands for a function

```console
cargo new --vcs=none --lib name-of-exercise
```

#### 3- try it yourself

Execute the below command for creating a **program** inside your repository

```console
cargo new --vcs=none hello_rust
```

Then adapt the `main.rs` so that it would display `Hello, Rust!`

You can test it with the below command inside the folder of your exercise.

```console
cargo run
```

This command will compile, and run the binary.

#### 4- return your solution

After that the `hello_rust` project is executing correctly, it needs to be uploaded to the repository with the following commands:

1. `git add hello_rust/`
2. `git commit -m "My very first rust commit"`
3. `git push`

Once these steps are applied, the project can now be submitted for grading on the platform by clicking on the "RUN HELLO_RUST..." button.

This action will run the tests on your submitted `hello_rust` folder.
