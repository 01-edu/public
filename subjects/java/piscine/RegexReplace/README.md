## RegexReplace

### Instructions

Create a file `RegexReplace.java`.

Write a function `removeUnits` that returns the string where the units `cm` and `€` are removed if they follow directly a number and followed by a space.
Write a function `obfuscateEmail` that returns a string where parts of email addresses are replaced by '*' if they follow the rules below:
- Hide from the username any character next to `-`,  `.` or `_` if they exist. Otherwise, hide 3 characters from the username if its length > 3
- If the remaining part after `@` is in the format `@<third level domain>.<second level domain>.<top level domain>`, then hide the third and top level domains, otherwise hide the second level domain and the top level domain if it is not included in `.com`, `.org` and `.net`.

### Expected Functions
```java
public class RegexReplace {
    public static String removeUnits(String s) {
        // your code here
    }
    
    public static String obfuscateEmail(String s) {
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
        System.out.println(RegexReplace.removeUnits("32cm et 50€"));
        System.out.println(RegexReplace.removeUnits("32 cm et 50 €"));
        System.out.println(RegexReplace.removeUnits("32cms et 50€!"));
        
        System.out.println(RegexReplace.obfuscateEmail("john.doe@example.com"));
        System.out.println(RegexReplace.obfuscateEmail("jann@example.co.org"));
        System.out.println(RegexReplace.obfuscateEmail("jackob@example.fr"));
    }
}
```
          
and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
32 et 50
32 cm et 50 €
32cms et 50€!
l lapin jou à la bel ball avec d animau rigolo pour gagner l bill bleu
john.***@*******.com
jan*@*******.co.***
jac***@*******.**
$ 
```

### Notions
[Pattern](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/regex/Pattern.html)  
