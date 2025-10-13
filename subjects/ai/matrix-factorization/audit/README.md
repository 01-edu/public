#### Matrix Factorization

##### Project Structure and Setup

###### Does the project follow the structure outlined in the README?

###### Is there a README.md file that explains how to run the code and the global approach?

###### Is there a `requirements.txt` or `environment.yml` file listing all necessary libraries and their versions?

###### Do the core files exist: `app.py`, `models/svd_model.py`, `models/pmf_model.py`, and `utils/recommendation.py`?

###### Do the main dependencies import without error?
```bash
python -c "import numpy, pandas, scipy, streamlit, matplotlib"
```

##### Data Processing and Exploratory Data Analysis

###### Is there an exploratory data analysis notebook describing insights from the MovieLens dataset?

###### Has the MovieLens dataset been properly loaded and preprocessed?

###### Has a user-item interaction matrix been created from the data?

###### Was a reproducible split used (e.g., `random_state = 42`)?

###### Does the normalized user–item matrix exist at `processed/user_item_matrix.csv`

##### Matrix Factorization Models

###### Has the Singular Value Decomposition (SVD) model been implemented using scipy.sparse.linalg.svds?

###### Has the Probabilistic Matrix Factorization (PMF) model been implemented?

###### Have both models been trained on the MovieLens dataset?

###### Is the SVD predicted rating matrix saved as `reports/svd_predictions.npy`?

###### Does the PMF implementation save a convergence plot (`reports/pmf_convergence.png`)?

###### Are the learned factor matrices (`U`, `V`) saved (e.g., under `reports/pmf_factors/`)?

##### Model Evaluation

###### Is the Root Mean Square Error (RMSE) calculated for both models on a test set?

###### Is the RMSE of the matrix factorization models lower than a benchmark collaborative filtering model?

###### Are there learning curves demonstrating that the models are not overfitting?

###### Is there an explanation of measures taken to prevent overfitting (e.g., regularization techniques)?

###### Is there a justification for when to stop training based on the learning curves?

###### Does `reports/model_metrics.json` exist with fields:
  ```json
  {
    "SVD_RMSE": ...,
    "PMF_RMSE": ...,
    "PMF_vs_SVD_improvement_%": ...
  }
  ```

###### Are the following thresholds met?
  * SVD RMSE ≤ 0.90
  * PMF RMSE ≤ 0.85
  * PMF improvement ≥ 5%
  * Are the plots saved? `reports/rmse_comparison.png` and `reports/predicted_vs_actual.png`.

##### Recommendation Generation

###### Is there a function that generates movie recommendations for a user based on both SVD and PMF models?

###### Does the recommendation system return the top 10 movie recommendations for a given user?

###### Does `utils/recommendation.py` expose:
```python
def generate_recommendations(user_id, model, top_n=10):
    ...
```

###### Are user-level outputs saved as `reports/user_<id>_recommendations.csv`

##### Model Interpretability

###### Is there an analysis of the key latent factors that drive recommendations (global interpretability)?

###### For each user, is there an explanation of why specific movies are recommended (local interpretability)?

###### Are matrix visualizations or similarity measures used to analyze the model and its predictions?

##### Visualization and User Interface

###### Is there a visualization showing why specific movies are recommended for a user?

###### Have recommendations and visualizations been generated for 3 users (2 from training set, 1 from test set)?

###### For the 2 users from the training set, is there an analysis of why the recommendations were accurate for one and less accurate for the other?

###### Are required visuals present in `reports/` with proper titles and labeled axes?
  * `pmf_convergence.png`
  * `rmse_comparison.png`
  * `predicted_vs_actual.png`
  * `user_comparison.png`

##### Streamlit Dashboard

###### Has a Streamlit dashboard been implemented?

###### Does the dashboard take a user ID as input and return recommendations and required visualizations?

###### Does `streamlit run app.py` launch the dashboard successfully?

###### Does the dashboard update recommendations dynamically on user ID input?

###### Does it handle invalid user IDs gracefully (error shown, no crash)?

##### Additional Considerations

###### Is the code well-documented and following these good coding practices:

- Consistent naming conventions
- Proper code formatting (e.g., PEP 8 for Python)
- Separation of concerns (e.g., modular functions and classes)
- Clear and concise comments and documentation

###### Are there any innovative approaches or additional features beyond the basic requirements?

###### Has proper error handling been implemented, especially for user inputs in the Streamlit dashboard?
