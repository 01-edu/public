#### Functional

###### Has the requirement for the allowed packages been respected? (Reminder for this project: only [standard packages](https://golang.org/pkg/))

> For consistency, use bash for the following tests.

##### Try passing as arguments `--align right something standard`

```
Usage: go run .  [OPTION] [STRING] [BANNER]

Example: go run . --align=right  something  standard
```

###### Does it display the same result as above?

##### Try passing as arguments `--align=right left standard`

###### Does it display the correct result at the right side?

##### Try passing as arguments `--align=left right standard `

###### Does it display the correct result at the left side?

##### Try passing as arguments `--align=center hello shadow`

###### Does it display the correct result at the center?

##### Try passing as arguments `--align=justify "1 Two 4" shadow`

###### Does it display the correct result justified?

##### Try passing as arguments `--align=right 23/32 standard`

###### Does it display the correct result at the right side?

##### Try passing as arguments `--align=right ABCabc123 thinkertoy`

###### Does it display the correct result at the right side?

##### Try passing as arguments `--align=center "#$%&\"" thinkertoy`

###### Does it display the correct result at the center?

##### Try passing as arguments `--align=left '23Hello World!' standard `

###### Does it display the correct result at the left side?

##### Try passing as arguments `--align=justify 'HELLO there HOW are YOU?!' thinkertoy`

###### Does it display the correct result justified?

##### Try passing as arguments `--align=right "a -> A b -> B c -> C" shadow `

###### Does it display the correct result at the right side?

##### Try reducing the terminal window and run `--align=right abcd shadow `

###### Does the representation adapt to the terminal size displaying the right result in the right side?

##### Try reducing the terminal window and run `--align=center ola standard `

###### Does the representation adapt to the terminal size displaying the right result in the center?

##### Try passing as arguments a random string with lower and upper case letters, and the align flag ("--align=") followed by a random alignment (left, right, center or justify).

###### Does it display the expected result?

##### Try passing as arguments a random string with lower case letters, numbers and spaces, and the align flag ("--align=") followed by a random alignment (left, right, center or justify).

###### Does it display the expected result?

##### Try passing as arguments a random string with special characters, and the align flag ("--align=") followed by a random alignment (left, right, center or justify).

###### Does it display the expected result?

##### Try passing as arguments a random string with lower, upper case, spaces and numbers letters, and the align flag ("--align=") followed by a random alignment (left, right, center or justify).

###### Does it display the expected result?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid Compilation, Cheating, Crashing, Leaks)

#### Basic

###### +Does the project run quickly and effectively (Favoring of recursive, no unnecessary data requests, etc.)?

###### +Is the output of the program well structured? Are the characters are correctly in line?

###### +Is there a test file for this code?

###### +Are the tests checking each possible case?

###### +Does the code obey the [good practices](../../../good-practices/README.md)?

#### Social

###### +Did you learn anything from this project?

###### +Would you recommend/nominate this program as an example for the rest of the school?
