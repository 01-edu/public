## StarOverride

### Instructions

In the `Star` class, let's add a new constructor with the following arguments : 
* the `name`
* the position `x`
* the position `y`
* the position `z`
* the `magnitude`  
It calls the constructor of the superclass `CelestialObject`.

We will override `hashCode` and `equals` methods, using `magnitude` property.

Finally, we override `toString` method. The returned String must have the following format : `<name> shines at the <magnitude> magnitude`

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        Star star = new Star();
        Star star2 = new Star();
        Star proxima = new Star("Proxima", 18.389, 832.32, 218, 0.4);
        
        System.out.println(star.toString());
        System.out.println(proxima.toString());
        System.out.println(star.equals(star2));
        System.out.println(star.equals(proxima));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
Soleil shines at the 0.000 magnitude
Proxima shines at the 0.400 magnitude
true
false
$ 
```

### Notions

[Override](https://docs.oracle.com/javase/tutorial/java/IandI/override.html)  