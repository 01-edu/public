## WeddingComplex

### Instructions

Create a file `WeddingComplex.java`.

Write a function `createBestCouple` that returns a map of name from the first map associated with a name from the second list.  
Each map argument will be composed as follows :
* Key: The name of the member to marry
* Value: The ordered list of preference containing members (keys) from the other map.

To objective of this function is to determine the best couple possible.  
The optimal solution occurs when, considering a particular member named `Alice`, if there exists a preferred partner `Bob` over the selected partner `Charly`, then `Bob` is happier with his chosen partner `Daphnee` compared to being with `Alice`.

> With an example : 
> * First map : 
>  * Naruto orders his preferences : Sakura > Hinata
>  * Sasuke orders his preferences : Sakura > Hinata
> * Second map : 
>  * Sakura orders her preferences : Sasuke > Naruto
>  * Hinata orders her preferences : Naruto > Sasuke
>
> A correct solution for this example is : 
> * Sasuke <-> Sakura (they are both with their first preference)
> * Naruto <-> Hinata (Naruto would prefere Sakura, but Sakura is happier with Sasuke, Hinata has her first choice)

All the test will be chosen to guarantee such a solution exists. More than one solution could exist.

In order to simplify, the list will always have the same size.

### Expected Functions

```java
import java.util.List;
import java.util.Map;

public class WeddingComplex {
    public static Map<String, String> createBestCouple(Map<String, List<String>> first, Map<String, List<String>> second) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.util.List;
import java.util.Map;

public class ExerciseRunner {

    public static void main(String[] args) {
        System.out.println(WeddingComplex.createBestCouple(
                Map.of("Naruto", List.of("Sakura", "Hinata"), "Sasuke", List.of("Sakura", "Hinata")),
                Map.of("Sakura", List.of("Sasuke", "Naruto"), "Hinata", List.of("Naruto", "Sasuke"))));
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
{Sasuke=Sakura, Naruto=Hinata}
$ 
```

### Notions
[Wedding Algorithm](https://fr.wikipedia.org/wiki/Algorithme_de_Gale_et_Shapley#Pseudo-code)  
[Map](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/Map.html)  