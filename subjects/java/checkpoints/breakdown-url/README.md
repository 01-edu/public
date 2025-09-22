## Breakdown URL

### Instructions

Create a class `BreakdownURL` with a method that uses a regular expression to parse a URL and return its components.

**Requirements**

- Extract the following components when present: `protocol`, `domain`, `port`, `path`, `query`.
- Assume the input URL is valid (well-formed).
- Return only the components that actually exist in the given URL (omit absent ones).

> 💡 Have you ever seen a URL without Path ???

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
