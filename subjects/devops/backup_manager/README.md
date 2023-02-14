## Backup manager

### Instructions

You will create two scripts that will manage and perform scheduled backups.

#### First script: backup_manager.py

The script `backup_manager.py` will orchestrate the backup service and if necessary update the file `backup_schedules.txt`. In order to do so it will accept the following command-line arguments:

- `start`:

  - Runs `backup_service.py` in the background.
  - Checks if the service is already running.
  - If the service is already running and you try to run it again it should send an error message to the log file.

- `stop` does the same as start but in order to sop or kill the process `backup_service.py`.
- `create [schedule]`:

  - Adds a new backup schedule in `backup_schedules.txt`.
  - This schedule will have a specific format which is `"file_name;hour:minutes;backup_file_name"`
  - If the schedule format is wrong, an error message should be sent to the log file.

```bash
$ python3 ./backup_manager.py create "wrong_format"
$ cat logs/backup_manager.log
[14/02/2023 15:07] Error: invalid schedule format
```

    - If the schedule format is correct and the schedule is created successfully, you should also send a the log file.

```bash
$ python3 ./backup_manager.py create "testing;15:44;backup_testing"
$ cat logs/backup_manager.log
[14/02/2023 15:44] Schedule created
```

- `list`:
  - Prints the scheduled backups in `backup_schedules.txt`, adding an index before each schedule.
  - Sends a message to the logs saying that the list is being accessed.

```bash
$ python3 ./backup_manager.py list
0: testing;15:46;backup_test
$ cat logs/backup_manager.log
[14/02/2023 15:46] Show schedule list
```

- `delete [index]`:

  - Delete the backup schedule at line `index` (starting by 0) in `backup_schedules.txt`.
  - If the schedule is deleted, you must send a message to the log file.
  - If the `index` does not exist or isn't valid, you must send an error message to the log file.

```bash
$ python3 ./backup_manager.py delete 1
0: testing; 15:54;backup_testing
$ cat logs/backup_manager.log
[14/02/2023 15:54] Error: index does not exist
[14/02/2023 15:54] Error: invalid index
[14/02/2023 15:54] Schedule deleted
```

- `backups`:

  - List the backups files in `./backups`.
  - If the `backups` folder is empty it should print an error message informing the user.
  - If the `backups` folder is not found it should be sent an error message to the log.
  - If it works correctly it should send a messafe to the lo file informing the user.

```bash
$ python3 ./backup_manager.py backups
backup_testing.tar
$ cat logs/backup_manager.log
[14/02/2023 16:05] Backup list
```

#### Second script: backup_service.py

The script `backup_service.py` will check the schedules in `backup_schedules.txt` and perform the daily backups at the proper time.

- The service will run in an infinite loop.
- On every cycle of the loop the service will check the schedules.
- The service will perform the backup if the current time (hour and minute) matches the time of the schedule.
- The format of the schedules is `path_to_save;time(hh:mm);backup_name`.
- The backups will be saved in `./backups`.
- The backups will be a compressed file (`.tar`) for the `path_to_save` directory.

> At the end of each loop it would be wise to make the service sleep for about 45 seconds to save processor cycles.

### Logs and error handling

Both scripts will have to use `try` and `except` to handle errors.
All actions and errors should be logged into a specific file in the `./log` directory.

- `backup_manager.py` will log into `./logs/backup_manager.log`.
- `backup_service.py` will log into `./logs/backup_service.log`.

You can choose to phrase your logs as you prefer, however it would be smart to start with a timestamp and try to be as specific as you can.

### Usage

Here is a little script to show how the backup system should work:

```bash
echo "--> Create 3 new schedules"
python3 ./backup_manager.py create "test;16:07;backup_test"
python3 ./backup_manager.py create "test1;16:07;personal_data"
python3 ./backup_manager.py create "test2;16:07;office_docs"
echo "--> Try to add a malformed schedule"
python3 ./backup_manager.py create "test;"
echo "--> Instruction list"
python3 ./backup_manager.py list

python3 ./backup_manager.py start
python3 ./backup_manager.py stop

echo "--> Instruction backups"
python3 ./backup_manager.py backups

echo "--> Content of the directory"
ls
echo "--> cat on ./logs/backup_manager.log"
cat ./logs/backup_manager.log
echo "--> cat on ./logs/backup_service.log"
cat ./logs/backup_service.log
echo "--> Content of ./backups"
ls ./backups
```

The output of this script should be similar to the following:

```console
--> Create 3 new schedules
--> Try to add a malformed schedule
--> Instruction list
0: test;16:07;backup_test
1: test1;16:07;personal_data
2: test2;16:07;office_docs
--> Instruction backups
personal_data.tar
office_docs.tar
backup_test.tar
--> Content of the directory
backup_manager.py  backups  backup_schedules.txt  backup_service.py  logs
--> cat on ./logs/backup_manager.log
[08/02/2023 16:07] New schedule added: test;16:07;backup_test
[08/02/2023 16:07] New schedule added: test1;16:07;personal_data
[08/02/2023 16:07] New schedule added: test2;16:07;office_docs
[08/02/2023 16:07] Error: malformed schedule: test;
[08/02/2023 16:07] Show schedules list
[08/02/2023 16:07] backup_service started
[08/02/2023 16:07] backup_service stopped
[08/02/2023 16:07] Show backups list
--> cat on ./logs/backup_service.log
[08/02/2023 16:07] Backup done for test in backups/backup_test.tar
[08/02/2023 16:07] Backup done for test1 in backups/personal_data.tar
[08/02/2023 16:07] Backup done for test2 in backups/office_docs.tar
--> Content of ./backups
backup_test.tar  office_docs.tar  personal_data.tar
```

### Hints

- Familiarize yourself with the Python, including the use of modules such as `subprocess` and `shlex`, and file `I/O` operations.
- To run a script from another python script you could use `subprocess.Popen` with the flag `start_new_session=True`.
- To kill a process you should find its process id and call `os.kill`.
- Play with the command `ps -A -f`, it will show a list of all active processes with the arguments attached to them and their process ids.
- The standard library `tarfile` is very useful when trying to work with archives.
- This starts to be quite a big project, try to be as modular as you can by creating single-task functions, this will save you a lot of time.

> There is many ways to run tasks in background and to perform actions on a specific time, the implementation proposed in this exercise is good for learning but should not be your preferred choice for production environments.

### References

- [Error handling in Python](https://docs.python.org/3.10/tutorial/errors.html)
- [Spawn a subprocess in Python](https://docs.python.org/3.10/library/subprocess.html)
- [Simple lexical analysis, shlex](https://docs.python.org/3/library/shlex.html)
- [Reading and writing files](https://docs.python.org/3/tutorial/inputoutput.html#reading-and-writing-files)
- [Set handlers for asynchronous events, signal](https://docs.python.org/3/library/signal.html)
- [Try and Except](https://pythonbasics.org/try-except/)
