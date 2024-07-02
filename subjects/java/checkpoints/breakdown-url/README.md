## URL Parsing

### Instructions

Create a class `UrlParser` that provides a method to parse and validate URLs using regex. The method should extract and return url components the `protocol`, `domain`, `port`, `path` and `query` parameters. The URL is always correct.

> give back in the map just the existing component.

### Expected Class

```java
public class UrlParser {
    public Map<String, String> parseURL(String url) {
        // Implementation to parse and validate URLs using regex
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
import java.util.Map;

public class ExerciseRunner {
    public static void main(String[] args) {
        UrlParser parser = new UrlParser();

        // Test case 1
        String url1 = "https://www.example.com:8080/path?name=value";
        Map<String, String> components1 = parser.parseURL(url1);
        System.out.println("Components of URL 1: " + components1);

        // Test case 2
        String url2 = "http://example.com/";
        Map<String, String> components2 = parser.parseURL(url2);
        System.out.println("Components of URL 2: " + components2);

    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Components of URL 1: {protocol=https, domain=www.example.com, port=8080, path=/path, query=name=value}
Components of URL 2: {protocol=http, domain=example.com, path=/}
$
```

This exercise guides the students in creating a method to parse and validate URLs using regex. It includes class definition, method signature, and usage example with expected output.
