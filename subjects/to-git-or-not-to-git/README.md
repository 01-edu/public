## to-git-or-not-to-git-?

### Instructions

Write in a file `to-git-or-not-to-git.sh` the command that isolates your Gitea `id`.

Only the numbers will appear.

Here is the base command that needs to be adapted with your username and more :

```
curl -s https://((DOMAIN))/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"choumi\"}}){id}}"}'
```

### Usage

```console
$ bash to-git-or-not-to-git.sh
231748
$
```
