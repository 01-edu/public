## Credit scoring

### Overview

The goal of this project is to implement a scoring model based on various source of data ([check data documentation](./readme_data.md)) that returns the probability of default. In a nutshell, credit scoring represents an evaluation of how well the bank's customer can pay and is willing to pay off debt. It is also required that you provide an explanation of the score. For example, your model returns that the probability that one client doesn't pay back the loan is very high (90%). The reason behind is that variable_xxx which represents the ability to pay back the past loan is low. The output interpretability will appear in a visualization.

### Role play

Hey there, future credit scoring expert! Ready to dive into the exciting world of predicting loan defaults? You're in for a treat! This project is all about building a nifty model that can help figure out how likely someone is to pay back their loan. Cool, right?

### Learning Objective

The ability to understand the underlying factors of credit scoring is important. Credit scoring is subject to more and more regulation, so transparency is key. And more generally, more and more companies prefer transparency to black box models.

Historical timeline of machine learning techniques applied to credit scoring

- [Machine Learning or Econometrics for Credit Scoring: Let’s Get the Best of Both Worlds](https://hal.archives-ouvertes.fr/hal-02507499v3/document)

### Instructions

#### Scoring model

There are 3 expected deliverables associated with the scoring model:

- An exploratory data analysis notebook that describes the insights you find out in the data set.
- The trained machine learning model with the features engineering pipeline:

  - Do not forget: **Coming up with features is difficult, time-consuming, requires expert knowledge. ‘Applied machine learning’ is basically feature engineering.**
  - The model is validated if the **AUC on the test set is at minimum 55%, ideally to 62% included (or in best cases higher than 62% if you can !)**.
  - The labelled test data is not publicly available. However, a Kaggle competition uses the same data. The procedure to evaluate test set submission is the same as the one used for the project 1.
  - Here are the [DataSets](https://assets.01-edu.org/ai-branch/project5/home-credit-default-risk.zip).

- A report on model training and evaluation:

  - Include learning curves (training and validation scores vs. training set size or epochs) to demonstrate that the model is not overfitting.
  - Explain the measures taken to prevent overfitting, such as early stopping or regularization techniques.
  - Justify your choice of when to stop training based on the learning curves.

#### Kaggle submission

The way the Kaggle platform works is explained in the challenge overview page. If you need more details, I suggest [this resource](https://towardsdatascience.com/getting-started-with-kaggle-f9138b35ae18) that gives detailed explanations.

- Create a username following that structure: username*01EDU* location_MM_YYYY. Submit the description profile and push it on the Git platform the first day of the week. Do not touch this file anymore.

- A text document `model_report.txt` that describes the methodology used to train the machine learning model :
  - Algorithm
  - Why the accuracy shouldn't be used in that case?
  - Limit and possible improvements

#### Model interpretability

This part hasn't been covered during the piscine. Take the time to understand this key concept.
There are different level of transparency:

- **Global**: understand important variables in a model. This answers the question: "What are the key variables to the model ? ". In that case it will tell if the revenue is more important than the age to the model for example. This allows to check that the model relies on important variables. No one wants his credit to be refused because of the weather in Lisbon !
- **Local**: each observation gets its own set of interpretability factors. This greatly increases its transparency. We can explain why a case receives its prediction and the contributions of the predictors. Traditional variable importance algorithms only show the results across the entire population but not on each individual case. The local interpretability enables us to pinpoint and contrast the impacts of the factors.

There are 2 tools you can use to analyse your model and its predictions: - Features importance (available if you use a Scikit Learn model) - [SHAP library](https://towardsdatascience.com/explain-your-model-with-the-shap-values-bc36aac4de3d)

Implement a program that takes as input the trained model, the customer id ... and returns:

- the score and the SHAP force plot associated with it
- Plotly visualization that show:
  - key variables describing the client and its loan(s)
  - comparison between this client and other clients

Choose the 3 clients of your choice, compute the score, run the visualizations on their data and save them.

- Take 2 clients from the train set:
  - 1 on which the model is correct and the other on which the model is wrong. Try to understand why the model got wrong on this client.
- Take 1 client from the test set

#### Bonus

Implement a dashboard (using [Dash](https://dash.plotly.com/)) that takes as input the customer id and that returns the score and the required visualizations.

### Project repository structure:

```
project
│   README.md
│   requirements.txt
│
└───data
│   │   ...
│
└───results
│   │
|   |───model (free format)
│   │   │   my_own_model.pkl
│   │   │   model_report.txt
│   │
|   |feature_engineering
│   │   │   EDA.ipynb
│   │
|   |───clients_outputs
|   |   |   client1_correct_train.pdf  (free format)
│   │   │   client2_wrong_train.pdf  (free format)
│   │   │   client_test.pdf   (free format)
│   │
|   |───dashboard (optional)
|   |   |   dashboard.py  (free format)
│   │   │   ...
|
|───scripts (free format)
│   │   train.py
│   │   predict.py
│   │   preprocess.py
```

- `README.md` introduces the project, how to run the code, and shows the username.
- `requirements.txt` contains all libraries required to run the code.
- `username.txt` contains the username, the last modified date of the file **has to correspond to the first day of the project**.
- `EDA.ipynb` contains the exploratory data analysis. This file should contain all steps of data analysis that contributed or not to improve the score of the model. It has to be commented so that the reviewer can understand the analysis and run it without any problem.

- `scripts` contains python file(s) that perform(s) the feature engineering, the model's training and prediction on the test set. It could also be one single Jupyter Notebook. It has to be commented to help the reviewers understand the approach and run the code without any bugs.

### Tips

Remember, creating a great credit scoring model is like baking a perfect cake - it takes the right ingredients, careful preparation, and a dash of creativity. You've got this!

### Resources

- [Interpreting machine learning models](https://towardsdatascience.com/interpretability-in-machine-learning-70c30694a05f)
