#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated.

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`

##### Run `python --version`

###### Does it print `Python 3.x`? x >= 9?

###### Does `import jupyter`, `import numpy`, `import pandas`, `import matplotlib` and `import sklearn` run without any error?

---

---

#### Exercise 1: Logistic regression with Scikit-learn

###### Is the predicted class for question 1 `0`?

###### Are the predicted probabilities for question 2 `[0.61450526 0.38549474]`?

###### Is the output for question 3 like this?

```console
Coefficient:
 [[0.81786797]]
Intercept:
 [-0.87522391]
Score:
 0.7142857142857143
```

---

---

#### Exercise 2: Sigmoid

###### Does the plot for question 1 look like this?

![alt text][ex2q1]

[ex2q1]: ../w2_day2_ex2_q1.png "Scatter plot"

---

---

#### Exercise 3: Decision boundary

##### The exercise is validated if all questions of the exercise are validated

###### Does the outputted plot for question 1 look like this?

![alt text][ex3q1]

[ex3q1]: ../w2_day2_ex3_q1.png "Scatter plot"

###### Are the coefficient and the intercept of the Logistic Regression for question 2 these?

```console
Intercept:  [-0.98385574]
Coefficient:  [[1.18866075]]
```

###### Does the plot for question 3 look like this?

![alt text][ex3q2]

[ex3q2]: ../w2_day2_ex3_q3.png "Scatter plot"

###### For question 4, does `predict_probability` output the same probabilities as `predict_proba`? Note that the values have to match one of the class probabilities, not both. To do so, compare the output with: `clf.predict_proba(X)[:,1]`. The shape of the arrays is not important.

###### Does `predict_class` output the same classes as `cfl.predict(X)` for question 5? The shape of the arrays is not important.

###### Does the plot for question 6 look like the plot below? As mentioned, it is not required to shift the class prediction to make the plot easier to understand.

![alt text][ex3q6]

[ex3q6]: ../w2_day2_ex3_q5.png "Scatter plot + Logistic regression + predictions"

###### Does the plot look like this for question 7?

![alt text][ex3q7]

[ex3q7]: ../w2_day2_ex3_q6.png "Logistic regression decision boundary"

---

---

#### Exercise 4: Train test split

##### The exercise is validated if all questions of the exercise are validated

###### Do X_train, y_train, X_test, y_test match the output below for question 1? The proportion of class `1` is **0.125** in the train set and **1.** in the test set.

```console
X_train:
 [[ 1  2]
 [ 3  4]
 [ 5  6]
 [ 7  8]
 [ 9 10]
 [11 12]
 [13 14]
 [15 16]]


y_train:
 [0. 0. 0. 0. 0. 0. 0. 1.]


X_test:
 [[17 18]
 [19 20]]


y_test:
 [1. 1.]
```

###### Is the proportion of class `1` **0.3** for both sets in question 2?

---

---

#### Exercise 5: Breast Cancer prediction

###### The exercise is validated if all questions of the exercise are validated

###### Is the proportion of class `Benign` 0.6552217453505007 for question 1? It means that if you always predict `Benign` your accuracy would be 66%.

###### Is the proportion of one of the classes approximately the same on the train and test set: ~0.65 for question 2? In my case:

- test: 0.6571428571428571
- train: 0.6547406082289803

###### Is this the output for question 3?

```console
# Train
Class prediction on train set:
 [4 2 4 2 2 2 2 4 2 2]

Probability prediction on train set:
 [0.99600415 0.00908666 0.99992744 0.00528803 0.02097154 0.00582772
 0.03565076 0.99515326 0.00788281 0.01065484]

Score on train set:
 0.9695885509838998

 #Test

 Class prediction on test set:
 [2 2 2 4 2 4 2 2 2 4]

Probability prediction on test set:
 [0.01747203 0.22495309 0.00698756 0.54020801 0.0015289  0.99862249
 0.33607994 0.01227679 0.00438157 0.99972344]

Score on test set:
 0.9642857142857143

```

Only the 10 first predictions are outputted. The score is computed on all the data in the folds.
For some reasons, you may have a different data splitting as mine. The requirement for this question is to have a score on the test set bigger than 92%.

If the score is 1, congratulate you peer, he's just leaked his first target. The target should be dropped from the X_train or X_test ;) !

###### Is the confusion matrix on the train set similar to this in question 4?

```console
array([[357,   9],
       [  8, 185]])
```

and if the confusion matrix on the test set is similar to:

```console
array([[90,  2],
       [ 3, 45]])
```

As said, for some reasons, the results may be slightly different from mine because of the data splitting. However, the values in the confusion matrix should be close to these results.

---

---

#### Bonus

#### Exercise 6: Multi-class (Optional)

##### The exercise is validated if all questions of the exercise are validated

###### +Does each classifier have as input a binary data as below for question 1?

```python
def train(X_train, y_train):
       clf = LogisticRegression()
       clf1 = LogisticRegression()
       clf2 = LogisticRegression()

       clf.fit(X_train, y_train == 0)
       clf1.fit(X_train, y_train == 1)
       clf2.fit(X_train, y_train == 2)

       return clf, clf1, clf2
```

###### +Are this the predicted classes on the test set for question 2?

```console
array([0, 0, 2, 1, 2, 0, 2, 1, 1, 1, 0, 1, 2, 0, 1, 1, 0, 0, 2, 2, 0, 0,
       0, 2, 2, 2, 0, 1, 0, 0])
```

Even if I had this warning `ConvergenceWarning: lbfgs failed to converge (status=1):` I noticed that `LogisticRegression` returns the same output.
