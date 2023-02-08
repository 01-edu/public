## job-regist

### Instruction

Create a script, `job-regist.sh`, that will be able to monitor a specific background job.

The script needs to track the status of the first jobs spawned by the current terminal and periodically save the status into a `job.log` file. It must append the new update without modifying the previous content.

Each update needs to be appended to the file with the current format: `YYYY-MM-DD hh:mm:ss - <job status>`

The script should stop running as soon as the job it is tracking ends.

### Usage

Here an example of how the script should behave:

```console
$ sleep 10 &
$ sleep 12 &
$ source job-regist.sh
$ cat job.log
2023-02-08 10:37:50 - [1]+  Running                 sleep 10 &
2023-02-08 10:37:51 - [1]+  Running                 sleep 10 &
2023-02-08 10:37:52 - [1]+  Running                 sleep 10 &
2023-02-08 10:37:53 - [1]+  Running                 sleep 10 &
2023-02-08 10:37:54 - [1]+  Running                 sleep 10 &
2023-02-08 10:37:55 - [1]+  Running                 sleep 10 &
2023-02-08 10:37:56 - [1]+  Running                 sleep 10 &
2023-02-08 10:37:57 - [1]+  Running                 sleep 10 &
2023-02-08 10:37:58 - [1]+  Running                 sleep 10 &
2023-02-08 10:37:59 - [1]+  Running                 sleep 10 &
$
```
