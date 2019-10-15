## cl-camp4

### Instructions

"someone familiar"

Create a file `myfamily.sh`, which will show a subject's family (key: relatives).

-   The quotes have to be removed.

-   The subject will be decided depending on his ID which will be contained in the environment variable HERO_ID.

*   Where to look : https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json

*   What to use : curl, jq and others...

### Usage

```console
student@ubuntu:~/piscine-go/test$ export HERO_ID=1
student@ubuntu:~/piscine-go/test$ ./myfamily.sh
Marlo Chandler-Jones (wife); Polly (aunt); Mrs. Chandler (mother-in-law); Keith Chandler, Ray Chandler, three unidentified others (brothers-in-law); unidentified father (deceased); Jackie Shorr (alleged mother; unconfirmed)
student@ubuntu:~/piscine-go/test$
```
