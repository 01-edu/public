## AdventureUtils

### Instructions

Let's add some useful methods to allow us to manipulate and monitor what happens with our characters.

Firstly, let's add something to see all our characters in one call : 
* Add a private static list of `Character`, `allCharacters`, which will contain all our characters. In the constructor, you need to add every new `Character` to this list.
* Add a static method, `printStatus`, which takes no parameters and returns a formatted String that lists `allCharaters`. The awaited format is as follows with new line at the end : 
```
------------------------------------------
Characters currently fighting : 
 - <character1.toString>
 - <character2.toString>
------------------------------------------
```
If there is no character created : 
```
------------------------------------------
Nobody's fighting right now !
------------------------------------------
```

Finally, add a static method, `fight`, which takes two `Character` objects and returns the winner in the fight.  
During the fight, the first one attacks, then the second one, then the first, then the second and so on until one of them reaches 0 in its currentHealth. When `currentHealth` for one of them reaches 0, the other one is the winner.

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {

    public static void main(String[] args) {
        System.out.println(Character.printStatus());

        Character aragorn = new Character("Aragorn", 20);
        Character uruk = new Character("Uruk", 15);

        System.out.println(Character.printStatus());

        Character winner = Character.fight(aragorn, uruk);

        System.out.println(winner.toString());
        System.out.println(Character.printStatus());
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
------------------------------------------
Nobody's fighting right now !
------------------------------------------
------------------------------------------
Characters currently fighting : 
 - Aragorn : 20/20
 - Uruk : 15/15
------------------------------------------
Aragorn : 11/20
------------------------------------------
Characters currently fighting : 
 - Aragorn : 11/20
 - Uruk : KO
------------------------------------------
$ 
```

### Notions
[Static](https://docs.oracle.com/javase/tutorial/java/javaOO/classvars.html)  
