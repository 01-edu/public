#### Functional

###### Has the requirement for the allowed packages been respected? (Reminder for this project: only [standard packages](https://golang.org/pkg/) )

##### Try passing as arguments `--color red "banana" `.

```
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
```

###### Does it display the same result as above?

##### Try passing as arguments `--color=red "hello world"`.

###### Does it display the expected result?

##### Try passing as arguments `--color=green "1 + 1 = 2"`.

###### Does it display the expected result?

##### Try passing as arguments `--color=yellow "(%&) ??"`.

###### Does it display the expected result?

##### Try specifying a substring to be colored (the second until the last letter).

###### Does it display the expected result (the corresponding substring with that color)?

##### Try specifying letter to be colored (the second letter).

###### Does it display the expected result (the corresponding letter with that color)?

##### Try specifying a substring to be colored (just two letters).

###### Does it display the expected result (the corresponding substring with that color)?

##### Try passing as arguments `--color=orange GuYs "HeY GuYs"`, in order to color `GuYs`.

###### Does it display the expected result?

##### Try passing as arguments `--color=blue B 'RGB()'`, in order to color just the `B`.

###### Does it display the expected result?

##### Try passing as arguments a random string with lower and upper case letters, and a random color in the color flag ("--color=").

###### Does it display the expected result?

##### Try passing as arguments a random string with lower case letters, numbers and spaces, and a random color in the color flag ("--color=").

###### Does it display the expected result?

##### Try passing as arguments a random string with special characters, and a random color in the color flag ("--color="), specifying one letter to be colored.

###### Does it display the expected result?

##### Try passing as arguments a random string with lower case letters, upper case letters, spaces and numbers with a random color in the color flag ("--color="), specifying a substring to be colored .

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

###### +Does the code obey the [good practices](../../../good-practices/README.md)?

#### Social

###### +Did you learn anything from this project?

###### +Would you recommend/nominate this program as an example for the rest of the school?

#### Bonus

##### Try to use different `--color` flag notations like: `--color=red`, `--color=#ff0000`, `--color=rgb(255, 0, 0)` or `--color=hsl(0, 100%, 50%)`.

###### +Is it possible to use 2 or more color flag notations?
