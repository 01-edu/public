## CountChar

### Instructions

In a file named `CountChar.java` write a function `count` that takes two parameters a string and a char, then return the count of that char in the given string.

### Expected Functions

```java
public class CountChar {
    public static int count(String s, char c) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(CountChar.count("Hello World !", ' '));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
2
$
```
