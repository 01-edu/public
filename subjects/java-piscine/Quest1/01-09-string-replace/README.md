## StringReplace

### Instructions

Create a file `StringReplace.java`.

Write a function called `StringReplace` that takes in three parameters: `original_string s`, `target`, and `replacement`. The function should replace all occurrences of target in `original_string s` with `replacement` and return the modified string.

### Expected Functions
```java
public class StringReplace {
    public static String replace(String s, target, replacement) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function : 
```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(StringContain.replace("javatpoint is a very good website", 'a', 'e'));
        System.out.println(StringContain.replace("my name is khan my name is java", "is","was"));
        System.out.println(StringContain.replace("hey i'am java", "I'am","was"));
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
jevetpoint is e very good website
my name was khan my name was java
hey i'am java
$ 
```

### Notions
[String](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/lang/String.html)