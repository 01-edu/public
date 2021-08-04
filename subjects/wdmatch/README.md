## wdmatch

### Instructions

Write a program that takes two `string` and checks whether it is possible to write the first `string` with characters from the second `string`. This rewrite must respect the order in which these characters appear in the second `string`.

If it is possible, the program displays the `string` followed by a newline (`'\n'`), otherwise it simply displays nothing.

If the number of arguments is different from 2, the program displays nothing.

### Usage

```console
$ go run . 123 123
123
$ go run . faya fgvvfdxcacpolhyghbreda
faya
$ go run . faya fgvvfdxcacpolhyghbred
$ go run . error rrerrrfiiljdfxjyuifrrvcoojh
$ go run . "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux
$ go run .
$
```
