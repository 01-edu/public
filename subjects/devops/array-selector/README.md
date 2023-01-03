## array-selector

### Instructions

Create a script `array-selector.sh`, which will have an array declared with the
following values (in this order): `red`, `blue`, `green`, `white`, `black`.

When executed, the script will try to print the element in the position specified in the first argument passed to the script.

The script should return `Error` when:

- the number of argument passed are different from one,

- the argument passed is not a number, 

- the argument passed is a number out of the array range.

Expected output:

```console
$ ./array-selector.sh 2
blue
$ ./array-selector.sh 5
black
$ ./array-selector.sh 6
Error
$ ./array-selector.sh 
Error
$
```

### Hints

`array` is a special variable type that can hold more than one value. Below an example on how to define an array, select an element in the specific position, and how find out the length of an array. Below an example `script.sh`.

```bash
#!/usr/bin/env bash

array=("one" "two" "three")

echo ${array[0]} # display element in postion 0
echo ${array[2]} # display element in position 2 
echo ${#array[@]} # display length of array
```

And its output:

```console
$ bash script.sh
one
three
3
```

[Bash conditional construct](https://www.gnu.org/software/bash/manual/bash.html#Conditional-Constructs) can be used to decide whether to execute a specific command. Below an example `script.sh`.

```bash
#!/usr/bin/env bash

if [[ 1 > 2 ]]; then
    echo "true"
else 
    echo "false"
fi
```

And its output:

```console
$ bash script.sh 
false
```

It is possible to combine several conditions with the **AND** (`&&`) and **OR** (`||`) logical operators. Below and example `script.sh`.

```bash
if [[ 1 > 2 ]] || [[ 1 == 1 ]]; then
    echo "true"
else
    echo "false"
fi

if [[ 1 > 2 ]] && [[ 1 == 1 ]]; then
    echo "true"
else
    echo "false"
fi
```

And its output:

```console
true
false
```
