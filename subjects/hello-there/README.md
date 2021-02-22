## Hello There

### Instructions

#### Install `Scoop` *(windows)* or `Brew` *(mac)*

Follow the instructions here:
- **scoop**: https://scoop.sh/ (windows)
- **brew**: https://brew.sh/ (mac)

> Linux people, you should know how to handle this


#### Install `git`, `vscode` and `nodejs`

```bash
# windows
scoop install git
git --version
scoop install nodejs
node --version
scoop bucket add extras
scoop install vscode
code --version

# mac
# git should already be installed
git --version
brew install node
node --version
brew install --cask visual-studio-code
code --version
```

#### Configure git and vscode

```bash
# Set your username and email
# (the double quotes "" are important)
git config --global user.name "Your name here"
git config --global user.email "your_email@example.com"

# Tell git to use vscode
git config --global core.editor code

# Add colors to the info, not important but fancy
git config --global color.ui true
```

> If you are adventurous, you can also setup your SSH key
> this will allow you to avoid typing your password everytime.


#### Write the code !

Once you have installed and configured the necessary tools,
create a `git` repository with a `hello.js` JS file
that display the text `Hello There` in the terminal.

> Don't forget to commit and push the file to the servers

### Usage

```console
student@ubuntu:~/[[ROOT]]/test$ node hello.js
hello there
student@ubuntu:~/[[ROOT]]/test$
```

### You will learn about

- Terminal
- Code Editor
- Git
- JS
