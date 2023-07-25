## StarGalaxy

### Instructions

Create a new class `Galaxy` in a file named `Galaxy.java`.

It has one private property : `celestialObjects` of type `List<CelestialObject>`.

It has one constructor with no parameters, which initialises `celestialObjects` property with an empty list.

We add a getter for `celestialObjects` property (`getCelestialObjects`).

We create a new method `addCelestialObject` with a `CelestialObject` argument. This method adds the object in parameter to the `celestialObjects` list.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.util.List;

public class ExerciseRunner {

    public static void main(String[] args) {
        Galaxy galaxy = new Galaxy();
        CelestialObject lune = new CelestialObject("Lune", -123.12, 392.238, 32.31);
        Star betelgeuse = new Star("Betelgeuse", 128.23, -12.82, 32.328, 1289.3);
        Planet naboo = new Planet("Naboo", 17.4389, 8349.1, 8943.92, betelgeuse);
        
        galaxy.addCelestialObject(lune);
        galaxy.addCelestialObject(betelgeuse);
        galaxy.addCelestialObject(naboo);
        
        List<CelestialObject> celestialObjects = galaxy.getCelestialObjects();
        
        for (CelestialObject celestialObject : celestialObjects) {
            System.out.println(celestialObject.toString());
        }
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
Lune is positioned at (-123,120, 392,238, 32,310)
Betelgeuse shines at the 1289.300 magnitude
Naboo circles around Betelgeuse at the 12220.902 AU
$ 
```

### Notions

[Polymorphism](https://docs.oracle.com/javase/tutorial/java/IandI/polymorphism.html)

