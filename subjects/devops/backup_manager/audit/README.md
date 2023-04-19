#### Functional

##### Before we start, let's make sure we have a fresh environment by running this command `rm -dr logs backups backup_schedules.txt`.

###### Can you confirm that the `backup_manager.py` and `backup_service.py` are present?

##### Run the following command `python3 ./backup_manager.py create "test2;18:15;backup_test2"`.

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
[13/02/2023 17:14] "Error: can't stop backup_service"
```

###### Can you confirm that the `backup_manager.log` file contains an error like the one above, stating that the service is already stopped?

##### Create a couple more backup schedules using the command `python3 ./backup_manager.py create "file;HH:MM;backup_name"`.

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

###### Was the task with the index `1` removed from the list like in the example above?

###### Did the task with the index `2` became the task with the index `1` after you removed the old one from the list?

##### Verify that the following command runs without errors: `python3 backup_manager.py start`.

###### Can you confirm that the `backup_service.py` process is running? (For example you could use `ps -ef | grep backup_service`). Check the example bellow:

```bash
$ python3 ./backup_manager.py start
$ -ps -ef | grep backup_service
user   1520028    2736  0 18:30 ?        00:00:00 python3 ./backup_service.py start &
```

##### Now run the command `python3 backup_manager.py start` again.

```bash
$ python3 ./backup_manager.py start
$ cat logs/backup_manager.log
...
[13/02/2023 17:14] Error: service already running
```

###### Can you confirm if the last log is stating that the service is already running? (You can use `cat logs/backup_manager.log` to confirm).

##### Run `python3 backup_manager.py stop` and then `rm -dr logs backups backup_schedules.txt`. Let's create the backup manually.

##### Run the following command by order:

```bash
mkdir testing testing2; touch testing/file1 testing/file2 testing/file3
python3 ./backup_manager.py create "testing;[Current_hour];backup_test"
python3 ./backup_manager.py create "testing;13:11;backup_test"
python3 ./backup_manager.py start
python3 ./backup_manager.py backups
```

Supposing the commands were run at 18:21, here is an example of the commands above:

```console
$ ls
backup_manager.py  backup_service.py  testing/
$ python3 ./backup_manager.py create "testing;18:21;backup_test"
$ python3 ./backup_manager.py create "testing2;13:11;passed_time_backup"
$ python3 ./backup_manager.py list
0: testing;18:21;backup_test
1: testing;13:11;passed_time_backup
$ python3 ./backup_manager.py start
$ python3 ./backup_manager.py backups
backup_test.tar
$ ls
backup_manager.py  backups/  backup_schedules.txt  backup_service.py  logs/ testing/
$ cat logs/backup_manager.log
[13/02/2023 18:21] Schedule created
[13/02/2023 18:21] Show backups list
[13/02/2023 18:21] Show backups list
$ cat backup_schedules.txt
testing;18:21;backup_test
$ cat logs/backup_service.log
[13/02/2023 18:21] Backup done for testing in backups/backup_test.tar
```

##### Follow the example above and then open the `backup_test.tar` file to ensure that the backup process was successful. Verify that the files are not empty or damaged and that it matches the original directory. The result should be similar to the following example:

```console
$ ls
backup_test.tar
$ tar -tvf backup_test.tar
drwxrwxr-x user/user     0 2023-04-19 10:23 testing/
-rw-rw-r-- user/user     0 2023-04-19 10:23 testing/1
-rw-rw-r-- user/user     0 2023-04-19 10:23 testing/2
```

###### Was the backup created successfully?

###### Can you confirm that the `passed_time_backup` was not created successfully because the time has already passed?

##### Run `python3 backup_manager.py stop` and then `rm -dr logs backups backup_schedules.txt`.

##### Create a `.zip` with some folders and files inside it and then replicate the steps you just did above to check if the backup is created successfully.

###### Was the backup created successfully? (open the backup and verify that the files are not empty or damaged and that it matches the original directory.)

### Error handling

##### Verify error handling for incorrect commands and incorrect arguments. For example, confirm that an error message is logged when attempting to run the command `python3 backup_manager.py invalid_command`.

```bash
$ python3 backup_manager.py invalid_command
$ cat logs/backup_service.log
...
[13/02/2023 18:21] Error: unknown instruction
```

###### Was the error message logged correctly?

##### Run `python3 backup_manager.py stop` and then `rm -dr logs backups backup_schedules.txt`. After that run the following commands:

```bash
python3 ./backup_manager.py stop
python3 ./backup_manager.py create "wrong_format"
python3 ./backup_manager.py backups
python3 ./backup_manager.py start
python3 ./backup_manager.py start
```

```console
$ cat logs/backup_manager.log
[14/02/2023 15:07] Error: can't stop backup_service
[14/02/2023 15:07] Error: malformed schedule: wrong_format
[14/02/2023 15:08] Error: can't find backups directory
[14/02/2023 15:08] Error: backup_service already running
$ cat logs/backup_service.log
[14/02/2023 15:11] Error: cannot open backup_schedules
$
```

###### Can you confirm that all error messages have been saved in the log files like in the example above?

##### Go to the code of the project and check if the creator used `try` and `except` to handle the errors.

###### Did he use `try` and `except` to handle the errors?
