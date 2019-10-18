## printrevcombprog

### Instructions

Write a program that prints in descending order on a single line all unique combinations of three different digits so that the first digit is greater than the second and the second is greater than the third.

These combinations are separated by a comma and a space.

### Usage

Here is an **incomplete** output :

```console
student@ubuntu:~/piscine-go/printcombprog$ go build
student@ubuntu:~/piscine-go/printcombprog$ ./printcombprog | cat -e
987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210$
student@ubuntu:~/piscine-go/printcombprog$
```

`999` or `000` are not valid combinations because the digits are not different.

`789` should not be shown because the first digit is not greater than the second.
