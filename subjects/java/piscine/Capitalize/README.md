## Capitalize

### Instructions

Create a file named `Capitalize.java`.

Write a function `capitalize` that reads the text from a file given as the first parameter and writes the result to a file given as the second parameter.

### Provided files

You can find the [input](input.txt) and its [result](result.txt) files to use for the test and to understand more what you have to do.

### Expected Functions

```java
import java.io.*;

public class Capitalize {
    public static void capitalize(String[] args) throws IOException {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.io.*;
import java.nio.file.Files;
import java.nio.file.Paths;

public class ExerciseRunner {
    public static void main(String[] args) throws IOException {
        Capitalize.capitalize(new String[]{"input", "output"});
        String expectedResult = new String(Files.readAllBytes(Paths.get("result")));
        String userOutput = new String(Files.readAllBytes(Paths.get("output")));
        System.out.println(expectedResult.equals(userOutput));
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
[File](https://docs.oracle.com/javase/7/docs/api/java/nio/file/Files.html)
[String](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/lang/String.html)
