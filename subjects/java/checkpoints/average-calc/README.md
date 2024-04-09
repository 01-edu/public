## AverageCalc

### Instructions

Write a function named `average` in a file named `AverageCalc` java. This function takes three parameters: start, end, and step. It returns the average of the numbers between start and end with each number being greater than the previous one by the step given."

### Expected Functions

```java
public class AverageCalc {
    public static int average(int start, int end, int step) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(MultiplicationTable.generate(1,5,1));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
3
$
```
