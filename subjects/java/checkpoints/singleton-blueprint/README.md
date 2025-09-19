## Singleton Blueprint

### Instructions

You are given an incomplete `Singleton` class. Complete the class to demonstrate your understanding of how the `Singleton` design pattern works. The `Singleton` pattern ensures that a class has only one instance and provides a global point of access to that instance.

### Expected Class

```java
public class Singleton {
    public Singleton instance;

    private Singleton() {
    }

    public Singleton get???() {
    }

    public String showMessage() {
        return "Hello, I am a singleton!"
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(Singleton.get???().showMessage());
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
