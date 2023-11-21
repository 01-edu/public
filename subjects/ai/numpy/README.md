# NumPy

The goal of this day is to understand practical usage of **NumPy**. **NumPy** is a commonly used Python data analysis package. By using **NumPy**, you can speed up your workflow, and interface with other packages in the Python ecosystem, like scikit-learn, that use **NumPy** under the hood. **NumPy** was originally developed in the mid 2000s, and arose from an even older package called Numeric. This longevity means that almost every data analysis or machine learning package for Python leverages **NumPy** in some way.

### Exercises of the day

- Exercise 0: Environment and libraries
- Exercise 1: Your first NumPy array
- Exercise 2: Zeros
- Exercise 3: Slicing
- Exercise 4: Random
- Exercise 5: Split, concatenate, reshape arrays
- Exercise 6: Broadcasting and Slicing
- Exercise 7: NaN
- Exercise 8: Wine
- Exercise 9: Football tournament

### Virtual Environment

- Python 3.x
- NumPy
- Jupyter or JupyterLab

_Version of NumPy I used to do the exercises: 1.18.1_.
I suggest to use the most recent one.

### Resources

- https://medium.com/fintechexplained/why-should-we-use-NumPy-c14a4fb03ee9
- https://numpy.org/doc/
- https://jakevdp.github.io/PythonDataScienceHandbook/

---

---

# Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries and to learn to launch a `jupyter notebook`. Jupyter notebooks are very convenient as they allow to write and test code within seconds. However, it really easy to implement instable and not reproducible code using notebooks. Keep the notebook and the underlying code clean. An article below detail when the Notebook should be used. Notebook can be used for most of the exercises of the piscine as the goal is to experiment A LOT. But no worries, you'll be asked to build a more robust structure for all the projects.

**Note:** For each quest, your first exercise will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python. However, for educational purpose you will install a specific version of Python in this exercise.
- the virtual environment you're the most comfortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recent versions of the libraries required

1. Create a virtual environment named `ex00`, with Python `3.8`, with the following libraries: `numpy`, `jupyter`. Save the installed packages in `requirements.txt` in the current directory.

2. Launch a `jupyter notebook` on port `8891` and create a notebook named `Notebook_ex00`. `JupyterLab` can be used instead of Jupyter Notebook here.

3. Put the text `H1 TITLE` as **heading level 1** and `H2 TITLE` as **heading level 2** in the first cell.

4. Run `print("Buy the dip ?")` in the second cell

### Resources:

- https://www.python.org/
- https://docs.conda.io/
- https://jupyter.org/
- https://numpy.org/
- https://towardsdatascience.com/jypyter-notebook-shortcuts-bf0101a98330
- https://odsc.medium.com/why-you-should-be-using-jupyter-notebooks-ea2e568c59f2
- https://stackoverflow.com/questions/50777849/from-conda-create-requirements-txt-for-pip3

---

---

# Exercise 1: Your first NumPy array

The goal of this exercise is to use many Python data types in **NumPy** arrays. **NumPy** arrays are intensively used in **NumPy** and **Pandas**. They are flexible and allow to use optimized **NumPy** underlying functions.

1. Create a NumPy array that contains: an integer, a float, a string, a dictionary, a list, a tuple, a set and a boolean. Add the following code at the end of your python file or in a cell of the jupyter notebook:

```python
for i in your_np_array:
    print(type(i))
```

---

---

# Exercise 2: Zeros

The goal of this exercise is to learn to create a NumPy array with 0s.

1. Create a NumPy array of dimension **300** with zeros without filling it manually
2. Reshape it to **(3,100)**

---

---

# Exercise 3: Slicing

The goal of this exercise is to learn NumPy indexing/slicing. It allows to access values of the NumPy array efficiently and without a for loop.

1. Create a NumPy array of dimension 1 that contains all integers from 1 to 100 ordered.
2. Without using a for loop and using the array created in Q1, create an array that contain all odd integers. The expected output is: `np.array([1,3,...,99])`. _Hint_: it takes one line
3. Without using a for loop and using the array created in Q1, create an array that contain all even integers reversed. The expected output is: `np.array([100,98,...,2])`. _Hint_: it takes one line

4. Using array of Q1, set the value of every 3 elements of the list (starting with the second) to 0. The expected output is: `np.array([[1,0,3,4,0,...,0,99,100]])`

---

---

# Exercise 4: Random

The goal of this exercise is to learn to generate random data.
In Data Science it is extremely useful to generate random data for many reasons:
Lack of real data, create a random benchmark, use varied data sets.
NumPy proposes a lot of options to generate random data. In statistics, assumptions are made on the distribution the data is from. All data distribution that can be generated randomly are described in the documentation. In this exercise we will focus on two distributions:

- Uniform: For example, if your goal is to generate a random number from 1 to 100 and that the probability that all the numbers is equal you'll need the uniform distribution. NumPy provides `randint` and `uniform` to generate uniform distribution

- Normal: The normal distribution is the most important probability distribution in statistics because it fits many natural phenomena.For example, if you need to generate a data sample that represents **Heights of 14 Year Old Girls** it can be done using the normal distribution. In that case, we need two parameters: the mean (1m51) and the standard deviation (0.0741m). NumPy provides `randn` to generate normal distribution (among other)

https://numpy.org/doc/stable/reference/random/generator.html

1. Set the seed to 888
2. Generate a **one-dimensional** array of size 100 with a normal distribution
3. Generate a **two-dimensional** array of size 8,8 with random integers from 1 to 10 - both included (same probability for each integer)
4. Generate a **three-dimensional** of size 4,2,5 array with random integers from 1 to 17 - both included (same probability for each integer)

---

---

# Exercise 5: Split, concatenate, reshape arrays

The goal of this exercise is to learn to concatenate and reshape arrays.

1. Generate an array with integers from 1 to 50: `array([1,...,50])`

2. Generate an array with integers from 51 to 100: `array([51,...,100])`

3. Using `np.concatenate`, concatenate the two arrays into: `array([1,...,100])`

4. Reshape the previous array into:

   ```console
   array([[  1,  ... ,  10],
               ...
       [ 91,  ... , 100]])
   ```

---

---

# Exercise 6: Broadcasting and Slicing

The goal of this exercise is to learn to access values of n-dimensional arrays efficiently.

1. Create an 2-dimensional array size 9,9 of 1s. Each value has to be an `int8`.
2. Using **slicing**, output this array:

   ```python
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

3. Using **broadcasting** create the ouptu matrix starting from these two arrays:

   ```python
   array_1 = np.array([1,2,3,4,5], dtype=int8)
   array_2 = np.array([1,2,3], dtype=int8)
   ...
   # output matrix
   array([[ 1,  2,  3],
          [ 2,  4,  6],
          [ 3,  6,  9],
          [ 4,  8, 12],
          [ 5, 10, 15]], dtype=int8)
   ```

https://jakevdp.github.io/PythonDataScienceHandbook/ (section: Computation on Arrays: Broadcasting)

---

---

# Exercise 7: NaN

The goal of this exercise is to learn to deal with missing data in NumPy and to manipulate NumPy arrays.

Let us consider a 2-dimensional array that contains the grades at the past two exams. Some of the students missed the first exam. As the grade is missing it has been replaced with a `NaN`.

1. Using `np.where` create a third column that is equal to the grade of the first exam if it exists and the second else. Add the column as the third column of the array.

**Using a for loop or if/else statement is not allowed in this exercise.**

```python
import numpy as np

generator = np.random.default_rng(123)
grades = np.round(generator.uniform(low = 0.0, high = 10.0, size = (10, 2)))
grades[[1,2,5,7], [0,0,0,0]] = np.nan
print(grades)
```

---

---

# Exercise 8: Wine

The goal of this exercise is to learn to perform a basic data analysis on real data using NumPy.

The data set that will be used for this exercise is the red wine data set.

https://archive.ics.uci.edu/ml/datasets/wine+quality

How to tell if a given 2D array has null columns?

1. Using `genfromtxt` load the data and reduce the size of the numpy array by optimizing the types. The sum of absolute differences between the original data set and the "memory" optimized one has to be smaller than 1.10**-3. I suggest to use `np.float32`. Check that the numpy array weights **76800 bytes\*\*.

2. Print 2nd, 7th and 12th rows as a two dimensional array

3. Is there any wine with a percentage of alcohol greater than 20% ? Return True or False

4. What is the average % of alcohol on all wines in the data set ? If needed, drop `np.nan` values

5. Compute the minimum, the maximum, the 25th percentile, the 50th percentile, the 75th percentile, the median (50th percentile) of the pH

6. Compute the average quality of the wines having the 20% least sulphates

7. Compute the mean of all variables for wines having the best quality. Same question for the wines having the worst quality

---

---

# Exercise 9: Football tournament

The goal of this exercise is to learn to use permutations, complex

A Football tournament is organized in your city. There are 10 teams and the director of the tournaments wants you to create a first round as exciting as possible. To do so, you are allowed to choose the pairs. As a former data scientist, you implemented a model based on teams' current season performance. This models predicts the score difference between two teams. You used this algorithm to predict the score difference for every possible pair.
The matrix returned is a 2-dimensional array that contains in (i,j) the score difference between team i and j. The matrix is in [model_forecasts.txt](data/model_forecasts.txt).

Using this output, what are the pairs that will give the most interesting matches ?

If a team wins 7-1 the match is obviously less exciting than a match where the winner wins 2-1.
The criteria that corresponds to **the pairs that will give the most interesting matches** is **the pairs that minimize the sum of squared differences**

The expected output is:

```console
 [[m1_t1 m2_t1 m3_t1 m4_t1 m5_t1]
  [m1_t2 m2_t2 m3_t2 m4_t2 m5_t2]]
```

- m1_t1 stands for match1_team1
- m1_t1 plays against m1_t2 ...

**Usage of for loop is not allowed, you may need to use the library** `itertools` **to create permutations**

https://docs.python.org/3.9/library/itertools.html
