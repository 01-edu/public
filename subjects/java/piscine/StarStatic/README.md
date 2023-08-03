## StarStatic

### Instructions

Let's add some computation.

Our objective is to compute the distance between celestial objects. As you may have guessed, the values x, y and z are the coordinates of the object. Their unit is AU (Astronomical Unit) which is 150 000 000 km.

We will add two `static` and `public` methods :

- `getDistanceBetween` which takes two CelestialObject as parameters and returns a double corresponding to the distance between the two objects.
- `getDistanceBetweenInKm` which takes two CelestialObject as parameters and returns a double corresponding to the distance in km between the two objects.

We add a public constant double property, named `KM_IN_ONE_AU` with the value of 150 000 000.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        CelestialObject defaultStar = new CelestialObject();
        CelestialObject earth = new CelestialObject("Terre", 1.0, 2.0, 2.0);

        System.out.println(CelestialObject.getDistanceBetween(defaultStar, earth));
        System.out.println(CelestialObject.getDistanceBetweenInKm(defaultStar, earth));
        System.out.println(CelestialObject.KM_IN_ONE_AU);
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
3.0
4.5E8
1.5E8
$
```

### Notions

[Class Member / Static](https://docs.oracle.com/javase/tutorial/java/javaOO/classvars.html)  
[Math](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/lang/Math.html)
