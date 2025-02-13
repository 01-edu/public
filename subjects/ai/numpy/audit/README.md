#### Exercise 0: Environment and libraries

##### Install the virtual environment with `requirements.txt`

##### Activate the virtual environment. If you used `conda`, run `conda activate ex00`

###### Does the shell specify the name `ex00` of the environment on the left?

##### Run `python --version`

###### Does it print `Python 3.x`? x >= 9

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

##### Add a cell and execute `type(your_numpy_array)`.

###### Is `your_numpy_array` identified as a NumPy array? It should display as `numpy.ndarray`.

##### Execute all the cells within the notebook or use `python main.py`.

###### Can you confirm that the types printed match the following:

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

###### For question 1, does the solution use `np.zeros` and is the shape of the array `(300,)`like bellow?

```console
[0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.]
```

###### For question 2, does the solution use `reshape` and is the shape of the array `(3, 100)` like bellow?

```console
[[0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0.]
 [0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0.]
 [0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0. 0.
  0. 0. 0. 0.]]
```

---

---

#### Exercise 3: Slicing

###### The exercise is validated if the solution doesn't involve a for loop or writing all integers from 1 to 100 and if the array is: `np.array([1,...,100])`. The list from 1 to 100 can be generated with an iterator: `range`. Are the previous requirements fulfilled?

###### For question 1, does the output look like bellow?

```console
[  1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18
  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36
  37  38  39  40  41  42  43  44  45  46  47  48  49  50  51  52  53  54
  55  56  57  58  59  60  61  62  63  64  65  66  67  68  69  70  71  72
  73  74  75  76  77  78  79  80  81  82  83  84  85  86  87  88  89  90
  91  92  93  94  95  96  97  98  99 100]
```

###### For question 2, does the output look like bellow?

```console
[ 1  3  5  7  9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47
 49 51 53 55 57 59 61 63 65 67 69 71 73 75 77 79 81 83 85 87 89 91 93 95
 97 99]
```

###### For question 3, does the output look like bellow?

```console
[100  98  96  94  92  90  88  86  84  82  80  78  76  74  72  70  68  66
  64  62  60  58  56  54  52  50  48  46  44  42  40  38  36  34  32  30
  28  26  24  22  20  18  16  14  12  10   8   6   4   2]
```

###### For question 4, does the output look like bellow?

```console
[  1   0   3   4   0   6   7   0   9  10   0  12  13   0  15  16   0  18
  19   0  21  22   0  24  25   0  27  28   0  30  31   0  33  34   0  36
  37   0  39  40   0  42  43   0  45  46   0  48  49   0  51  52   0  54
  55   0  57  58   0  60  61   0  63  64   0  66  67   0  69  70   0  72
  73   0  75  76   0  78  79   0  81  82   0  84  85   0  87  88   0  90
  91   0  93  94   0  96  97   0  99 100]
```

---

---

#### Exercise 4: Random

> Note: For this exercise, as the results may change depending on the version of the package or the OS, I give the code to correct the exercise. If the code is correct and the output is not the same as mine, it is accepted.

###### For question 1, does the solution contain `np.random.seed(888)`?

###### For question 2, does the solution contain `np.random.randn(100)`?

###### For question 3, does the solution contain `np.random.randint(1,11,(8,8))`?

```console
    Given the NumPy version and the seed, this is my output:

    array([[ 7,  4,  8, 10,  2,  1,  1, 10],
        [ 4,  1,  7,  4,  3,  5,  2,  8],
        [ 3,  9,  7,  4,  9,  6, 10,  5],
        [ 7, 10,  3, 10,  2,  1,  3,  7],
        [ 3,  2,  3,  2, 10,  9,  5,  4],
        [ 4,  1,  9,  7,  1,  4,  3,  5],
        [ 3,  2, 10,  8,  6,  3,  9,  4],
        [ 4,  4,  9,  2,  8,  5,  9,  5]])
```

###### For question 4, does the solution contain `np.random.randint(1,18,(4,2,5))`?

```console
    Given the NumPy version and the seed, this is my output:

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

###### Run the exercise and check if the output is the same as bellow:

```console
[ 1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24
 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48
 49 50]
[ 51  52  53  54  55  56  57  58  59  60  61  62  63  64  65  66  67  68
  69  70  71  72  73  74  75  76  77  78  79  80  81  82  83  84  85  86
  87  88  89  90  91  92  93  94  95  96  97  98  99 100]
[  1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18
  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36
  37  38  39  40  41  42  43  44  45  46  47  48  49  50  51  52  53  54
  55  56  57  58  59  60  61  62  63  64  65  66  67  68  69  70  71  72
  73  74  75  76  77  78  79  80  81  82  83  84  85  86  87  88  89  90
  91  92  93  94  95  96  97  98  99 100]
[[  1   2   3   4   5   6   7   8   9  10]
 [ 11  12  13  14  15  16  17  18  19  20]
 [ 21  22  23  24  25  26  27  28  29  30]
 [ 31  32  33  34  35  36  37  38  39  40]
 [ 41  42  43  44  45  46  47  48  49  50]
 [ 51  52  53  54  55  56  57  58  59  60]
 [ 61  62  63  64  65  66  67  68  69  70]
 [ 71  72  73  74  75  76  77  78  79  80]
 [ 81  82  83  84  85  86  87  88  89  90]
 [ 91  92  93  94  95  96  97  98  99 100]]
```

###### Can you confirm that the student didn't just printed the actual result?

---

---

#### Exercise 6: Broadcasting and Slicing

###### Run the exercise and check if the output is the same as bellow:

```console
[[1 1 1 1 1 1 1 1 1]
 [1 1 1 1 1 1 1 1 1]
 [1 1 1 1 1 1 1 1 1]
 [1 1 1 1 1 1 1 1 1]
 [1 1 1 1 1 1 1 1 1]
 [1 1 1 1 1 1 1 1 1]
 [1 1 1 1 1 1 1 1 1]
 [1 1 1 1 1 1 1 1 1]
 [1 1 1 1 1 1 1 1 1]]

[[1 1 1 1 1 1 1 1 1]
 [1 0 0 0 0 0 0 0 1]
 [1 0 1 1 1 1 1 0 1]
 [1 0 1 0 0 0 1 0 1]
 [1 0 1 0 1 0 1 0 1]
 [1 0 1 0 0 0 1 0 1]
 [1 0 1 1 1 1 1 0 1]
 [1 0 0 0 0 0 0 0 1]
 [1 1 1 1 1 1 1 1 1]]

[[ 1  2  3]
 [ 2  4  6]
 [ 3  6  9]
 [ 4  8 12]
 [ 5 10 15]]

```

##### Check the solution for cheating like:

- The values of the array have been changed one by one manually.
- The usage of the for loop, which is not allowed.
- Printing the full output given in the readme.

###### Can you confirm that there was no cheating in the solution?

---

---

#### Exercise 7: NaN

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

---

---

#### Exercise 8: Wine

###### Was the text file successfully loaded into a NumPy array using `genfromtxt('winequality-red.csv', delimiter=';')` and optimized for memory usage, weighing `76800` bytes or less?

Use this in the solution to confirm:

```Python

# Check the optimized data size
optimized_size = optimized_data.nbytes

# Verify if the dataset size criterion is met
if optimized_size <= 76800:
    print("Data optimized successfully.")
else:
    print("Optimization criteria not met.")
```

##### For question 2:

"Display the 2nd, 7th, and 12th rows as a two-dimensional array. Exclude `np.nan` values if present."

###### Is the output in line with the data present in the provided dataset in the subject?

##### For question 3:

"Determine if there is any wine in the dataset with an alcohol percentage greater than 20%. Return True or False."

###### Is the answer `False`?

##### For question 4:

"Calculate the average alcohol percentage across all wines in the dataset. Exclude `np.nan` values if present."

###### Is the answer `10.422984`?

##### For question 5:

"Compute various statistical measures (minimum, maximum, 25th percentile, 50th percentile, 75th percentile and the mean for the pH values)."

###### Check if you have the correct results as bellow?

```console
    25 percentile:  3.21
    50 percentile:  3.31
    75 percentile:  3.40
    mean:  3.31
    min:  2.74
    max:  4.01
```

> _Note: Using `percentile` or `median` may give different results depending on the duplicate values in the column. If you do not have my results please use `percentile`._

##### For question 6:

"Find the average quality score of wines with the 20% least sulphate content."

###### Is the answer ~`5.2`?

##### For question 7:

Compute the mean of all variables for wines with the best quality. Also, do the same for wines with the worst quality.

###### Is the output for the best wines the following?

```console
[ 8.566666    0.4233333   0.39111114  2.5777776   0.06844445 13.277778
 33.444443    0.99521226  3.2672222   0.76777774 12.094444    8.        ]
```

###### Is the output for the bad wines the following?

```console
[ 8.359999    0.8845      0.17099999  2.6350002   0.12249999 11.
 24.9         0.997464    3.398       0.57000005  9.955       3.        ]
```

---

---

#### Exercise 9: Football tournament

###### Is the output the following?

```console
[[0 1 2 3 5]
[8 7 9 4 6]]
```
