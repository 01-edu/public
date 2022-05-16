## Hello There 👋

### Installation

For each programs, make sure you are downloading a version that works for your system

#### Install `vscode`

- Download the installer https://code.visualstudio.com/download

> If you already have and love a code editor, feel free to use it,
> but keep in mind that we will use VSCode as reference

VSCode is a code editor, it will give you an interface to write, test and submit your code.

#### Install `git`

- Download the installer https://git-scm.com/downloads
- While installing, Git will ask a bunch of questions, it is recommended to change those:
	- Choosing the default editor (pick VSCode)
	- Configuring the line ending conversions (Choose : Checkout as-is, commit Unix-style line endings)
	- Otherwise stick with the defaults unless you know what you are doing.

Git is a versioning tool, we use it to upload your solutions on the school server.

#### Install `nodejs`

- Download the installer https://nodejs.org/en/download/current/

NodeJS will allow to execute JavaScript code outside of the browser, useful for testing your code.


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
create in your [Gitea](<https://((DOMAIN))/git>) account the repository named `((ROOT))` with a `hello-there.js`
JS file that is a program that displays:
- the exact text `Hello There !`
- any `Number` 
- and a `Boolean`.

In order to work in your repository and put files in it, you need to clone it first. 

If your username was `choumi` this is the command that will need to be used:

```
git clone https://((DOMAIN))/git/choumi/((ROOT)).git
```
To execute it, open a Unix shell (e.g. Git Bash on Windows), you are going to type commands in it.
This command needs to be adapted with **your own username**.

> Don't forget to commit and push the file to the servers

### Optional (git setting to avoid typing the password every time)

The `((ROOT))` repository will be the folder where all the exercises must be uploaded.

In order to avoid writing your username and password every time you `git push` an exercise, 
tell Git to remember your password (like a web browser would) with the below command:

```
git config --global credential.helper store
```

### Recommendation

Videos designed to give **hints** are assigned to each quest. It is strongly suggested to watch them as you go.

### You will learn about

- Code Editor
- Git
- JS
