## IntOperations

### Instructions

Create a file `IntOperation.java`.

Write a function `addTwoIntegers` that returns the sum of two integers passed as parameter.
Write a function `substractTwoIntegers` that returns subtraction of two integers passed as parameter.
Write a function `multiplyTwoIntegers` that returns the multiplication of two integers passed as parameter.
Write a function `divideTwoIntegers` that returns the Euclidian division of two integers passed as parameter. We ignore the divide by 0 case.

### Expected Functions
```java
public class IntOperation {
    public static int addTwoIntegers(int a, int b) {
        // your code here
    }
    public static int subtractTwoIntegers(int a, int b) {
        // your code here
    }
    public static int multiplyTwoIntegers(int a, int b) {
        // your code here
    }
    public static int divideTwoIntegers(int a, int b) {
        // your code here
    }
}
```

# Usage

Here is a possible ExerciseRunner.java to test your function : 
```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println("Add");
        System.out.println(IntOperation.addTwoIntegers(1, 2));
        System.out.println(IntOperation.addTwoIntegers(15, 68));
        
        System.out.println("Subtract");
        System.out.println(IntOperation.subtractTwoIntegers(4, 1));
        System.out.println(IntOperation.subtractTwoIntegers(10, 23));
        
        System.out.println("Multiply");
        System.out.println(IntOperation.multiplyTwoIntegers(3, 6));
        System.out.println(IntOperation.multiplyTwoIntegers(12, 11));
        
        System.out.println("Divide");
        System.out.println(IntOperation.divideTwoIntegers(8, 2));
        System.out.println(IntOperation.divideTwoIntegers(13, 4));
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
Add
3
83
Subtract
3
-13
Multiply
18
132
Divide
4
3
$ 
```

# Notions
[Primitive Links](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/datatypes.html)  
[Operations](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/op1.html)