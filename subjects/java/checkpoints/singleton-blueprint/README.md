## Singleton Blueprint

### Instructions

You are given an incomplete singleton class. Complete the class to demonstrate your understanding of how the Singleton design pattern works. The Singleton pattern ensures that a class has only one instance and provides a global point of access to that instance.

### Expected Class

```java
public class Singleton {
    public static Singleton instance;

    private Singleton() {
        // Initialization code
    }

    public static Singleton get???() {
        // Implementation to return the single instance
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        // Get the only object available
        Singleton singleton = Singleton.getInstance();

        // Show the message
        singleton.showMessage();
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Hello, I am a singleton!
$
```
