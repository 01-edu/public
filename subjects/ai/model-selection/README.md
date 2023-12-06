# Model selection

If you finished yesterday's exercises you should be able to train several Machine Learning algorithms and to choose one returned by GridSearchCV.
GridSearchCV returns the model that gives the best score on the test set. Yesterday, as I told you, I changed the **cv** parameter to compute the GridSearch with a train set and a test set.

It means that the selected model is based on one single measure. What if, by luck, we predict correctly on that section ? What if the best model is bad ? What if I could have selected a better model ?

We will answer these questions today ! The topics we will cover are the one of the most important in Machine Learning.

### Exercises of the day

- Exercise 0: Environment and libraries
- Exercise 1: K-Fold
- Exercise 2: Cross validation (k-fold)
- Exercise 3: GridsearchCV
- Exercise 4: Validation curve and Learning curve

### Virtual Environment

- Python 3.x
- NumPy
- Pandas
- Jupyter or JupyterLab
- Scikit-learn
- Matplotlib

_Version of Pandas I used to do the exercises: 1.0.1_.
I suggest to use the most recent one.

### **Resources**

**Must read before to start the exercises**

### Biais-Variance trade off, aka Underfitting/Overfitting:

- https://machinelearningmastery.com/gentle-introduction-to-the-bias-variance-trade-off-in-machine-learning/

- https://jakevdp.github.io/PythonDataScienceHandbook/05.03-hyperparameters-and-model-validation.html

### Cross-validation

- https://algotrading101.com/learn/train-test-split/

---

---

# Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

**Note:** For each quest, your first exercise will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python.
- the virtual environment you're the most comfortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recent versions of the libraries required

1. Create a virtual environment named `ex00`, with a version of Python >= `3.8`, with the following libraries: `pandas`, `numpy`, `jupyter`, `matplotlib` and `scikit-learn`.

---

---

# Exercise 1: K-Fold

The goal of this exercise is to learn to use `KFold` to split the data set in a k-fold cross validation. Most of the time you won't use this function to split your data because this function is used by others as `cross_val_score` or `cross_validate` or `GridSearchCV` ... . But, this allows to understand the splitting and to create a custom one if needed.

```python
X = np.array(np.arange(1,21).reshape(10,-1))
y = np.array(np.arange(1,11))
```

1. Using `KFold`, perform a 5-fold cross validation. For each fold, print the train index and test index. The expected output is:

   ```console
   Fold:  1
   TRAIN: [2 3 4 5 6 7 8 9] TEST: [0 1]

   Fold:  2
   TRAIN: [0 1 4 5 6 7 8 9] TEST: [2 3]

   Fold:  3
   TRAIN: [0 1 2 3 6 7 8 9] TEST: [4 5]

   Fold:  4
   TRAIN: [0 1 2 3 4 5 8 9] TEST: [6 7]

   Fold:  5
   TRAIN: [0 1 2 3 4 5 6 7] TEST: [8 9]
   ```

---

---

# Exercise 2: Cross validation (k-fold)

The goal of this exercise is to learn how to use cross validation. After reading the articles you should be able to explain why we need to cross-validate the models. We will firstly focus on Linear Regression to reduce the computation time. We will be using `cross_validate` to run the cross validation. Note that `cross_val_score` is similar but the `cross_validate` calculates one or more scores and timings for each CV split.

Preliminary:

- Import California Housing data set and split it in a train set and a test set (10%). Fit a linear regression on the data set. _The goal is to focus on the cross validation, that is why the code to fit the Linear Regression is given._

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
```

1. Cross validate the Pipeline using `cross_validate` with 10 folds. Print the scores on each validation sets, the mean score on the validation sets and the standard deviation on the validation sets. The expected output is:

```console
Scores on validation sets:
 [0.62433594 0.61648956 0.62486602 0.59891024 0.59284295 0.61307055
 0.54630341 0.60742976 0.60014575 0.59574508]

Mean of scores on validation sets:
 0.60201392526743

Standard deviation of scores on validation sets:
 0.0214983822773466

```

**Note: It may be confusing that the key of the dictionary that returns the results on the validation sets is `test_score`. Sometimes, the validation sets are called test sets. In that case, we run the cross validation on X_train. It means that the scores are computed on sets in the initial train set. The X_test is not used for the cross-validation.**

- https://scikit-learn.org/stable/modules/generated/sklearn.model_selection.cross_validate.html

- https://machinelearningmastery.com/how-to-configure-k-fold-cross-validation/

---

---

# Exercise 3: GridsearchCV

The goal here is to utilize GridSearchCV for running a grid search, making predictions, and scoring on a test set.

Preliminary:

- Import California Housing dataset, split it into a train and a test set (10%), and fit a linear regression on the dataset.

```python
# imports
from sklearn.datasets import fetch_california_housing
from sklearn.model_selection import train_test_split

# data
housing = fetch_california_housing()
X, y = housing['data'], housing['target']
# split data train test
X_train, X_test, y_train, y_test = train_test_split(X,
                                                    y,
                                                    test_size=0.1,
                                                    shuffle=True,
                                                    random_state=43)
```

1. Run `GridSearchCV` with the following settings:

   - Using all CPUs, perform 5-fold cross-validation.
   - Scoring metric: MSE (Mean Squared Error)
   - Model: Random Forest

   Hyperparameters to search:

   - `max_depth`: range between 1 and 20 (minimum 3 values)
   - `n_estimators`: range between 1 and 100 (minimum 3 values)

   This computation might take a few minutes to run.

_Hint_: The name of the metric to put in the parameter `scoring` is `neg_mean_squared_error`. The smaller the MSE is, the better the model is. At the contrary, The greater the R2 is the better the model is. `GridSearchCV` chooses the best model by selecting the one that maximized the score on the validation sets. And, in mathematic, maximizing a function or minimizing its opposite is equivalent. More details:

- https://stackoverflow.com/questions/21443865/scikit-learn-cross-validation-negative-values-with-mean-squared-error

2. Extract the best fitted estimator, print its parameters, its score on the validation set, and display `cv_results_`.

3. Compute the score on the test set.

**WARNING: For classification tasks using AUC score, an error or warning might occur if a fold contains only one class, rendering the AUC unable to be computed due to its definition.**

---

---

# Exercise 4: Validation curve and Learning curve

The goal of this exercise is to learn to analyze the model's performance with two tools:

- Validation curve
- Learning curve

For this exercise we will use a dataset of 100k data points to give you an idea of the computation time you can expect during projects.

Preliminary:

- Using make_classification from sklearn, generate a binary data set with 100k data points and with 30 features.

```python
X, y = make_classification(n_samples=100000,
                        n_features= 30,
                        n_informative=10,
                        flip_y=0.2 )
```

1. Plot the validation curve, using all CPUs, with 5 folds. The goal is to focus again on max_depth between 1 and 20.
   You may need to increase the window (example: between 1 and 50 ) if you notice that other values of max_depth could have returned better results. This may take few minutes.

I do not expect that you implement all the plot from scratch, you'd better leverage the code here:

- https://scikit-learn.org/stable/auto_examples/model_selection/plot_validation_curve

The plot should look like this:

![alt text][logo_ex5q1]

[logo_ex5q1]: ./w2_day5_ex5_q1.png 'Validation curve '

The interpretation is that from max_depth=10, the train score keeps increasing but the test score (or validation score) reaches a plateau. It means that choosing max_depth = 20 may lead to have an over fitted model.

Note: Given the time computation is is not possible to plot the validation curve for all parameters. It is useful to plot it for parameters that control the over fitting the most.

More details:

- https://chrisalbon.com/machine_learning/model_evaluation/plot_the_validation_curve/

2. Let us assume the gridsearch returned `clf = RandomForestClassifier(max_depth=12)`. Let's check if the models under fits, over fit or fits correctly. Plot the learning curve. These two resources will help you a lot to understand how to analyze the learning curves and how to plot them:

- https://machinelearningmastery.com/learning-curves-for-diagnosing-machine-learning-model-performance/

- https://scikit-learn.org/stable/auto_examples/model_selection/plot_learning_curve.html#sphx-glr-auto-examples-model-selection-plot-learning-curve-py

- **Re-use the function in the second resource**, change the cross validation to a classic 10-folds, run the learning curve data computation on all CPUs and plot the three plots as shown below.

![alt text][logo_ex5q2]

[logo_ex5q2]: ./w2_day5_ex5_q2.png 'Learning curve '

- **Note Plot Learning Curves**: The learning curves is detailed in the first resource.

- **Note Plot Scalability of the model**: This plot shows the relationship between the time to train the model and the number of rows in the data. In that case the relationship is linear.

- **Note Performance of the model**: This plot shows wether it worths increasing the training time by adding data to increase the score. It would worth to add data to increase the score if the curve hasn't reach a plateau yet. In that case, increasing the training time by 10 units increases the score by less than 0.001.
