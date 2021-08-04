## linear-stats

This time, like in the first exercise, you will have to build a program that prints various statistical calculations. Now, you will focus on the `Linear Regression Line` and the `Pearson Correlation Coefficent`.

### Instructions

Your program must be able to read from a file and print the result of each statistic mentioned above. In other words, your program must be able to read the data present in the path passed as argument. The data in the file will be presented as below:

```console
189
113
121
114
145
110
...
```

This data represents a graph in which the values of the x axis are the number of the lines (0, 1, 2, 3, 4, 5 ...) and the values of the y axis are the actual numbers (189, 113, 121, 114, 145, 110...).

To run your program, a command similar to this one will be used if your project is made in Go:

```sh
>$ go run your-program.go data.txt
```

After reading the file, your program should print the `Linear Regression Line` and the `Pearson Correlation Coefficient` in the following format:

```console
Linear Regression Line: y = <value>x + <value>
Pearson Correlation Coefficient: <value>
```

The values in between the single angle quotation marks (`< >`) should be a decimal number. The values for the `Linear Regression Line` should have 6 decimal places, while the `Pearson Correlation Coefficient` value should have 10 decimal places.

#### Testing

Your program will be tested by an auditor who will run a program provided by us, that creates a random data set of numbers and prints the result. The auditor task is to compare how your program performed.

You can choose the language in which you want to build your program.

This program will help you learn about:

- Statistical and Probabilities Calculations
  - [Linear Regression Line](https://en.wikipedia.org/wiki/Linear_regression)
  - [Pearson Correlation Coefficient](https://en.wikipedia.org/wiki/Pearson_correlation_coefficient)
