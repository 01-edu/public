#### Linear regression with Scikit Learn

#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`

##### Run `python --version`

###### Does it print `Python 3.x`? x >= 8

###### Do `import jupyter`, `import numpy`, `import pandas`, `import matplotlib` and `import sklearn` run without any error?

---

---

#### Exercise 1: Scikit-learn estimator

###### For question 1, is the output the following?

```python
    array([[3.96013289]])
```

###### For question 2, is the output the following?

```output
    Coefficients:  [[0.99667774]]
    Intercept:  [-0.02657807]
    Score:  0.9966777408637874
```

---

---

#### Exercise 2: Linear regression in 1D

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, does the plot look like the following?

![alt text][q1]

[q1]: ../w2_day1_ex2_q1.png "Scatter plot"

###### For question 2, is the equation of the fitted line the following? `y = 42.619430291366946 * x + 99.18581817296929`

###### For question 3, does the plot look like the following?

![alt text][q3]

[q3]: ../w2_day1_ex2_q3.png "Scatter plot + fitted line"

###### For question 4, is the outputted prediction for the first 10 values the following?

```python
array([ 83.86186727, 140.80961751, 116.3333897 ,  64.52998689,
        61.34889539, 118.10301628,  57.5347917 , 117.44107847,
       108.06237908,  85.90762675])
```

###### For question 5, is the MSE returned `114.17148616819485`?

###### For question 6, is the MSE returned `2854.2871542048706`?

---

---

#### Exercise 3: Train test split

###### For question 1, do X_train, y_train, X_test, y_test match this output?

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
 [1 2 3 4 5 6 7 8]


X_test:
 [[17 18]
 [19 20]]


y_test:
 [ 9 10]
```

---

---

#### Exercise 4: Forecast diabetes progression

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is the output of `y_train.values[:10]` and `y_test.values[:10]` the following?

```python
    print(y_train.values[:10])
    [202.  55. 202.  42. 214. 173. 118.  90. 129. 151.]
    print(y_test.values[:10])
    [ 71.  72. 235. 277. 109.  61. 109.  78.  66. 192.]
```

###### For question 2, are the coefficients and the intercept the following?

```console
    [('age', -60.40163046086952),
    ('sex', -226.08740652083418),
    ('bmi', 529.383623302316),
    ('bp', 259.96307686274605),
    ('s1', -859.121931974365),
    ('s2', 504.70960058378813),
    ('s3', 157.42034928335502),
    ('s4', 226.29533600601638),
    ('s5', 840.7938070846119),
    ('s6', 34.712225788519554),
    ('intercept', 152.05314895029233)]
```

###### For question 3, is the output of `predictions_on_test[:10]`?

```console
    array([[111.74351759],
        [ 98.41335251],
        [168.36373195],
        [255.05882934],
        [168.43764643],
        [117.60982186],
        [198.86966323],
        [126.28961941],
        [117.73121787],
        [224.83346984]])
```

###### For question 4, is the mse on the **train set** `2888.326888` and the mse on the **test set** `2858.255153`?

---

---

#### Exercise 5: Gradient Descent (Optional)

##### The exercise is validated if all questions of the exercise are validated.

###### +For question 1, does the outputted plot looks like?

![alt text][ex5q1]

[ex5q1]: ../w2_day1_ex5_q1.png "Scatter plot "

###### +For question 2, is the output `11808.867339751561`?

###### +For question 3, is `grid.shape` equal to `(640000,2)`?

###### +For question 4, are the 10 first values of losses the following?

```console
array([158315.41493175, 158001.96852692, 157689.02212209, 157376.57571726,
    157064.62931244, 156753.18290761, 156442.23650278, 156131.79009795,
    155821.84369312, 155512.39728829])
```

###### +For question 5, does the outputted plot look like the following?

![alt text][ex5q5]

[ex5q5]: ../w2_day1_ex5_q5.png "MSE"

###### +For question 6, is the point returned the following?

`array([42.5, 99. ])`. It means that `a= 42.5` and `b=99`.

###### +For question 7, are the coefficients returned the following?

```console
Coefficients (a): 42.61943031121358
Intercept (b): 99.18581814447936
```

###### +For question 8, is the outputted plot the following?

![alt text][ex5q8]

[ex5q8]: ../w2_day1_ex5_q8.png "MSE + Gradient descent"

###### +For question 9, are the coefficients and intercept returned the following?

```console
Coefficients:  [42.61943029]
Intercept:  99.18581817296929
```
