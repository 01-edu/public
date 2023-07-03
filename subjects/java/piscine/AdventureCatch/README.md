## AdventureCatch

### Instructions

Now, we will throw and catch some exceptions.

For the methods `attack`, `takeDamage` and `heal`, if the character is already dead, you must throw a `DeadCharacterException`. Do this for all the `Character` subclasses.  

The `fight` method should not throw an exception, so you will need to catch the exception in this method.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        Weapon excalibur = new Weapon("Excalibur", 7);
        Weapon baton = new Weapon("Baton", 3);
        Templar arthur = new Templar("Arthur", 30, 5, 3, excalibur);
        Sorcerer merlin = new Sorcerer("Merlin", 28, 2, baton);

        try {
            arthur.takeDamage(50);
        } catch (DeadCharacterException e) {
            System.out.println(e.getMessage());
        }

        try {
            arthur.takeDamage(2);
        } catch (DeadCharacterException e) {
            System.out.println(e.getMessage());
        }
        try {
            arthur.attack(merlin);
        } catch (DeadCharacterException e) {
            System.out.println(e.getMessage());
        }
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
The templar Arthur is dead.
The templar Arthur is dead.
$ 
```

### Notions
[Try](https://docs.oracle.com/javase/tutorial/essential/exceptions/try.html)  
[Catch](https://docs.oracle.com/javase/tutorial/essential/exceptions/catch.html)  
[Declaring thrown](https://docs.oracle.com/javase/tutorial/essential/exceptions/declaring.html)  
[Throwing](https://docs.oracle.com/javase/tutorial/essential/exceptions/throwing.html)  