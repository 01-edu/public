#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated

##### Install the virtual environment with `requirements.txt`

##### Activate the virtual environment. If you used `conda`, run `conda activate ex00`

###### Does the shell specify the name `ex00` of the environment on the left?

##### Run `python --version`

###### Does it print `Python 3.8.x`? x could be any number from 0 to 9

###### Does `import jupyter` and `import numpy` run without any error?

###### Have you used the following command `jupyter notebook --port 8891`?

###### Is there a file named `Notebook_ex00.ipynb` in the working directory?

###### Is the following markdown code executed in a markdown cell in the first cell?

```
# H1 TITLE
## H2 TITLE
```

###### Does the second cell contain `print("Buy the dip ?")` and return `Buy the dip ?` in the output section?

---

---

#### Exercise 1: Your first NumPy array

##### Add cell and run `type(your_numpy_array)`

###### Is the your_numpy_array an NumPy array? It can be checked with that should be equal to `numpy.ndarray`.

##### Run all the cells of the notebook or `python main.py`

###### Are the types printed are as follows?

```
<class 'int'>
<class 'float'>
<class 'str'>
<class 'dict'>
<class 'list'>
<class 'tuple'>
<class 'set'>
<class 'bool'>
```

##### Delete all the cells you added for the audit and restart the notebook

---

---

#### Exercise 2: Zeros

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, does the solution use `np.zeros` and is the shape of the array `(300,)`?

###### For question 2, does the solution use `reshape` and is the shape of the array `(3, 100)`?

---

---

#### Exercise 3: Slicing

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is validated if the solution doesn't involve a for loop or writing all integers from 1 to 100 and if the array is: `np.array([1,...,100])`. The list from 1 to 100 can be generated with an iterator: `range`. Were the previous requirements fulfilled?

###### For question 2, is the solution `integers[::2]`?

###### For question 3, is the solution `integers[::-2]`?

###### For question 4, is the array `np.array([1,0,3,4,0,...,0,99,100])`? There are at least two ways to get this results without for loop. The first one uses `integers[1::3] = 0` and the second involves creating a boolean array that indexes the array:

```python
mask = (integers+1)%3 == 0
integers[mask] = 0
```

---

---

#### Exercise 4: Random

##### The exercise is validated if all questions of the exercise are validated

##### For this exercise, as the results may change depending on the version of the package or the OS, I give the code to correct the exercise. If the code is correct and the output is not the same as mine, it is accepted.

###### For question 1, is the solution `np.random.seed(888)`?

###### For question 2, is the output of the solution the same as `np.random.randn(100)`? The value of the first element is `0.17620087373662233`.

###### For question 3, is the solution `np.random.randint(1,11,(8,8))`?

```console
    Given the NumPy version and the seed, you should have this output:

    array([[ 7,  4,  8, 10,  2,  1,  1, 10],
        [ 4,  1,  7,  4,  3,  5,  2,  8],
        [ 3,  9,  7,  4,  9,  6, 10,  5],
        [ 7, 10,  3, 10,  2,  1,  3,  7],
        [ 3,  2,  3,  2, 10,  9,  5,  4],
        [ 4,  1,  9,  7,  1,  4,  3,  5],
        [ 3,  2, 10,  8,  6,  3,  9,  4],
        [ 4,  4,  9,  2,  8,  5,  9,  5]])
```

###### For question 4, is the solution `np.random.randint(1,18,(4,2,5))`?

```console
    Given the NumPy version and the seed, you should have this output:

    array([[[14, 16,  8, 15, 14],
            [17, 13,  1,  4, 17]],

        [[ 7, 15,  2,  8,  3],
            [ 9,  4, 13,  9, 15]],

        [[ 5, 11, 11, 14, 10],
            [ 2,  1, 15,  3,  3]],

        [[ 3, 10,  5, 16, 13],
            [17, 12,  9,  7, 16]]])
```

---

---

#### Exercise 5: Split, concatenate, reshape arrays

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is the generated array based on an iterator as `range` or `np.arange`? Check that 50 is part of the array.

###### For question 2, is the generated array based on an iterator as `range` or `np.arange`? Check that 100 is part of the array.

###### For question 3, is the array concatenated this way `np.concatenate(array1,array2)`?

###### For question 4, is the result the following?

```console
    array([[  1,  ... ,  10],
                ...
        [ 91,  ... , 100]])
```

The easiest way is to use `array.reshape(10,10)`.

https://jakevdp.github.io/PythonDataScienceHandbook/ (section: The Basics of NumPy Arrays)

---

---

#### Exercise 6: Broadcasting and Slicing

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is the output the same as the following?

`np.ones([9,9], dtype=np.int8)`

###### For question 2, is the output the following?

```console
    array([[1, 1, 1, 1, 1, 1, 1, 1, 1],
        [1, 0, 0, 0, 0, 0, 0, 0, 1],
        [1, 0, 1, 1, 1, 1, 1, 0, 1],
        [1, 0, 1, 0, 0, 0, 1, 0, 1],
        [1, 0, 1, 0, 1, 0, 1, 0, 1],
        [1, 0, 1, 0, 0, 0, 1, 0, 1],
        [1, 0, 1, 1, 1, 1, 1, 0, 1],
        [1, 0, 0, 0, 0, 0, 0, 0, 1],
        [1, 1, 1, 1, 1, 1, 1, 1, 1]], dtype=int8)
```

##### The solution of question 2 is not accepted if the values of the array have been changed one by one manually. The usage of the for loop is not allowed neither.

Here is an example of a possible solution:

```console
        x[1:8,1:8] = 0
        x[2:7,2:7] = 1
        x[3:6,3:6] = 0
        x[4,4] = 1
```

---

---

#### Exercise 7: NaN

##### The exercise is validated if all questions of the exercise are validated

###### Without having used a for loop or having filled the array manually, is the output the following?

```console
[[ 7.  1.  7.]
[nan  2.  2.]
[nan  8.  8.]
[ 9.  3.  9.]
[ 8.  9.  8.]
[nan  2.  2.]
[ 8.  2.  8.]
[nan  6.  6.]
[ 9.  2.  9.]
[ 8.  5.  8.]]
```

There are two steps in this exercise:

- Create the vector that contains the grade of the first exam if available or the second. This can be done using `np.where`:

```python
    np.where(np.isnan(grades[:, 0]), grades[:, 1],     grades[:, 0])
```

- Add this vector as third column of the array. Here are two ways:

```python
    np.insert(arr = grades, values = new_vector, axis = 1, obj = 2)

    np.hstack((grades, new_vector[:, None]))
```

---

---

#### Exercise 8: Wine

##### The exercise is validated if all questions of the exercise are validated

###### Has the text file successfully been loaded in a NumPy array with `genfromtxt('winequality-red.csv', delimiter=',')` and the reduced arrays weights **76800 bytes**?

###### Is the output the following?

```python
array([[ 7.4   ,  0.7   ,  0.    ,  1.9   ,  0.076 , 11.    , 34.    ,
        0.9978,  3.51  ,  0.56  ,  9.4   ,  5.    ],
      [ 7.4   ,  0.66  ,  0.    ,  1.8   ,  0.075 , 13.    , 40.    ,
        0.9978,  3.51  ,  0.56  ,  9.4   ,  5.    ],
      [ 6.7   ,  0.58  ,  0.08  ,  1.8   ,  0.097 , 15.    , 65.    ,
        0.9959,  3.28  ,  0.54  ,  9.2   ,  5.    ]])
```

This slicing gives the answer `my_data[[1,6,11],:]`.

###### Is the answer False? There are many ways to get the answer: find the maximum or check values greater than 20.

###### Is the answer 10.422983114446529?

###### Is the answer the following?

```console
    pH stats
    25 percentile:  3.21
    50 percentile:  3.31
    75 percentile:  3.4
    mean:  3.3111131957473416
    min:  2.74
    max:  4.01
```

    > *Note: Using `percentile` or `median` may give different results depending on the duplicate values in the column. If you do not have my results please use `percentile`.*

###### Is the answer ~`5.2`? The first step is to get the percentile 20% of the column `sulphates`, then create a boolean array that contains `True` of the value is smaller than the percentile 20%, then select this rows with the column quality and compute the `mean`.

###### Is the output for the best wines the following?

```python
array([ 8.56666667,  0.42333333,  0.39111111,  2.57777778,  0.06844444,
    13.27777778, 33.44444444,  0.99521222,  3.26722222,  0.76777778,
    12.09444444,  8.        ])
```

###### Is the output for the bad wines the following?

```python
array([ 8.36    ,  0.8845  ,  0.171   ,  2.635   ,  0.1225  , 11.      ,
    24.9     ,  0.997464,  3.398   ,  0.57    ,  9.955   ,  3.      ])
```

This can be done in three steps: Get the max, create a boolean mask that indicates rows with max quality, use this mask to subset the rows with the best quality and compute the mean on the axis 0.

---

---

#### Exercise 9: Football tournament

###### Is the output the following?

```console
 [[2 9 6 0 8]
  [1 3 5 7 4]]
```
