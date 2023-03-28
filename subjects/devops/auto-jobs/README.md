## auto-jobs

### Instructions

In this exercise, you will create four files called `task1`, `task2`, `task3`, and `task4`, which will contain the necessary formulas to set up a group of scheduled tasks in the group

##### task1:

Time: `Every Friday, at 20:41`
Command: `echo "01"`

##### task2:

Time: `Every minute`
Command: `sh /home/user01/check`

##### task3:

Time: `Everyday midnight at 12 AM`
Command: `sh /home/user01/backup`

##### task4:

Time: `After Every Reboot`
Command: `01exec`

All file content must be in this format:

```console
$ cat task-example | cat -e
* * * * * {command}$
$
```

Once it is done, use the command below to create the file `auto-jobs.tar` to be submitted.

```console
$ tar -cf auto-jobs.tar task1 task2 task3 task4
$ ls
task1 task2 task3 task4 auto-jobs.tar
```

**Only `auto-jobs.tar` should be submitted.**

### Hints

Linux Cron utility is an effective way to schedule a routine background job at a specific time and/or day on an on-going basis.

You can use the `crontab` command to manage your jobs. This command can be called in four different ways:

`crontab -l`: List the jobs for the current user
`crontab -r`: Remove all jobs for the current users.
`crontab -e`: Edit jobs for the current user.

##### Linux Crontab Format:

`MIN HOUR DOM MON DOW CMD`

##### Table: Crontab Fields (Linux Crontab Syntax):

| Field | Description  |               Allowed Value |
| ----: | ------------ | --------------------------: |
|   MIN | Minute field |                     0 to 59 |
|  HOUR | Hour field   |                     0 to 23 |
|   DOM | Day of month |                        1-31 |
|   MON | Month field  |                        1-12 |
|   DOW | Day Of Week  |                         0-6 |
|   CMD | Command      | Any command to be executed. |

##### Examples:

1. Schedule a job every hour at the fifth minute, every day: `5 * * * * {command}`

2. Schedule a job 5 minutes after midnight every day: `5 0 * * * {command}`

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [an editor for cron schedule expressions](https://crontab.guru/).
