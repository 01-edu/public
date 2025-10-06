## burial

### Instructions

Create a script `burial.sh` which will execute the job `sleep 2` in the background and then prints the list of running jobs using this command `jobs -l | awk '{print $1, $3, $4, $5, $6}'`.

Expected output:

```console
$ ./burial.sh
[1]+ Running sleep 2 &
$
```

### Hints

To run a job in the foreground and then send it to the background, you can use the `fg` and `bg` commands.

The `fg` command allows you to bring a background job to the foreground, so that it becomes the active job in your terminal. The `bg` command allows you to continue running the job in the background.

The `&` symbol is used in job control to run a command in the background as well.

To run a command in the background, simply append an & at the end of the command. For example:

`command &`

This will run the command in the background, allowing you to continue using the terminal while the command is running.

You can use the `jobs` command to list all background jobs that are currently running in your shell. This will show you a list of all background jobs, along with their job numbers and status.

```console
$ jobs
[1]+ Running ping google.com &
[2]- Running sleep 100 &
$ fg %2
sleep 100 &
^Z
[2]+  Stopped                 sleep 100
$ bg %2
[2]+ sleep 100 &
$
```

This will bring the second background job to the foreground. You can use the job number or the job's command line arguments to specify which job to bring to the foreground or send to the background.

Once a job is in the foreground, you can use the `Ctrl + Z` keyboard shortcut to suspend it and send it to the background, or you can use the `Ctrl + C` to finish/close the job.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
