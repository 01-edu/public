## Hello There 👋

### Installation

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

### Values

There are 3 types of values, and they can model the **WORLD** !\
_(In the end it's all `1`'s and `0`'s for the computer)_, but us, **humans**, need
an easier way of representing stuff.

> We can all agree that `11010001100101110110011011001101111` is not a very
> friendly way to say `'hello'` !

#### Numbers 🔢

- Whole numbers: `1`, `23`, `232139283`
- Negative numbers are prefixed with `-`: `-1`, `-1231`
- Decimal numbers: `3.14`, `-2.53343` etc...

Use them for _quantities_ like in daily life.

#### Booleans ✖️ / ✔️

- Something is `true`
- or `false`

They represent a truth, an answer to a closed-ended question _(anything that can
be answered with yes or no)_:

- Is paris the capital of France ? `true`
- Are you born before 1723 ? `false`
- Is your screen turned on ? `true` _(most likely)_

#### Strings 🆒

- `'Hello'`
- `'This is some text'`

A string is a sequence of characters used to represent text, it needs
**delimiters** to define its _begining_ and _end_.\
Delimiters are matching quotes, either `` ` ``, `"` or `'`.

### Using `console.log`

To display output from a script into a console, use the function `console.log`:

```js
console.log() // <- will show an empty line
```

Add any value between the parentheses to see it appear when the script is
executed.

> It is very important to use this often to validate that our code is valid. The
> more it is tested, the easier it is to understand what's going on in the code
> !
>
> In doubt, `console.log` everything, don't be shy, they are for free.

### Instructions

Once you have installed and configured the necessary tools,
create a `git` repository named `[[ROOT]]` with a `hello-there.js`
JS file that is a program that displays the exact text `Hello There !`, any `Number` and a
`Boolean`.

> Don't forget to commit and push the file to the servers

### You will learn about

- Terminal
- Code Editor
- Git
- JS
