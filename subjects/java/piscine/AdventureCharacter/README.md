## AdventureCharacter

### Instructions

In the following quest, we will work with the same files and classes. You should keep them from one exercise to the following.

Create a file `Character.java`.

Create a public class named `Character`.  
The class must contain three properties : 
* maxHealth (int) : with a getter and no setter. This property is not updatable (`final` keyword).
* currentHealth (int) : with a getter and no setter.
* name (String) : with a getter and no setter. This property is not updatable (`final` keyword).

Create a constructor with two parameters (name and maxHealth) : the currentHealth is initialized with the value of maxHealth.

Override `toString` method, Must have the format `<name> : <currentHealth>/<maxHealth>`. If the currentHealth is 0, the format is `<name> : KO`.

Implement two methods : 
* `takeDamage`, with an integer parameter, that will subtract the amount in parameter from `currentHealth`. `currentHealth` can't be lower than 0.
* `attack`, with a parameter of type `Character`, that will call `takeDamage` of the parameter with a default value : `9`.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        Character aragorn = new Character("Aragorn", 20);
        Character uruk = new Character("Uruk", 5);
        
        System.out.println(aragorn.toString());
        System.out.println(uruk.toString());
        
        aragorn.attack(uruk);

        System.out.println(uruk.toString());
        
        aragorn.takeDamage(12);

        System.out.println(aragorn.toString());
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
Aragorn : 20/20
Uruk : 5/5
Uruk : KO
Aragorn : 8/20
$ 
```

### Notions
[Class](https://docs.oracle.com/javase/tutorial/java/javaOO/classdecl.html)  
[Property](https://docs.oracle.com/javase/tutorial/java/javaOO/variables.html)  
[Override](https://docs.oracle.com/javase/tutorial/java/IandI/override.html)  