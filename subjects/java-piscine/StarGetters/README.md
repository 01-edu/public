## StarGetters

### Instructions

Now, we will update the accessibility of the properties.

In order to still have access to them, we need to implement getters and setters for each property :

- `getX` and `setX` for the `x` property
- `getY` and `setY` for the `y` property
- `getZ` and `setZ` for the `z` property
- `getName` and `setName` for the `name` property

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        CelestialObject defaultStar = new CelestialObject();
        System.out.println(defaultStar.getX());
        System.out.println(defaultStar.getY());
        System.out.println(defaultStar.getZ());
        System.out.println(defaultStar.getName());

        defaultStar.setName("Terre");
        defaultStar.setX(0.43);
        defaultStar.setY(0.98);
        defaultStar.setZ(1.43);
        System.out.println(defaultStar.getX());
        System.out.println(defaultStar.getY());
        System.out.println(defaultStar.getZ());
        System.out.println(defaultStar.getName());
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
0.0
0.0
0.0
Soleil
0.43
0.98
1.43
Terre
$
```

### Notions

[Method](https://docs.oracle.com/javase/tutorial/java/javaOO/methods.html)  
[Property](https://docs.oracle.com/javase/tutorial/java/javaOO/variables.html)
