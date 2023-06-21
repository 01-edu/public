## AdventurAbstract

### Instructions

Let's change some things : make the `Character` as `abstract`. You can now try to instantiate a Character object, it will fail :)

Change the `attack` and `takeDamage` methods : put them as `abstract` too.

Now, if you try to launch some example, it will fail. Indeed, you need to implement both methods in all subclasses. Do this as follow : 

Remember that in all cases, the `currentHealth` should not be lower than 0.

* For Monster class : 
  * the `attack` method should deal 7 damages to the character as parameter.
  * the `takeDamage` method should take 80% of all the damages, rounded to the inferior integer.
* For the `Sorcerer` class : 
  * the `attack` method should heal itself (using the `heal` method) then deal 10 damages to the character as parameter.
  * the `takeDamage` method should take all the damages.
* For the `Templar` class : 
  * the `attack` method should heal itself (using the `heal` method) then deal 6 damages to the character as parameter.
  * the `takeDamage` method should take the damage as parameter minus the shield value.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        Templar alistair = new Templar("Alistair", 18, 2, 4);
        Sorcerer morrigan = new Sorcerer("Morrigan", 21, 5);
        Monster dragon = new Monster("Dragon", 12);

        dragon.attack(alistair);
        dragon.attack(morrigan);
        
        alistair.attack(dragon);
        morrigan.attack(dragon);

        System.out.println(Character.printStatus());
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
------------------------------------------
Characters currently fighting :  
 - Alistair is a strong Templar with 17 HP. It can heal 2 HP and has a shield of 4.
 - Morrigan is a sorcerer with 19 HP. It can heal 5 HP.
 - Dragon is a monster and is dead
------------------------------------------
$ 
```

### Notions
[Abstract](https://docs.oracle.com/javase/tutorial/java/IandI/abstract.html)  