## ListEquals

### Instructions

Create a file `ListEquals.java`.

Write a function `areListsEqual` that returns true if the lists in parameters are equal. Returns false otherwise.

### Expected Functions

```java
import java.util.List;

public class ListEquals {
    public static boolean areListsEqual(List<String> list1, List<String> list2) {
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
        System.out.println(ListEquals.areListsEqual(List.of("Alice", "Bob", "Charly", "Emily"), List.of("Alice", "Bob", "Charly", "Emily")));
        System.out.println(ListEquals.areListsEqual(List.of("Alice", "Bob", "Charly", "Emily"), List.of("Alice", "Bob", "Emily", "Charly")));
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
