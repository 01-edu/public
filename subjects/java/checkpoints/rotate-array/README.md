## RotateArray

### Instructions

In a file named `RotateArray.java` write a function `rotate` that rotates the given array specified in the parameters by the number given in the second parameter and returns it.

> Note: The rotation count may have a negative value!!!

### Expected Functions

```java
public class RotateArray {
    public static Integer[] rotate(Integer[] arr, int rotationCount) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
import java.io.*;
import java.util.Arrays;

public class ExerciseRunner {
    public static void main(String[] args) throws IOException {
        System.out.println(Arrays.toString(RotateArray.rotate(new Integer[]{4, 2, 1, 3}, 1)));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
[3, 4, 2, 1]
$
```
