## Armstrong Number

### Instructions

Write a program to check if a given number is an Armstrong number. An Armstrong number of n digits is an integer such that the sum of its digits raised to the power n is equal to the number itself.

Example: 153 is an Armstrong number bacause `1^3 + 5^3 + 3^3 = 153`

### Expected Class

```java
public class ArmstrongNumber {
    public boolean isArmstrong(int number) {
        // Implementation to check if the given number is an Armstrong number
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        ArmstrongNumber armstrongNumber = new ArmstrongNumber();

        // Test case 1
        int number1 = 153;
        System.out.println("Is " + number1 + " an Armstrong number? " + armstrongNumber.isArmstrong(number1)); // Expected output: true

        // Test case 2
        int number2 = 123;
        System.out.println("Is " + number2 + " an Armstrong number? " + armstrongNumber.isArmstrong(number2)); // Expected output: false

        // Test case 3
        int number3 = 9474;
        System.out.println("Is " + number3 + " an Armstrong number? " + armstrongNumber.isArmstrong(number3)); // Expected output: true
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Is 153 an Armstrong number? true
Is 123 an Armstrong number? false
Is 9474 an Armstrong number? true
$
```
