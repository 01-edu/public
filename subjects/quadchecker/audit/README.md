#### Quadrangle Checker

> **_NOTE: If you are an admin and you want to test this project, follow the instructions [in the this subject](https://github.com/01-edu/go-tests/blob/master/raid-testing.md) before you proceed to the questions._**

##### Open the repository of the project and check the submitted files

###### Can you confirm that the `"go.mod"` file is present and you can run the program?

##### Try running the program: `"./quadA 3 3 | ./quadchecker"`

```
[quadA] [3] [3]
```

###### Does the program returns the value above?

##### Try running the program: `"./quadB 3 3 | ./quadchecker"`

```
[quadB] [3] [3]
```

###### Does the program returns the value above?

##### Try running the program: `"./quadC 1 1 | ./quadchecker"`

```
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

###### Does the program returns the value above?

##### Try running the program: `"./quadE 1 2 | ./quadchecker"`

```
[quadC] [1] [2] || [quadE] [1] [2]
```

###### Does the program returns the value above?

##### Try running the program: `"./quadE 2 1 | ./quadchecker"`

```
[quadD] [2] [1] || [quadE] [2] [1]
```

###### Does the program returns the value above?

##### Try running the program: `"./quadC 2 1 | ./quadchecker"`

```
[quadC] [2] [1]
```

###### Does the program returns the value above?

##### Try running the program: `"./quadD 1 2 | ./quadchecker"`

```
[quadD] [1] [2]
```

###### Does the program returns the value above?

##### Try running the program: `"./quadE 1 1 | ./quadchecker"`

```
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

###### Does the program returns the value above?

##### Try running the program: `"echo 0 0 | ./quadchecker"`

```
Not a quad function
```

###### Does the program returns the value above?

##### Try running the program: `"echo -n "o--" | ./quadchecker"`

```
Not a quad function
```

###### Does the program returns the value above?

##### Try running the program: `"echo -n "/****" | ./quadchecker"`

```
Not a quad function
```

###### Does the program returns the value above?

##### Try running the program: `"echo -n "ABBB" | ./quadchecker"`

```
Not a quad function
```

###### Does the program returns the value above?

##### Try running the program: `"echo -n "ABBBA"$'\n'"B"$'\n'"B" | ./quadchecker"`

```
Not a quad function
```

###### Does the program returns the value above?

##### Try running the program: `"echo -n "o--o"$'\n'"|"$'\n'"o" | ./quadchecker"`

```
Not a quad function
```

###### Does the program returns the value above?
