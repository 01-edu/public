## StarInheritance

### Instructions

Create a new class `Star` in a file named `Star.java`.

This class inherits from `CelestialObject`. We add a private double property `magnitude`, with public getter and setter.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        Star star = new Star();
        
        System.out.println(star.getName());
        System.out.println(star.getX());
        System.out.println(star.getY());
        System.out.println(star.getZ());
        System.out.println(star.getMagnitude());
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
Soleil
0.0
0.0
0.0
0.0
$ 
```

### Notions

[Subclasses](https://docs.oracle.com/javase/tutorial/java/IandI/subclasses.html)  