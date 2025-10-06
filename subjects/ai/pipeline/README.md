## Pipeline

### Overview

Today we will focus on the data preprocessing and discover the Pipeline object from scikit learn.

### Role play

You are a data scientist working for a large e-commerce company. The marketing team has provided you with a dataset containing customer information and purchase history. However, the data is messy - it contains categorical variables, missing values, and features on different scales. Your task is to preprocess this data and prepare it for a machine learning model that will predict customer lifetime value.

### Learning Objective

1. Manage categorical variables with Integer encoding and One Hot Encoding.
2. Impute the missing values.
3. Reduce the dimension of the data.
4. Scale the data.

- The **step 1** is always necessary. Models use numbers, for instance string data can't be processed raw.
- The **steps 2** is always necessary. Machine learning models use numbers, missing values do not have mathematical representations, that is why the missing values have to be imputed.
- The **step 3** is required when the dimension of the data set is high. The dimension reduction algorithms reduce the dimensionality of the data either by selecting the variables that contain most of the information (SelectKBest) or by transforming the data. Depending on the signal in the data and the data set size the dimension reduction is not always required. This step is not covered because of its complexity. The understanding of the theory behind is important. However, I suggest to give it a try during the projects.
- The **step 4** is required when using some type of Machine Learning algorithms. The Machine Learning algorithms that require the feature scaling are mostly KNN (K-Nearest Neighbors), Neural Networks, Linear Regression, and Logistic Regression. The reason why some algorithms work better with feature scaling is that the minimization of the loss function may be more difficult if each feature's range is completely different.

These steps are sequential. The output of step 1 is used as input for step 2 and so on; and, the output of step 4 is used as input for the Machine Learning model.
Scikitlearn proposes an object: Pipeline.

As we know, the model evaluation methodology requires splitting the data set in a train set and test set. **The preprocessing is learned/fitted on the training set and applied on the test set**.

This object takes as input the preprocessing transforms and a Machine Learning model. Then this object can be called the same way a Machine Learning model is called. This is pretty practical because we do not need anymore to carry many objects.

### Exercises of the day

- **Exercise 0:** Environment and libraries
- **Exercise 1:** Imputer 1
- **Exercise 2:** Scaler
- **Exercise 3:** One hot Encoder
- **Exercise 4:** Ordinal Encoder
- **Exercise 5:** Categorical variables
- **Exercise 6:** Pipeline

### Virtual Environment

- Python 3.x
- NumPy
- Pandas
- Matplotlib
- Scikit Learn
- Jupyter or JupyterLab

_Version of Scikit Learn we used to do the exercises: 0.22_.
We suggest using the most recent one. Scikit Learn 1.0 is finally available after ... 14 years.

---

---

### Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

**Note:** For each quest, your first exercise will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python.
- the virtual environment you're the most comfortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recent versions of the libraries required

1. Create a virtual environment named `ex00`, with a version of Python >= `3.9`, with the following libraries: `pandas`, `numpy`, `jupyter`, `matplotlib` and `scikit-learn`.

---

---

### Exercise 1: Imputer 1

The goal of this exercise is to learn how to use an `Imputer` to fill missing values on basic example.

```python
train_data = [[7, 6, 5],
              [4, np.nan, 5],
              [1, 20, 8]]
```

1. Fit the `SimpleImputer` on the data. Print the `statistics_`. Check that the statistics match `np.nanmean(train_data, axis=0)`.

2. Fill the missing values in `train_data` using the fitted `imputer` and `transform`.

3. Fill the missing values in `test_data` using the fitted `imputer` and `transform`.

```python
test_data = [[np.nan, 1, 2],
             [7, np.nan, 9],
             [np.nan, 2, 4]]
```

---

---

### Exercise 2: Scaler

The goal of this exercise is to learn to scale a data set. There are various scaling techniques, we will focus on `StandardScaler` from scikit learn.

We will use a tiny data set for this exercise that we will generate by ourselves:

```python
X_train = np.array([[ 1., -1.,  2.],
                     [ 2.,  0.,  0.],
                     [ 0.,  1., -1.]])
```

1. Fit the `StandardScaler` on the data and scale X_train using `fit_transform`. Compute the `mean` and `std` on `axis 0`.

2. Scale the test set using the `StandardScaler` fitted on the train set.

```python
X_test = np.array([[ 2., -1.,  1.],
                     [ 3.,  3.,  -1.],
                     [ 1.,  1., 1.]])
```

**WARNING:
If the data is split in train and test set, it is extremely important to apply the same scaling the test data. As the model is trained on scaled data, if it takes as input unscaled data, it returns incorrect values.**

Resources:

- [Machine Learning](https://medium.com/technofunnel/what-when-why-feature-scaling-for-machine-learning-standard-minmax-scaler-49e64c510422)

- [Preprocessing](https://scikit-learn.org/stable/modules/preprocessing.html)

---

---

### Exercise 3: One hot Encoder

The goal of this exercise is to learn how to deal with Categorical variables using the `OneHot` Encoder.

```python
X_train = [['Python'], ['Java'], ['Java'], ['C++']]
```

1. Using `OneHotEncoder` with `handle_unknown='ignore'`, fit the One Hot Encoder and transform X_train. The expected output is:

   |     | ('C++',) | ('Java',) | ('Python',) |
   | --: | -------: | --------: | ----------: |
   |   0 |        0 |         0 |           1 |
   |   1 |        0 |         1 |           0 |
   |   2 |        0 |         1 |           0 |
   |   3 |        1 |         0 |           0 |

   To get this output create a DataFrame from the transformed X*train and the attribute `categories*`.

2. Transform X_test using the fitted One Hot Encoder on the train set.

```python
X_test = [['Python'], ['Java'], ['C'], ['C++']]
```

The expected output is:

    |    |   ('C++',) |   ('Java',) |   ('Python',) |
    |---:|-----------:|------------:|--------------:|
    |  0 |          0 |           0 |             1 |
    |  1 |          0 |           1 |             0 |
    |  2 |          0 |           0 |             0 |
    |  3 |          1 |           0 |             0 |

- [Resource 1](https://scikit-learn.org/stable/modules/generated/sklearn.preprocessing.OneHotEncoder.html)

---

---

### Exercise 4: Ordinal Encoder

The goal of this exercise is to learn how to deal with Categorical variables using the Ordinal Encoder.

In that case, we want the model to consider that: **good > neutral > bad**

```python
X_train = [['good'], ['bad'], ['neutral']]
```

1. Fit the `OrdinalEncoder` by specifying the categories in the following order: `categories=[['bad', 'neutral', 'good']]`. Transform the train set. Print the `categories_`

2. Transform the X_test using the fitted Ordinal Encoder on train set.

```python
X_test = [['good'], ['good'], ['bad']]
```

_Note: In the version 0.22 of Scikit-learn, the Ordinal Encoder doesn't handle new values in the test set. But it will be possible in the version 0.24 !_

---

---

### Exercise 5: Categorical variables

The goal of this exercise is to learn how to deal with Categorical variables with Ordinal Encoder, Label Encoder and One Hot Encoder. For this exercise I strongly suggest using a recent version of `sklearn >= 0.24.1` to avoid issues with the Ordinal Encoder.

Preliminary:

- Load the [breast-cancer.csv](./data/breast-cancer.csv) file
- Drop `Class` column
- Drop NaN values
- Split the data in a train set and test set (test set size = 20% of the total size) with `random_state=43`.

1. Count the number of unique values per feature in the train set.

2. Identify the variables ordinal variables, nominal variables and the target. Compute a One Hot Encoder transformation on the test set for all categorical features (no ordinal) in the following order `['node-caps' , 'breast', 'breast-quad', 'irradiat']`. Here are the assumptions made on the variables:

```console
age: Ordinal
['90-99' > '80-89' > '70-79' > '60-69' > '50-59' > '40-49' > '30-39' > '20-29' > '10-19']

menopause: Ordinal
['ge40'> 'premeno' >'lt40']

tumor-size: Ordinal
['55-59' > '50-54' > '45-49' > '40-44' > '35-39' > '30-34' > '25-29' > '20-24' > '15-19' > '10-14' > '5-9' > '0-4']

inv-nodes: Ordinal
['36-39' > '33-35' > '30-32' > '27-29' > '24-26' > '21-23' > '18-20' > '15-17' > '12-14' > '9-11' > '6-8' > '3-5' > '0-2']

node-caps: One Hot
['yes' 'no']

deg-malig: Ordinal
[3 > 2 > 1]

breast: One Hot
['left' 'right']

breast-quad: One Hot
['right_low' 'left_low' 'left_up' 'central' 'right_up']

irradiat: One Hot
['yes' 'no']

Class: Target (One Hot)
['recurrence-events' 'no-recurrence-events']
```

- Fit on the train set

- Transform the test set

Example of expected output:

```console
# One Hot encoder on: ['node-caps' , 'breast', 'breast-quad', 'irradiat']

input: ohe.transform(X_test[ohe_cols])[:10]
output:
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

input: ohe.get_feature_names_out(ohe_cols)
output:
array(['node-caps_no', 'node-caps_yes', 'breast_left', 'breast_right',
       'breast-quad_central', 'breast-quad_left_low',
       'breast-quad_left_up', 'breast-quad_right_low',
       'breast-quad_right_up', 'irradiat_no', 'irradiat_yes'],
      dtype=object)

```

3. Create one Ordinal encoder for all Ordinal features in the following order `["menopause", "age", "tumor-size","inv-nodes", "deg-malig"]` on the test set. The documentation of Scikit-learn is not clear on how to perform this on many columns at the same time. Here's a **hint**:

If the ordinal data set is (subset of two columns, but I keep all rows for this example):

    |    | menopause     |   deg-malig |
    |---:|:--------------|------------:|
    |  0 | premeno       |           3 |
    |  1 | ge40          |           1 |
    |  2 | ge40          |           2 |
    |  3 | premeno       |           3 |
    |  4 | premeno       |           2 |

The first step is to create a dictionary or a list - the most recent version of sklearn take as input lists:

```console
dict_ = {0: ['lt40', 'premeno' , 'ge40'], 1:[1,2,3]}
```

Then to instantiate an `OrdinalEncoder`:

```console
oe = OrdinalEncoder(dict_)
```

Now that you have enough information:

- Fit on the train set
- Transform the test set

4. Use a `make_column_transformer` to combine the two Encoders.

- Fit on the train set
- Transform the test set

_Hint: Check the first resource_

**Note: The version 0.22 of Scikit-learn can't handle `get_feature_names` on `OrdinalEncoder`. If the column transformer contains an `OrdinalEncoder`, the method returns this error**:

```console
AttributeError: Transformer ordinalencoder (type OrdinalEncoder) does not provide get_feature_names.
```

**It means that if you want to use the Ordinal Encoder, you will have to create a variable that contains the columns name in the right order. This step is not required in that exercise**

Resources:

- [Resource 2](https://towardsdatascience.com/guide-to-encoding-categorical-features-using-scikit-learn-for-machine-learning-5048997a5c79)

---

---

### Exercise 6: Pipeline

The goal of this exercise is to learn to use the Scikit-learn object: Pipeline. The data set: used for this exercise is the `iris` data set.

Preliminary:

- Run the code below.

  ```console
  iris = load_iris()
  X, y = iris['data'], iris['target']

  #add missing values
  X[[1,20,50,100,135], 0] = np.nan
  X[[2,5,88,135], 1] = np.nan
  X[[4,15], 2] = np.nan
  X[[40,135], 3] = np.nan
  ```

- Split the data set in a train set and test set (33%), fit the Pipeline on the train set and predict on the test set. Use `random_state=43`.

The pipeline you will implement has to contain 3 steps:

- Imputer (median)
- Standard Scaler
- LogisticRegression

1. Train the pipeline on the train set and predict on the test set. Give the score of the model on the test set.

---

---

### Resources

#### Step 3

- [Dimensionality reduction](https://www.geeksforgeeks.org/machine-learning/dimensionality-reduction/)

#### Step 4

- [Feature Scaling](https://medium.com/@societyofai/simplest-way-for-feature-scaling-in-gradient-descent-ae0aaa383039#:~:text=Feature%20scaling%20is%20an%20idea,of%20convergence%20of%20gradient%20descent)

#### Pipeline

- [Pipeline](https://scikit-learn.org/stable/modules/generated/sklearn.pipeline.Pipeline.html)
