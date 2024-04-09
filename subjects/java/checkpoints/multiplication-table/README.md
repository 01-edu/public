## MultiplicationTable

### Instructions

in a file named `MultiplicationTable.java` Write a function `generate` which takes a number as a parameter, and prints the multiplication table for this number.

### Expected Functions

```java
public class MultiplicationTable {
    public static void generate(int num) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        MultiplicationTable.generate(5);
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
5 x 4 = 20
5 x 5 = 25
5 x 6 = 30
5 x 7 = 35
5 x 8 = 40
5 x 9 = 45
5 x 10 = 50
$
```