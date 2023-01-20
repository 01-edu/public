## json-researcher

### Instructions

Create the script `json-researcher.sh` which will show the "name" and the "power" of the superhero with `id: 1`

- Where to look : https://((DOMAIN))/assets/superhero/all.json

- What to use : `curl` and `jq` and `grep`

- The output should be exactly like the example bellow but with the expected name

```console
$./json-researcher.sh | cat -e
  "name": "<...>",$
    "power": <...>,$
$
```

### Hints

With `curl` you can get the content of the file from the url:

```console
$curl https://assets.01-edu.org/devops-branch/HeadTail.txt
<...>
$
```

`jq` is a command-line utility that can slice, filter, and transform the components of a JSON file

```console
$cat fruit.json
{
  "name": "apple",
  "color": "green",
  "price": 1.0
}
$jq '.color' fruit.json
"green"
$
```

`grep` is a command-line utility for searching plain-text data sets for lines that match a regular expression.

```console
$cat fruit.json
{
  "name": "apple",
  "color": "green",
  "price": 1.0
}
$cat fruit.json | grep color
  "color": "green",
$
```

You may need to use pipes `(|)` and the logical operator `&&`.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
