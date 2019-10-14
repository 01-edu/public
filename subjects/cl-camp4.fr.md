## cl-camp4

### Instructions

"quelqu'un de familier"

Créer un fichier `myfamily.sh`, qui affichera la famille d'un individu (clef: relatives).

-   Les guillemets doivent être enlevés.

-   L'individu sera choisi en fonction de son ID qui sera contenu dans la variable d'environment HERO_ID.

*   Où chercher : https://01.alem.school/assets/superhero/all.json

*   Quoi utiliser : `curl`, `jq` et d'autres...

### Utilisation

```console
student@ubuntu:~/piscine-go/test$ export HERO_ID=1
student@ubuntu:~/piscine-go/test$ ./myfamily.sh
Marlo Chandler-Jones (wife); Polly (aunt); Mrs. Chandler (mother-in-law); Keith Chandler, Ray Chandler, three unidentified others (brothers-in-law); unidentified father (deceased); Jackie Shorr (alleged mother; unconfirmed)
student@ubuntu:~/piscine-go/test$
```
