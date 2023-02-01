# Keras 2

The goal of this day is to learn to use Keras to build Neural Networks and train them on small data sets. This helps to understand the specifics of networks for classification and regression.

Note:

The audit will provide the code and output because it is not straightforward to reproduce results using Keras. There are many source of randomness. Even if all the seeds are fixed to a constant they may be other source of randomness. https://machinelearningmastery.com/reproducible-results-neural-networks-keras/

### Exercises of the day

- Exercise 0: Environment and libraries
- Exercise 1: Regression - Optimize
- Exercise 2: Regression example
- Exercise 3: Multi classification - Softmax
- Exercise 4: Multi classification - Optimize
- Exercise 5: Multi classification example

### Virtual Environment

- Python 3.x
- NumPy
- Pandas
- Jupyter or JupyterLab
- Keras

_Version of Keras I used to do the exercises: 2.4.3_.
I suggest to use the most recent one.

### **Resources**

- https://machinelearningmastery.com/tutorial-first-neural-network-python-keras/

---

---

# Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

**Note:** For each quest, your first exercice will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python.
- the virtual environment you're the most confortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recents versions of the libraries required

1. Create a virtual environment named with a version of Python >= `3.8`, with the following libraries: `pandas`, `numpy`, `jupyter` and `keras`.

---

---

# Exercise 1: Regression - Optimize

The goal of this exercise is to learn to set up the optimization for a regression neural network. There's no code to run in that exercise. In W2D2E3, we implemented a neural network designed for regression. We will be using this neural network:

    ```
    model = keras.Sequential()
    model.add(Dense(8, input_shape=(5,), activation= 'sigmoid'))
    model.add(Dense(4, activation= 'sigmoid'))
    model.add(Dense(1, activation= 'linear'))

    ```

As a reminder, the main difference between a regression and classification neural network's architecture is the output layer activation function.

1. Fill this chunk of code to set up the optimization part of the regression neural network:

```
model.compile(
  optimizer='adam',
  loss='',#TODO1
  metrics=[''] #TODO2
)
```

Hint:

- Mean Squared Error (MSE) and Mean Absolute Error (MAE) are common loss functions used for regression problems. Mean Absolute Error is less sensitive to outliers. Different loss functions are used for classification problems. Similarly, evaluation metrics used for regression differ from classification.

https://keras.io/api/metrics/regression_metrics/

---

---

# Exercise 2: Regression example

The goal of this exercise is to learn to train a neural network to perform a regression on a data set.
The data set is Auto MPG Dataset and the go is to build a model to predict the fuel efficiency of late-1970s and early 1980s automobiles. To do this, provide the model with a description of many automobiles from that time period. This description includes attributes like: cylinders, displacement, horsepower, and weight.

https://www.tensorflow.org/tutorials/keras/regression

1. Preprocess the data set as follow:

   - Drop the columns: **model year**, **origin**, **car name**
   - Split train test without shuffling the data. Keep 20% for the test set.
   - Scale the data using Standard Scaler

2. Train a neural network on the train set and predict on the test set. The neural network should have 2 hidden layers and the loss should be **mean_squared_error**. The expected **mean absolute error** on the test set is maximum 10.
   _Hint_: inscrease the number of epochs
   **Warning**: Do no forget to evaluate the neural network on the **SCALED** test set.

---

---

# Exercise 3: Multi classification - Softmax

The goal of this exercise is to learn to a neural network architecture for multi-class data. This is an important type of problem on which to practice with neural networks because the three class values require specialized handling. A multi-classification neural network uses as output layer a **softmax** layer. The **softmax** activation function is an extension of the sigmoid as it is designed to output the probabilities to belong to each class in a multi-class problem. This output layer has to contain as much neurons as classes in the multi-classification problem. This article explains in detail how it works. https://developers.google.com/machine-learning/crash-course/multi-class-neural-networks/softmax

Let us assume we want to classify images and we know they contain either apples, bears, candies, eggs or dogs (extension of the example in the link above).

1. Create the architecture for a multi-class neural network with the following architecture and return `print(model.summary())`:

   - 5 inputs variables
   - hidden layer 1: 16 neurons and sigmoid as activation function
   - hidden layer 2: 8 neurons and sigmoid as activation function
   - output layer: The number of neurons and the activation function should be adapted to this multi-classification problem.

---

---

# Exercise 4: Multi classification - Optimize

The goal of this exercise is to learn to optimize a multi-classification neural network. As learnt previously, the loss function used in binary classification is the log loss - also called in Keras `binary_crossentropy`. This function is defined for binary classification and can be extended to multi-classfication. In Keras, the extended loss that supports multi-classification is `binary_crossentropy`. There's no code to run in that exercise.

1. Fill the chunk of code below in order to optimize the neural network defined in the previous exercise. Choose the adapted loss, adam as optimizer and the accuracy as metric.

```
model.compile(loss='',#TODO1
              optimizer='', #TODO2
              metrics=['']) #TODO3
```

---

---

# Exercise 5 Multi classification example

The goal of this exercise is to learn to use a neural network to classify a multiclass data set. The data set used is the Iris data set which allows to classify flower given basic features as flower's measurement.

Preliminary:

- [Load the dataset from `sklearn`.](https://scikit-learn.org/stable/auto_examples/datasets/plot_iris_dataset.html)
- Split train test. Keep 20% for the test set. Use `random_state=1`.
- Scale the data using Standard Scaler

1. Use the `LabelBinarizer` from Sckit-learn to create a one hot encoding of the target. As you know, the output layer of a multi-classification neural network shape is equal to the number of classes. The output layer expects to have a target with the same shape as its output layer.

2. Train a neural network on the train set and predict on the test set. The neural network should have 1 hidden layers. The expected **accuracy** on the test set is minimum 90%.
   _Hint_: inscrease the number of epochs
   **Warning**: Do no forget to evaluate the neural network on the **SCALED** test set.
