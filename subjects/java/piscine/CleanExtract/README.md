## CleanExtract

### Instructions

Create a file `CleanExtract.java`.

You are given a single string consisting of several substrings; each separated by a `|` character. Unfortunately, some of these substrings contain unnecessary characters that need to be removed in order to form a well-structured output string.

For each substring, you need to extract the portion of the substring which is between the first and last `.`. The leading and trailing white spaces must be removed. The strings should then be joined together, using a space to separate each substring.

Write a function `CleanExtract` to solve this problem.

### Expected Functions

```java
public class CleanExtract {
    public static String extract(String s) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(CleanExtract.extract("The|. quick brown. | what do you ..| .fox .|. Jumps over the lazy dog. ."));
        System.out.println(CleanExtract.extract("  | Who am .I  | .love coding,  |  |.  Coding is fun . | ...  "));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
The quick brown fox Jumps over the lazy dog.
I love coding, Coding is fun .
$
```

### Notions

[String](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/lang/String.html)
[Array](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/arrays.html)
[Conditional statement](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/if.html)
[Loop statement](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/for.html)
