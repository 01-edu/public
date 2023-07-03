## StarUtils

### Instructions

Now, some utils method for our class.

First, the `toString` method which will return a literal version of our class. The format is the following : `<name> is positioned at (<x>, <y>, <z>)`. The printed double will be with 3 decimals.

Then, the `equals(Object object)` method which will return trus if all properties of the object as parameters are equals to the current object.

As we have rewritten the `equals` method, we need to rewrite the `hashCode` method. This method returns an integer. If two objects are equals (using the `equals` method), then the results of their hashCode method should be equals.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        CelestialObject celestialObject = new CelestialObject();
        CelestialObject earth = new CelestialObject("Terre", 1.0, 2.0, 2.0);
        CelestialObject earth1 = new CelestialObject("Terre", 1.0, 2.0, 2.0);

        System.out.println(earth.toString());
        System.out.println(earth.equals(earth1));
        System.out.println(earth.equals(celestialObject));

        System.out.println(earth.hashCode());
        System.out.println(celestialObject.hashCode());
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Terre is positioned at (1.000, 2.000, 2.000)
true
false
2129490293
-1811995559
$
```

### Notions

[Equals](<https://docs.oracle.com/javase/10/docs/api/java/lang/Object.html#equals(java.lang.Object)>)  
[HashCode](<https://docs.oracle.com/javase/10/docs/api/java/lang/Object.html#hashCode()>)  
[ToString](<https://docs.oracle.com/javase/10/docs/api/java/lang/Object.html#toString()>)
