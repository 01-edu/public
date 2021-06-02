## Introduction

### Instructions

#### 1- get-ready

Create in your Gitea (https://git.[[DOMAIN]]) account the repository named `[[ROOT]]`.

This repository will be the folder where all the exercices must be uploaded.

Once created, clone that repository on your desktop.

If your username was `choumi` this is the command that will need to be used :

`git clone git@git.[[DOMAIN]]:choumi/[[ROOT]].git`

This command needs to be adapted with **your own username**.

If the `git clone` gives you an authenticity of host error, your SSH key must be configured.
Follow the steps below.

#### SSH Configuration

Execute the following commands:

```console
mkdir -p ~/.ssh
ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N ''
cat ~/.ssh/id_ed25519.pub
```

- Copy the result and paste it in the content field of adding an SSH key in your Gitea settings : https://git.[[DOMAIN]]/user/settings/keys

- Confirm by clicking on the add key button.

Once this is done the git clone command should work now.

#### 2- set

Once the repository is created, write your first shell program called `hello.sh`

When executed this program must print `Hello choumi!`
Where `choumi` is your username

##### Usage

If the username is `choumi` :

```console
$ ./hello.sh
Hello choumi!
$
```

#### 3- go-say-hello

After that the `hello.sh` is executing correctly, it needs to be uploaded to the repository with the following commands :

1. `git add hello.sh`
2. `git commit -m "My very first commit"`
3. `git push`

Once these steps are applied, the file can now be submitted for grading on the platform by clicking on the "RUN INTRODUCTION TEST..." button.

This action will run the tests on your submitted `hello.sh` file.

### Recommendation

Videos designed to give **hints** are assigned to each quest. It is strongly suggested to watch them as you go.
