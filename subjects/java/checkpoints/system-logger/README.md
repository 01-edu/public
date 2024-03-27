## SystemLogger

### Instructions

In a file named `SystemLog.java` write a function `systemLog` that takes a String message as a parameter and returns the given message concatenated with 'System Log: '

### Expected Functions

```java
public class SystemLog {
    public static String systemLog(String message) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(SystemLog.systemLog('message'));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
System Log: message
$
```
