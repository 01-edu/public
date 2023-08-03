## StarProperties

### Instructions

In the following quest, we will work with the same files and classes. You should keep them from one exercise to the following.

Create a file `CelestialObject.java`.

Create a public class named `CelestialObject`.  
The class must contain four properties : 
* x (double)
* y (double)
* z (double)
* name (String)

The (x, y, z) properties are the coordinates of the object.

For the moment, you should declare the properties as public.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        CelestialObject celestialObject = new CelestialObject();
        System.out.println(celestialObject.x);
        System.out.println(celestialObject.y);
        System.out.println(celestialObject.z);
        System.out.println(celestialObject.name);
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
null
$ 
```

### Notions
[Class](https://docs.oracle.com/javase/tutorial/java/javaOO/classdecl.html)  
[Property](https://docs.oracle.com/javase/tutorial/java/javaOO/variables.html)  