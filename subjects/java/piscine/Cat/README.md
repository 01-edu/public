## Cat

### Instructions

Create a file named `Cat.java`.

Write a function `cat` that reads the content of the file given as the argument, and writes it's content to the standard output.

> ⚠️ The files can have some binary content.

> 💡 Be aware of how much you read at once.



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
        PrintStream stdout = System.out;

        ByteArrayOutputStream outputStream = new ByteArrayOutputStream();
        PrintStream printStream = new PrintStream(outputStream);
        System.setOut(printStream);

        Cat.cat(new String[]{"input.txt"});
        String output = outputStream.toString();
        // Reset out to stdout
        System.setOut(stdout);
        System.out.println(output.equals("test input file\n"));

        ByteArrayOutputStream outputStream2 = new ByteArrayOutputStream();
        PrintStream printStream2 = new PrintStream(outputStream2);
        System.setOut(printStream2);

        Cat.cat(new String[]{});
        String output2 = outputStream2.toString();
        // Reset out to stdout
        System.setOut(stdout);
        System.out.println(output2.equals(""));
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
[File](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/nio/file/Files.html)
[Standard Output](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/io/PrintStream.html)
[IOException](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/io/IOException.html)
