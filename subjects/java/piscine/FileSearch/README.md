## FileSearch

### Instructions

Create a file `FileSearch.java`.

Write a function `searchFile` that returns the path of the file in parameter. You need to start searching from a folder named `documents` at the root of the current folder.  
The file to search can be located inside nested subfolders.

### Expected Functions
```java
public class FileSearch {
    public static String searchFile(String fileName) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function
: 
```java
import java.io.IOException;

public class ExerciseRunner {
    public static void main(String[] args) throws IOException {
        System.out.println(FileSearch.searchFile("file.txt"));
    }
}
```

Create the following arborescence : 
```
\ documents
| --- file.csv
| --- \ rep
      | --- example.txt
      | --- file2.csv
      | --- file2.xml
      | --- \ directory33
            | --- file.txt
            | --- test.png
      | --- \ directory22
            | --- test.gif
| --- \ directory435
      | ---test33.xls
          
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
document/rep/directory33/file.txt
$ 
```

### Notions
[File](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/io/File.html)  