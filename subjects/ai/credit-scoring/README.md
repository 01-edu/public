# Credit scoring

The goal of this project is to implement a scoring model based on various source of data (check data documentation) that returns the probability of default. In a nutshell, credit scoring represents an evaluation of how well the bank's customer can pay and is willing to pay off debt. It is also required that you provide an explanation of the score. For example, your model returns that the probability that one client doesn't pay back the loan is very high (90%). The reason behind is that variable_xxx which represents the ability to pay back the past loan is low. The output interpretability will appear in a visualization.

The ability to understand the underlying factors of credit scoring is important. Credit scoring is subject to more and more regulation, so transparency is key. And more generaly, more and more companies prefer transparency to black box models.

### Resources

Historical timeline of machine learning techniques applied to credit scoring

- https://hal.archives-ouvertes.fr/hal-02507499v3/document
- https://www.kaggle.com/c/home-credit-default-risk/data

# Deliverables

### Scoring model

The are 3 expected deliverables associated with the scoring model:

- An exploratory data analysis notebook that describes the insights you find out in the data set.
- The trained machine learning model with the features engineering pipeline:

  - Do not forget: **Coming up with features is difficult, time-consuming, requires expert knowledge. ‘Applied machine learning’ is basically feature engineering.**
  - The model is validated if the **AUC on the test set is higher than 75%**.
  - The labelled test data is not publicly available. However a Kaggle competition uses the same data. The procedure to evaluate test set submission is the same as the one used for the project 1.

### Kaggle submission

The way the Kaggle platform works is explained in the challenge overview page. If you need more details, I suggest this resource that gives detailed explanations.

- https://towardsdatascience.com/getting-started-with-kaggle-f9138b35ae18

- Create a username following that structure: username*01EDU* location_MM_YYYY. Submit the description profile and push it on the Git platform the first day of the week. Do not touch this file anymore.

- A text document that describes the methodology used to train the machine learning model:
  - Algorithm
  - Why the accuracy shouldn't be used in that case ?
  - Limit and possible improvements

### Model interpretability

This part hasn't been covered during the piscine. Take the time to understand this key concept.
There are different level of transparency:

- **Global**: understand important variables in a model. This answers the question: "What are the key variables to the model ? ". In that case it will tell if the revenue is more important than the age to the model for example. This allows to check that the model relies on important variables. No one wants his credit to be refused because of the weather in Lisbon !
- **Local**: each observation gets its own set of interpretability factors. This greatly increases its transparency. We can explain why a case receives its prediction and the contributions of the predictors. Traditional variable importance algorithms only show the results across the entire population but not on each individual case. The local interpretability enables us to pinpoint and contrast the impacts of the factors.

There are 2 tools you can use to analyse your model and its predictions: - Features importance (available if you use a Scikit Learn model) - [SHAP library](https://towardsdatascience.com/explain-your-model-with-the-shap-values-bc36aac4de3d)

Implement a program that takes as input the trained model, the customer id ... and returns:

- the score and the SHAP force plot associated with it
- Plotly visualisations that show:
  - key variables describing the client and its loan(s)
  - comparison between this client and other clients

Choose the 3 clients of your choice, compute the score, run the visualizations on their data and save them.

- Take 2 clients from the train set:
  - 1 on which the model is correct and the other on which the model is wrong. Try to understand why the model got wrong on this client.
- Take 1 client from the test set

### Optional

Implement a dashboard (using Dash) that takes as input the customer id and that returns the score and the required visualizations.

- https://stackoverflow.com/questions/54292226/putting-html-output-from-shap-into-the-dash-output-layout-callback

### Deliverables

```
project
│   README.md
│   environment.yml
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

- `README.md` introduces the project and shows the username.
- `environment.yml` contains all libraries required to run the code.
- `username.txt` contains the username, the last modified date of the file **has to correspond to the first day of the project**.
- `EDA.ipynb` contains the exploratory data analysis. This file is should contain all steps of data analysis that contributed or not to improve the score of the model. It has to be commented so that the reviewer can understand the analysis and run it without any problem.

- `scripts` contains python file(s) that perform(s) the feature engineering, the model's training and prediction on the test set. It could also be one single Jupyter Notebook. It has to be commented to help the reviewers understand the approach and run the code without any bugs.

### Useful resources

- https://towardsdatascience.com/interpretability-in-machine-learning-70c30694a05f

### Files needed for this project

[File 1](https://assets.01-edu.org/ai-branch/project5/project05-20221024T130417Z-001.zip)
[File 2](https://assets.01-edu.org/ai-branch/project5/project05-20221024T130417Z-002.zip)
