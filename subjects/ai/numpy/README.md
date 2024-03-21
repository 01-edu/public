## NumPy

The goal of this day is to understand practical usage of **NumPy**. **NumPy** is a commonly used Python data analysis package. By using **NumPy**, you can speed up your workflow, and interface with other packages in the Python ecosystem, like scikit-learn, that use **NumPy** under the hood. **NumPy** was originally developed in the mid 2000s, and arose from an even older package called Numeric. This longevity means that almost every data analysis or machine learning package for Python leverages **NumPy** in some way.

### Virtual Environment

- Python 3.x
- NumPy
- Jupyter or JupyterLab

_Version of NumPy I used to do the exercises: 1.18.1_.
I suggest to use the most recent one.

### Resources

- [Why Should We Use NumPy](https://medium.com/fintechexplained/why-should-we-use-NumPy-c14a4fb03ee9)
- [NumPy Documentation](https://numpy.org/doc/)
- [Python Data Science Handbook](https://jakevdp.github.io/PythonDataScienceHandbook/)

---

---

## Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries and to learn to launch a `jupyter notebook`. Jupyter notebooks are very convenient as they allow to write and test code within seconds. However, it really easy to implement instable and not reproducible code using notebooks. Keep the notebook and the underlying code clean. Notebook can be used for most of the exercises of the piscine as the goal is to experiment a lot. But no worries, you'll be asked to build a more robust structure for all the projects.

**Note:** For each quest, your first exercise will be to set up the virtual environment with the required libraries.

I suggest utilizing:

- The **latest stable version** of Python for your work. However, in this exercise, you'll install and use a specific Python version for educational purposes.
- Choose a virtual environment that aligns with your familiarity. Common choices among Data Science practitioners are `virtualenv` and `conda`.
- Install the most recent versions of the required libraries to ensure compatibility and access to the latest features

1. Begin by creating a virtual environment named `ex00` that utilizes Python version `3.8`. Install the required libraries `numpy` and `jupyter`. Save the installed packages to a file named `requirements.txt`, located in the current directory.

2. Launch a `jupyter` notebook or `JupyterLab` on port `8891`. Create a new notebook named `Notebook_ex00`.

3. In the first cell of the notebook, set `H1 TITLE` as a **heading level 1** and `H2 TITLE` as a **heading level 2**.

4. Execute `print("Buy the dip ?")` in the second cell to display the message.

### Resources:

- [python](https://www.python.org/)
- [Conda Documentation](https://docs.conda.io/)
- [jupyter](https://jupyter.org/)
- [numpy](https://numpy.org/)
- [Jupyter Notebook Shortcuts](https://towardsdatascience.com/jypyter-notebook-shortcuts-bf0101a98330)
- [Why You Should be Using Jupyter Notebooks](https://odsc.medium.com/why-you-should-be-using-jupyter-notebooks-ea2e568c59f2)

---

---

## Exercise 1: Your first NumPy array

The objective of this exercise is to familiarize yourself with incorporating various Python data types into **NumPy** arrays. **NumPy** arrays play a vital role in both **NumPy** and **Pandas**, offering flexibility and optimized functionalities.

1. Create a NumPy array that contains: an `integer`, a `float`, a `string`, a `dictionary`, a `list`, a `tuple`, a `set` and a `boolean`. Add the following code at the end of your python file or in a cell of the jupyter notebook:

```python
for i in your_np_array:
    print(type(i))
```

---

---

## Exercise 2: Zeros

The goal of this exercise is to learn to create a NumPy array with 0s.

1. Create a NumPy array of dimension **300** with zeros without filling it manually
2. Reshape it to **(3,100)**

---

---

## Exercise 3: Slicing

The goal of this exercise is to learn NumPy indexing/slicing. It allows to access values of the NumPy array efficiently and without a for loop.

1. Create a NumPy array of dimension 1 that contains all integers from 1 to 100 ordered.

2. Without using a for loop and using the array created in Q1, create an array that contain all odd integers. The expected output is:

```console
[ 1  3  5  7  9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47
 49 51 53 55 57 59 61 63 65 67 69 71 73 75 77 79 81 83 85 87 89 91 93 95
 97 99]
```

3. Without using a for loop and using the array created in Q1, create an array that contain all even integers reversed. The expected output is:

```console
[100  98  96  94  92  90  88  86  84  82  80  78  76  74  72  70  68  66
  64  62  60  58  56  54  52  50  48  46  44  42  40  38  36  34  32  30
  28  26  24  22  20  18  16  14  12  10   8   6   4   2]
```

4. Using array of Q1, set the value of every 3 elements of the list (starting with the second) to 0. The expected output is:

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

## Exercise 4: Random

The goal of this exercise is to learn to generate random data.
In Data Science it is extremely useful to generate random data for many reasons:
Lack of real data, create a random benchmark, use varied data sets.
NumPy proposes a lot of options to generate random data. In statistics, assumptions are made on the distribution the data is from. All data distribution that can be generated randomly are described in the documentation. In this exercise we will focus on two distributions:

- Uniform: For example, if your goal is to generate a random number from 1 to 100 and that the probability that all the numbers is equal you'll need the uniform distribution. NumPy provides `randint` and `uniform` to generate uniform distribution

- Normal: The normal distribution is the most important probability distribution in statistics because it fits many natural phenomena.For example, if you need to generate a data sample that represents **Heights of 14 Year Old Girls** it can be done using the normal distribution. In that case, we need two parameters: the mean (1m51) and the standard deviation (0.0741m). NumPy provides `randn` to generate normal distribution (among other)

[Random Generator](https://numpy.org/doc/stable/reference/random/generator.html)

1. Set the seed to 888
2. Generate a **one-dimensional** array of size 100 with a normal distribution
3. Generate a **two-dimensional** array of size 8,8 with random integers from 1 to 10 - both included (same probability for each integer)
4. Generate a **three-dimensional** of size 4,2,5 array with random integers from 1 to 17 - both included (same probability for each integer)

---

---

## Exercise 5: Split, concatenate, reshape arrays

The goal of this exercise is to learn to concatenate and reshape arrays.

1. Generate an array with integers from 1 to 50: `array([1,...,50])`

2. Generate an array with integers from 51 to 100: `array([51,...,100])`

3. Using `np.concatenate`, concatenate the two arrays into: `array([1,...,100])`

4. Reshape the previous array into:

   ```console
   [[  1   2   3   4   5   6   7   8   9  10]
    [ 11  12  13  14  15  16  17  18  19  20]
                        ...
    [ 81  82  83  84  85  86  87  88  89  90]
    [ 91  92  93  94  95  96  97  98  99 100]]
   ```

Print what you've created in the previous steps.

---

---

## Exercise 6: Broadcasting and Slicing

The goal of this exercise is to learn to access values of n-dimensional arrays efficiently.

**Using a for loop is not allowed in this exercise.**

1. Generate a 2-dimensional array of size 9x9, with all elements initialized to 1 and of type `int8`.
2. Using **slicing**, create the following array:

   ```python
   array([[1, 1, 1, 1, 1, 1, 1, 1, 1],
       [1, 0, 0, 0, 0, 0, 0, 0, 1],
       [1, 0, 1, 1, 1, 1, 1, 0, 1],
       [1, 0, 1, 0, 0, 0, 1, 0, 1],
       [1, 0, 1, 0, 1, 0, 1, 0, 1],
       [1, 0, 1, 0, 0, 0, 1, 0, 1],
       [1, 0, 1, 1, 1, 1, 1, 0, 1],
       [1, 0, 0, 0, 0, 0, 0, 0, 1],
       [1, 1, 1, 1, 1, 1, 1, 1, 1]], dtype=np.int8)
   ```

3. Using **broadcasting** create an output matrix based on the following two arrays:

   ```python
   array_1 = np.array([1,2,3,4,5], dytpe=np.int8)
   array_2 = np.array([1,2,3], dytpe=np.int8)
   ```

Expected output:

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

### Resources

[Computation on Arrays: Broadcasting](https://jakevdp.github.io/PythonDataScienceHandbook/)

---

---

## Exercise 7: NaN

The goal of this exercise is to handle missing data in NumPy and manipulate arrays effectively.

Let's consider a 2-dimensional array containing grades from the last two exams. Some students missed the first exam, so their grades are replaced with `NaN`.

To simulate this scenario, we'll create a mock dataset using NumPy. Here's a snippet of code to generate this dataset:

```python
import numpy as np

generator = np.random.default_rng(123)
grades = np.round(generator.uniform(low = 0.0, high = 10.0, size = (10, 2)))
grades[[1,2,5,7], [0,0,0,0]] = np.nan
print(grades)
```

This code returns:

```console
[[ 7.  1.]
 [nan  2.]
 [nan  8.]
 [ 9.  3.]
 [ 8.  9.]
 [nan  2.]
 [ 8.  2.]
 [nan  6.]
 [ 9.  2.]
 [ 8.  5.]]
```

1. Using `np.where`, create a third column that takes the grade of the first exam if available; otherwise, it uses the grade from the second exam. Add this column as the third column of the array.

**Using a for loop or if/else statement is not allowed in this exercise.**

Expected output:

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

## Exercise 8: Wine

The goal of this exercise is to perform fundamental data analysis on real data using NumPy.

The dataset chosen for this task was the [red wine dataset](./data/winequality-red.csv). You can find more info [HERE](./data/)

1. Load the data using `genfromtxt`, specifying the delimiter as ';', and optimize the numpy array size by reducing the data types. Use `np.float32` and verify that the resulting numpy array weighs **76800 bytes**.

2. Display the 2nd, 7th, and 12th rows as a two-dimensional array. Exclude `np.nan` values if present.

3. Determine if there is any wine in the dataset with an alcohol percentage greater than 20%. Return True or False.

4. Calculate the average alcohol percentage across all wines in the dataset. Exclude `np.nan` values if present.

5. Compute various statistical measures (minimum, maximum, 25th percentile, 50th percentile, 75th percentile and the mean for the pH values).

> _Note: Using `percentile` or `median` may give different results depending on the duplicate values in the column. If you do not have my results please use `percentile`._

6. Find the average quality score of wines with the 20% least sulphate content.

**Tip:** The first step is to get the percentile 20% of the column `sulphates`, then create a boolean array that contains `True` of the value is smaller than the percentile 20%, then select this rows with the column quality and compute the `mean`.

7. Compute the mean of all variables for wines with the best quality. Also, do the same for wines with the worst quality.

**Tip:** This can be done in three steps: Get the max, create a boolean mask that indicates rows with max quality, use this mask to subset the rows with the best quality and compute the mean on the axis 0.

---

## Exercise 9: Football tournament

This exercise focuses on utilizing permutations and complex computations.

A Football tournament is underway in your city involving 10 teams. The tournament director seeks an engaging first round and has delegated the pairing decisions to you.

Leveraging your expertise as a former data scientist, you've developed a predictive model based on teams' current season performance. This model forecasts the score difference between any two teams.

The model generates a 2-dimensional array stored in [model_forecasts.txt](data/model_forecasts.txt). Each (i, j) entry in this matrix signifies the predicted score difference between Team i and Team j.

The objective is to determine the pairs that will result in the most interesting matches.

If a team wins 7-1 the match is obviously less exciting than a match where the winner wins 2-1.
The criteria that corresponds to **the pairs that will give the most interesting matches** is **the pairs that minimize the sum of squared differences**

The expected output is:

```console
[[m1_t1 m2_t1 m3_t1 m4_t1 m5_t1]
 [m1_t2 m2_t2 m3_t2 m4_t2 m5_t2]]
```

- m1_t1 stands for match1_team1
- m1_t1 plays against m1_t2 ...

**Usage of for loop is not allowed, you may need to use the library [itertools](https://docs.python.org/3.9/library/itertools.html) to create permutations.**
