## WeddingComplex

### Instructions

Create a file `WeddingComplex.java`.

Write a function `createBestCouple` that returns a map of name from the first map associated with a name from the second list.  
Each argument will be composed as follow : 
* The name of the member to marry
* The ordered list of preference of member of the second list.  

To objective of this function is to determine the best couple possible.  
The best solution is when for a given member (will be called Alice) : If there is a preferred partner (will be called Bob) than the chosen partner (will be called Charly), Bob is happier with its chosen partner (called Daphnee) than with Alice.

> With an example : 
> * First list : 
>  * Naruto orders his preferences : Sakura > Hinata
>  * Sasuke orders his preferences : Sakura > Hinata
> * Second list : 
>  * Sakura orders her preferences : Sasuke > Naruto
>  * Hinata orders her preferences : Naruto > Sasuke
>
> A correct solution for this example is : 
> * Sasuke <-> Sakura (they are both with their first choice)
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