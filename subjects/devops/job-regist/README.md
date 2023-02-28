## job-regist

### Instruction

Create a script, `job-regist.sh`, that will be able to launch and monitor a specific job.

The script needs to launch a specific binary and track its status saving it into a `job.log` file every second. It must append the new update without modifying the previous content.

Each update needs to be appended to the file with the current format: `YYYY-MM-DD hh:mm:ss - <job status>`

The script should stop running as soon as the job it is tracking ends.

If the script is used with a wrong number of argument it should exit with a status code `1` and print the following error message: `"Error: wrong number of arguments!"`.

### Usage

Here an example of how the script should behave:

```console
$ cat job.sh
sleep 2
$ bash job-regist.sh job.sh
$ cat exp.log
2023-02-28 09:02:09 - [1]+  Running                 ...
2023-02-28 09:02:10 - [1]+  Running                 ...
$
```
