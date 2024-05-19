#### General

###### Was the project written in a compiled programming language?

#### Functional

##### Try to run the command `"ls -lRr / 2>1 >/dev/null &"` then run the command `"jobs"`.

```
[1]+  Running                 ls -lRr / 2>1 >/dev/null &
```

###### Can you confirm that the program displayed a list with the status of all jobs like in the example above?

##### Try to run the command `"jobs -l"`.

```
[1]+ 13612 Running                 ls -lRr / 2>1 >/dev/null &
```

###### Can you confirm that the program added the process ID to the normal information given in the command `"jobs"` like in the example above?

##### Try to run the command `"jobs -p"`.

```
13612
```

###### Can you confirm that the program only displays the process ID like in the example above?

##### Try to run the command `"sleep 50000 &"` then run `"python &"` and press enter without any input in the last command.

```
[1]   Running                 ls -lRr / 2>1 >/dev/null &
[2]-  Running                 sleep 50000 &
[3]+  Stopped                 python
```

###### Run the command `"jobs"`. Can you confirm that the program displays the list with the status of all jobs and that one of them is "Stopped" like the example above?

##### Try to run the command `"jobs -r"`.

```
[1]   Running                 ls -lRr / 2>1 >/dev/null &
[2]-  Running                 sleep 50000 &
```

###### Can you confirm that the program only displays the list with running jobs like in the example above?

##### Try to run the command `"jobs -s"`.

```
[3]+  Stopped                 python
```

###### Can you confirm that the program only displays the list with stopped jobs like in the example above?

##### Try to run the command `"kill 7764"`(the process ID must be yours this is just an example).

```
[2]-  Terminated              sleep 50000
```

###### Can you confirm that the program killed and displayed the process with the given id like in the example above?

##### Try to run the command `"kill %1"`.

```
[1]   Terminated              ls -lRr / 2>1 >/dev/null
```

###### Can you confirm that the program killed and displayed the first process like in the example above?

##### Close the program and run it again. Try to run the commands `"ls -lRr / 2>1 >/dev/null &"`, `"sleep 50000 &"` and then run `"fg"`.

```
sleep 50000

```

###### Can you confirm that the program brings the background job to the foreground like in the example above?

##### Try to run the command `"fg"` then stop the process with the `"Ctrl + Z"`.

```
sleep 50000
^Z
[2]+  Stopped                 sleep 50000
```

###### Can you confirm that the program brings the background job to the foreground and after you press `"Ctrl + Z"` the process stops like in the example above?

##### Try to run the command `"bg"`.

```
[2]+ sleep 50000 &
```

###### Run `"jobs"`. Can you confirm that the program started the process in the background like in the example above?
