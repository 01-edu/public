## AdventureTemplar

### Instructions

We will now implement our second interface through a new class : `Templar` in a file named `Templar.java`.

This class implements `Healer` and `Tank` interfaces and inherits from `Character`.  
It has two new attributes :

- a non modifiable integer `healCapacity`,
- a non modifiable integer `shield`.

It has a constructor with four parameters :
- `name`,
- `maxHealth`,
- `healCapacity`,
- `shield`

From the `Healer` interface :

- `getHealCapacity` returns the property `healCapacity`
- `heal` adds `healCapacity` to the currentHealth of the `Character` in parameter. Beware that `currentHealth` can't be greater the `maxHealth`.

From `Tank` interface :

- `getShield` returns the property `shield`.

You will override `toString` method with the following format : `<name> is a strong Templar with <currentHealth> HP. It can heal <healCapacity> HP and has a shield of <shield>.`
If its `currentHeal` is equal to 0, the format is `<name> has been beaten, even with its <shield> shield. So bad, it could heal <healCapacity> HP.`

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        Templar alistair = new Templar("Alistair", 20, 5, 4);
        Templar roderick = new Templar("Roderick", 10, 3, 2);

        Character.fight(alistair, roderick);

        alistair.heal(alistair);

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
 - Alistair is a strong Templar with 16 HP. It can heal 5 HP and has a shield of 4.
 - Roderick has been beaten, even with its 2 shield. So bad, it could heal 3 HP.
------------------------------------------
$
```

### Notions

[Implementation](https://docs.oracle.com/javase/tutorial/java/IandI/usinginterface.html)  
[Inheritance](https://docs.oracle.com/javase/tutorial/java/IandI/subclasses.html)
