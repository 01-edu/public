#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated.

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`.

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 8

##### Do `import jupyter`, `import numpy`, `import pandas`, `import matplotlib` and `import sklearn` run without any error?

---

---

#### Exercise 1: MSE Scikit-learn

###### Is the Mean Squared Error (MSE) calculated using the `sklearn.metrics` library?

###### Is the Mean Squared Error (MSE) correctly computed for the given `y_true` and `y_pred` values, and does the calculated MSE match the expected value?

---

---

#### Exercise 2: Accuracy Scikit-learn

###### Is the accuracy computed using the `sklearn.metrics` library?

###### Is the accuracy correctly calculated for the given `y_true` and `y_pred` values, and does the calculated accuracy match the expected value?

---

---

#### Exercise 3: Regression

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, are the predictions on the train set and test set the following?

```console
    #10 first values Train
array([1.54505951, 2.21338527, 2.2636205 , 3.3258957 , 1.51710076,
    1.63209319, 2.9265211 , 0.78080924, 1.21968217, 0.72656239])

```

```console
#10 first values Test

array([ 1.82212706,  1.98357668,  0.80547979, -0.19259114,  1.76072418,
        3.27855815,  2.12056804,  1.96099917,  2.38239663,  1.21005304])

```

###### For question 2, Do the results match the following output?

```console
r2 on the train set:  0.3552292936915783
MAE on the train set:  0.5300159371615256
MSE on the train set:  0.5210784446797679

r2 on the test set:  0.30265471284464673
MAE on the test set:  0.5454023699809112
MSE on the test set:  0.5537420654727396
```

This result shows that the model has slightly better results on the train set than the test set. That's frequent since it is easier to get a better grade on an exam we studied than an exam that is different from what was prepared. However, the results are not good: r2 ~ 0.3. Fitting non linear models as the Random Forest on this data may improve the results. That's the goal of the exercise 5.

---

---

#### Exercise 4: Classification

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, are the predictions on the train set and test set the following?

```console
    # 10 first values Train
    array([1, 0, 1, 1, 1, 0, 0, 1, 1, 0])

    # 10 first values Test
    array([1, 1, 0, 0, 0, 1, 1, 1, 0, 0])
```

###### For question 2, do the results match this output?

```console
F1 on the train set:  0.9911504424778761
Accuracy on the train set:  0.989010989010989
Recall on the train set:  0.9929078014184397
Precision on the train set:  0.9893992932862191
ROC_AUC on the train set:  0.9990161111794368


F1 on the test set:  0.9801324503311258
Accuracy on the test set:  0.9736842105263158
Recall on the test set:  0.9866666666666667
Precision on the test set:  0.9736842105263158
ROC_AUC on the test set:  0.9863247863247864
```

###### For question 2, do the results match the confusion matrix on the test set? It should be:

```console
array([[37,  2],
        [ 1, 74]])
```

###### For question 3, Does the ROC AUC plot look like the plot below?

![alt text][logo_ex4]

[logo_ex4]: ../w2_day4_ex4_q3.png 'ROC AUC '

Having a 99% ROC AUC is not usual. The data set we used is easy to classify. On real data sets, always check if there's any leakage while having such a high ROC AUC score.

---

---

#### Exercise 5: Machine Learning models

###### For question 1, are the scores outputted close to the scores below? Some of the algorithms use random steps (random sampling used by the `RandomForest`). I used `random_state = 43` for the Random Forest, the Decision Tree and the Gradient Boosting.

```console
# Linear regression

TRAIN
r2 on the train set:  0.34823544284172625
MAE on the train set:  0.533092001261455
MSE on the train set:  0.5273648371379568

TEST
r2 on the test set:  0.3551785428138914
MAE on the test set:  0.5196420310323713
MSE on the test set:  0.49761195027083804


# SVM

TRAIN
r2 on the train set:  0.6462366150965996
MAE on the train set:  0.38356451633259875
MSE on the train set:  0.33464478671339165

TEST
r2 on the test set:  0.6162644671183826
MAE on the test set:  0.3897680598426786
MSE on the test set:  0.3477101776543003


# Decision Tree

TRAIN
r2 on the train set:  0.9999999999999488
MAE on the train set:  1.3685733933909677e-08
MSE on the train set:  6.842866883530944e-14

TEST
r2 on the test set:  0.6263651902480918
MAE on the test set:  0.4383758696244002
MSE on the test set:  0.4727017198871596


# Random Forest

TRAIN
r2 on the train set:  0.9705418471542886
MAE on the train set:  0.11983836612191189
MSE on the train set:  0.034538356420577995

TEST
r2 on the test set:  0.7504673649554309
MAE on the test set:  0.31889891600404635
MSE on the test set:  0.24096164834441108


# Gradient Boosting

TRAIN
r2 on the train set:  0.7395782392433273
MAE on the train set:  0.35656543036682264
MSE on the train set:  0.26167490389525294

TEST
r2 on the test set:  0.7157456298013534
MAE on the test set:  0.36455447680396397
MSE on the test set:  0.27058170064218096

```

It is important to notice that the Decision Tree overfits very easily. It learns easily the training data but is not able to extrapolate on the test set. This algorithm is not used a lot because of its overfitting ability.

However, Random Forest and Gradient Boosting propose a solid approach to correct the overfitting (in that case the parameters `max_depth` is set to None that is why the Random Forest overfits the data). These two algorithms are used intensively in Machine Learning Projects.

---

---

#### Exercise 6: Grid Search

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is the code that runs the `gridsearch` like the following? (the parameters may change):

```python
parameters = {'n_estimators':[10, 50, 75],
            'max_depth':[3,5,7],
            'min_samples_leaf': [10,20,30]}

rf = RandomForestRegressor()
gridsearch = GridSearchCV(rf,
                        parameters,
                        cv = [(np.arange(18576), np.arange(18576,20640))],
                        n_jobs=-1)
gridsearch.fit(X, y)
```

###### For question 2, is the function as below?

```python
def select_model_verbose(gs):

    return gs.best_estimator_, gs.best_params_, gs.best_score_
```

In my case, the `gridsearch` parameters are not interesting. Even if I reduced the over-fitting of the Random Forest, the score on the test is lower than the score on the test returned by the Gradient Boosting in the previous exercise without optimal parameters search.

###### For question 3, is the code used the following?

```python
model, best_params, best_score = select_model_verbose(gridsearch)
model.predict(new_point)
```
