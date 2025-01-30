### Forest Prediction

### Overview

The goal of this project is to use cartographic variables to classify forest categories. You will have to analyze the data, create features and to train a machine learning model on the cartographic data to make it as accurate as possible.

### Role Play

You are a data scientist working for an environmental conservation agency. Your team has been tasked with developing a machine learning model to classify forest cover types using cartographic data. This model will be used to assist in forest management and conservation efforts. You need to analyze the data, engineer relevant features, select and train the best model, and ultimately make predictions on new data. Your work will directly impact decision-making in forest preservation strategies.

### Learning Objectives

By the end of this project, you will be able to:

1. Perform exploratory data analysis (EDA) on complex environmental datasets.
2. Engineer relevant features from cartographic variables to improve model performance.
3. Implement a robust model selection process using cross-validation and grid search.
4. Work with various machine learning algorithms including Gradient Boosting, KNN, Random Forest, SVM, and Logistic Regression.
5. Evaluate model performance using appropriate metrics such as accuracy and confusion matrices.
6. Visualize model performance through learning curves and other relevant plots.
7. Handle potential overfitting issues and ensure model generalization.
8. Use pickle to save and load trained machine learning models.
9. Make predictions on unseen data and interpret the results in the context of forest classification.
10. Document your process and results clearly for stakeholders in the conservation agency.

### Instructions

#### Data

The input files are `train.csv`, `test.csv` and `covtype.info`:

- `train.csv`
- `test.csv`
- `covtype.info`

The train data set is used to **analyze the data and calibrate the models**. The goal is to get the accuracy as high as possible on the test set. The test set will be available at the end of the last day to prevent from the overfitting of the test set.

The data is described in `covtype.info`.

#### 1. EDA and feature engineering

- Create a Jupyter Notebook to analyze the data sets and perform EDA (Exploratory Data Analysis). This notebook will not be evaluated.

- _Hint: Examples of interesting features_

  - `Distance to hydrology = sqrt((Horizontal_Distance_To_Hydrology)^2 + (Vertical_Distance_To_Hydrology)^2)`
  - `Horizontal_Distance_To_Fire_Points - Horizontal_Distance_To_Roadways`

#### 2. Model Selection

The model selection approach is a key step because, t should return the best model and guaranty that the results are reproducible on the final test set. The goal of this step is to make sure that the results on the test set are not due to test set overfitting. It implies to split the data set as shown below:

```console
DATA
└───TRAIN FILE (0)
│   └───── Train (1)
│   |           Fold0:
|   |                  Train
|   |                  Validation
|   |           Fold1:
|   |                   Train
|   |                   Validation
... ...         ...
|   |
|   └───── Test (1)
│
└─── TEST FILE (0) (available last day)

```

**Rules:**

- Split train test
- Cross validation: at least 5 folds
- Grid search on at least 5 different models:
  - Gradient Boosting, KNN, Random Forest, SVM, Logistic Regression. _Remember that for some model scaling the data is important and for others it doesn't matter._
- Train accuracy score < **0.98**. Train set (0). Write the result in the `README.md`
- Test (last day) accuracy > **0.65**. Test set (0). Write the result in the `README.md`
- Display the confusion matrix for the best model in a DataFrame. Precise the index and columns names (True label and Predicted label)
- Plot the learning curve for the best model
- Save the trained model as a [pickle](https://docs.python.org/3/library/pickle.html) file

> Advice: As the grid search takes time, I suggest preparing and test the code. Once you are confident it works, run the gridsearch at night and analyse the results

**Hint**: The confusion matrix shows the misclassifications class per class. Try to detect if the model misclassifies badly one class with another. Then, do some research on the internet on the two forest cover types, find the differences and create some new features that underline these differences. More generally, the methodology of a models learning is a cycle with several iterations. More details [here](https://serokell.io/blog/machine-learning-testing)

#### 3. Predict (last day)

Once you have selected the best model and you are confident it will perform well on new data, you're ready to predict on the test set:

- Load the trained model
- Predict on the test set and compute the accuracy
- Save the predictions in a csv file
- Add your score on the `README.md`

### Project repository structure:

The structure of the project is:

```console
project
│   README.md
│   requirements.txt
│
└───data
│   │   train.csv
│   |   test.csv (not available first day)
|   |   covtype.info
│
└───notebook
│   │   EDA.ipynb
|
|───scripts
|   │   preprocessing_feature_engineering.py
|   │   model_selection.py
│   |   predict.py
│
└───results
    │   plots
    │   test_predictions.csv
    │   best_model.pkl

```

### Tips

1. **Data Exploration:**

   - Start with a thorough EDA. Understand the distribution of each feature and the target variable.
   - Look for correlations between features and the target class.

2. **Feature Engineering:**

   - Consider creating composite features like the suggested distance calculations.
   - Look for domain-specific insights that could inform feature creation.

3. **Data Preprocessing:**

   - Handle any missing values appropriately.
   - Consider scaling features for algorithms sensitive to feature scales (e.g., SVM, KNN).

4. **Model Selection:**

   - Use stratified k-fold cross-validation to ensure each fold represents the overall class distribution.
   - Start with a wide range of hyperparameters in your grid search, then narrow down for a more focused search.

5. **Avoiding Overfitting:**

   - Pay close attention to the difference between training and validation accuracy.
   - Use regularization techniques where appropriate.

6. **Handling Class Imbalance:**

   - Check if the classes are balanced. If not, consider techniques like oversampling, undersampling, or using class weights.

7. **Feature Importance:**

   - Analyze feature importance from tree-based models to gain insights and potentially reduce the feature set.

8. **Ensemble Methods:**

   - Consider using ensemble methods to combine predictions from multiple models.

9. **Performance Evaluation:**

   - Don't rely solely on accuracy. Consider metrics like precision, recall, and F1-score, especially if classes are imbalanced.

10. **Computational Efficiency:**

    - For time-consuming operations like grid search, consider using parallelization if available.

11. **Interpretability:**

    - Try to balance model performance with interpretability, especially since this is for an environmental agency that may need to explain decisions.

12. **Version Control:**

    - Keep track of different versions of your features and models. This helps in comparing different approaches.

13. **Documentation:**

    - Document your process, decisions, and findings clearly. This is crucial for reproducibility and for communicating with stakeholders.

14. **Final Testing:**
    - Reserve your test set strictly for final evaluation. Don't use it for any decision-making during the model development process.

Remember, the goal is not just to achieve high accuracy, but to create a robust, generalizable model that provides valuable insights for forest classification and management.

### Resources

- [Assets](https://assets.01-edu.org/ai-branch/piscine-ai/raid02/raid02-20221024T133335Z-001.zip)
