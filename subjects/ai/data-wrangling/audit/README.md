#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated.

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`.

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 9

###### Does `import jupyter`, `import numpy` and `import pandas`, `matplotlib`, `tabulate` run without any error?

---

---

#### Exercise 1: Concatenate

###### Is the outputted DataFrame as below for question 1?

    |    | letter   |   number |
    |---:|:---------|---------:|
    |  0 | a        |        1 |
    |  1 | b        |        2 |
    |  2 | c        |        1 |
    |  3 | d        |        2 |

###### Check that the index is RangeIndex(start=0, stop=4, step=1) and no manual index assignment

---

---

#### Exercise 2: Merge

##### The exercise is validated if all questions of the exercise are validated.

###### Does the output for question 1 look as below?

    |    |   id | Feature1_x   | Feature2_x   | Feature1_y   | Feature2_y   |
    |---:|-----:|:-------------|:-------------|:-------------|:-------------|
    |  0 |    1 | A            | B            | K            | L            |
    |  1 |    2 | C            | D            | M            | N            |

###### Does the output for question 2 look as below?

    |    |   id | Feature1_df1   | Feature2_df1   | Feature1_df2   | Feature2_df2   |
    |---:|-----:|:---------------|:---------------|:---------------|:---------------|
    |  0 |    1 | A              | B              | K              | L              |
    |  1 |    2 | C              | D              | M              | N              |
    |  2 |    3 | E              | F              | nan            | nan            |
    |  3 |    4 | G              | H              | nan            | nan            |
    |  4 |    5 | I              | J              | nan            | nan            |
    |  5 |    6 | nan            | nan            | O              | P              |
    |  6 |    7 | nan            | nan            | Q              | R              |
    |  7 |    8 | nan            | nan            | S              | T              |

    Note: Check that the suffixes are set using the suffix parameters rather than manually changing the columns' name.

---

---

#### Exercise 3: Merge MultiIndex

##### The exercise is validated if all questions of the exercise are validated.

###### Is the outputted DataFrame's shape `(1305, 5)` and `merged.head()` returns a table as below for question 1? One of the answers that returns the correct DataFrame is `market_data.merge(alternative_data, how='left', left_index=True, right_index=True)`

|                                                      |      Open |    Close | Close_Adjusted |     Twitter |    Reddit |
| :--------------------------------------------------- | --------: | -------: | -------------: | ----------: | --------: |
| (Timestamp('2021-01-01 00:00:00', freq='B'), 'AAPL') | 0.0991792 | -0.31603 |       0.634787 | -0.00159041 |   1.06053 |
| (Timestamp('2021-01-01 00:00:00', freq='B'), 'FB')   | -0.123753 |  1.00269 |       0.713264 |   0.0142127 | -0.487028 |
| (Timestamp('2021-01-01 00:00:00', freq='B'), 'GE')   |  -1.37775 | -1.01504 |         1.2858 |    0.109835 |   0.04273 |
| (Timestamp('2021-01-01 00:00:00', freq='B'), 'AMZN') |   1.06324 | 0.841241 |      -0.799481 |   -0.805677 |  0.511769 |
| (Timestamp('2021-01-01 00:00:00', freq='B'), 'DAI')  | -0.603453 | -2.06141 |      -0.969064 |     1.49817 |  0.730055 |

###### For question 2, are the numbers that are missing in the DataFrame equal to 0 and `filled_df.sum().sum() == merged_df.sum().sum()` gives: `True`?

###### Check that the resulting index is still MultiIndex and no reset_index was used.

---

---

#### Exercise 4: Groupby Apply

##### The exercise is validated if all questions of the exercise are validated and if the for loop hasn't been used. The goal is to use `groupby` and `apply`.

###### Is the output for question 1 the following?

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

###### Is the output for question 2 a Pandas Series or DataFrame with the first 11 rows equal to the output below? The code below gives a solution.

    |    |   sequence |
    |---:|-----------:|
    |  0 |       1.45 |
    |  1 |       2    |
    |  2 |       3    |
    |  3 |       4    |
    |  4 |       5    |
    |  5 |       6    |
    |  6 |       7    |
    |  7 |       8    |
    |  8 |       9    |
    |  9 |       9.55 |
    | 10 |      11.45 |

```python
    def winsorize(df_series, quantiles):
    """
        df: pd.DataFrame or pd.Series
        quantiles: list [0.05, 0.95]

    """
    min_value = np.quantile(df_series, quantiles[0])
    max_value = np.quantile(df_series, quantiles[1])

    return df_series.clip(lower = min_value, upper = max_value)


    df.groupby("group")[['sequence']].apply(winsorize, [0.05,0.95])
```

- [https://towardsdatascience.com/how-to-use-the-split-apply-combine-strategy-in-pandas-groupby-29e0eb44b62e](https://pandas.pydata.org/docs/user_guide/groupby.html)

---

---

#### Exercise 5: Groupby Agg

###### Is the output for question 1 as below? The columns don't have to be MultiIndex. A solution could be `df.groupby('product').agg({'value':['min','max','mean']})`

| product      | ('value', 'min') | ('value', 'max') | ('value', 'mean') |
| :----------- | ---------------: | ---------------: | ----------------: |
| chair        |            22.89 |            32.12 |            27.505 |
| mobile phone |              100 |           111.22 |            105.61 |
| table        |            20.45 |            99.99 |             51.22 |

---

---

#### Exercise 6: Unstack

###### Is the output similar (as the values are generated randomly, it's obvious the audit doesn't require to match the values below) to what `unstacked_df.head()`returns for question 1?

    | Date                |   ('Prediction', 'AAPL') |   ('Prediction', 'AMZN') |   ('Prediction', 'DAI') |   ('Prediction', 'FB') |   ('Prediction', 'GE') |
    |:--------------------|-------------------------:|-------------------------:|------------------------:|-----------------------:|-----------------------:|
    | 2021-01-01 00:00:00 |                 0.382312 |                -0.072392 |               -0.551167 |             -0.0585555 |                1.05955 |
    | 2021-01-04 00:00:00 |                -0.560953 |                 0.503199 |               -0.79517  |             -3.23136   |                1.50271 |
    | 2021-01-05 00:00:00 |                 0.211489 |                 1.84867  |                0.287906 |             -1.81119   |                1.20321 |

###### Is the answer for question 2: `unstacked.plot(title = 'Stocks 2021')`? The title can be anything else.
