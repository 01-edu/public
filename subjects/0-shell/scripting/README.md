## scripting

### Objectives

You must follow the same [principles](https://public.01-edu.org/subjects/0-shell/) as the first subject.

In `scripting` you must find a way to make your `0-shell` run `bash` scripts.

By now you should be familiar with the basics in an interactive shell. Now your purpose is to learn how to put those commands together to create scripts and run them in your own interpreter.

Most users that think of `bash` think of it as a prompt and a command line. That is `bash` in interactive mode. BASH can also run in non-interactive mode, as when executing scripts. We can use scripts to automate certain logic. Scripts are basically lists of commands (just like the ones you can type on the command line), but stored in a file. When a script is executed, all these commands are (generally) executed sequentially, one after another.

You will have to create your own valid script called `create-dir.sh`. This script must be able to create a new directory by checking if it exists or not. So, you will have to run the script, where it will ask the user for a "Directory name", then the user will give a name to the directory that he wants to create and press enter. The script will check if the directory exists or not and if it does not exist, the directory will be created and the message "Directory created" will be displayed. If the directory already exists, nothing will be created and the message "Directory exist" will be displayed.

Your program must be able not just to run `.sh` files but to run the scripts directly on the interpreter. For example, it must be able to pass for loops, functions, manipulate data etc...:

```
for ((i = 0 ; i < 100 ; i++)); do
  echo $i
done
```

```
STR="HELLO WORLD!"
echo ${STR,}   #=> "hELLO WORLD!" (lowercase 1st letter)
echo ${STR,,}  #=> "hello world!" (all lowercase)

STR="hello world!"
echo ${STR^}   #=> "Hello world!" (uppercase 1st letter)
echo ${STR^^}  #=> "HELLO WORLD!" (all uppercase)
```

```
myfunc() {
    echo "hello $NAME"
}
```
### Instructions

- You have to create your own script.
- The `0-shell` must be able to read and execute `bash` scripts.
- The project has to be written in a compiled language like (C, Rust or other), no semi compiled language like (Pearl and others) are allowed.
- The code must respect the [good practices](https://public.01-edu.org/subjects/good-practices/)

This project will help you learn about:

- Scripting/ scripting language
- Commands syntax

### Usage

```
student$ ./0-shell
$ bash create-dir.sh
Enter directory name
Example
Directory created
$
Enter directory name
Example
Directory exist
$ exit
student$
```
```
student$ ./0-shell
$ for ((i = 0 ; i < 5 ; i++)); do echo $i; done
0
1
2
3
4
$ exit
student$
```
