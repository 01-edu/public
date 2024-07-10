## Age Finder

### Instructions

Create a class `AgeFinder` that provides a method to calculate the age from a given date. The date will be provided in the format `yyyy-MM-dd`.

In case of any error the method `calculateAge` should return `-1`

> 💡 Have you ever seen someone with a negative age???

### Expected Class

```java
import java.time.LocalDate;
import java.time.Period;
import java.time.format.DateTimeFormatter;

public class AgeFinder {
    public int calculateAge(String date) {
        // Implementation to calculate the age from the given date
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        AgeFinder AgeFinder = new AgeFinder();

        // Test case 1
        String date1 = "2000-01-01";
        System.out.println("Age: " + AgeFinder.calculateAge(date1));

        // Test case 2
        String date2 = "1990-06-15";
        System.out.println("Age: " + AgeFinder.calculateAge(date2));

        // Test case 3
        String date3 = "2010-12-25";
        System.out.println("Age: " + AgeFinder.calculateAge(date3));
    }
}
```

### Expected Output

```shell
$ date
Tue Jul  9 03:58:06 PM UTC 2024
$ javac *.java -d build
$ java -cp build ExerciseRunner
Age: 24
Age: 34
Age: 13
$
```
