## Factory

### Instructions

Now let's implement the Factory Design Pattern

```mermaid
classDiagram
class Driver {
    <<abstract>>
    +createTransport()* Transport
}
class CarDriver {
  +createTransport() Transport
}
class PlaneDriver {
  +createTransport() Transport
}

class Transport {
    <<interface>>
    +getDistance()* int
}
class Car {
  +getDistance() int
}
class Plane {
  +getDistance() int
}

class DriverFactory {
  + getDriver(Type: String) Driver
  
}

class TransportFactory {
  + getTransport(Type: String) Transport
}

Driver <|-- CarDriver
Driver <|-- PlaneDriver
Transport <|-- Car
Transport <|-- Plane
```

Here is the matching class diagram. Create the matching classes in the matching files.

The method createTransport of CarDriver should build and return a Car using TransportFactory, and for the PlaneDriver class, it should build and return a Plane.

The car should return 600 and the plane should return 10000.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args)  {

        System.out.println(DriverFactory.getDriver("Car").createTransport().getDistance());
        System.out.println(DriverFactory.getDriver("Plane").createTransport().getDistance());
    }
}
```
          
and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
600
10000
$
```

### Notions
[Class diagram](https://fr.wikipedia.org/wiki/Diagramme_de_classes)  

