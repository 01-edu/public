## Breakdown URL

### Instructions

Create a class `BreakdownURL` that provides a method to parse and validate URLs using regex.The method should extract and return the following URL components: `protocol`, `domain`, `port`, `path` and `query`
Assume the URL is always correct.
The method should extract and return URL components the parameters. The URL is always correct.

> Give back in the map just the existing component.

### Expected Class

```java
public class BreakdownURL {
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
        BreakdownURL parser = new BreakdownURL();

        // Test case 1
        String URL1 = "https://www.example.com:8080/path?name=value";
        Map<String, String> components1 = parser.parseURL(URL1);
        System.out.println("Components of URL 1: " + components1);

        // Test case 2
        String URL2 = "http://example.com/";
        Map<String, String> components2 = parser.parseURL(URL2);
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
