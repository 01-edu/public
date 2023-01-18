## in-the-dark

### Instructions

Create a script `in-the-dark.sh` which will execute a job in the background that will run the command `ls` which will recursively list all files and directories in the current directory and its subdirectories. The output will be piped into the `grep` command, which will search for all files with the `.txt` extension. After the search is complete, you must redirect to a file "new" the following string "Search complete". You also need to List all background jobs.

You must do all these steps running only one job!

Expected output:

```console
$ ./in-the-dark.sh
[1]+  Running                 "Your running job, this is just an example of output"
news_amazon.txt
model_forecasts.txt
Ecommerce_purchases.txt
breast_cancer_readme.txt
shadow.txt
standard.txt
thinkertoy.txt
$ cat new
Search complete
$
```

### Hints

To run the `ls` command recursively in the background on a lot of files, you can use the `-R` option which will recursively list all files and directories in the current directory and its subdirectories. The output must be piped into the grep command. Here is a similar example.

```console
$ ls
is-winner
printalphabetalt
printcomb2
question_mark
onlyz
organize_garage
tetris-optimizer
unzipstring
$ ls | grep "z"
onlyz
organize_garage
tetris-optimizer
unzipstring
```

The `&&` operator is used to combine multiple commands, and execute them as a single command. It is often used to run multiple commands in sequence, where the second command only runs if the first command succeeds.

```console
$ touch newfile.txt && chmod 755 newfile.txt
$  ls -l
-rwxr-xr-x   1 user user       0 jan  2 17:29 newfile.txt
$
```

This will run the `touch` command, which will create a new file called "newfile.txt". If the touch command succeeds, the `chmod` command will run and change the permissions of "newfile.txt" to allow the owner to read, write, and execute the file, and allow everyone else to read and execute the file.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
