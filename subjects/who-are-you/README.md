## who-are-you

### Instructions

"You just woke up in a dark alley...
You can not remember who you are...
The only thought that comes to your mind is a tag that says: `subject Id: 70`"

Create the file `who-are-you.sh` which will remind you who you are by showing your name only.

- Where to look : https://((DOMAIN))/assets/superhero/all.json

- What to use : `curl` and `jq`

- The output should be exactly like the example bellow but with the expected name

```console
$./who-are-you.sh | cat -e
"name"$
$
```
