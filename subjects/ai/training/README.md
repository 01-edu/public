## Training

### Overview

Today we will learn how to train and evaluate a machine learning model. You'll learn how to choose the right Machine Learning metric depending on the problem you are solving and to compute it. A metric gives an idea of how good the model performs. Depending on working on a classification problem or a regression problem the metrics considered are different. It is important to understand that all metrics are just metrics, not the truth.

### Role Play

You are a machine learning engineer at a tech startup that specializes in predictive analytics for various industries. Your team has been tasked with developing and evaluating machine learning models for two key projects:

1. A real estate price prediction tool for the California housing market.
2. A medical diagnostics system for breast cancer detection.

### Learning Objectives

- Regression:
  - **R2**, **Mean Square Error**, **Mean Absolute Error**
- Classification:
  - **F1 score**, **accuracy**, **precision**, **recall** and **AUC scores**. Even if it not considered as a metric, the **confusion matrix** is always useful to understand the model performance.

Warning: **Imbalanced data set**

Let us assume we are predicting a rare event that occurs less than 2% of the time. Having a model that scores a good accuracy is easy, it doesn't have to be "smart", all it has to do is to always predict the majority class. Depending on the problem it can be disastrous. For example, working with real life data, breast cancer prediction is an imbalanced problem where predicting the majority leads to disastrous consequences. That is why metrics as AUC are useful. Before to compute the metrics, read carefully this article to understand the role of these metrics.

You'll learn to train other types of Machine Learning models than linear regression and logistic regression. You're not supposed to spend time understanding the theory. I recommend to do that during the projects. Today, read the Scikit-learn documentation to have a basic understanding of the models you use. Focus on how to use correctly those Machine Learning models with Scikit-learn.

You'll also learn what is a grid-search and how to use it to train your machine learning models.

### Exercises of the day

- Exercise 0: Environment and libraries
- Exercise 1: MSE Scikit-learn
- Exercise 2: Accuracy Scikit-learn
- Exercise 3: Regression
- Exercise 4: Classification
- Exercise 5: Machine Learning models
- Exercise 6: Grid Search

### Virtual Environment

- Python 3.x
- NumPy
- Pandas
- Matplotlib
- Scikit Learn
- Jupyter or JupyterLab

_Version of Scikit Learn I used to do the exercises: 0.22_. I suggest to use the most recent one. Scikit Learn 1.0 is finally available after ... 14 years.

### Resources

#### Metrics

- https://medium.com/analytics-vidhya/different-metrics-to-evaluate-the-performance-of-a-machine-learning-model-90acec9e8726

#### Imbalance datasets

- https://stats.stackexchange.com/questions/260164/auc-and-class-imbalance-in-training-test-dataset

#### Gridsearch

- https://www.dremio.com/wiki/grid-search/

---

---

### Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

**Note:** For each quest, your first exercise will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python.
- the virtual environment you're the most comfortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recent versions of the libraries required

1. Create a virtual environment named `ex00`, with a version of Python >= `3.9`, with the following libraries: `pandas`, `numpy`, `jupyter`, `matplotlib` and `scikit-learn`.

---

---

### Exercise 1: MSE Scikit-learn

The goal of this exercise is to learn to use `sklearn.metrics` to compute the mean squared error (MSE).

1. Compute the MSE using `sklearn.metrics` on `y_true` and `y_pred` below:

```python
y_true = [91, 51, 2.5, 2, -5]
y_pred = [90, 48, 2, 2, -4]
```

---

---

### Exercise 2: Accuracy Scikit-learn

The goal of this exercise is to learn to use `sklearn.metrics` to compute the accuracy.

1. Compute the accuracy using `sklearn.metrics` on `y_true` and `y_pred` below:

```python
y_pred = [0, 1, 0, 1, 0, 1, 0]
y_true = [0, 0, 1, 1, 1, 1, 0]
```

---

---

### Exercise 3: Regression

The goal of this exercise is to learn to evaluate a machine learning model using many regression metrics.

Preliminary:

- Import California Housing data set and split it in a train set and a test set (10%). Fit a linear regression on the data set. _The goal is focus on the metrics, that is why the code to fit the Linear Regression is given._

```python
# imports
from sklearn.datasets import fetch_california_housing
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.preprocessing import StandardScaler
from sklearn.impute import SimpleImputer
from sklearn.pipeline import Pipeline
# data
housing = fetch_california_housing()
X, y = housing['data'], housing['target']
# split data train test
X_train, X_test, y_train, y_test = train_test_split(X,
                                                    y,
                                                    test_size=0.1,
                                                    shuffle=True,
                                                    random_state=13)
# pipeline
pipeline = [('imputer', SimpleImputer(strategy='median')),
            ('scaler', StandardScaler()),
            ('lr', LinearRegression())]
pipe = Pipeline(pipeline)
# fit
pipe.fit(X_train, y_train)
```

1. Predict on the train set and test set

2. Compute R2, Mean Square Error, Mean Absolute Error on both train and test set

---

---

### Exercise 4: Classification

The goal of this exercise is to learn to evaluate a machine learning model using many classification metrics.

Preliminary:

- Import Breast Cancer data set and split it in a train set and a test set (20%). Fit a logistic regression on the data set. _The goal is focus on the metrics, that is why the code to fit the logistic Regression is given._

```python
from sklearn.linear_model import LogisticRegression
from sklearn.datasets import load_breast_cancer
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler

X , y = load_breast_cancer(return_X_y=True)
X_train, X_test, y_train, y_test = train_test_split(
    X, y, test_size=0.20, random_state=43)
scaler = StandardScaler()
X_train_scaled = scaler.fit_transform(X_train)
classifier = LogisticRegression()
classifier.fit(X_train_scaled, y_train)
```

1. Predict on the train set and test set

2. Compute F1, accuracy, precision, recall, roc_auc scores on the train set and test set. Print the confusion matrix on the test set results.

**Note: AUC can only be computed on probabilities, not on classes.**

3. Plot the AUC curve for on the test set using roc_curve of scikit learn. There many ways to create this plot. It should look like this:

![alt text][logo_ex4]

[logo_ex4]: ./w2_day4_ex4_q3.png "ROC AUC "

- https://scikit-learn.org/1.1/modules/generated/sklearn.metrics.plot_roc_curve.html

---

---

### Exercise 5: Machine Learning models

The goal of this exercise is to have an overview of the existing Machine Learning models and to learn to call them from scikit learn.
We will focus on:

- SVM/SVC
- Decision Tree
- Random Forest (Ensemble learning)
- Gradient Boosting (Ensemble learning, Boosting techniques)

All these algorithms exist in two versions: regression and classification. Even if the logic is similar in both classification and regression, the loss function is specific to each case.

It is really easy to get lost among all the existing algorithms. This article is very useful to have a clear overview of the models and to understand which algorithm use and when. https://www.geeksforgeeks.org/choosing-a-suitable-machine-learning-algorithm/

Preliminary:

- Import California Housing data set and split it in a train set and a test set (10%). Fit a linear regression on the data set. _The goal is to focus on the metrics, that is why the code to fit the Linear Regression is given._

```python
# imports
from sklearn.datasets import fetch_california_housing
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.preprocessing import StandardScaler
from sklearn.impute import SimpleImputer
from sklearn.pipeline import Pipeline
# data
housing = fetch_california_housing()
X, y = housing['data'], housing['target']
# split data train test
X_train, X_test, y_train, y_test = train_test_split(X,
                                                    y,
                                                    test_size=0.1,
                                                    shuffle=True,
                                                    random_state=43)
# pipeline
pipeline = [('imputer', SimpleImputer(strategy='median')),
            ('scaler', StandardScaler()),
            ('lr', LinearRegression())]
pipe = Pipeline(pipeline)
# fit
pipe.fit(X_train, y_train)

```

1. Create 5 pipelines with 5 different models as final estimator (keep the imputer and scaler unchanged):
   1. Linear Regression
   2. SVM
   3. Decision Tree (set `random_state=43`)
   4. Random Forest (set `random_state=43`)
   5. Gradient Boosting (set `random_state=43`)

Take time to have basic understanding of the role of the basic hyperparameter and their default value.

- For each algorithm, print the R2, MSE and MAE on both train set and test set.

---

---

### Exercise 6: Grid Search

The goal of this exercise is to learn how to make an exhaustive search over specified parameter values for an estimator. This is very useful because the hyperparameter which are the parameters of the model impact the performance of the model.

The scikit learn object that runs the Grid Search is called GridSearchCV. We will learn tomorrow about the cross validation. For now, let us set the parameter **cv** to `[(np.arange(18576), np.arange(18576,20640))]`.
This means that GridSearchCV splits the data set in a train and test set.

Preliminary:

- Load the California Housing data set. As precised, this time, there's no need to split the data set in train set and test set since GridSearchCV does it.

You will have to run a Grid Search on the Random Forest on at least the hyperparameter that are mentioned below. It doesn't mean these are the only hyperparameter of the model. If possible, try at least 3 different values for each hyperparameter.

1. Run a Grid Search with `n_jobs` set to `-1` to parallelize the computations on all CPUs. The hyperparameter to change are: n_estimators, max_depth, min_samples_leaf. It may take

Now, let us analyse the grid search's results in order to select the best model.

2. Write a function that takes as input the Grid Search object and that returns the best model **fitted**, the best set of hyperparameter and the associated score:

   ```python
   def select_model_verbose(gs):

       return trained_model, best_params, best_score
   ```

3. Use the trained model to predict on a new point:

```python
new_point = np.array([[3.2031, 52., 5.47761194, 1.07960199, 910., 2.26368159, 37.85, -122.26]])
```

How do we know the best model returned by GridSearchCV is good enough and stable ? That is what we will learn tomorrow !

**WARNING: Some combinations of hyper parameters are not possible. For example using the SVM, the kernel linear has no parameter gamma.**

**Note**:

- GridSearchCV can also take a Pipeline instead of a Machine Learning model. It is useful to combine some Imputers or Dimension reduction techniques with some Machine Learning models in the same Pipeline.
- It may be useful to check on Kaggle if some Kagglers share their Grid Searches.

Ressources:

- https://scikit-learn.org/stable/modules/generated/sklearn.model_selection.GridSearchCV.html

- https://stackoverflow.com/questions/38555650/try-multiple-estimator-in-one-grid-search

- https://www.dremio.com/wiki/grid-search/

- https://elutins.medium.com/grid-searching-in-machine-learning-quick-explanation-and-python-implementation-550552200596

- https://scikit-learn.org/stable/auto_examples/model_selection/plot_grid_search_digits.html