## scripting

### Objectives

You must follow the same [principles](https://public.01-edu.org/subjects/0-shell/) as the first subject.

In `scripting` you must find a way to make your `0-shell` run shell scripts.

By now you should be familiar with the basics in an interactive shell. Now your purpose is to learn how to put those commands together to create scripts and run them in your own interpreter.

Some examples of common scripting languages are [ash](https://en.wikipedia.org/wiki/Almquist_shell), [bash](<https://en.wikipedia.org/wiki/Bash_(Unix_shell)>) and [dash](https://wiki.archlinux.org/index.php/Dash). But you can choose another one as long as it is a [Unix shell](https://en.wikipedia.org/wiki/Unix_shell).

You will have to create your own valid script called `create-dir`. This script must be able to create a new directory by checking if it exists or not. So, you will have to run the script, where it will ask the user for a "Directory name", then the user will give a name of the directory that he wants to create and press enter. The script will check whether the directory exists or not, and if it does not exist, the directory will be created and the message "Directory created" will be displayed. If the directory already exists, nothing will be created and the message "Directory exist" will be displayed.

Your program must be able not just to run script files but also to run the scripts directly on the interpreter. For example, it must be able to pass for loops, functions, manipulate data etc...:

```
for ((i = 0 ; i < 100 ; i++)); do
  echo $i
done
```

```
myfunc() {
    echo "hello $NAME"
}
```

### Instructions

- You have to create your own script.
- The `0-shell` must be able to read and execute scripts.
- The project has to be written in **Rust**.
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
