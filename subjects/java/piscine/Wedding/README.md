## Wedding

### Instructions

Create a file `Wedding.java`.

Write a function `createCouple` that returns a map of names which associates randomly a name from the first list to a name of the second list.  
If the lists have different sizes, some names from the bigger list will be ignored.

### Expected Functions

```java
import java.util.Set;
import java.util.Map;

public class Wedding {
    public static Map<String, String> createCouple(Set<String> first, Set<String> second) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.util.Set;

public class ExerciseRunner {

    public static void main(String[] args) {
        System.out.println(Wedding.createCouple(Set.of("Pikachu", "Dracaufeu", "Tortank"), Set.of("Legolas", "Aragorn", "Gimli")));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
{Pikachu=Legolas, Tortank=Gimli, Dracaufeu=Aragorn}
$
```

### Notions

[Set](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/List.html)  
[Map](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/Map.html)
