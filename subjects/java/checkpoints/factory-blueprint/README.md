## Factory Blueprint

### Instructions

You are given an incomplete Factory design pattern implementation with some incorrect parts. Complete and fix the class to demonstrate your understanding of how the Factory design pattern works.

The method `showDetails` should print the class name 'ConcreteProductA' for `ConcreteProductA` and so on.

### Expected Classes

```java
// Product interface
public interface Product {
    void showDetails();
}

// ConcreteProductA class
public class ConcreteProductA {
    ...
}

// ConcreteProductB class
public class ConcreteProductB {
    ...
}

// Factory class
public class Factory {
    public Product createProduct(String type) { // the type parametre accept two values `A` and `B`

    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your classes:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        Factory factory = new Factory();

        // Handle invalid product type
        Object invalidProduct = factory.createProduct("C");
        if (invalidProduct != null) {
            invalidProduct.showDetails();
        } else {
            System.out.println("Invalid product type");
        }
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Invalid product type
$
```
