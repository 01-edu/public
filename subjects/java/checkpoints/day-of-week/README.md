## Day Of Week

### Instructions

Create a class `DayOfWeekFinder` that provides a method to find the date of the next specified day of the week from a given date. The date will be provided in the format `yyyy-MM-dd`, and the day of the week will be provided as a string (e.g., "Monday", "Tuesday").

### Expected Class

```java
import java.time.DayOfWeek;
import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;

public class DayOfWeekFinder {
    public String findNextDayOfWeek(String startDate, String dayOfWeek) {
        ...
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        DayOfWeekFinder finder = new DayOfWeekFinder();

        // Test case 1
        String startDate1 = "2023-06-22";
        String dayOfWeek1 = "Monday";
        System.out.println("Next " + dayOfWeek1 + " after " + startDate1 + ": " + finder.findNextDayOfWeek(startDate1, dayOfWeek1));

        // Test case 2
        String startDate2 = "2023-06-22";
        String dayOfWeek2 = "Friday";
        System.out.println("Next " + dayOfWeek2 + " after " + startDate2 + ": " + finder.findNextDayOfWeek(startDate2, dayOfWeek2));

        // Test case 3
        String startDate3 = "2023-06-22";
        String dayOfWeek3 = "Sunday";
        System.out.println("Next " + dayOfWeek3 + " after " + startDate3 + ": " + finder.findNextDayOfWeek(startDate3, dayOfWeek3));

        // Error case: invalid date format
        String startDate4 = "invalid-date";
        String dayOfWeek4 = "Monday";
        System.out.println("Next " + dayOfWeek4 + " after " + startDate4 + ": " + finder.findNextDayOfWeek(startDate4, dayOfWeek4));

        // Error case: invalid day of the week
        String startDate5 = "2023-06-22";
        String dayOfWeek5 = "Funday";
        System.out.println("Next " + dayOfWeek5 + " after " + startDate5 + ": " + finder.findNextDayOfWeek(startDate5, dayOfWeek5));
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Next Monday after 2023-06-22: 2023-06-26
Next Friday after 2023-06-22: 2023-06-23
Next Sunday after 2023-06-22: 2023-06-25
Next Monday after invalid-date: Error
Next Funday after 2023-06-22: Error
$
```
