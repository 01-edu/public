## Pandas

### Overview

This set of exercises focuses on using Pandas, a powerful library for data manipulation and analysis in Python. You'll learn to create and manipulate DataFrames, work with real-world datasets, handle missing values, and perform various data operations. The exercises cover key Pandas functionalities including data loading, cleaning, transformation, and basic analysis.

### Role Play

You are a data analyst at a multinational energy company. Your team has been tasked with analyzing various datasets to improve operational efficiency and customer service.

Your manager emphasizes the importance of clean, efficient code and clear explanations of your methods and findings. You'll need to present your results to both technical team members and non-technical executives, so focus on creating clear visualizations and concise summaries of your insights.

### Learning Objectives

The goal of this day is to understand practical usage of **Pandas**.
As **Pandas** in intensively used in Data Science, other days of the piscine will be dedicated to it.

Not only is the **Pandas** library a central component of the data science toolkit but it is used in conjunction with other libraries in that collection.

**Pandas** is built on top of the NumPy package, meaning a lot of the structure of NumPy is used or replicated in **Pandas**. Data in **Pandas** is often used to feed statistical analysis in SciPy, plotting functions from Matplotlib, and machine learning algorithms in Scikit-learn.

Most of the topics we will cover today are explained and describes with examples in the first resource. The number of exercises is low on purpose: Take the time to understand the chapter 5 of the resource, even if there are 40 pages.

### Exercises of the day

- **Exercise 0:** Environment and libraries
- **Exercise 1:** Your first DataFrame
- **Exercise 2:** Electric power consumption
- **Exercise 3:** E-commerce purchases
- **Exercise 4:** Handling missing values

### Virtual Environment

- Python 3.x
- NumPy
- Pandas
- Jupyter or JupyterLab

_Version of Pandas we used to do the exercises: 1.0.1_.
We suggest using the most recent one.

---

---

### Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

**Note:** For each quest, your first exercise will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python.
- the virtual environment you're the most comfortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recent versions of the libraries required

1. Create a virtual environment named `ex00`, with a version of Python >= `3.9`, with the following libraries: `pandas`, `numpy` and `jupyter`.

---

---

### Exercise 1: Your first DataFrame

The goal of this exercise is to learn to create basic Pandas objects.

1. Create a DataFrame as below this using two ways:

   - From a NumPy array
   - From a Pandas Series

   |     | color | list    | number |
   | --: | :---- | :------ | -----: |
   |   1 | Blue  | [1, 2]  |    1.1 |
   |   3 | Red   | [3, 4]  |    2.2 |
   |   5 | Pink  | [5, 6]  |    3.3 |
   |   7 | Grey  | [7, 8]  |    4.4 |
   |   9 | Black | [9, 10] |    5.5 |

2. Print the types for every column and the types of the first value of every column

---

---

### Exercise 2: Electric power consumption

The goal of this exercise is to learn to manipulate real data with Pandas.

The data set used is [**Individual household electric power consumption**](https://assets.01-edu.org/ai-branch/piscine-ai/household_power_consumption.txt)

1. Delete the columns `Time`, `Sub_metering_2` and `Sub_metering_3`
2. Set `Date` as index
3. Create a function that takes as input the DataFrame with the data set and returns a DataFrame with updated types:

   ```python
       def update_types(df):
           #TODO
           return df
   ```

4. Use `describe` to have an overview on the data set
5. Delete the rows with missing values
6. Modify `Sub_metering_1` by adding 1 to it and multiplying the total by 0.06. If x is a row the output is: (x+1)\*0.06
7. Select all the rows for which the Date is greater or equal than 2008-12-27 and `Voltage` is greater or equal than 242
8. Print the 88888th row.
9. What is the date for which the `Global_active_power` is maximal ?
10. Sort the first three columns by descending order of `Global_active_power` and ascending order of `Voltage`.
11. Compute the daily average of `Global_active_power`.

---

---

### Exercise 3: E-commerce purchases

The goal of this exercise is to learn to manipulate real data with Pandas. This exercise is less guided since the exercise 2 should have given you a nice introduction.

The data set used is [**E-commerce purchases**](./data/Ecommerce_purchases.txt).

Questions:

1. How many rows and columns are there?
2. What is the average Purchase Price?
3. What were the highest and lowest purchase prices?
4. How many people have English `'en'` as their Language of choice on the website?
5. How many people have the job title of `"Lawyer"` ?
6. How many people made the purchase during the `AM` and how many people made the purchase during `PM` ?
7. What are the 5 most common Job Titles?
8. Someone made a purchase that came from Lot: `"90 WT"` , what was the Purchase Price for this transaction?
9. What is the email of the person with the following Credit Card Number: `4926535242672853`
10. How many people have American Express as their Credit Card Provider and made a purchase above `$95` ?
11. How many people have a credit card that expires in `2025`?
12. What are the top 5 most popular email providers/hosts (e.g. gmail.com, yahoo.com, etc...)

---

---

### Exercise 4: Handling missing values

The goal of this exercise is to learn to handle missing values. In the previous exercise we used the first techniques: filter out the missing values. We were lucky because the proportion of missing values was low. But in some cases, dropping the missing values is not possible because the filtered data set would be too small.

This article explains the different types of missing data and how they should be handled.

- [Resource 1](https://towardsdatascience.com/data-cleaning-with-python-and-pandas-detecting-missing-values-3e9c6ebcf78b)

"**It’s important to understand these different types of missing data from a statistics point of view. The type of missing data will influence how you deal with filling in the missing values.**"

For this exercise, use [this dataset](./data/iris.csv).

- Preliminary: Drop the `flower` column. Then try to fill missing values with different strategies:

1. Fill the missing values with a different "strategy" for each column:

   `sepal_length` -> `mean`

   `sepal_width` -> `median`

   `petal_length`, `petal_width` -> `0`

2. Fill the missing values using the median of the associated column using `fillna`.

- Bonus questions:
  - Filling the missing values by 0 or the mean of the associated column is common in Data Science. In that case, explain why filling the missing values with 0 or the mean is a bad idea.
  - Find a special row ;-).

### Resources

- [The ultimate Pandas resource](https://bedford-computing.co.uk/learning/wp-content/uploads/2015/10/Python-for-Data-Analysis.pdf)

- [Pandas documentation](https://pandas.pydata.org/docs/)

- [Pandas resource 1](https://jakevdp.github.io/PythonDataScienceHandbook/)

- [Pandas resource 2](https://pandas.pydata.org/Pandas_Cheat_Sheet.pdf)

- [Pandas resource 3](https://www.learndatasci.com/tutorials/python-pandas-tutorial-complete-introduction-for-beginners/)

- [Pandas resource 4](https://jakevdp.github.io/PythonDataScienceHandbook/03.04-missing-values.html)
