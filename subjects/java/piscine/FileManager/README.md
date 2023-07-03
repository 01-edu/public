## FileManager

### Instructions

Create a file `FileManager.java`.

Write a function `createFile` that creates a file with the name and the content as parameter.
Write a function `getContentFile` that returns the content of the file as parameter.  
Write a function `deleteFile` that deletes the file as parameter.

### Expected Functions
```java
public class FileManager {
    public static void createFile(String fileName, String content) throws IOException {
        // your code here
    }
    public static String getContentFile(String fileName) throws IOException {
        // your code here
    }
    public static void deleteFile(String fileName) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function : 
```java
import java.io.IOException;

public class ExerciseRunner {
    public static void main(String[] args) throws IOException {
        FileManager.createFile("file.txt", "Lorem ipsum");
        System.out.println(FileManager.getContentFile("file.txt"));
        FileManager.deleteFile("file.txt");
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
Lorem ipsum
$ 
```

### Notions
[File](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/io/File.html)  