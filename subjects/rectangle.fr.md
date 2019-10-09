## Rectangle

### Instructions

Considérer qu'un point est défini par ses coordonnées et qu'un rectangle est défini par les points de son coin du haut à gauche et son coin du bas à droite.

-   Définir deux structures nommées, `point` et `rectangle`.

-   La structure `point` doit avoir deux variables, `x` et `y`, de type `int`.

-   La structure `rectangle` doit avoir deux variables, `upLeft` et `downRight` de type `point`.

-   Le but est de faire un programme qui:
    -   Avec une slice de points donnée de taille `n` retournes le plus petit rectangle qui contient tous les points dans le vecteur de points0. Le nom de cette fonction est `defineRectangle`.
    -   Et qui calcules et affiche l'airethe de ce rectangle défini.

### Main et fonctions attendues pour ce programme

```go
func defineRectangle(ptr *point, n int) *rectangle {

}

func calArea(ptr *rectangle) int {

}

func main() {
	vPoint := []point{}
	rectangle := &rectangle{}
	n := 7

	for i := 0; i < n; i++ {
		val := point{
			x: i%2 + 1,
			y: i + 2,
		}
		vPoint = append(vPoint, val)
	}
	rectangle = defineRectangle(vPoint, n)
	fmt.Println("area of the rectangle:", calArea(rectangle))
}
```

### Expected output

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
area of the rectangle: 6
student@ubuntu:~/piscine-go/test$
```
