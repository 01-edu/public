#### Functional

###### Has the requirement for the allowed packages been respected? (Reminder for this project: only [standard packages](https://golang.org/pkg/))

##### Try passing as arguments `"banana standard --align right"`

```
Usage: go run . [STRING] [BANNER] [OPTION]

EX: go run . something standard --align=right
```

###### Does it display the correct result as above?

##### Try passing as arguments `"left standard --align=right"`

###### Does it display the correct result at the right side?

##### Try passing as arguments `"right standard --align=left"`

###### Does it display the correct result at the left side?

##### Try passing as arguments `"hello shadow --align=center"`

###### Does it display the correct result at the center?

##### Try passing as arguments `""1 Two 4" shadow --align=justify"`

###### Does it display the correct result justified?

##### Try passing as arguments `"23/32 standard --align=right"`

###### Does it display the correct result at the right side?

##### Try passing as arguments `"ABCabc123 thinkertoy --align=right"`

###### Does it display the correct result at the right side?

##### Try passing as arguments `"#$%&" thinkertoy --align=center"`

###### Does it display the correct result at the center?

##### Try passing as arguments `""23Hello World!" standard --align=left"`

###### Does it display the correct result at the left side?

##### Try passing as arguments `""HELLO there HOW are YOU?!" thinkertoy --align=justify"`

###### Does it display the correct result justified?

##### Try passing as arguments `""a -> A b -> B c -> C" shadow --align=right"`

###### Does it display the correct result at the right side?

##### Try reducing the terminal window and run `"abcd shadow --align=right"`

###### Does the representation adapt to the terminal size displaying the right result in the right side?

##### Try reducing the terminal window and run `"ola standard --align=center"`

###### Does the representation adapt to the terminal size displaying the right result in the center?

##### Try passing as arguments a random string with lower and upper case letters, and the align flag ("--align=") followed by a random alignment (left, right, center or justify).

###### Does it display the expected result?

##### Try passing as arguments a random string with lower case letters, numbers and spaces, and the align flag ("--align=") followed by a random alignment (left, right, center or justify).

###### Does it display the expected result?

##### Try passing as arguments a random string with special characters, and the align flag ("--align=") followed by a random alignment (left, right, center or justify).

###### Does it display the expected result?

##### Try passing as arguments a random string with lower, upper case, spaces and numbers letters, and the align flag ("--align=") followed by a random alignment (left, right, center or justify).

###### Does it display the expected result?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### Basic

###### +Does the project runs quickly and effectively (Favoring of recursive, no unnecessary data requests, etc.)?

###### +Is the output of the program well structured? Does any letter seems to be out of line?

###### +Is there a test file for this code?

###### +Are the tests checking each possible case?

###### +Does the code obey the [good practices](../../good-practices/README.md)?

#### Social

###### +Did you learn anything from this project?

###### +Would you recommend/nominate this program as an example for the rest of the school?
