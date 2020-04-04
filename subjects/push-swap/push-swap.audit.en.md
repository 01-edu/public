#### Functional

##### Try to run `"./push_swap"`.

###### Does it display nothing?

##### Try to run `"./push_swap 2 1 3 6 5 8"`.

###### Does it display a valid solution and less than 9 instructions?

##### Try to run `"./push_swap 0 1 2 3 4 5"`.

###### Does it display nothing?

##### Try to run `"./push_swap 0 one 2 3"`.

```
Error
```

###### Does it display the right result as above?

##### Try to run `"./push_swap 1 2 2 3"`.

```
Error
```

###### Does it display the right result as above?

##### Try to run `"./push_swap <5 random numbers>"` with 5 random numbers instead of the tag.

###### Does it display a valid solution and less than 12 instructions?

##### Try to run `"./push_swap <5 random numbers>"` with 5 different random numbers instead of the tag.
###### Does it still displays a valid solution and less than 12 instructions?

##### Try to run `"./checker "` and input nothing.

###### Does it display nothing?

##### Try to run `"./checker 0 one 2 3"`.

```
Error
```

###### Does it display the right result as above?

##### Try to run `"echo -e "sa\npb\nrrr\n" | ./checker 0 9 1 8 2 7 3 6 4 5"`.

```
KO
```

###### Does it display the right result as above?

##### Try to run `"echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker 0 9 1 8 2"`.

```
OK
```

###### Does it display the right result as above?

##### Try to run `"ARG=("4 67 3 87 23"); ./push_swap $ARG | ./checker $ARG"`.

```
OK
```

###### Does it display the right result as above?

#### General

##### Try to run `"ARG=("<100 random numbers>"); ./push_swap $ARG"` with 100 random different numbers instead of the tag.

###### +Does it display less than 700 commands?

##### Try to run `"ARG=("<100 random numbers>"); ./push_swap $ARG | ./checker $ARG"` with the same 100 random different numbers as before instead of the tag.

```
OK
```

###### +Does it display the right result as above?

#### Basic

###### +Does the code obey the [good practices](https://public.01-edu.org/subjects/good-practices.en)?

###### +Is there a test file for this code?

###### +Are the tests checking each possible case?

#### Social

###### +Did you learn anything from this project?

###### +Would you recommend/nominate this program as an example for the rest of the school?
