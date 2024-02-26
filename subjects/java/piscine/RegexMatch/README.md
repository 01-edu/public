## RegexMatch

### Instructions

Create a file `RegexMatch.java`.

Write a function `containsOnlyAlpha` that returns `true` if the string as parameter contains only alpha characters.  
Write a function `startWithLetterAndEndWithNumber` that returns `true` if the string as parameter starts with one letter and ends with one number.  
Write a function `containsAtLeast3SuccessiveA` that returns `true` if the string as parameter contains at least 3 successive A.

### Expected Functions
```java
public class RegexMatch {
    public static boolean containsOnlyAlpha(String s) {
        // your code here
    }
    
    public static boolean startWithLetterAndEndWithNumber(String s) {
        // your code here
    }
    
    public static boolean containsAtLeast3SuccessiveA(String s) {
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
        System.out.println(RegexMatch.containsOnlyAlpha("azejkdfhjsdf"));
        System.out.println(RegexMatch.containsOnlyAlpha("azejkd fhjsdf"));
        System.out.println(RegexMatch.startWithLetterAndEndWithNumber("asjd jd34jds jkfd6f5"));
        System.out.println(RegexMatch.startWithLetterAndEndWithNumber("asjd jd34jds jkfd6."));
        System.out.println(RegexMatch.containsAtLeast3SuccessiveA("sdjkAAAAAsdjksj"));
        System.out.println(RegexMatch.containsAtLeast3SuccessiveA("sdjkAAsdaaasdjksj"));
    }
}
```
          
and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
true
false
true
false
true
false
$ 
```

### Notions
[Pattern](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/regex/Pattern.html)  