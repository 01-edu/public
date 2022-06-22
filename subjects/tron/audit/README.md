#### Functional

##### Open the project, and remove all the files except `ai.js`.

##### Copy the game engine ([index.html](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/game_students/index.html)) to the root of the student's project.

##### Create a simple web server by running the following command. You will first need to `cd` to the root of the student's project.

```sh
$ python3 -m http.server
Serving HTTP on :: port 8000 (http://[::]:8000/)
```

##### Open a web browser, and go to the specified port. In the above case, that will be http://localhost:8000.

##### You can find the source code for each AI opponent here. You'll need to place each `.js` file at the root of the student's project: http://localhost:8000/?ai=https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/

##### Battle `ai.js` versus `random.js`. It is best out of 3. Delete the seed before each go. http://localhost:8000/?ai=random.js+ai.js

###### Did `ai.js` win `random.js` at least 2 times out of the 3 games?

##### Battle `ai.js` versus `right.js`. It is best out of 3. Delete the seed before each go. http://localhost:8000/?ai=right.js+ai.js

###### Did `ai.js` win `right.js` at least 2 times out of the 3 games?

##### Battle `ai.js` versus `snail.js`. It is best out of 3. Delete the seed before each go. http://localhost:8000/?ai=snail.js+ai.js

###### Did `ai.js` win `snail.js` at least 2 times out of the 3 games?

###### Does the code avoid [deep nesting](https://testing.googleblog.com/2017/06/code-health-reduce-nesting-reduce.html)?

##### Battle `ai.js` versus `hard.js`. It is best out of 3. Delete the seed before each go. http://localhost:8000/?ai=hard.js+ai.js

###### Did `ai.js` win `hard.js` at least 2 times out of the 3 games?

###### At any point, did the AI player complete the game without crashing because of too much CPU usage?

#### Bonus

##### Battle `ai.js` versus `licence-to-kill.js`. It is best out of 3. Delete the seed before each go. http://localhost:8000/?ai=licence-to-kill.js+ai.js

###### +Did `ai.js` win `licence-to-kill.js` at least 2 times out of the 3 games?

##### If you have an AI, and are prepared for battle. Modify the URL to battle against your AI. Best out of 3.

###### +Did the audited AI win against your AI?
