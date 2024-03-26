## HelloWorld

### Instructions

In a file named `HelloWorld.java` write a function `helloWorld` that returns the string 'Hello World !'.

### Expected Functions

```java
public class HelloWorld {
    public static String helloWorld() {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(HelloWorld.helloWorld());
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Hello World !
$
```