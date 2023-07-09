## MapEquals

### Instructions

Create a file `MapEquals.java`.

Write a function `areMapsEqual` that returns `true` if the maps as parameters are equal. Returns `false` otherwise.

### Expected Functions

```java
import java.util.Map;

public class MapEquals {
    public static boolean areMapsEqual(Map<String, Integer> map1, Map<String, Integer> map2) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function:

```java
import java.util.Map;

public class ExerciseRunner {
    public static void main(String[] args) {
        Map<String, Integer> map1 = Map.of("Alice", 1, "Bob", 2, "Charly", 3);
        Map<String, Integer> map2 = Map.of("Alice", 1, "Bob", 2, "Charly", 3);
        System.out.println(MapEquals.areMapsEqual(map1, map2)); // Expected Output: true

        Map<String, Integer> map3 = Map.of("Alice", 1, "Bob", 2, "Charly", 3);
        Map<String, Integer> map4 = Map.of("Alice", 1, "Bob", 2, "Emily", 3);
        System.out.println(MapEquals.areMapsEqual(map3, map4)); // Expected Output: false
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

[Map](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/Map.html)
