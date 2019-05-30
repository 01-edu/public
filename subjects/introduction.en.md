## Introduction

### Instructions

#### 1- get-ready

Create in your github account the repository named `piscine-go`.

Once created, clone that repository on your desktop.
If your github username was `kigiri` this is the command that will need to be used :

`git clone https://github.com/kigiri/piscine-go`

This command needs to be adapted with **your own github username**.

This repository will be the folder where all the exercices must be submitted.

#### 2- set

Once the repository is properly created, write your first shell program called `hello.sh`

When executed this program must print `hello {username}!`
Where `{username}` is your GitHub username

##### Usage

If the `{username}` is `kigiri` :

```console
user@host:~/piscine$ ./hello.sh
Hello kigiri!
user@host:~/piscine$
```

#### 3- go-say-hello

After that the `hello.sh` is executing properly, it needs to be uploaded to the repository with the following commands :

1. `git add hello.sh`
2. `git commit -m "My very first commit"`
3. `git push origin master`

Once these steps are executed properly, the file can now be submitted for grading on the platform by clicking on the `submit` button.

This action will run the tests on your submitted `hello.sh` file.
