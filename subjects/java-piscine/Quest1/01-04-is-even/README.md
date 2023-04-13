## IsEven

### Instructions

Create a file `IsEven.java`.

Write a function `isEven` that returns true if the int passed in parameter is even, otherwise returns false.

### Expected Functions

```java
public class IsEven {
    public static boolean isEven(int a) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        System.out.println(IsEven.isEven(2));
        System.out.println(IsEven.isEven(26));
        System.out.println(IsEven.isEven(57));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
true
true
false
$
```

# Notions

[Primitive Links](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/datatypes.html)  
[Logic operations](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/op2.html)  
[Operations](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/op1.html)
