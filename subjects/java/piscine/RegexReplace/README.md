## RegexReplace

### Instructions

Create a file `RegexReplace.java`.

Write a function `removeUnits` that returns the string where the units `cm` and `€` are removed if they follow directly a number and followed by a space.  
Write a function `removeFeminineAndPlural` that returns the string where the mark of feminine and plural are removed from a word : 
 - If a word ends with a mark of plural `s or x`, remove it.
 - If it ends with an `e` or if `e` is followed by a mark of plural `s`, remove both `e and s`.
 - If it ends with `le` following `el`, remove `le`.

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
