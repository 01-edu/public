## to-git-or-not-to-git-?

### Instructions

Write in a file `to-git-or-not-to-git.sh` the command that isolates your `gitHub id`.
Only the numbers will appear.

adapt this command 

```curl -s "https://01.alem.school/api/graphql-engine/v1/graphql" --data '{"query":"{user(where:{githubLogin:{_eq:\"'$USERNAME'\"}}){id}}"}'
```

### Usage

```console
$ ./to-git-or-not-to-git.sh
231748
$
`
```
