#### Functional

###### Has the requirement for the allowed packages been respected?

##### Run both my-ls and the system command `ls` with no arguments.

###### Does it display the same files and/or folders in the same order?

##### Run both my-ls and the system command `ls` with the arguments: `"<file name>"`.

###### Does it display the same file?

##### Run both my-ls and the system command `ls` with the arguments: `"<directory name>"`.

###### Does it display the same files and/or folders in the same order?

##### Run both my-ls and the system command `ls` with the flag: `"-l"`.

###### Does it display the same files and/or folders with the same display?

##### Run both my-ls and the system command `ls` with the arguments: `"-l <file name>"`.

###### Does it display the same file with the same display?

##### Run both my-ls and the system command `ls` with the arguments: `"-l <directory name>"`.

###### Does it display the same files and/or folders with the same display?

##### Run both my-ls and the system command `ls` with the flag: `"-l /usr/bin"`.

###### Does it display the same files and/or folders with the same display? Be aware of symbolic links.

##### Run both my-ls and the system command `ls` with the flag: `"-R"`, in a directory with folders in it.

###### Does it display the same files and/or folders?

##### Run both my-ls and the system command `ls` with the flag: `"-a"`.

###### Does it display the same files and/or folders in the same order?

##### Run both my-ls and the system command `ls` with the flag: `"-r"`.

###### Does it display the same files and/or folders in the same order?

##### Run both my-ls and the system command `ls` with the flag: `"-t"`.

###### Does it display the same files and/or folders in the same order?

##### Run both my-ls and the system command `ls` with the flag: `"-la"`.

###### Does it display the same files and/or folders in the same order?

##### Run both my-ls and the system command `ls` with the arguments: `"-l -t <directory name>"`.

###### Does it display the same files and/or folders in the same order?

##### Run both my-ls and the system command `ls` with the arguments: `"-lRr <directory name>"`, in which the directory chosen contains folders.

###### Does it display the same files and/or folders in the same order?

##### Run both my-ls and the system command ls with the arguments: `"-l <directory name> -a <file name>"`

###### Is the output displayed the same way?

##### Run both my-ls and the system command ls with the arguments: `"-lR <directory name>///<sub directory name>/// <directory name>/<sub directory name>/"`

###### Is the output displayed the same way? Number of `/` must be the same.

##### Run both my-ls and the system command ls with the arguments: `"-la /dev"`

###### Does it display the same files and/or folders with the same display? Do not pay attention to ACL permission flag.

##### Run both my-ls and the system command ls with the arguments: `"-alRrt <directory name>"`, in which the directory chosen contains folders and files within folders. Time of modification of all files within that folder must be the same.

###### Is the displayed output the same?

##### Create directory with `-` name and run both my-ls and the system command ls with the arguments: `"-"`

###### Is the displayed output the same?

##### Create file and link for this file and run both my-ls-1 and the system command ls with the arguments: `"-l <symlink file>/"`

###### Is the displayed output the same? Pay attention to `/` at the end.

##### Create file and link for this file and run both my-ls-1 and the system command ls with the arguments: `"-l <symlink file>"`

###### Is the displayed output the same?

##### Create directory that contains files and link for this directory and run both my-ls-1 and the system command ls with the arguments: `"-l <symlink dir>/"`

###### Is the displayed output the same? Pay attention to `/` at the end.

##### Create directory that contains files and link for this directory and run both my-ls-1 and the system command ls with the arguments: `"-l <symlink dir>"`

###### Is the displayed output the same?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### General

###### +Does the program with colors as in the ls command?

###### +Does the program has other flags except for the mandatory ones?

##### Try running the program with `"-R ~"` and with the command time before the program name (ex: "time ./my-ls-1 -R ~").

###### +Is the real time less than 1,5 seconds?

#### Basic

###### +Does the code obey the [good practices](../../good-practices/README.md)?

###### +Is there a test file for this code?

###### +Are the tests checking each possible case?

#### Social

###### +Did you learn anything from this project?

###### +Would you recommend/nominate this program as an example for the rest of the school?
