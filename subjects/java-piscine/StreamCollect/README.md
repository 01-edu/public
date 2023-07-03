## StreamCollect

### Instructions

Create a file `StreamCollect.java`.

Create a function `mapByFirstLetter` which maps the first letter (to upper case) to the String.  
Create a function `getMaxByModulo4` which return a map associating the modulo 4 to the max of the integer mathching the modulo 4.  
##Create a function `orderAndConcatWithSharp` which joins the alphabetically sorted Strings with a ' # ' between each. The result String starts with a '{' and ends with a '}'.

### Expected Functions

```java
public class StreamCollect {
    public static Map<Object, List<String>> mapByFirstLetter(Stream<String> s) {
        // your code here
    }

    public static Map<Integer, Optional<Integer>> getMaxByModulo4(Stream<Integer> s) {
        // your code here
    }

    public static String orderAndConcatWithSharp(Stream<String> s) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.util.stream.Stream;

public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(StreamCollect.mapByFirstLetter(Stream.of("Bonjour", "le", "monde !", "bonsoir")));
        System.out.println(StreamCollect.getMaxByModulo4(Stream.of(5, 12, 32, 4, 9, 17, 98, 424, 97, 5843, 48354)));
        System.out.println(StreamCollect.orderAndConcatWithSharp(Stream.of("Hello", "how are you ?", "where do you live ?", "Bordeaux")));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
{B=[Bonjour, bonsoir], L=[le], M=[monde !]}
{0=Optional[424], 1=Optional[97], 2=Optional[48354], 3=Optional[5843]}
##{Bordeaux # Hello # how are you ? # where do you live ?}
$
```

### Notions

[Stream](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/stream/Stream.html)
