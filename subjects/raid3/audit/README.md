#### Raid 3

##### Try running the program: `"./raid1a 3 3 | ./raid3"`

```
[raid1a] [3] [3]
```

###### Does the program returns the value above?

##### Try running the program: `"./raid1b 3 3 | ./raid3"`

```
[raid1b] [3] [3]
```

###### Does the program returns the value above?

##### Try running the program: `"./raid1c 1 1 | ./raid3"`

```
[raid1c] [1] [1] || [raid1d] [1] [1] || [raid1e] [1] [1]
```

###### Does the program returns the value above?

##### Try running the program: `"./raid1e 1 2 | ./raid3"`

```
[raid1c] [1] [2] || [raid1e] [1] [2]
```

###### Does the program returns the value above?

##### Try running the program: `"./raid1e 2 1 | ./raid3"`

```
[raid1d] [2] [1] || [raid1e] [2] [1]
```

###### Does the program returns the value above?

##### Try running the program: `"./raid1c 2 1 | ./raid3"`

```
[raid1c] [2] [1]
```

###### Does the program returns the value above?

##### Try running the program: `"./raid1d 1 2 | ./raid3"`

```
[raid1d] [1] [2]
```

###### Does the program returns the value above?

##### Try running the program: `"./raid1e 1 1 | ./raid3"`

```
[raid1c] [1] [1] || [raid1d] [1] [1] || [raid1e] [1] [1]
```

###### Does the program returns the value above?

##### Try running the program: `"echo 0 0 | ./raid3"`

```
Not a Raid function
```

###### Does the program returns the value above?
