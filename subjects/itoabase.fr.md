## itoabase

### Instructions

Écrire une fonction qui:

- convertit une valeur `int` en `string` en utilisant la base spécifiée en argument
- et qui retourne cette `string`

Cette base est exprimée comme un `int`, de 2 à 16. Les caractères compris dans la base sont les chiffres de 0 à 9, suivis des lettres majuscules de A à F.

Par exemple, la base `4` sera équivalente à "0123" et la base `16` sera équivalente à "0123456789ABCDEF".

Si la valeur est négative, la `string` résultante doit être précédée d'un signe moins `-`.

Seuls des arguments valables seront testés.

### Fonction attendue

```go
func ItoaBase(value, base int) string {

}
```
