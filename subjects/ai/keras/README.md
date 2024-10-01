## Keras

### Overview

This exercise focuses on using Keras to build and train neural networks. Keras is a high-level deep learning API that runs on top of TensorFlow, designed for fast experimentation with deep neural networks. You'll learn to create sequential models, add dense layers, design network architectures, and optimize your models. 

### Role Play

You are a machine learning engineer at a cutting-edge AI startup. Your team has been tasked with developing a neural network model to predict breast cancer diagnoses. The company wants to leverage the power of deep learning to improve early detection rates. Your job is to build, train, and optimize a neural network using Keras and TensorFlow. You'll need to demonstrate your understanding of neural network architectures, the Keras API, and best practices in deep learning model development.

### Learning Objectives

The goal of this day is to learn to use Keras to build Neural Networks. As explained on Keras website, Keras is a deep learning API written in Python, running on top of the machine learning platform TensorFlow. It was developed with a focus on enabling fast experimentation. Being able to go from idea to result as fast as possible is key to doing good research.
And, TensorFlow was created by the Google Brain team, TensorFlow is an open source library for numerical computation and large-scale machine learning. TensorFlow bundles together a slew of machine learning and deep learning (aka neural networking) models and algorithms and makes them useful by way of a common metaphor. It uses Python to provide a convenient front-end API for building applications with the framework, while executing those applications in high-performance C++.

There are two ways to build Keras models: sequential and functional.The sequential API allows you to create models layer-by-layer for most problems. It is limited in that it does not allow you to create models that share layers or have multiple inputs or outputs. The exercises focuses on the usage of the sequential API.

Note:

The audit will provide the code and output because it is not straightforward to reproduce results using Keras. There are many source of randomness. Even if all the seeds are fixed to a constant they may be other source of randomness. https://machinelearningmastery.com/reproducible-results-neural-networks-keras/

### Exercises of the day

- Exercise 0: Environment and libraries
- Exercise 1: Sequential
- Exercise 2: Dense
- Exercise 3: Architecture
- Exercise 4: Optimize

### Virtual Environment

- Python 3.x
- NumPy
- Pandas
- Jupyter or JupyterLab
- Keras

_Version of Keras I used to do the exercises: 2.4.3_.
I suggest to use the most recent one.

### Resources

- https://machinelearningmastery.com/tutorial-first-neural-network-python-keras/

---

---

### Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

**Note:** For each quest, your first exercice will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python.
- the virtual environment you're the most confortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recents versions of the libraries required

1. Create a virtual environment named with a version of Python >= `3.9`, with the following libraries: `pandas`, `numpy`, `jupyter`, and `keras`.

---

---

### Exercise 1: Sequential

The goal of this exercise is to learn to call the object `Sequential`.

1. Put the object Sequential in a variable named `model` and print the variable `model`.

---

---

### Exercise 2: Dense

The goal of this exercise is to learn to create layers of neurons. Keras proposes options to create custom layers. The neural networks build in these exercises do not require custom layers. `Dense` layers do the job. A dense layer is simply a layer where each unit or neuron is connected to each neuron in the next layer. As seen yesterday, there are three main types of layers: input, hidden and output. The **input layer** that specifies the number of inputs (features) is not represented as a layer in Keras. However, `Dense` has a parameter `input_dim` that gives the number of inputs in the previous layer. The output layer as any hidden layer can be created using `Dense`, the only difference is that the output layer contains one single neuron.

1. Create a `Dense` layer with these parameters and return the output of `get_config`:

   - First hidden layer connected to 5 input variables.
   - 8 neurons
   - sigmoid as activation function

2. Create a `Dense` layer with these parameters and return the output of `get_config`:

   - Hidden layer (not the first one)
   - 4 neurons
   - sigmoid as activation function

3. Create a `Dense` layer with these parameters and return the output of `get_config`:

   - Output layer
   - 1 neuron
   - sigmoid as activation function

---

---

### Exercise 3: Architecture

The goal of this exercise is to combine the layers and to create a neural network.

1. Create a neural network for regression with the following architecture and return `print(model.summary())`:

   - 5 inputs variables
   - hidden layer 1: 8 neurons and sigmoid as activation function
   - hidden layer 2: 4 neurons and sigmoid as activation function
   - output layer: 1 neuron. Find the adapted activation function

---

---

### Exercise 4: Optimize

The goal of this exercise is to learn to train the neural network. Once the architecture of the neural network is set there are two steps to train the neural network:

- `compile`: The compilation step aims to set the loss function, to choose the algoithm to minimize the chosen loss function and to choose the metric the model outputs.

  - The **optimizer**. We’ll stick with a pretty good default: the Adam gradient-based optimizer. Keras has many other optimizers you can look into as well.
  - The **loss function**. Depending on the problem to solve: classification or regression Keras proposes different loss functions. In classification Keras distinguishes between `binary_crossentropy` (2 classes) and `categorical_crossentropy` (>2 classes), so we’ll use the former.
  - The **metric(s)**. A list of metrics. Depending on the problem to solve: classification or regression Keras proposes different loss functions. For example for classification the metric can be the accuracy.

- `fit`: Training a model in Keras literally consists only of calling fit() and specifying some parameters. There are a lot of possible parameters, but we’ll only manually supply a few:

  - The **training data**, commonly known as X and Y, respectively.
  - The **number of epochs** (iterations over the entire dataset) to train for.
  - The **batch size** (number of samples per gradient update) to use when training.

  This article gives more details about **epoch** and **batch size**: https://machinelearningmastery.com/difference-between-a-batch-and-an-epoch/

1. Create the following neural network (classification):

   - Set the right number of inputs variables
   - hidden layer 1: 10 neurons and sigmoid as activation function.
   - hidden layer 2: 5 neurons and sigmoid as activation function.
   - output layer: 1 neuron and sigmoid as activation function.
   - Choose the accuracy metric, the adam optimizer, the adapted loss and epoch smaller than 50.

   Import the breast cancer data set from `sklearn.datasets` using `load_breast_cancer` and train the neural network on the data set.

2. Scale the data using `StandardScaler` from `sklearn.preprocessing`. Train the neural network again.
