## Matrix Factorization Recommender System

### Overview

<center>
    <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQiRyPw-WJ-TI66GKXPoKpDgv3mO7RWhH14VA&s"/>
</center>

This project implements an advanced movie recommender system using matrix factorization techniques. The system is built with **Singular Value Decomposition (SVD)** and **Probabilistic Matrix Factorization (PMF)** to provide personalized recommendations based on user preferences. It also includes an interactive dashboard using **Streamlit** that allows users to input their user ID and receive recommendations, while visualizing the differences between SVD and PMF predictions.

### Learning Objective

The goal of this project is to understand and apply advanced matrix factorization techniques in building recommender systems. By comparing two different models (SVD and PMF), you'll gain insights into their performance and how they generate recommendations. You'll also learn to integrate these models into an interactive interface using **Streamlit** for real-time interaction.

### Instructions

#### Data Loading and Preprocessing

1. **Download the [MovieLens Dataset](https://grouplens.org/datasets/movielens/1m/)** (ratings, users, and movies).
2. Preprocess the dataset to remove null values and prepare it for matrix factorization.
3. Create a user-item interaction matrix from the data.

#### Singular Value Decomposition (SVD) Model

1. Implement the SVD algorithm using the **scipy.sparse.linalg.svds** function for matrix factorization.
2. Train the SVD model on the MovieLens dataset to generate predicted ratings for all users.

#### Probabilistic Matrix Factorization (PMF) Model

1. Implement the PMF algorithm.
2. Train the PMF model and visualize the model's convergence (e.g., plot Mean Squared Error over iterations).

#### Model Comparison and Evaluation

1. Compare the performance of SVD and PMF using evaluation metrics such as **Mean Squared Error (MSE)**.
2. Provide visual comparisons between the models using **matplotlib** to plot predicted vs. actual ratings.

#### Recommendation Generation

1. Implement a function that generates movie recommendations for a user based on the predicted ratings from both the SVD and PMF models.
2. Display top-rated movies for users and compare recommendations from both models.

#### Analysis and Visualization

1. Provide visualizations comparing SVD and PMF predictions for the same user.
2. Offer insights into how the models differ in recommending movies for specific users based on their ratings history.

#### Streamlit Dashboard

1. Build an interactive **Streamlit** dashboard that takes a **user ID** as input and returns the following:
   - List of top-rated movies for the user.
   - Movie recommendations from both the **SVD** and **PMF** models.
   - Visual comparison of the SVD vs. PMF predictions for the user.
2. Ensure real-time interaction, with recommendations and visualizations updating dynamically based on user input.

### Project Repository Structure

```
matrix-factorization-project/
│
├── data/
│   ├── ratings.dat
│   ├── users.dat
│   └── movies.dat
│
├── models/
│   ├── svd_model.py
│   ├── pmf_model.py
│
├── utils/
│   ├── data_loader.py
│   ├── matrix_creation.py
│   ├── recommendation.py
│
├── app.py
├── requirement.txt
├── Movie_Recommender_System.ipynb
└── README.md
```

- **data/**: Contains the raw MovieLens dataset files.
- **models/**: Holds the matrix factorization models (`svd_model.py` for SVD, `pmf_model.py` for PMF). These files contain the logic for training and making predictions with each model.
- **utils/**: Includes utility functions for loading and processing the data, creating user-item matrices, and generating recommendations.
- **app.py**: The core of the project, this file creates the interactive Streamlit dashboard, enabling users to input a user ID and get movie recommendations and model comparison visualizations.
- **Movie_Recommender_System.ipynb**: A notebook for initial experiments, data exploration, and visualization of the model training and recommendations.
- **README.md**: Project documentation with an overview of the recommender system, instructions for setup and running the dashboard, and additional resources.

### Timeline (1-2 weeks)

**Week 1:**

- **Days 1-2:** Load and preprocess the dataset, create user-item interaction matrix.
- **Days 3-4:** Implement and train the SVD model.
- **Days 5-7:** Implement and train the PMF model, visualize MSE vs. iterations for PMF.

**Week 2:**

- **Days 1-2:** Compare SVD and PMF models, evaluate using MSE.
- **Days 3-4:** Implement recommendation generation for both models.
- **Days 5-7:** Build the Streamlit dashboard, create visualizations, and finalize the project.

### Tips

Remember, a great recommender system needs to understand both the users and the content. Keep in mind the trade-off between model complexity and interpretability. Here are some additional considerations:

- Code Quality: Is the code well-documented and following good coding practices? Clearly labeled functions, docstrings, and a consistent style will make the project easier to understand and maintain.
- Error Handling: Include checks and error messages for unexpected user inputs on the Streamlit dashboard.
- Efficiency: For large datasets, consider optimizing your code to avoid memory bottlenecks, especially in matrix operations.
  Parameter Tuning: Experiment with model parameters for SVD and PMF to find the best-performing configuration.
- User Experience: Consider adding features like a search bar for movies or a "surprise me" option for diverse recommendations.

### Resources

- [Matrix Factorization Techniques for Recommender Systems](https://www.diva-portal.org/smash/get/diva2:633561/FULLTEXT01.pdf) - A comprehensive overview of matrix factorization methods.
- [Singular Value Decomposition and Its Applications in Recommender Systems](https://www.cs.uic.edu/~liub/KDD-cup-2007/proceedings/Regular-Paterek.pdf) - Detailed explanation of SVD.
- [Probabilistic Matrix Factorization](https://www.cs.toronto.edu/~amnih/papers/bpmf.pdf) - Original paper on PMF by Salakhutdinov and Mnih.
- [MovieLens Dataset](https://grouplens.org/datasets/movielens/1m/) - Official page for the MovieLens 1M dataset.
- [Scipy sparse.linalg.svds Documentation](https://docs.scipy.org/doc/scipy/reference/generated/scipy.sparse.linalg.svds.html) - Reference for SVD.
- [Pandas Documentation](https://pandas.pydata.org/docs/) - Guide for data manipulation with Pandas.
- [Matplotlib Documentation](https://matplotlib.org/stable/contents.html) - Guide for creating visualizations.
- [Streamlit Documentation](https://docs.streamlit.io/) - Guide for creating interactive dashboards.
