## Factorial

### Instructions

In a file named `Factorial.java` write a function `factorial` that calculates the factorial of a non-negative integer given as a parameter.

> The factorial of a non-negative integer n (denoted as n!) is the product of all positive integers less than or equal to n. For example, 5! = 5 × 4 × 3 × 2 × 1 = 120.

### Expected Functions

```java
public class Factorial {
    public static Integer factorial(Integer n) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(Factorial.factorial(3));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
6
$
```
