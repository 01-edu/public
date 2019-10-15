## sortList

### Instructions

Écrire une fonction qui doit :

-   Trier la liste donnée en paramètre en utilisant la fonction cmp pour sélectionner l'ordre à appliquer,

-   Retourner un pointeur au premier élément de la liste triée.

Les duplications doivent rester.

Les inputs seront toujours valides.

Le `type NodeList` doit être utilisé.

Les fonctions passées comme `cmp` retourneront toujours `true` si `a` et `b` sont dans le bon ordre, sinon elles retourneront `false`.

### Fonction et structure attendues

```go
type Nodelist struct {
	Data int
	Next *Nodelist
}

func SortList (l *NodeList, cmp func(a,b int) bool) *NodeList{

}
```

-   Par exemple, la fonction suivante utilisée comme `cmp` triera la liste dans l'ordre croissant :

```go
func ascending(a, b int) bool{
	if a <= b {
		return true
	} else {
		return false
	}
}
```
