## files-details

### Instructions

Create a script `files-details.sh`, which will allow you to see the files inside the folder `hard-perms`, this way:

Expected output:

```console
$ ./files-details.sh
dr-------x 2 user user 4096 dez 13 17:50 0
-r------w- 1 user user    0 dez 13 17:51 1
-rw----r-- 1 user user    0 dez 13 17:51 2
drwxrwxrwx 2 user user 4096 dez 13 17:51 3
-r-x--x--- 1 user user    0 dez 13 17:51 4
-r--rw---- 1 user user    0 dez 13 17:51 5
-r--rw---- 1 user user    0 dez 13 17:51 6
-r-x--x--- 1 user user    0 dez 13 17:51 7
-rw----r-- 1 user user    0 dez 13 17:51 8
-r------w- 1 user user    0 dez 13 17:51 9
dr-------x 2 user user 4096 dez 13 17:50 A
```

### Hints

- `ls -l --time-style=TIME_STYLE`

  - Time conversion specifiers you need to know:

    - `%R`, 24-hour hour and minute. Same as ‘%H:%M’.

  - Date conversion specifiers you need to know:

    - `%F`, full date in ISO 8601 format; like ‘%+4Y-%m-%d’ except that any flags or field width override the ‘+’ and (after subtracting 6) the ‘4’. This is a good choice for a date format, as it is standard and is easy to sort in the usual case where years are in the range 0000…9999.

- You can use `sed` to remove the first line of the output.

```console
$ ./files-details.sh        # without using sed
total
-rw-rw-r-- 2022-12-20 03:10:18 README.md
$ ./files-details.sh        # using sed
-rw-rw-r-- 2022-12-20 03:10:18 README.md
$
```

- `awk`.

You can specify specific column names to display or include in the awk output using the field numbers. For example:

```console
$ ls -l | awk '{print $1, $2, $3, $4, $5, $6, $7, $8, $9, $10}' # print all the given fields
total 4
-rw-rw-r-- 1 user user 1989 dez 20 15:19 README.md
$ ls -l | awk '{print $1, $2, $4, $5, $7, $8, $10}'             # print all the given fields
total 4
-rw-rw-r-- 1 user 2350 20 15:25
$
```

awk ‘{print $1}’ emp_records.txt
awk {print $1, $6, $7, $8, $9, $10}'

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [Time conversion specifiers](https://www.gnu.org/software/coreutils/manual/html_node/Time-conversion-specifiers.html).
- [Date conversion specifiers](https://www.gnu.org/software/coreutils/manual/html_node/Date-conversion-specifiers.html).
- [awk](https://www.gnu.org/software/gawk/manual/html_node/Print-Examples.html).
