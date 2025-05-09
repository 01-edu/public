## SortArgs

### Instructions

Create a file named `SortArgs.java`.

Write a function `sort` that sorts the given array specified in the parameters and prints it to the standard output. The elements should be separated by spaces followed by a a new line character. All the given elements are valid numbers.


### Expected Functions

```java
public class SortArgs {
    public static void sort(String[] args) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.io.*;

public class ExerciseRunner {
    public static void main(String[] args) throws IOException {
        ByteArrayOutputStream outputStream = new ByteArrayOutputStream();
        PrintStream printStream = new PrintStream(outputStream);

        var defaultOut = System.out;

        System.setOut(printStream);
        SortArgs.sort(new String[]{"4", "2", "1", "3"});
        System.setOut(defaultOut);

        String output = outputStream.toString();
        System.out.println(output.equals("1 2 3 4\n"));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
true
$
```

### Notions

[Command-Line Arguments](https://docs.oracle.com/javase/tutorial/essential/environment/cmdLineArgs.html)
[Conditional statement](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/if.html)
[String](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/lang/String.html)
[Array](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/arrays.html)
[File](https://docs.oracle.com/javase/7/docs/api/java/nio/file/Files.html)
