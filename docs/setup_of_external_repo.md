# ADDITION OF A REPOSITORY OF EXERCISE PROCEDURE

## INTRODUCTION

This document is a guide on how to add your own exercises from your own repository.
This guide assumes that you have understood the files structures of the repository [public](https://github.com/01-edu/public).
It will only address the settings part of this task. 
Should you need more information regarding the file structure of the exercises, check the [addition of exercise procedure](https://github.com/01-edu/public/blob/master/docs/addition_of_exercise_draft.md).

## PREQUISITES

-A github account where your exercises repository will be stored [https://github.com/](https://github.com/).

-A dockerHub account [https://hub.docker.com/](https://hub.docker.com/). 


## I- SETUP OF YOUR GUTHUB REPOSITORY THRU A FORK

Instead of starting with an empty repository, for the very first time we recommend forking the official [public](https://github.com/01-edu/public). repository.

The advantages of this way:

This will give you a base to work on regarding the file architecture of an exercise repository with examples.

You will also be able to customize already existing exercises to your needs. 

With just a quick modification the repo will be ready to be linked. 

## **1. Fork the 01 public repository**

Once logged into your github account, go to:
https://github.com/01-edu/public
From there fork the public repo to your account (the button is on the top corner right side)

## **2. Remove the CNAME file from the forked repo**

Git clone the repo and push the deletion or simply delete it directly from github

This is the occasion to push a new test exercise if you have one already written.

## **3. Publish the repository on github pages**

- Go to the settings tab of your exercises repository.
- On the option page find the GitHub Pages section.
- Please see below the settings to follow. 
(Please not that it might take up to 10 mins for your page to be pusblished)

<img width="1280" alt="Capture d’écran " src="img/adding-exercises-repository/1.png">

## **4. Take note of the path of an exercise subject you added

Example:
If user Frenchris, added an exercise to the forked called how-2-go
This is the path where the README.md would be. 
https://frenchris.github.io/public/subjects/how-2-go/

Note that you do not keep the README.md at the end of the path

This path should be added to the attribute “subject” of type string in the object attribute of the new exercise. 
 

## II- SETUP OF YOUR DOCKER REPOSITORY 

- 1. Sign in your docker hub account the 01 public repository.

- 2. In your account, go your settings and link your github account.

- 3. Create a repository named “test” and make sure that your github account is linked. 
If you see this image,

<img width="160" alt="Capture d’écran " src="img/adding-exercises-repository/2.png">

It means your github account is correctly linked.

- 4. In the Builds tab configure the automated build settings as below (for the go tests).

<img width="1280" alt="Capture d’écran " src="img/adding-exercises-repository/3.png">


- 5. Once the build is complete (it can take 5 to 15 mins). Go back to the  attributes of the exercise,
Add the attribute **testImage (type string)**
Fill it with the name of the repository,
**In this example: frenchris/test** 

- 6. Once your exercise has both the attributes completed correctly, the exercise is viable and can be tested on the server which was selected for its addition. 

- 7. As a reminder to test the exercise it is suggested to follow these steps:
a. `Create` a custom `Quest-test` object
b. `Adding` the new `exercise` object as a child to the newly created `Quest-test` object
c. `Create` a `Piscine-test` object
d. `Adding` the new `Quest-test` as a child to the newly created `Piscine-test` object
e. `Adding` the `Piscine-test` to the `campus` object as **first child** 
f. Go to the event page and launch the newly created `Piscine-test`. (You may need to refresh the page 2-3 times before the `campus/Piscine-test` option appears)
g. Once the event is launched, use the event page to add yourself as a student in the launched event `Piscine-test`
h. You can now try the exercise. If everything is well set, the subject should be loaded and, when you submit a correct solution, the exercise should pass.