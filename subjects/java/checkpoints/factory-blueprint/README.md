## Factory Blueprint

### Instructions

You are given an incomplete Factory design pattern implementation with some incorrect parts. Complete and fix the class to demonstrate your understanding of how the Factory design pattern works.

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
    public Product createProduct(String type) {

    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your classes:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        Factory factory = new Factory();

        ConcreteProductA productA = factory.createProduct("A");
        if (productA != null) {
            productA.showDetails();
        } else {
            System.out.println("Invalid product type");
        }

        ConcreteProductA productB = factory.createProduct("B");
        if (productB != null) {
            productB.showDetails();
        } else {
            System.out.println("Invalid product type");
        }

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
This is ConcreteProductA.
This is ConcreteProductB.
Invalid product type
$
```
