## Chifoumi

### Prerequisite

Create a file `ChifoumiAction.java` and paste the following content : 

```java
public enum ChifoumiAction {
    ROCK, PAPER, SCISSOR
}
```

### Instructions

Create a file `Chifoumi.java`.

Write a function `getActionBeatenBy` that returns the action beaten by the action in parameter : 
- `ROCK` will return `SCISSOR`
- `PAPER` will return `ROCK`
- `SCISSOR` will return `PAPER`

Use the `switch` operator.

### Expected Functions
```java
public class Chifoumi {
    public static ChifoumiAction getActionBeatenBy(ChifoumiAction chifoumiAction) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function : 
```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(Chifoumi.getActionBeatenBy(ChifoumiAction.ROCK));
        System.out.println(Chifoumi.getActionBeatenBy(ChifoumiAction.PAPER));
        System.out.println(Chifoumi.getActionBeatenBy(ChifoumiAction.SCISSOR));
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
SCISSOR
ROCK
PAPER
$ 
```

### Notions
[Switch statement](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/switch.html)  
[Enum](https://docs.oracle.com/javase/tutorial/java/javaOO/enum.html)