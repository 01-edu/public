## Builder Blueprint

### Instructions

You are given an incomplete Builder design pattern implementation with some incorrect parts. Complete and fix the classes to demonstrate your understanding of how the Builder design pattern works.

Regex component required:

- Any Character: `.`
- Digit: `\d`
- Whitespace : `\s`
- Word Character: `\w`

### Expected Classes

```java
// Regex class
public class Regex {
    private StringBuilder pattern;

    public Regex(List<String> component) {
       ...
    }

    public String getPattern() {
        return pattern.toString();
    }
}

// Builder interface
public ... RegexBuilder {
    void buildLiteral(String literal);
    void buildAnyCharacter();
    void buildDigit();
    void buildWhitespace();
    void buildWordCharacter();
    Regex getResult();
}

// ConcreteRegexBuilder class
public class ConcreteRegexBuilder {
    private Regex regex = new Regex();
    List<String> component;

    ...

    @Override
    public Regex getResult() {
        // Missing return statement
    }
}

// RegexDirector class
public class RegexDirector {
    private RegexBuilder builder;

    public void setBuilder(RegexBuilder builder) {
        this.builder = builder;
    }

    public Regex construct() {
        builder.buildLiteral("Hello");
        builder.buildWhitespace();
        builder.buildWordCharacter();
        builder.buildAnyCharacter();
        return builder.getResult();
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your classes:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        RegexDirector director = new RegexDirector();
        RegexBuilder builder = new ConcreteRegexBuilder();

        director.setBuilder(builder);
        Regex regex = director.construct();

        System.out.println(regex);
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Hello\s\w.
$
```
