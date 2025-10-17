## Backup manager

### Logs and error handling

All scripts will have to use `try` and `except` to handle errors.

The scripts should log (aka write) all actions and errors into a specific file in the `./logs` directory.

- The script `backup_manager.py` will log into `./logs/backup_manager.log`.
- The script `backup_service.py` will log into `./logs/backup_service.log`.

You can choose to phrase your logs as you prefer, however it would be smart to start with a timestamp and try to be as specific as you can.

> Check the examples in the list of arguments for `backup_manager.py` and the usage section to see what the logs should looks like.

### Instructions

You will create two scripts that will manage and perform scheduled backups.

#### First script: backup_manager.py

The script `backup_manager.py` will orchestrate the backup service and if necessary update the file `backup_schedules.txt`. In order to do so it will accept the following command-line arguments:

- `start`: Runs `backup_service.py` in the background.

  - Example of logs:
  - `[dd/mm/yyyy hh:mm] backup_service started`
  - `[dd/mm/yyyy hh:mm] Error: can't start backup_service`
  - `[dd/mm/yyyy hh:mm] Error: backup_service already running`

- `stop`: Kills `backup_service.py` process.

  - Example of logs:
  - `[dd/mm/yyyy hh:mm] backup_service stopped`
  - `[dd/mm/yyyy hh:mm] Error: can't stop backup_service`
  - `[dd/mm/yyyy hh:mm] Error: backup_service not running`

- `create [schedule]`: Adds a new backup schedule in `backup_schedules.txt`.

  - This schedule will have a specific format which is `"file_name;hour:minutes;backup_file_name"`.

  - Example of logs:
    - `[dd/mm/yyyy hh:mm] New schedule added: test2;16:07;office_docs`
    - `[dd/mm/yyyy hh:mm] Error: malformed schedule: test;`

- `list`: Prints the scheduled backups in `backup_schedules.txt`, adding an index before each schedule.

  - Example of logs:
    - `[dd/mm/yyyy hh:mm] Show backups list`
    - `[dd/mm/yyyy hh:mm] Error: can't find backup_schedules.txt`

- `delete [index]`: Delete the backup schedule at line `index` (starting by 0) in `backup_schedules.txt`.

  - Example of logs:
    - `[dd/mm/yyyy hh:mm] Schedule at index 3 deleted`
    - `[dd/mm/yyyy hh:mm] Error: can't find schedule at index 3`
    - `[dd/mm/yyyy hh:mm] Error: can't find backup_schedules.txt`

- `backups`: List the backups files in `./backups`.

  - Example of logs:
    - `[dd/mm/yyyy hh:mm] Show backups list`
    - `[dd/mm/yyyy hh:mm] Error: can't find backups directory`

#### Second script: backup_service.py

The script `backup_service.py` will check the schedules in `backup_schedules.txt` and perform the daily backups at the proper time.

- The service will run in an infinite loop.
- On every cycle of the loop the service will check the schedules.
- The service will perform the backup if the current time (hour and minute) matches the time of the schedule.
- The format of the schedules is `path_to_save;time(hh:mm);backup_name`.
- The backups will be saved in `./backups`.
- The backups will be a compressed file (`.tar`) for the `path_to_save` directory.

> At the end of each loop it would be wise to make the service sleep for about 45 seconds to save processor cycles.

### Usage

Here is a little script to show how the backup system should work (assuming the script is ran at 16:07):

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
[dd/mm/yyyy 16:07] New schedule added: test;16:07;backup_test
[dd/mm/yyyy 16:07] New schedule added: test1;16:07;personal_data
[dd/mm/yyyy 16:07] New schedule added: test2;16:07;office_docs
[dd/mm/yyyy 16:07] Error: malformed schedule: test;
[dd/mm/yyyy 16:07] Show schedules list
[dd/mm/yyyy 16:07] backup_service started
[dd/mm/yyyy 16:07] backup_service stopped
[dd/mm/yyyy 16:07] Show backups list
--> cat on ./logs/backup_service.log
[dd/mm/yyyy 16:07] Backup done for test in backups/backup_test.tar
[dd/mm/yyyy 16:07] Backup done for test1 in backups/personal_data.tar
[dd/mm/yyyy 16:07] Backup done for test2 in backups/office_docs.tar
--> Content of ./backups
backup_test.tar  office_docs.tar  personal_data.tar
```

### Hints

- Familiarize with Python, including modules such as `subprocess`, `shlex` and file `I/O` operations.
- To run a script from another Python script you could use `subprocess.Popen` with the flag `start_new_session=True`.
- To kill a process you should find its process id and call `os.kill`.
- Play with the command `ps -A -f`, it will show a list of all active processes with the arguments attached to them and their process ids.
- The standard library `tarfile` is very useful when trying to work with archives.
- This starts to be quite a big project, try to be as modular as possible by creating single-task functions, this will save you a lot of time.

> There is many ways to run tasks in background and to perform actions on a specific time, the implementation proposed in this exercise is good for learning but should not be your preferred choice for production environments.

**.tar files**

In Linux, tar is a command-line utility used to create, manipulate and extract .tar (tape archive) files. Here are some common operations you can perform with .tar files using the tar command:

- To create a .tar file:
  `tar -cvf archive.tar file1 file2 directory1/`
- To extract the contents of a .tar file:
  `tar -xvf archive.tar`
- To extract a specific file from a .tar file:
  `tar -xvf archive.tar file1`
- To compress a .tar file using gzip:
  `tar -cvzf archive.tar.gz file1 file2 directory1/`
- To list the contents of a tar file:
  `tar tf archive.tar`
- To list the contents of a tar archive in a long listing format:
  `tar -tvf archive.tar`

### References

- [Error handling in Python](https://docs.python.org/3.10/tutorial/errors.html)
- [Spawn a subprocess in Python](https://docs.python.org/3.10/library/subprocess.html)
- [Simple lexical analysis, shlex](https://docs.python.org/3/library/shlex.html)
- [Reading and writing files](https://docs.python.org/3/tutorial/inputoutput.html#reading-and-writing-files)
- [Set handlers for asynchronous events, signal](https://docs.python.org/3/library/signal.html)
- [Try and Except](https://pythonbasics.org/try-except/)
