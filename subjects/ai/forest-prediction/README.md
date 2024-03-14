## Forest Prediction

The goal of this project is to use cartographic variables to classify forest categories. You will have to analyse the data, create features and to train a machine learning model on the cartographic data to make it as accurate as possible.

### Data

The input files are `train.csv`, `test.csv` and `covtype.info`:

- `train.csv`
- `test.csv`
- `covtype.info`

The train data set is used to **analyse the data and calibrate the models**. The goal is to get the accuracy as high as possible on the test set. The test set will be available at the end of the last day to prevent from the overfitting of the test set.

The data is described in `covtype.info`.

### Structure

The structure of the project is:

```console
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
    │   plots
    │   test_predictions.csv
    │   best_model.pkl

```

### 1. EDA and feature engineering

- Create a Jupyter Notebook to analyse the data sets and perform EDA (Exploratory Data Analysis). This notebook will not be evaluated.

- _Hint: Examples of interesting features_

  - `Distance to hydrology = sqrt((Horizontal_Distance_To_Hydrology)^2 + (Vertical_Distance_To_Hydrology)^2)`
  - `Horizontal_Distance_To_Fire_Points - Horizontal_Distance_To_Roadways`

### 2. Model Selection

The model selection approach is a key step because, t should return the best model and guaranty that the results are reproducible on the final test set. The goal of this step is to make sure that the results on the test set are not due to test set overfitting. It implies to split the data set as shown below:

```console
DATA
└───TRAIN FILE (0)
│   └───── Train (1)
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
└─── TEST FILE (0) (available last day)

```

**Rules:**

- Split train test
- Cross validation: at least 5 folds
- Grid search on at least 5 different models:
  - Gradient Boosting, KNN, Random Forest, SVM, Logistic Regression. _Remember that for some model scaling the data is important and for others it doesn't matter._
- Train accuracy score < **0.98**. Train set (0). Write the result in the `README.md`
- Test (last day) accuracy > **0.65**. Test set (0). Write the result in the `README.md`
- Display the confusion matrix for the best model in a DataFrame. Precise the index and columns names (True label and Predicted label)
- Plot the learning curve for the best model
- Save the trained model as a [pickle](https://docs.python.org/3/library/pickle.html) file

> Advice: As the grid search takes time, I suggest preparing and test the code. Once you are confident it works, run the gridsearch at night and analyse the results

**Hint**: The confusion matrix shows the misclassifications class per class. Try to detect if the model misclassifies badly one class with another. Then, do some research on the internet on the two forest cover types, find the differences and create some new features that underline these differences. More generally, the methodology of a models learning is a cycle with several iterations. More details [here](https://serokell.io/blog/machine-learning-testing)

### 3. Predict (last day)

Once you have selected the best model and you are confident it will perform well on new data, you're ready to predict on the test set:

- Load the trained model
- Predict on the test set and compute the accuracy
- Save the predictions in a csv file
- Add your score on the `README.md`

### Files needed for this project

[link](https://assets.01-edu.org/ai-branch/piscine-ai/raid02/raid02-20221024T133335Z-001.zip)
