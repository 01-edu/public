## to-git-or-not-to-git-?

### Instructions

Write in a file `to-git-or-not-to-git.sh` the command that isolates your Gitea `id`.

Only the numbers will appear.

> [Local API documentation](https://git.01.alem.school/api/swagger)

Here is the base command that needs to be adapted with your username and more :

```
curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"choumi\"}}){id}}"}'
```

### Usage

```console
$ ./to-git-or-not-to-git.sh
231748
$
`
```
