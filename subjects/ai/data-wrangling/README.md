## Data Wrangling

### Overview

Data wrangling is one of the crucial tasks in data science and analysis.

### Role Play

You are a newly hired data analyst at a major e-commerce company. Your first assignment is to clean and prepare various datasets for analysis. The company's data comes from multiple sources and in different formats. Your manager has tasked you with combining these datasets, dealing with missing or inconsistent data, and preparing summary reports. You'll need to use your data wrangling skills to transform raw data into a format suitable for analysis and visualization.

### Learning Objectives

- **Data Sorting:** To rearrange values in ascending or descending order.
- **Data Filtration:** To create a subset of available data.
- **Data Reduction:** To eliminate or replace unwanted values.
- **Data Access:** To read or write data files.
- **Data Processing:** To perform aggregation, statistical, and similar operations on specific values.

As explained before, Pandas is an open-source library, specifically developed for data science and analysis. It is built upon the NumPy package (to handle numeric data in tabular form) and has inbuilt data structures to ease the process of data manipulation, aka data munging/wrangling.

### Exercises of the Day

- **Exercise 0:** Environment and libraries
- **Exercise 1:** Concatenate
- **Exercise 2:** Merge
- **Exercise 3:** Merge MultiIndex
- **Exercise 4:** Groupby Apply
- **Exercise 5:** Groupby Agg
- **Exercise 6:** Unstack

### Virtual Environment

- Python 3.x
- NumPy
- Pandas
- Matplotlib
- Jupyter or JupyterLab
- Tabulate

_Version of Pandas we used to do the exercises: 1.0.1_.
We suggest using the most recent one.

---

---

### Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

**Note:** For each quest, your first exercice will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python.
- the virtual environment you're the most confortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recents versions of the libraries required

1. Create a virtual environment named `ex00`, with a version of Python >= `3.9`, with the following libraries: `pandas`, `numpy` ,`tabulate` and `jupyter`.

---

---

### Exercise 1: Concatenate

The goal of this exercise is to learn to concatenate DataFrames. The logic is the same for the Series.

Here are the two DataFrames to concatenate:

```python
df1 = pd.DataFrame([['a', 1], ['b', 2]],
                   columns=['letter', 'number'])
df2 = pd.DataFrame([['c', 1], ['d', 2]],
                   columns=['letter', 'number'])
```

1. Concatenate this two DataFrames on index axis and reset the index. The index of the outputted should be `RangeIndex(start=0, stop=4, step=1)`. **Do not change the index manually**.

---

---

### Exercise 2: Merge

The goal of this exercise is to learn to merge DataFrames
The logic of merging DataFrames in Pandas is quite similar as the one used in SQL.

Here are the two DataFrames to merge:

```python
#df1

df1_dict = {
        'id': ['1', '2', '3', '4', '5'],
        'Feature1': ['A', 'C', 'E', 'G', 'I'],
        'Feature2': ['B', 'D', 'F', 'H', 'J']}

df1 = pd.DataFrame(df1_dict, columns = ['id', 'Feature1', 'Feature2'])

#df2
df2_dict = {
        'id': ['1', '2', '6', '7', '8'],
        'Feature1': ['K', 'M', 'O', 'Q', 'S'],
        'Feature2': ['L', 'N', 'P', 'R', 'T']}

df2 = pd.DataFrame(df2_dict, columns = ['id', 'Feature1', 'Feature2'])
```

1. Merge the two DataFrames to get this output:

   |     |  id | Feature1_x | Feature2_x | Feature1_y | Feature2_y |
   | --: | --: | :--------- | :--------- | :--------- | :--------- |
   |   0 |   1 | A          | B          | K          | L          |
   |   1 |   2 | C          | D          | M          | N          |

2. Merge the two DataFrames to get this output:

   |     |  id | Feature1_df1 | Feature2_df1 | Feature1_df2 | Feature2_df2 |
   | --: | --: | :----------- | :----------- | :----------- | :----------- |
   |   0 |   1 | A            | B            | K            | L            |
   |   1 |   2 | C            | D            | M            | N            |
   |   2 |   3 | E            | F            | nan          | nan          |
   |   3 |   4 | G            | H            | nan          | nan          |
   |   4 |   5 | I            | J            | nan          | nan          |
   |   5 |   6 | nan          | nan          | O            | P            |
   |   6 |   7 | nan          | nan          | Q            | R            |
   |   7 |   8 | nan          | nan          | S            | T            |

---

---

### Exercise 3: Merge MultiIndex

The goal of this exercise is to learn to merge DataFrames with MultiIndex.
Use the code below to generate the DataFrames. `market_data` contains fake market data. In finance, the market is available during the trading days (business days). `alternative_data` contains fake alternative data from social media. This data is available every day. But, for some reasons the Data Engineer lost the last 15 days of alternative data.

1. Using `market_data` as the reference, merge `alternative_data` on `market_data`

   ```python
   #generate days
   all_dates = pd.date_range('2021-01-01', '2021-12-15')
   business_dates = pd.bdate_range('2021-01-01', '2021-12-31')

   #generate tickers
   tickers = ['AAPL', 'FB', 'GE', 'AMZN', 'DAI']

   #create indexs
   index_alt = pd.MultiIndex.from_product([all_dates, tickers], names=['Date', 'Ticker'])
   index = pd.MultiIndex.from_product([business_dates, tickers], names=['Date', 'Ticker'])

   # create DFs
   market_data = pd.DataFrame(index=index,
                           data=np.random.randn(len(index), 3),
                           columns=['Open','Close','Close_Adjusted'])

   alternative_data = pd.DataFrame(index=index_alt,
                                   data=np.random.randn(len(index_alt), 2),
                                   columns=['Twitter','Reddit'])
   ```

`reset_index` is not allowed for this question

2. Fill missing values with 0

- https://medium.com/swlh/merging-dataframes-with-pandas-pd-merge-7764c7e2d46d

---

---

### Exercise 4: Groupby Apply

The goal of this exercise is to learn to group the data and apply a function on the groups.
The use case we will work on is computing

1.  Create a function that uses `pandas.DataFrame.clip` and that replace extreme values by a given percentile. The values that are greater than the upper percentile 80% are replaced by the percentile 80%. The values that are smaller than the lower percentile 20% are replaced by the percentile 20%. This process that correct outliers is called **winsorizing**.
    I recommend to use NumPy to compute the percentiles to make sure we used the same default parameters.

```python
            def winsorize(df, quantiles):
                """
                    df: pd.DataFrame
                    quantiles: list
                        ex: [0.05, 0.95]
                """
                #TODO
                return
```

Here is what the function should output:

```python
            df = pd.DataFrame(range(1,11), columns=['sequence'])
            print(winsorize(df, [0.20, 0.80]).to_markdown())

```

        |    |   sequence |
        |---:|-----------:|
        |  0 |        2.8 |
        |  1 |        2.8 |
        |  2 |        3   |
        |  3 |        4   |
        |  4 |        5   |
        |  5 |        6   |
        |  6 |        7   |
        |  7 |        8   |
        |  8 |        8.2 |
        |  9 |        8.2 |

2.  Now we consider that each value belongs to a group. The goal is to apply the **winsorizing to each group**. In this question we use winsorizing values that are common: `[0.05,0.95]` as percentiles. Here is the new data set:

    ```python
    groups = np.concatenate([np.ones(10), np.ones(10)+1,  np.ones(10)+2, np.ones(10)+3, np.ones(10)+4])

    df = pd.DataFrame(data= zip(groups,
                                range(1,51)),
                    columns=["group", "sequence"])
    ```

    The expected output (first rows) is:

    |     | sequence |
    | --: | -------: |
    |   0 |     1.45 |
    |   1 |        2 |
    |   2 |        3 |
    |   3 |        4 |
    |   4 |        5 |
    |   5 |        6 |
    |   6 |        7 |
    |   7 |        8 |
    |   8 |        9 |
    |   9 |     9.55 |
    |  10 |    11.45 |

---

---

### Exercise 5: Groupby Agg

The goal of this exercise is to learn to compute different type of aggregations on the groups. This small DataFrame contains products and prices.

|     |  value | product      |
| --: | -----: | :----------- |
|   0 |  20.45 | table        |
|   1 |  22.89 | chair        |
|   2 |  32.12 | chair        |
|   3 | 111.22 | mobile phone |
|   4 |  33.22 | table        |
|   5 |    100 | mobile phone |
|   6 |  99.99 | table        |

1. Compute the min, max and mean price for each product in one single line of code. The expected output is:

| product      | ('value', 'min') | ('value', 'max') | ('value', 'mean') |
| :----------- | ---------------: | ---------------: | ----------------: |
| chair        |            22.89 |            32.12 |            27.505 |
| mobile phone |              100 |           111.22 |            105.61 |
| table        |            20.45 |            99.99 |             51.22 |

Note: The columns don't have to be MultiIndex

---

---

### Exercise 6: Unstack

The goal of this exercise is to learn to unstack a MultiIndex
Let's assume we trained a machine learning model that predicts a daily score on the companies (tickers) below. It may be very useful to unstack the MultiIndex: plot the time series, vectorize the backtest, ...

```python
business_dates = pd.bdate_range('2021-01-01', '2021-12-31')

#generate tickers
tickers = ['AAPL', 'FB', 'GE', 'AMZN', 'DAI']

#create indexs
index = pd.MultiIndex.from_product([business_dates, tickers], names=['Date', 'Ticker'])

# create DFs
market_data = pd.DataFrame(index=index,
                        data=np.random.randn(len(index), 1),
                        columns=['Prediction'])

```

1. Unstack the DataFrame.

The first 3 rows of the DataFrame should like this:

| Date                | ('Prediction', 'AAPL') | ('Prediction', 'AMZN') | ('Prediction', 'DAI') | ('Prediction', 'FB') | ('Prediction', 'GE') |
| :------------------ | ---------------------: | ---------------------: | --------------------: | -------------------: | -------------------: |
| 2021-01-01 00:00:00 |               0.382312 |              -0.072392 |             -0.551167 |           -0.0585555 |              1.05955 |
| 2021-01-04 00:00:00 |              -0.560953 |               0.503199 |              -0.79517 |             -3.23136 |              1.50271 |
| 2021-01-05 00:00:00 |               0.211489 |                1.84867 |              0.287906 |             -1.81119 |              1.20321 |

2. Plot the 5 times series in the same plot using Pandas built-in visualization functions with a title.

### Resources

- [Python Handbook](https://jakevdp.github.io/PythonDataScienceHandbook/)

- [Pandas Cheatsheet](https://pandas.pydata.org/Pandas_Cheat_Sheet.pdf)

- [Pandas tutorial](https://www.learndatasci.com/tutorials/python-pandas-tutorial-complete-introduction-for-beginners/)

- [Pandas iteration](https://towardsdatascience.com/different-ways-to-iterate-over-rows-in-a-pandas-dataframe-performance-comparison-dc0d5dcef8fe)
