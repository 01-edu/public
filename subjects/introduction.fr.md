## Introduction

### Instructions

#### 1- get-ready

Créer dans votre compte `git` le repository nommé `[[ROOT]]`.

Une fois créé, cloner ce repository sur votre desktop.
Si votre nom d'utilisateur `git` était `choumi` voici la commande qui devrait être utilisé :

`git clone https://[[DOMAIN]]/git/choumi/[[ROOT]]`

Cette commande doit être adaptée avec **votre propre username**.

Ce repository sera le dossier où tous les exercises doivent être uploadés.

#### 2- set

Une fois que le repository est créé, écrire votre premier programme shell nommé `hello.sh`

Quand ce programme est éxécuté il doit afficher `Hello {username}!`
Où `{username}` est votre `username`

##### Utilisation

Si l'`{username}` est `choumi` :

```console
user@host:~/[[ROOT]]$ ./hello.sh
Hello choumi!
user@host:~/[[ROOT]]$
```

#### 3- go-say-hello

Après que `hello.sh` s'exécute proprement, il doit être uploadé dans le repository avec les commandes ci-dessous :

1. `git add hello.sh`
2. `git commit -m "My very first commit"`
3. `git push origin master`

Une fois ces étapes appliquées, le fichier peut maintenant être soumis pour correction sur la platforme en cliquant le bouton `submit`.

Cette action fera passer les tests sur le fichier rendu `hello.sh`.
