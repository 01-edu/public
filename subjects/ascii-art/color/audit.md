#### Functional

###### Has the requirement for the allowed packages been respected? (Reminder for this project: only [standard packages](https://golang.org/pkg/))

##### Try passing as arguments `"banana" --color red`.

```
Usage: go run . [STRING] [OPTION]

EX: go run . something --color=<color>
```

###### Does it display the same result as above?

##### Try passing as arguments `"hello world" --color=red`.

###### Does it display the expected result?

##### Try passing as arguments `"1 + 1 = 2" --color=green`.

###### Does it display the expected result?

##### Try passing as arguments `"(%&) ??" --color=yellow`.

###### Does it display the expected result?

##### Try specifying a set of letters to be colored (the second until the last letter).

###### Does it display the expected result (the corresponding set of letters with that color)?

##### Try specifying letter to be colored (the second letter).

###### Does it display the expected result (the corresponding letter with that color)?

##### Try specifying a set of letters to be colored(just two letters).

###### Does it display the expected result (the corresponding set of letters with that color)?

##### Try passing as arguments `"HeY GuYs" --color=orange`, in order to color `GuYs`.

###### Does it display the expected result?

##### Try passing as arguments `"RGB()" --color=blue`, in order to color just the B.

###### Does it display the expected result?

##### Try passing as arguments a random string with lower and upper case letters, and a random color in the color flag ("--color=").

###### Does it display the expected result?

##### Try passing as arguments a random string with lower case letters, numbers and spaces, and a random color in the color flag ("--color=").

###### Does it display the expected result?

##### Try passing as arguments a random string with special characters, and a random color in the color flag ("--color="), specifying one letter to be colored.

###### Does it display the expected result?

##### Try passing as arguments a random string with lower case letters, upper case letters, spaces and numbers with a random color in the color flag ("--color="), specifying a set of letters to be colored.

###### Does it display the expected result?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### General

###### +Is it easy/intuitive to specify letter(s) to be colored?

###### +Can you use more than one color in the same string?

#### Basic

###### +Does the project run quickly and effectively (favoring of recursive, no unnecessary data requests, etc.)?

###### +Is the output of the program well structured? Are the characters displayed correctly in line?

###### +Is there a test file for this code?

###### +Are the tests checking each possible case?

###### +Does the code obey the [good practices](../../good-practices/README.md)?

#### Social

###### +Did you learn anything from this project?

###### +Would you recommend/nominate this program as an example for the rest of the school?

#### Bonus

##### Try to use different `--color` flag notations like: `--color=red`, `--color=#ff0000`, `--color=rgb(255, 0, 0)` or `--color=hsl(0, 100%, 50%)`.

###### +Is it possible to use 2 or more color flag notations?
