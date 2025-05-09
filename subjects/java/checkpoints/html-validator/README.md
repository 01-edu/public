## HTML Validator

### Instructions

Create a class `HTMLValidator` that provides a method to validate whether a given HTML string is correctly formatted. The method should support a few basic HTML tags (`<html>`, `<body>`, `<div>`, `<p>`, `<b>`, `<i>`, `<h1>`, `<h2>`). If the HTML is valid, the method should return `true`. If the HTML is invalid, the method should return `false`.

### Expected Class

```java
public class HTMLValidator {
    public boolean validateHTML(String html) {
        // Implementation to validate if the given HTML is correctly formatted
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        HTMLValidator validator = new HTMLValidator();

        // Test case 1: Valid HTML
        String html1 = "<html><body><h1>Hello, World!</h1></body></html>";
        System.out.println("Is HTML valid? " + validator.validateHTML(html1)); // Expected output: true

        // Test case 2: Invalid HTML (missing closing tag)
        String html2 = "<html><body><h1>Hello, World!</body></html>";
        System.out.println("Is HTML valid? " + validator.validateHTML(html2)); // Expected output: false

        // Test case 3: Invalid HTML (incorrect nesting)
        String html3 = "<html><body><h1><div>Hello, World!</h1></div></body></html>";
        System.out.println("Is HTML valid? " + validator.validateHTML(html3)); // Expected output: false

        // Test case 4: Valid HTML with multiple tags
        String html4 = "<html><body><div><p>This is a <b>bold</b> word and this is <i>italic</i>.</p></div></body></html>";
        System.out.println("Is HTML valid? " + validator.validateHTML(html4)); // Expected output: true
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Is HTML valid? true
Is HTML valid? false
Is HTML valid? false
Is HTML valid? true
$
```
