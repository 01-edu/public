## Cat

### Instructions

Create a file named `Cat.java`.

Write a function `cat` that reads from a file given as parameter and writes to the standard output.

### Expected Functions

```java
import java.io.*;

public class Cat {
    public static void cat(String[] args) throws IOException {
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
        System.setOut(printStream);

        Cat.cat(new String[]{"input"});
        String output = outputStream.toString();
        System.out.println(outputStream.toString().equals("test input file\n"));

        Cat.cat(new String[]{});
        String output = outputStream.toString();
        System.out.println(outputStream.toString().equals(""));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
true
true
$
```

### Notions

[Command-Line Arguments](https://docs.oracle.com/javase/tutorial/essential/environment/cmdLineArgs.html)
[File](https://docs.oracle.com/javase/7/docs/api/java/nio/file/Files.html)
[Standard Output](https://docs.oracle.com/javase/7/docs/api/java/io/PrintStream.html)
