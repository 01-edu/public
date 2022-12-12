# Forest Cover Type Prediction

The goal of this project is to use cartographic variables to classify forest categories. You will have to analyse the data, create features and to train a machine learning model on the cartographic data to make it as accurate as possible.

### Preliminary

###### Does the structure of the project is as below ?


The expected structure of the project is:

```
project
│   README.md
│   environment.yml
│
└───data
│   │   train.csv
│   |   test.csv (not available first day)
|   |   covtype.info
│
└───notebook
│   │   EDA.ipynb
|
|───scripts
|   │   preprocessing_feature_engineering.py
|   │   model_selection.py
│   |   predict.py
│
└───results
    │   confusion_matrix_heatmap.png
    │   learning_curve_best_model.png
    │   test_predictions.csv
    │   best_model.pkl

```

###### Does the readme file contain a description of the project, explain how to run the code from an empty environment, give a summary of the implementation of each python file, especially details on the feature engineering which is a key step ?


###### Does the environment contain all libraries used and their versions that are necessary to run the code ?



### 1. Preprocessing and features engineering:



## 2. Model selection and predict

### Data splitting

###### Does data splitting (cross-validation) structure as follow ?

```
DATA
└───TRAIN FILE (0)
│   └───── Train (1):
│   |           Fold0:
|   |                  Train
|   |                  Validation
|   |           Fold1:
|   |                   Train
|   |                   Validation
... ...         ...
|   |
|   └───── Test (1)
│
└─── TEST FILE (0)(available last day)

```

##### The train set (0) id divised in a train set (1) and test set (1). The ratio is less than 33%.
##### The cross validation splits the train set (1) is at least 5 folds. If the cross validation is stratified that's a good point but it is not a requirement.

### Gridsearch

##### It contains at least these 5 different models: Gradient Boosting, KNN, Random Forest, SVM, Logistic Regression.

There are many options:
- 5 grid searches on 1 model
- 1 grid search on 5 models
- 1 grid search on a pipeline that contains the preprocessing
- 5 grid searches on a pipeline that contains the preprocessing

### Training

###### Is the **target is removed from the X** matrix ?

### Results

##### Run predict.py on the test set, check that: Test (last day) accuracy > **0.65**.

##### Train accuracy score < **0.98**.
It can be checked on the learning curve. If you are not sure, load the model, load the training set (0), score on the training set (0).

##### The confusion matrix is represented as a DataFrame. Example:
![alt text][confusion_matrix]

[confusion_matrix]: ../images/w2_weekend_confusion_matrix.png "Confusion matrix "

##### The learning curve for the best model is plotted. Example:

![alt text][logo_learning_curve]

[logo_learning_curve]: ../images/w2_weekend_learning_curve.png "Learning curve "

Note: The green line on the plot shows the accuracy on the validation set not on the test set (1) and not on the test set (0).

###### Is the trained model saved as a pickle file ?
