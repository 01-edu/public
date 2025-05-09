## Factorial Master

### Instructions

Create an abstract class `Factorial` that will serve as the base class for calculating the factorial of a number. This class should have an abstract method `calculate` that will be implemented by its child classes.

Implement two child classes:

- `IterativeFactorial` which implements the factorial calculation using an iterative approach.
- `RecursiveFactorial` which implements the factorial calculation using a recursive approach.

> Note: The factorial of 0 is 1, according to the [convention](https://www.chilimath.com/lessons/intermediate-algebra/zero-factorial/)

#### Formula for factorial calculation:

n! = n × (n−1) × (n−2) × … × 1

And an example:
5! = 5 × 4 × 3 × 2 × 1 = 120

### Expected Abstract Class

```java
public abstract class Factorial {
    public abstract long calculate(int n);
}

public class IterativeFactorial extends Factorial {
    @Override
    public long calculate(int n) {
        // iterative factorial calculation
    }
}

public class RecursiveFactorial extends Factorial {
    @Override
    public long calculate(int n) {
        // recursive factorial calculation
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your classes:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        Factorial iterativeFactorial = new IterativeFactorial();
        Factorial recursiveFactorial = new RecursiveFactorial();

        // Test iterative factorial
        int number = 5;
        long iterativeResult = iterativeFactorial.calculate(number);
        System.out.println("Iterative Factorial of " + number + " is: " + iterativeResult); // Expected output: 120

        // Test recursive factorial
        long recursiveResult = recursiveFactorial.calculate(number);
        System.out.println("Recursive Factorial of " + number + " is: " + recursiveResult); // Expected output: 120
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Iterative Factorial of 5 is: 120
Recursive Factorial of 5 is: 120
$
```
