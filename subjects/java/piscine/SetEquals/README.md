## SetEquals

### Instructions

Create a file `SetEquals.java`.

Write a function `areSetsEqual` that returns `true` if the sets as parameters are equal. Returns `false` otherwise.

### Expected Functions

```java
import java.util.Set;

public class SetEquals {
    public static boolean areSetsEqual(Set<String> set1, Set<String> set2) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function:

```java
import java.util.HashSet;
import java.util.Set;
import java.util.TreeSet;

public class ExerciseRunner {

    public static void main(String[] args) {
        Set<String> set1 = new TreeSet<>();
        set1.add("Alice");
        set1.add("Bob");
        set1.add("Charly");

        Set<String> set2 = new TreeSet<>();
        set2.add("Bob");
        set2.add("Charly");
        set2.add("Alice");

        System.out.println(SetEquals.areSetsEqual(set1, set2)); // Expected Output: true

        Set<String> set3 = new HashSet<>();
        set3.add("Alice");
        set3.add("Bob");
        set3.add("Charly");

        Set<String> set4 = new HashSet<>();
        set4.add("Alice");
        set4.add("Bob");
        set4.add("Emily");

        System.out.println(SetEquals.areSetsEqual(set3, set4)); // Expected Output: false
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

[Set](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/List.html)
