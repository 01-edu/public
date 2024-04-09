## DoOperation

### Instructions

In a file named `DoOperation.java` write a function `operate` that returns the result of the given arithmetic operation specified in the parameters. The arguments should be passed in the following order:

- The first argument is the left operand
- The second argument is the operation sign
- The third argument is the right operand

### Expected Functions

```java
public class DoOperation {
    public static String operate(String[] args) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(DoOperation.operate(new String[]{"1","+","2"}));
        System.out.println(DoOperation.operate(new String[]{"1","-","1"}));
        System.out.println(DoOperation.operate(new String[]{"1","%","0"}));
        System.out.println(DoOperation.operate(args));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
3
0
Error
it depend on the input.
$
```
