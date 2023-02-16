# Reviews

## Context

Only unblocked bonus exercises can be reviewed, which are usually the last items of a quest.
Bonus exercises are available in quests 02 / 03 / 04 / 05 / 07 / 08 / 09.
<img width="1073" alt="Screenshot 2019-10-17 at 02.43.31" src="img/reviews/Screenshot 2019-10-17 at 02.43.31.jpg">

Making a review involves 2 students who are going to review each other's exercise.
Each user has to :

- **check** the code of the other one on Gitea. **You must add your reviewer as contributor in your repository so he/she can see your code**.
- **bet** if it is going to fail or succeed in the Review interface.
- **run** his own exercise in the Review interface to give a result to compare with the bet, and state if it was right or wrong.

## Usage

To submit an exercise to review and be reviewed, go to the `Review panel` on the right of the quest interface and click on the button to submit it :

<img width="1073" alt="Screenshot 2019-10-17 at 01.59.32" src="img/reviews/Screenshot 2019-10-17 at 01.59.32.jpg">

Then, when a match with another student who also wants to review is found, you will be notified :

<img width="1073" alt="Screenshot 2019-10-17 at 02.03.11" src="img/reviews/Screenshot 2019-10-17 at 02.03.11.jpg">

Once you and the other student have confirmed the match, you can go and check the code of the other student on their repository in the component `Your review` :

<img width="1073" alt="Screenshot 2019-10-17 at 02.05.46" src="img/reviews/Screenshot 2019-10-17 at 02.05.46.jpg">

When you have checked their code, those 4 steps will have to be completed to achieve the review :

- You have to bet if the exercise of the other student will fail or succeed after running the tester
- The other student has to do the same with your exercise and make his own bet, that you will see on the component `Review of your exercise`
  <img width="1073" alt="Screenshot 2019-10-17 at 02.20.18" src="img/reviews/Screenshot 2019-10-17 at 02.20.18.jpg">

- Once the other student's bet is set, the tester button unblocks and you have to run it on your exercise to output the result determining if your exercise has failed or succeeded ; this result is compared to the other student's bet to determine if it was wrong or right
- The other student also has to run the tester on their exercise to determine their result and so whether your bet was wrong or right
  <img width="1073" alt="Screenshot 2019-10-17 at 02.22.30" src="img/reviews/Screenshot 2019-10-17 at 02.22.30.jpg">

A bet is considered succeeded if it is equal to the tester output :
<img width="1073" alt="Screenshot 2019-10-17 at 02.24.04" src="img/reviews/Screenshot 2019-10-17 at 02.24.04.jpg">

You can have those 4 cases :

| bet     | tester output | bet result    |
| ------- | ------------- | ------------- |
| fail    | failed        | **succeeded** |
| fail    | succeeded     | **failed**    |
| succeed | succeeded     | **succeeded** |
| succeed | failed        | **failed**    |