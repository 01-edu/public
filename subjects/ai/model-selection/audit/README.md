#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated.

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`.

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 9

###### Do `import jupyter`, `import numpy`, `import pandas`, `import matplotlib` and `import sklearn` run without any error?

---

---

#### Exercise 1: K-Fold

###### For question 1, is the output of the 5-fold cross validation the following?

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

#### Exercise 2: Cross validation (k-fold)

###### For question 1, is the output the following?

```console
Scores on validation sets:
 [0.62433594 0.61648956 0.62486602 0.59891024 0.59284295 0.61307055
 0.54630341 0.60742976 0.60014575 0.59574508]

Mean of scores on validation sets:
 0.60201392526743

Standard deviation of scores on validation sets:
 0.0214983822773466

```

The model is consistent across folds: it is stable. That's a first sign that the model is not over-fitted. The average R2 is 60% that's a good start ! To be improved...

---

---

#### Exercise 3: GridsearchCV

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is the code that runs the grid search similar to the following?

```python
parameters = {'n_estimators':[10, 50, 75],
            'max_depth':[4, 7, 10]}

rf = RandomForestRegressor()
gridsearch = GridSearchCV(rf,
                        parameters,
                        cv = 5,
                        n_jobs=-1,
                        scoring='neg_mean_squared_error')

gridsearch.fit(X_train, y_train)
```

The answers that uses another list of parameters are accepted too !

###### For question 2, were these attributes used?

```python
print(gridsearch.best_score_)
print(gridsearch.best_params_)
print(gridsearch.cv_results_)
```

The best score is -0.29028202683007526, that means that the MSE is ~0.29, it doesn't give any information since this metric is arbitrary. This score is the average of `neg_mean_squared_error` on all the validation sets.

The best models params are `{'max_depth': 10, 'n_estimators': 75}`.

Note that if the parameters used are different, the results should be different.

###### For question 3, was the fitted estimator used to compute the score on the test set: `gridsearch.score(X_test, y_test)`? The MSE score is ~0.27. The score I got on the test set is close to the score I got on the validation sets. It means the models is not over fitted.

---

---

#### Exercise 4: Validation curve and Learning curve

###### For question 1, does the outputted plot look like the plot below? The two important points to check are: The training score has to converge towards `1` and the cross-validation score reaches a plateau around `0.9` from `max_depth = 10`

![alt text][logo_ex5q1]

[logo_ex5q1]: ../w2_day5_ex5_q1.png "Validation curve "

The code that generated the data in the plot is:

```python
from sklearn.model_selection import validation_curve

clf = RandomForestClassifier()
param_range = np.arange(1,30,2)
train_scores, test_scores = validation_curve(clf,
                                            X,
                                            y,
                                            param_name="max_depth",
                                            param_range=param_range,
                                            scoring="roc_auc",
                                            n_jobs=-1)
```

###### For question 2, do the outputted plots look like the following?

![alt text][logo_ex5q2]

[logo_ex5q2]: ../w2_day5_ex5_q2.png "Learning curve "
