## TimeTracker

### Instructions

Write a `ProjectTime` class to track the duration of a project using start and end times. You will be asked to add getter and setter methods for the `startTime` and `endTime` attributes. When you set one of them, you must update the `hoursLogged` attribute. The `hoursLogged` getter should output the hours in the following format:

- Less than 120 minutes: `hoursLogged` m
- Less than 120 hours: `hoursLogged` h
- Less than 120 days: `hoursLogged` d
- More than 120 days: `hoursLogged` mo

If there is any error, the number of `hoursLogged` should be handled accordingly.

> 💡 The `hoursLogged` shouldn't be a negative number except -1 in case of errors.

### Expected Functions

```java
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class ProjectTime {
    private String startTime;
    private String endTime;
    private float hoursLogged;

    public ProjectTime(String start, String end);

    public void setStartTime();
    public void setEndTime();

    public String getStartTime();
    public String getEndTime();

    public String getHoursLogged();
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {

        ProjectTime shortProject = new ProjectTime("2023-05-14 09:00", "2023-05-14 09:30");
        System.out.println("Short Project Total Logged Time: " + shortProject.getHoursLogged());

        ProjectTime overnightProject = new ProjectTime("2023-05-14 20:00", "2023-05-15 08:00");
        System.out.println("Overnight Project Total Logged Time: " + overnightProject.getHoursLogged());

        ProjectTime fullDayProject = new ProjectTime("2023-05-14 00:00", "2023-05-15 00:00");
        System.out.println("Full Day Project Total Logged Time: " + fullDayProject.getHoursLogged());

        ProjectTime errorProject = new ProjectTime("2023-05-14", "2023-05-15 08:00");
        System.out.println("Error Project Total Logged Time: " + errorProject.getHoursLogged());
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Short Project Total Logged Time: 30 m
Overnight Project Total Logged Time: 12 h
Full Day Project Total Logged Time: 24 h
Error Project Total Logged Time: -1
$
```

### Tip — How To Use SimpleDateFormat

```java
// Creation of the date parser
SimpleDateFormat DATE_FORMAT = new SimpleDateFormat("yyyy-MM-dd HH:mm");

// parse and get the date
Date date = DATE_FORMAT.parse(startTime);

// getTime in milisecond
Long timeInMilisecond = date.getTime();
```
