## Strlen

### Instructions

In a file named `Strlen.java` write a function `strlen` that returns the length of the string passed as a parameter.

### Expected Functions

```java
public class Strlen {
    public static int strlen(String s) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(Strlen.strlen("Hello World !"));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
13
$
```
