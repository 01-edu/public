#### Functional

###### Can you confirm that the `backup_manager.py` and `backup_service.py` are present?

##### Run the following command `python3 ./backup_manager.py create "test2;18:15;backup_test2"`

###### Can you confirm that the `backup_schedules.txt` was created?

##### Run the following command `cat backup_schedules.txt`.

```bash
$ cat backup_schedules.txt
test2;18:15;backup_test2
```

###### Was the schedule created properly, with the same format as the example above?

##### Run the following command `python3 ./backup_manager.py stop`.

###### Can you confirm that the `logs` folder was created and inside of it is the `backup_manager.log`?

##### Now run the command `cat logs/backup_manager.log`.

```bash
$ cat logs/backup_manager.log
[13/02/2023 17:14] Error: cannot kill the service
```

###### Can you confirm that the `backup_manager.log` file contains an error like the one above, stating that the service is already stopped?

##### Create a couple more backup schedules using the command `python3 ./backup_manager.py create "file;HH:MM;backup_name"`

##### Run the command `python3 ./backup_manager.py list`.

```bash
$ python3 ./backup_manager.py list
0: test2;18:15;backup_test2
1: test3;18:16;backup_test3
2: test4;18:17;backup_test4
```

###### Did you get a result like the one above, with the schedule you created attached with an index at the beginning, on your terminal?

##### Run the command `python3 ./backup_manager.py delete 1`.

```bash
$ python3 ./backup_manager.py list
0: test2;18:15;backup_test2
1: test4;18:17;backup_test4
```

###### Was the task with the index `1` removed from the list? Use `python3 ./backup_manager.py list`to.

##### Verify that the following command runs without errors: `python3 backup_manager.py start`.

###### Can you confirm that the `backup_service.py` process is running by using the command `ps -ef | grep backup_service`?

##### Now run the command `python3 backup_manager.py start` again.

```bash
$ python3 ./backup_manager.py start
$ cat logs/backup_manager.log
[13/02/2023 17:14] Error: service already running
```

###### Can you confirm bu running the command `cat logs/backup_manager.log` if the log file contains an error like the one above, stating that the service is already running?

##### Run `python3 backup_manager.py stop` and delete everything that was created until now, the `logs` folder and the `backup_schedules.txt`. Lets create the backup manually.

##### Run the following command by order:

    1- mkdir testing; cd testing; touch 1 2 3; cd ../
    2- python3 ./backup_manager.py create "testing;"Current_hour";backup_test".
    3- python3 ./backup_manager.py start
    4- python3 ./backup_manager.py backups

```bash
$ mkdir testing; cd testing; touch 1 2 3; cd ../
$ ls
backup_manager.py  backup_service.py  testing/
$ python3 ./backup_manager.py create "testing;18:21;backup_test"
$ python3 ./backup_manager.py list
0: testing;18:21;backup_test
$ python3 ./backup_manager.py start
$ python3 ./backup_manager.py backups
backup_test.tar
$ ls
backup_manager.py  backups/  backup_schedules.txt  backup_service.py  logs/ testing/
$ cat logs/backup_manager.log
[13/02/2023 18:20] Schedule created
[13/02/2023 18:20] Show schedule list
[13/02/2023 18:21] Backup list
$ cat backup_schedules.txt
testing;18:21;backup_test
$ cat logs/backup_service.log
[13/02/2023 18:21] Backup done for testing in backups/backup_test.tar
```

##### Can you confirm that the backup was created successfully? Follow the example above.

### Error handling

##### Verify error handling for incorrect commands and incorrect arguments. For example, confirm that an error message is displayed when attempting to run the command `python3 backup_manager.py invalid_command`.

```bash
$ python3 backup_manager.py invalid_command
Error: unknown instruction
```

###### Was the error message displayed?

##### Run `python3 backup_manager.py stop` and delete everything that was created until now, the `logs` folder and the `backup_schedules.txt`. After that run the following commands:

    1- python3 backup_manager.py stop
    2- python3 ./backup_manager.py create "wrong_format"
    3- python3 ./backup_manager.py backups
    4- python3 ./backup_manager.py start
    5- python3 ./backup_manager.py start

```bash
$ cat logs/backup_manager.log
[14/02/2023 15:07] Error: cannot kill the service
[14/02/2023 15:07] Error: invalid schedule format
[14/02/2023 15:08] Error: [Errno 2] No such file or directory: "./backups"
[14/02/2023 15:08] Error: service already running
$ cat logs/backup_service.log
[14/02/2023 15:11] Error: cannot open backup_schedules

```

###### Can you confirm that all error messages have been saved in the log files like in the example above?
