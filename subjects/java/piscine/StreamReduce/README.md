## StreamReduce

### Instructions

Create a file `StreamReduce.java`.

Create a function `sumAll` which returns the sum of integers in the stream.  
Create a function `divideAndAddElements` which sums the result of the division, of all the integers in the stream, by the divider.  


### Expected Functions
```java
public class StreamReduce {
    public static Integer sumAll(Stream<Integer> s) {
        // your code here
    }

    public static Integer divideAndAddElements(Stream<Integer> s, int divider) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.io.IOException;
import java.util.stream.Stream;

public class ExerciseRunner {
    public static void main(String[] args) throws IOException {
        System.out.println(StreamReduce.sumAll(Stream.of(3, 5, 7, 10)));
        System.out.println(StreamReduce.sumAll(Stream.of()));
        System.out.println(StreamReduce.divideAndAddElements(Stream.of(3, 5, 7, 10), 2));
        System.out.println(StreamReduce.divideAndAddElements(Stream.of(), 2));
    }
}
```
          
and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
25
0
11
0
$ 
```

### Notions
[Stream](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/stream/Stream.html)  
[Reduce](https://www.baeldung.com/java-stream-reduce)  
