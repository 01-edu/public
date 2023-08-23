## CatInFile

### Instructions

Create a file named `CatInFile.java`.

Write a function `cat` that reads from standard input and write to file given as parameter.

> ⚠️ You may read from the input some binary content.

> 💡 Be aware of how much you read at once.

### Expected Functions

```java
import java.io.*;

public class CatInFile {
    public static void cat(String[] args) throws IOException {
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
        String input = "input file test\n";
        ByteArrayInputStream in = new ByteArrayInputStream(input.getBytes());
        System.setIn(in);
        CatInFile.cat(new String[]{"output"});
        String fileContent = new String(Files.readAllBytes(Paths.get("output")));
        System.out.println(fileContent.equals(input));
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
[Standard Input](https://docs.oracle.com/javase/8/docs/api/java/io/InputStream.html)
