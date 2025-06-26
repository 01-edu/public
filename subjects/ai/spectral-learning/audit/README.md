#### Spectral Learning

##### Preliminary

###### Does the structure of the project match the setup outlined in the README, with well-organized folders for data, models, utilities, and documentation?

###### Does the README provide a comprehensive overview, including setup instructions, objectives, and usage guidelines?

###### Is there a `requirements.txt` or `environment.yml` file listing all necessary libraries and their versions?

###### Do text files or documentation explain the methods and underlying concepts of the implemented algorithms?

#### Data Processing and Exploratory Data Analysis

###### Is the dataset effectively loaded and preprocessed, handling null values and ensuring consistency?

###### Is the high-dimensional dataset transformed into a suitable feature matrix for matrix decomposition with proper normalization and scaling?

#### Spectral Learning Models

###### Is PCA implemented from scratch, calculating eigenvalues and eigenvectors of the covariance matrix without relying on external libraries like scikit-learn?

###### Is SVD implemented from scratch, avoiding libraries such as `scipy.svds` or similar tools, ensuring an understanding of matrix factorization?

###### Functions for matrix operations (e.g., eigenvalue computation, SVD decomposition) should be reusable across both PCA and SVD scripts.

###### Are both PCA and SVD applied to the feature matrix, generating reduced-dimensional representations?

#### Model Evaluation

###### Is the **variance explained** by PCA and SVD calculated, assessing the effectiveness of each dimensionality reduction method?

###### Are metrics such as clustering accuracy or separability scores used to evaluate the reduced dimensions?

###### Are learning curves or variance-explained plots provided to demonstrate model convergence and the sufficiency of principal components?

###### Are overfitting prevention strategies discussed (e.g., regularization) and applied to the dimensionality reduction techniques, if applicable?

###### Is the number of dimensions chosen for PCA and SVD justified based on variance retention or other criteria?

#### Dimensionality Reduction and Feature Extraction

###### Is there a function that allows extraction and visualization of key principal components, supporting feature interpretability?

###### Are latent components analyzed to provide insights into the underlying structure of the data captured through dimensionality reduction?

#### Visualization and Interpretation

###### Are 2D or 3D visualizations provided to show the separation of data points in the reduced dimensional space?

###### Is a side-by-side comparison of variance explained by PCA and SVD presented for interpretability?

###### Are clustering patterns in reduced dimensions visualized to demonstrate the effectiveness of dimensionality reduction?

###### Are multiple visual examples, such as clustering subsets of the data, used to showcase robustness?

#### Additional Considerations

###### Is the code well-documented, with comments explaining each function, following best practices for readability and maintainability?

###### Are additional techniques such as **non-linear dimensionality reduction** (e.g., t-SNE or UMAP) explored for complex data patterns or advanced visualizations?

###### Is error handling implemented in data loading, matrix decomposition functions, and visualization methods to ensure stability across various dataset inputs?
