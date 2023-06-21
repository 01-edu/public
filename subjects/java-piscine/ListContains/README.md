## ListContains

### Instructions

Create a file `ListContains.java`.

Write a function `containsValue` that returns true if the list as parameter contains the element as parameter, false otherwise.

### Expected Functions

```java
import java.util.List;

public class ListContains {
    public static boolean containsValue(List<Integer> list, Integer value) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.util.List;

public class ExerciseRunner {

    public static void main(String[] args) {
        System.out.println(ListContains.containsValue(List.of(9, 13, 8, 23, 1, 0, 89), 8));
        System.out.println(ListContains.containsValue(List.of(9, 13, 8, 23, 1, 0, 89), 10));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
true
false
$
```

### Notions

[List](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/List.html)
