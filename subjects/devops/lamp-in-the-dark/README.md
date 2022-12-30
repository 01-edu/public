## lamp-in-the-dark

### Instructions

Create a script `lamp-in-the-dark.sh` which will list all background jobs that are currently running in your shell.

Add the following to your script:

```console
sleep 4 &
sleep 3 &
sleep 2 &
```

Expected output:

```console
# The processes running here are just an example
$ ./lamp-in-the-dark.sh
[1]   Running                 sleep 4 &
[2]-  Running                 sleep 3 &
[3]+  Running                 sleep 2 &
$
```

### Hints

You can use the `jobs` command to list all background jobs that are currently running in your shell. This will show you a list of all background jobs, along with their job numbers and status.

```console
$ jobs
[1]+ Running ping google.com &
[2]- Running sleep 100 &
```

- The flag `-l` show status and process IDs of all jobs.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
