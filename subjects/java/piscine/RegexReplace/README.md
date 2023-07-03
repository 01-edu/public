## RegexReplace

### Instructions

Create a file `RegexReplace.java`.

Write a function `removeUnits` that returns the string where the units `cm` and `€` is removed if it followed directly a number and followed by a space.  
Write a function `removeFeminineAndPlural` that returns the string where the mark of feminine and plural is removed from word : 
 - if a word ends a mark of plural (with s or x), remove it.
 - if it ends with an e or if the e is followed by the mark of plural, remove the it
 - if it ends with `le` following `el` (or if `le` is followed by plural), remove it.

### Expected Functions
```java
public class RegexReplace {
    public static String removeUnits(String s) {
        // your code here
    }
    
    public static String removeFeminineAndPlural(String s) {
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
        System.out.println(RegexReplace.removeFeminineAndPlural("le lapin joue à la belle balle avec des animaux rigolos pour gagner les billes bleues"));
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
$ 
```

### Notions
[Pattern](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/regex/Pattern.html)  
