## Kaggle Titanic

### Overview

The goal of this **1 week** project is to get the highest possible score on a Data Science competition. More precisely you will have to predict who survived the Titanic crash.

![alt text][titanic]

[titanic]: titanic.jpg "Titanic"

#### Kaggle

Kaggle is an online community of data scientists and machine learning practitioners. Kaggle allows users to find and publish data sets, explore and build models in a web-based data-science environment, work with other data scientists and machine learning engineers, and enter competitions to solve data science challenges. It’s a crowd-sourced platform to attract, nurture, train and challenge data scientists from all around the world to solve data science, machine learning and predictive analytics problems.

#### Titanic - Machine Learning from Disaster

One of the first Kaggle competition I did was: Titanic - Machine Learning from Disaster. This is a not-to-be-missed Kaggle competition.

You can see more [here](https://www.kaggle.com/c/titanic)

The sinking of the Titanic is one of the most infamous shipwrecks in history. On April 15, 1912, during her maiden voyage, the widely considered “unsinkable” RMS Titanic sank after colliding with an iceberg. Unfortunately, there were not enough lifeboats for everyone onboard, resulting in the death of 1502 out of 2224 passengers and crew.

While there was some element of luck involved in surviving, it seems some groups of people were more likely to survive than others.

### Role play

Ahoy, data explorer! Ready to set sail on the most thrilling voyage of your data science career? Welcome aboard the Kaggle Titanic challenge! You're about to embark on a journey through time, back to that fateful night in 1912.
Your mission, should you choose to accept it (and let's face it, you're already hooked), is to dive deep into the passenger manifest and uncover the secrets of survival. Who lived? Who perished? And most importantly, can you build a model that predicts it all?

### Learning Objective

In this challenge, you have to build a predictive model that answers the question: **“what sorts of people were more likely to survive?”** using passenger data (ie name, age, gender, socio-economic class, etc). **You will have to submit your prediction on Kaggle**.

### Instructions

#### Preliminary

The way the Kaggle platform works is explained in the challenge overview page. If you need more details, I suggest this [resource](https://www.kaggle.com/code/alexisbcook/getting-started-with-kaggle) that gives detailed explanations.

- Create a username following this structure: username*01EDU* location_MM_YYYY. Submit the description profile and push it on GitHub the first day of the week. Do not touch this file anymore.

- It is possible to have different personal accounts merged in a team for one single competition.

#### Scores

In order to validate the project you will have to score at least **78.9% accuracy on the Leaderboard**:

- 78.9% accuracy is the minimum score to validate the project.

Scores indication:

- 78.9% difficult - minimum required
- 80% very difficult: smart feature engineering needed
- More than 83%: excellent that corresponds to the top 2% on Kaggle
- More than 85%: cheating

#### Cheating

It is impossible to get 100%. Who would have predicted that Rose wouldn't let [Jack on the door](https://www.reddit.com/r/titanic/comments/14i0v5j/for_all_the_newbies_proof_its_not_a_door/?rdt=35268) ?

All people having 100% of accuracy on the Leaderboard cheated, there's no point to compare with them or to cheat. The Kaggle community estimates that having more than 85% is almost considered as cheated submissions as they are element of luck involved in the surviving.

**You can't use external data sets than the ones provided in that competition.**

#### The key points

- **Feature engineering**:
  Put yourself in the shoes of an investigator trying to understand what happened exactly in that boat during the crash. Do not hesitate to watch the movie to try to find as many insights as possible. Without a smart the feature engineering there's no way to validate the project ;-)

- The leaderboard evaluates on test data for which you don't have the labels. It means that there's no point to over fit the train set. Check the over fitting on the train set by dividing the data and by cross-validating the accuracy.

### Project repository structure

```console
project
│   README.md
│   environment.yml
│   username.txt
│
└───data
│   │   train.csv
│   |   test.csv
|   |   gender_submission.csv
│
└───notebook
│   │   main.ipynb


```

- `README.md` introduction of the project, shows the username, describes the features engineering and the best score on the **leaderboard**. Note the score on the test set using the exact same pipeline that led to the best score on the leaderboard.

- `environment.yml` contains all libraries required to run the code.

- `username.txt` contains the username, the last modified date of the file **has to correspond to the first day of the project**.

- `main.ipynb` This file (single Jupyter Notebook) should contain all steps of data analysis that contributed or not to improve the accuracy, the feature engineering, the model's training and prediction on the test set. It has to be commented to help the reviewers understand the approach and run the code without any bugs.
- **Submit your predictions on the Kaggle's competition platform**. Check your ranking and score in the leaderboard.

### Tips

Don't try to build the perfect model the first day. Iterate a lot and test your assumptions:

Iteration 1:

- Predict all passengers die

Iteration 2

- Fit a logistic regression with a basic feature engineering

Iteration 3:

- Perform an EDA. Make assumptions and check them. Example: What if first class passengers survived more. Check the assumption through EDA and create relevant features to help the model capture the information.

Iteration 4:

- Good luck !
