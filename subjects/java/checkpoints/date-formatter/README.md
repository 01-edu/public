## Date Formatter

### Instructions

Write a class `DateFormatter` that has three private attributes:

- `date`: the date to be formatted (as a date).
- `format`: the format to convert the date to.
- `formattedDate`: contains the date converted into the given format.

Create getters for all the attributes and setters for `date` and `format`. The conversion should happen when you set the `format` or the `date`. If the `format` is not correct, the conversion shouldn't happen. Additionally, create methods to output the date in the specified `format`, and constructors with only the `date`, with `date` and `format`, and a default constructor.

#### The accepted date formats are:

- `DD/MM/YYYY`
- `DD Month YYYY`
- `DD.MM.YYYY`

#### The accepted date formats are:

- `DD/MM/YYYY`
- `DD Month YYYY`
- `DD.MM.YYYY`

### Expected Functions

```java
import java.text.SimpleDateFormat;
import java.util.Date;

public class DateFormatter {
    private long timestamp;
    private String formattedDate;
    private String format;

    ...
}
```

### Usage

Here is a possible ExerciseRunner.java to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        DateFormatter df = new DateFormatter(1656374400, "DD/MM/YYYY");
        System.out.println(df.getFormattedDate());

        df.setFormat("dd.MM.yyyy");
        System.out.println(df.getFormattedDate());

        df.setDate(1672531200);
        System.out.println(df.getFormattedDate());

        df.setFormat("DD month yyyy");
        System.out.println(df.getFormattedDate());
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
28/06/2022
28.06.2022
31.12.2022
31 December 2022
$
```
