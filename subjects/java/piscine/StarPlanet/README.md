## StarPlanet

### Instructions

Create a new class `Planet` in a file named `Planet.java`.

This class inherits `CelestialObject`.  
It has one private property : `centerStar` of type `Star`.

There are two constructors :

- One with no parameters, which calls the superclass constructor with no parameters. In this case, the `centerStar` property is initialised with the default Star constructor.
- The other with many parameters :
  - `name`
  - `x`
  - `y`
  - `z`
  - `centerStar`
    which calls the superclass constructor with full parameters.

We add a getter and a setter for the centerStar property (`getCenterStar` and `setCenterStar`).

We override the `hashCode` and `equals`, using the `centerStar` property.

Finally, we override `toString` method. The returned String must have the following format : `<name> circles around <centerStar.name> at the <distanceWithCenterStar> AU`.
The `distanceWithCenterStar` is computed with the help of `distanceBetween` method taking as parameter the planet and its center star.##

# Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        Planet earth = new Planet();
        Planet naboo = new Planet("Naboo", 17.4389, 8349.1, 8943.92, new Star("Betelgeuse", 128.23, -12.82, 32.328, 1289.3));

        System.out.println(naboo.toString());
        System.out.println(earth.toString());
        System.out.println(naboo.getCenterStar().toString());
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Naboo circles around Betelgeuse at the 12220.902 AU
Soleil circles around Soleil at the 0.000 AU
Betelgeuse shines at the 1289.300 magnitude
$
```
