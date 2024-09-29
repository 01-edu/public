#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated.

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`.

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 8

###### Do `import jupyter`, `import numpy`, `import pandas`, `import matplotlib` and `import sklearn` run without any error?

---

---

#### Exercise 1: Imputer 1

##### The exercise is validated if all questions of the exercise are validated.

###### For question 1 is validated if the `imp_mean.statistics_` returns:

```console
    array([ 4., 13.,  6.])
```

###### For question 2, is the filled train set the following?

```console
    array([[ 7.,  6.,  5.],
        [ 4., 13.,  5.],
        [ 1., 20.,  8.]])
```

###### For question 3, is the filled test set the following?

```console
    array([[ 4.,  1.,  2.],
        [ 7., 13.,  9.],
        [ 4.,  2.,  4.]])
```

---

---

#### Exercise 2: Scaler

##### The exercise is validated if all questions of the exercise are validated.

###### For question 1, is the scaled train set as below? And by definition, the mean on the axis 0 should be `array([0., 0., 0.])` and the standard deviation on the axis 0 should be `array([1., 1., 1.])`.

```console
array([[ 0.        , -1.22474487,  1.33630621],
       [ 1.22474487,  0.        , -0.26726124],
       [-1.22474487,  1.22474487, -1.06904497]])
```

###### For question 2, is the scaled test set the following?

```console
array([[ 1.22474487, -1.22474487,  0.53452248],
       [ 2.44948974,  3.67423461, -1.06904497],
       [ 0.        ,  1.22474487,  0.53452248]])
```

---

---

#### Exercise 3: One hot Encoder

##### The exercise is validated if all questions of the exercise are validated.

###### For question 1, is the output the following?

    |    |   ('C++',) |   ('Java',) |   ('Python',) |
    |---:|-----------:|------------:|--------------:|
    |  0 |          0 |           0 |             1 |
    |  1 |          0 |           1 |             0 |
    |  2 |          0 |           1 |             0 |
    |  3 |          1 |           0 |             0 |

###### For question 2, is the output the following?

    |    |   ('C++',) |   ('Java',) |   ('Python',) |
    |---:|-----------:|------------:|--------------:|
    |  0 |          0 |           0 |             1 |
    |  1 |          0 |           1 |             0 |
    |  2 |          0 |           0 |             0 |
    |  3 |          1 |           0 |             0 |

---

---

#### Exercise 4: Ordinal Encoder

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is the output of the Ordinal Encoder on the train set the following?

```console
array([[2.],
       [0.],
       [1.]])
```

Check that `enc.categories_` returns`[array(['bad', 'neutral', 'good'], dtype=object)]`.

###### For question 2, is the output of the Ordinal Encoder on the test set the following?

```console
array([[2.],
       [2.],
       [0.]])
```

---

---

#### Exercise 5: Categorical variables

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, are the number of unique values per feature outputted the following?

```console
age             6
menopause       3
tumor-size     11
inv-nodes       6
node-caps       2
deg-malig       3
breast          2
breast-quad     5
irradiat        2
dtype: int64
```

###### For question 2, is the transformed test set by the `OneHotEncoder` fitted on the train set as below? Make sure the transformer takes as input a dataframe with the columns in the order defined `['node-caps' , 'breast', 'breast-quad', 'irradiat']` :

```console
#First 10 rows:

array([[1., 0., 1., 0., 0., 1., 0., 0., 0., 1., 0.],
       [1., 0., 1., 0., 0., 1., 0., 0., 0., 1., 0.],
       [0., 1., 1., 0., 0., 1., 0., 0., 0., 0., 1.],
       [0., 1., 1., 0., 0., 1., 0., 0., 0., 0., 1.],
       [1., 0., 1., 0., 0., 0., 1., 0., 0., 1., 0.],
       [1., 0., 1., 0., 0., 0., 0., 1., 0., 1., 0.],
       [1., 0., 0., 1., 0., 0., 0., 0., 1., 1., 0.],
       [1., 0., 0., 1., 0., 1., 0., 0., 0., 1., 0.],
       [1., 0., 1., 0., 0., 0., 0., 1., 0., 0., 1.],
       [1., 0., 0., 1., 0., 1., 0., 0., 0., 1., 0.]])

```

###### For question 3, is the transformed test set by the `OrdinalEncoder` fitted on the train set as below with the columns ordered as `["menopause", "age", "tumor-size","inv-nodes", "deg-malig"]`?

```console
#First 10 rows:

array([[2., 5., 2., 0., 1.],
       [2., 5., 2., 0., 0.],
       [2., 5., 4., 5., 2.],
       [1., 4., 5., 1., 1.],
       [2., 5., 5., 0., 2.],
       [1., 2., 1., 0., 1.],
       [1., 2., 8., 0., 1.],
       [2., 5., 2., 0., 0.],
       [2., 5., 5., 0., 2.],
       [1., 2., 3., 0., 0.]])

```

###### For question 4, is the column transformer transformed that is fitted on the X_train, transformed the X_test as below?

```console
# First 2 rows:

array([[1., 0., 1., 0., 0., 1., 0., 0., 0., 1., 0., 2., 5., 2., 0., 1.],
       [1., 0., 1., 0., 0., 1., 0., 0., 0., 1., 0., 2., 5., 2., 0., 0.]])
```

---

---

#### Exercise 6: Pipeline

###### For question 1, are the predictions on the test set the following?

```console
array([0, 0, 2, 1, 2, 0, 2, 1, 1, 1, 0, 1, 2, 0, 1, 1, 0, 0, 2, 2, 0, 0,
       0, 2, 2, 2, 0, 1, 0, 0, 1, 0, 1, 1, 2, 2, 1, 2, 1, 1, 1, 2, 1, 2,
       0, 1, 1, 1, 1, 1])
```

and is the score on the test set **98%**?

**Note: Keep in mind that having a 98% accuracy is not common when working with real life data. Every time you have a score > 97% check that there's no leakage in the data. On financial data set, the ratio signal to noise is low. Trying to forecast stock prices is a difficult problem. Having an accuracy higher than 70% should be interpreted as a warning to check data leakage!**
