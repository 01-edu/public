## printcombprog

### Instructions

Write a program that prints in ascending order on a single line all unique combinations of three different digits so that the first digit is lower than the second and the second is lower than the third.

These combinations are separated by a comma and a space.

### Usage

Here is an **incomplete** output :

```console
student@ubuntu:~/piscine-go/printcombprog$ go build
student@ubuntu:~/piscine-go/printcombprog$ ./printcombprog | cat -e
012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789$
student@ubuntu:~/piscine-go/printcombprog$
```

`000` or `999` are not valid combinations because the digits are not different.

`987` should not be shown because the first digit is not less than the second.
