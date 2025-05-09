## Perfect Number

### Instructions

Write a program to check if a given number is a perfect number. A perfect number is a positive integer that is equal to the sum of its proper positive divisors, excluding itself.

### Expected Class

```java
public class PerfectNumber {
    public boolean isPerfectNumber(int number) {
        // Implementation to check if the given number is a perfect number
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        PerfectNumber perfectNumber = new PerfectNumber();

        int number1 = 6;
        System.out.println("Is " + number1 + " a perfect number? " + perfectNumber.isPerfectNumber(number1));

        int number2 = 28;
        System.out.println("Is " + number2 + " a perfect number? " + perfectNumber.isPerfectNumber(number2));

        int number3 = 12;
        System.out.println("Is " + number3 + " a perfect number? " + perfectNumber.isPerfectNumber(number3));
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Is 6 a perfect number? true
Is 28 a perfect number? true
Is 12 a perfect number? false
$
```
