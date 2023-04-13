## FloatOperations

### Instructions

Create a file `FloatOperation.java`.

Write a function `divideTwoFloats` that returns the float division of two floats passed as parameter.
Write a function `addTwoFloats` that returns the sum of two floats passed as parameter.

### Expected Functions

```java
public class FloatOperation {
    public static float addTwoFloats(float a, float b) {
        // your code here
    }
    public static float divideTwoFloats(float a, float b) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println("Add");
        System.out.println(FloatOperation.addTwoFloats(1.0f, 2.5f));
        System.out.println(FloatOperation.addTwoFloats(15.18f, 68.28347f));

        System.out.println("Divide");
        System.out.println(FloatOperation.divideTwoFloats(7.0f, 2.0f));
        System.out.println(FloatOperation.divideTwoFloats(5.6f, 6.9f));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Add
3.5
83.46347
Divide
3.5
0.8115942
$
```

# Notions

[Primitive Links](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/datatypes.html)  
[Operations](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/op1.html)
