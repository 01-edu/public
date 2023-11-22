#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated.

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 8

###### Do `import jupyter`, `import numpy` and `import pandas` run without any error?

---

---

#### Exercise 1: Your first DataFrame

##### The exercise is validated if all questions of the exercise are validated.

###### Is the solution of question 1 the same as the "model" DataFrame? Check that the index is not 1,2,3,4,5.

###### For question 2, are the columns' types as below and are the types of the first value of the columns as below?

```console
    <class 'pandas.core.series.Series'>
    <class 'pandas.core.series.Series'>
    <class 'pandas.core.series.Series'>
```

```console
        <class 'str'>
        <class 'list'>
        <class 'float64'>
```

---

---

#### Exercise 2: Electric power consumption

##### The exercise is validated if all questions of the exercise are validated

###### For the solution of question 1 is `drop`  used with `axis=1`.`inplace=True`? It may be useful to avoid to affect the result to a variable. A solution that could be accepted too (even if it's not a solution I recommend) is `del`.

###### For question 2 does the DataFrame return the output below? If the type of the index is not `dtype='datetime64[ns]'` the solution is not accepted. I recommend to use `set_index` with `inplace=True` to do so.

```python
        Input: df.head().index

        Output:

        DatetimeIndex(['2006-12-16', '2006-12-16','2006-12-16', '2006-12-16','2006-12-16'],
        dtype='datetime64[ns]', name='Date', freq=None)
```

###### For question 3, are all the types `float64` as below? The preferred solution is `pd.to_numeric` with `coerce=True`.

```python
        Input: df.dtypes

        Output:

            Global_active_power      float64
            Global_reactive_power    float64
            Voltage                  float64
            Global_intensity         float64
            Sub_metering_1           float64
            dtype: object

```

###### For question 4, is `df.describe()` being used?

###### For question 5 is `dropna` being used and is the number of missing values equal to 0? It is important to notice that 25979 rows contain missing values (for a total of 129895). `df.isna().sum()` allows to check the number of missing values and `df.dropna()` with `inplace=True` allows to remove the rows with missing values.

###### For question 6, were one of the two approaches below used?

```python
        #solution 1
        df.loc[:,'A'] = (df['A'] + 1) * 0.06

        #solution 2
        df.loc[:,'A'] = df.loc[:,'A'].apply(lambda x: (x+1)*0.06)

```

    You may wonder `df.loc[:,'A']` is required and if `df['A'] = ...` works too. **The answer is no**. This is important in Pandas. Depending on the version of Pandas, it may return a warning. The reason is that you are affecting a value to a **copy** of the DataFrame and not in the DataFrame.
    More details: https://stackoverflow.com/questions/20625582/how-to-deal-with-settingwithcopywarning-in-pandas

###### For question 7, is the output of `print(filtered_df.head().to_markdown())` as below and is the number of rows equal to **449667**?

    | Date                |   Global_active_power |   Global_reactive_power |
    |:--------------------|----------------------:|------------------------:|
    | 2008-12-27 00:00:00 |                 0.996 |                   0.066 |
    | 2008-12-27 00:00:00 |                 1.076 |                   0.162 |
    | 2008-12-27 00:00:00 |                 1.064 |                   0.172 |
    | 2008-12-27 00:00:00 |                 1.07  |                   0.174 |
    | 2008-12-27 00:00:00 |                 0.804 |                   0.184 |

###### For question 8, is the output the following?

```console
        Global_active_power        0.254
        Global_reactive_power      0.000
        Voltage                  238.350
        Global_intensity           1.200
        Sub_metering_1             0.000
        Name: 2007-02-16 00:00:00, dtype: float64

```

###### For question 9, is the output `Timestamp('2009-02-22 00:00:00')`?

###### For question 10 is the output of `print(sorted_df.tail().to_markdown())` the following?

    | Date                |   Global_active_power |   Global_reactive_power |   Voltage |
    |:--------------------|----------------------:|------------------------:|----------:|
    | 2008-08-28 00:00:00 |                 0.076 |                       0 |    234.88 |
    | 2008-08-28 00:00:00 |                 0.076 |                       0 |    235.18 |
    | 2008-08-28 00:00:00 |                 0.076 |                       0 |    235.4  |
    | 2008-08-28 00:00:00 |                 0.076 |                       0 |    235.64 |
    | 2008-12-08 00:00:00 |                 0.076 |                       0 |    236.5  |

###### For question 11, is the output the following? The solution is based on `groupby` which creates groups based on the index `Date` and aggregates the groups using the `mean`.

```console
    Date
    2006-12-16    3.053475
    2006-12-17    2.354486
    2006-12-18    1.530435
    2006-12-19    1.157079
    2006-12-20    1.545658
                    ...
    2010-12-07    0.770538
    2010-12-08    0.367846
    2010-12-09    1.119508
    2010-12-10    1.097008
    2010-12-11    1.275571
    Name: Global_active_power, Length: 1433, dtype: float64
```

---

---

#### Exercise 3: E-commerce purchases

##### The exercise is validated if all questions of the exercise are validated.

##### To validate this exercise all answers should return the expected numerical value given in the correction AND use Pandas. For example using NumPy to compute the mean doesn't respect the philosophy of the exercise which is to use Pandas.

###### For question 1, does the solution contain **10000 entries** and **14 columns**? There many solutions based on: shape, info, describe.

###### For question 2, is the answer **50.34730200000025**?

    Even if `np.mean` gives the solution, `df['Purchase Price'].mean()` is preferred

###### For question 3, are the min `0`and the max `99.989999999999995`?

###### For question 4, is the answer **1098**?

###### For question 5, is the answer **30**?

###### For question 6, are there `4932` people that made the purchase during the `AM` and `5068` people that made the purchase during `PM`? There are many ways to get the solution but the goal of this question was to make you use `value_counts`.

###### For question 7, is the answer as below? There are many ways to the solution but the goal of this question was to use `value_counts`.

    Interior and spatial designer    31

    Lawyer                           30

    Social researcher                28

    Purchasing manager               27

    Designer, jewellery              27

###### For question 8, is the purchase price **75.1**?

###### For question 9, is the email address **bondellen@williams-garza.com**?

###### For question 10, is the answer **39**? The preferred solution is based on this: `df[(df['A'] == X) & (df['B'] > Y)]`.

###### For question 11, is the answer **1033**? The preferred solution is based on the usage of `apply` on a `lambda` function that slices the string that contains the expiration date.

###### For question 12, is the answer as below? The preferred solution is based on the usage of `apply` on a `lambda` function that slices the string that contains the email. The `lambda` function uses `split` to split the string on `@`. Finally, `value_counts` is used to count the occurrences.

    - hotmail.com     1638
    - yahoo.com       1616
    - gmail.com       1605
    - smith.com         42
    - williams.com      37

---

---

#### Exercise 4: Handling missing values

##### The exercise is validated if all questions of the exercise are validated (except the bonus question)

###### For question 1, are the two steps implemented in that order? First, convert the numerical columns to `float` and then fill the missing values. The first step may involve `pd.to_numeric(df.loc[:,col], errors='coerce')`. The second step is validated if you eliminated all missing values. However there are many possibilities to fill the missing values. Here is one of them:

    example:

```python
    df.fillna({
        0:df.sepal_length.mean(),
        2:df.sepal_width.median(),
        3:0,
        4:0
    })
```

###### For question 2, is the solution `df.loc[:,col].fillna(df[col].median())` or any equivalent formula?

###### +The solution of bonus question is accepted if you find out this answer: Once we filled the missing values as suggested in the first question, `df.describe()` returns this interesting summary. We notice that the mean is way higher than the median. It means that there are maybe some outliers in the data. The quantile 75 and the max confirm that: 75% of the flowers have a sepal length smaller than 6.4 cm, but the max is 6900 cm. If you check on the internet you realise this small flower can't be that big. The outliers have a major impact on the mean which equals to 56.9. Filling this value for the missing value is not correct since it doesn't correspond to the real size of this flower. That is why in that case the best strategy to fill the missing values is the median. The truth is that I modified the data set ! But real data sets ALWAYS contains outliers. Always think about the meaning of the data transformation! If you fill the missing values by zero, it means that you consider that the length or width of some flowers may be 0. It doesn't make sense.

|       | sepal_length | sepal_width | petal_length | petal_width |
| :---- | -----------: | ----------: | -----------: | ----------: |
| count |          146 |         141 |          120 |         147 |
| mean  |      56.9075 |     52.6255 |      15.5292 |     12.0265 |
| std   |      572.222 |     417.127 |       127.46 |     131.873 |
| min   |         -4.4 |        -3.6 |         -4.8 |        -2.5 |
| 25%   |          5.1 |         2.8 |        2.725 |         0.3 |
| 50%   |         5.75 |           3 |          4.5 |         1.3 |
| 75%   |          6.4 |         3.3 |          5.1 |         1.8 |
| max   |         6900 |        3809 |         1400 |        1600 |

###### +Has the presence of negative values and huge values been detected? A good data scientist always check abnormal values in the dataset. **YOU SHOULD ALWAYS TRY TO UNDERSTAND YOUR DATA**. Print the row with index 122 ;-) This week, we will have the opportunity to focus on the data pre-processing to understand how the outliers can be handled.
