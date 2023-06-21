## AdventureWeapon

### Instructions

We create now the weapon for our adventurers. Create a new class `Weapon` in a new file `Weapon.java`.

It has two attributes : 
* a String called `name` with a getter.
* an integer called `damage` with a getter.

The constructor has these 2 attributes as parameters.

Let's overwrite the `toString` method for this class : it returns a string under the format `<name> deals <damage> damages`

Now, add a weapon attribute to the `Character` class, in its constructor and in the constructor of all subclasses.

In all subclasses, in the `attack` method, you need to use the `damage` when calling the `takeDamage` method. It the character has no weapon, use the previous defined damage per subclasses.
 
Update the `toString` method of all subclasses by using the concatenation of the current value with the following string : `He has the weapon <weapon.toString>`.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        Weapon narsil = new Weapon("Narsil", 15);
        Weapon baguette = new Weapon("Baguette magique", 20);
        Weapon massue = new Weapon("Massue", 8);
        Monster troll = new Monster("Troll", 30,  massue);
        Sorcerer dumbledore = new Sorcerer("Dumbledore", 25, 5, baguette);
        Templar alistair = new Templar("Alistair", 18, 2, 3, narsil);

        Character.fight(alistair, troll);

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
 - Troll is a monster and is dead. He has the weapon Massue deals 8 damages.
 - Dumbledore is a sorcerer with 25 HP. It can heal 5 HP. He has the weapon Baguette magique deals 20 damages.
 - Alistair is a strong Templar with 12 HP. It can heal 2 HP and has a shield of 3. He has the weapon Narsil deals 15 damages.
------------------------------------------
$ 
```
