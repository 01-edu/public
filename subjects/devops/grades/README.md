## grades

### Instructions

Create a script called `grades.sh` that takes, as an argument, the number of students to be evaluated. Once the number is specified, the script should prompt you to enter each student’s **name** and **grade**.

Grades must be numeric values only, within the range of 0 to 100.

For each student, the script should return an evaluation based on their grade as follows:

- If the student grade is anything equal or greater than 90 you will return the string `"<name>: You did an excellent job!"`:

```console
$ ./grades.sh 1
Student Name #1: Sara
Student Grade #1: 90
Sara: You did an excellent job!
$
```

- If the student grade is anything equal or greater than 70 you will return the string `"<name>: You did a good job!"`:

```console
$ ./grades.sh 1
Student Name #1: Sara
Student Grade #1: 75
Sara: You did a good job!
$
```

- If the student grade is anything equal or greater than 50 you will return the string `"<name>: You need a bit more effort!"`:

```console
$ ./grades.sh 1
Student Name #1: Sara
Student Grade #1: 51
Sara: You need a bit more effort!
$
```

- If the student grade is anything lower than 50 you will return the string `"<name>: You had a poor performance!"`:

```console
$ ./grades.sh 1
Student Name #1: Sara
Student Grade #1: 34
Sara: You had a poor performance!
$
```

### Error handling

All errors will print a specific message on **stderr** (ending with a newline) and returns a specific non-zero value:

- Wrong number of arguments: `"Error: expect 1 argument only!"`, exit with `1`.

- If the student grade is not a number or is greater than 100: `Error: The grade <grade> is not a valid input. Only numerical grades between 0 and 100 are accepted.`, exit with `1`.

```console
$ ./grades.sh 2
Student Name #1: Sara
Student Grade #1: 101
Error: The grade '101' is not a valid input. Only numerical grades between 0 and 100 are accepted.
$ ./grades.sh 2
Student Name #1: Bob
Student Grade #1: ten
Error: The grade 'ten' is not a valid input. Only numerical grades between 0 and 100 are accepted.
$ ./grades.sh 1 2 3
Error: expect 1 argument only!
$
```

### Usage

```console
$ ./grades.sh 4
Student Name #1: Sara   # introduced by the user
Student Grade #1: 34    # introduced by the user
Student Name #2: Norman
Student Grade #2: 56
Student Name #3: James
Student Grade #3: 78
Student Name #4: Albert
Student Grade #4: 90
Sara: You had a poor performance!
Norman: You need a bit more effort!
James: You did a good job!
Albert: You did an excellent job!
$
```

### Hints

In bash, you can use the `read` command to read input from the user and store it in a variable like so:

`read var_name`

You can also use the `-p` option to specify a prompt that will be displayed to the user before reading the input.

Use loops to go through each student. For example the for loop is very helpful.

```bash
for element in sequence
do
  # commands to be executed
done
```

You can use an array to store multiple pieces of data. To create an array, you can use the `declare` command with the `-a` option:

`declare -a name_array`

To add elements to the array, you can use the `+=` operator:

`name_array+=("Adding this to the array")`

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
