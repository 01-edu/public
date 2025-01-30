#### Matrix Factorization

##### Project Structure and Setup

###### Does the project follow the structure outlined in the README?

###### Is there a README.md file that explains how to run the code and the global approach?

###### Is there a `requirements.txt` or `environment.yml` file listing all necessary libraries and their versions?

##### Data Processing and Exploratory Data Analysis

###### Is there an exploratory data analysis notebook describing insights from the MovieLens dataset?

###### Has the MovieLens dataset been properly loaded and preprocessed?

###### Has a user-item interaction matrix been created from the data?

##### Matrix Factorization Models

###### Has the Singular Value Decomposition (SVD) model been implemented using scipy.sparse.linalg.svds?

###### Has the Probabilistic Matrix Factorization (PMF) model been implemented?

###### Have both models been trained on the MovieLens dataset?

##### Model Evaluation

###### Is the Root Mean Square Error (RMSE) calculated for both models on a test set?

###### Is the RMSE of the matrix factorization models lower than a benchmark collaborative filtering model?

###### Are there learning curves demonstrating that the models are not overfitting?

###### Is there an explanation of measures taken to prevent overfitting (e.g., regularization techniques)?

###### Is there a justification for when to stop training based on the learning curves?

##### Recommendation Generation

###### Is there a function that generates movie recommendations for a user based on both SVD and PMF models?

###### Does the recommendation system return the top 10 movie recommendations for a given user?

##### Model Interpretability

###### Is there an analysis of the key latent factors that drive recommendations (global interpretability)?

###### For each user, is there an explanation of why specific movies are recommended (local interpretability)?

###### Are matrix visualizations or similarity measures used to analyze the model and its predictions?

##### Visualization and User Interface

###### Is there a visualization showing why specific movies are recommended for a user?

###### Have recommendations and visualizations been generated for 3 users (2 from training set, 1 from test set)?

###### For the 2 users from the training set, is there an analysis of why the recommendations were accurate for one and less accurate for the other?

##### Streamlit Dashboard

###### Has a Streamlit dashboard been implemented?

###### Does the dashboard take a user ID as input and return recommendations and required visualizations?

##### Additional Considerations

###### Is the code well-documented and following these good coding practices:

- Consistent naming conventions
- Proper code formatting (e.g., PEP 8 for Python)
- Separation of concerns (e.g., modular functions and classes)
- Clear and concise comments and documentation

###### Are there any innovative approaches or additional features beyond the basic requirements?

###### Has proper error handling been implemented, especially for user inputs in the Streamlit dashboard?
