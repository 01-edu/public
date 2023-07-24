## AdventureMonster

### Instructions

Create a new class `Monster` in a new file `Monster.java`.

This class inherits from `Character`.

It has one constructor, with the same parameters as `Character` (`name` and `maxHealth`).

You need to override `toString` method, use the following format : 
* if the monster is still alive : `<name> is a monster with <currentHealth> HP`.
* Otherwise : `<name> is a monster and is dead`.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        Character aragorn = new Character("Aragorn", 20);
        Monster slime = new Monster("Slime", 15);

        System.out.println(Character.printStatus());

        Character winner = Character.fight(aragorn, slime);

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
 - Aragorn : 20/20
 - Slime is a monster with 15 HP
------------------------------------------
------------------------------------------
Characters currently fighting : 
 - Aragorn : 11/20
 - Slime is a monster and is dead
------------------------------------------
$ 
```

### Notions
[Inheritance](https://docs.oracle.com/javase/tutorial/java/IandI/subclasses.html)  