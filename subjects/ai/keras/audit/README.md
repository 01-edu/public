#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`

##### Run `python --version`

###### Does it print `Python 3.x`? x >= 9

###### Does `import jupyter`, `import numpy`, `import pandas`, and `import keras` run without any error?

---

---

#### Exercise 1: Sequential

###### For question 1, does the output end with `keras.engine.sequential.Sequential object at xxx`?

---

---

#### Exercise 2: Dense

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, do the fields `batch_input_shape`, `units` and `activation` match this output?

```
{'name': 'dense_7',
'trainable': True,
'batch_input_shape': (None, 5),
'dtype': 'float32',
'units': 8,
'activation': 'sigmoid',
'use_bias': True,
'kernel_initializer': {'class_name': 'GlorotUniform',
'config': {'seed': None}},
'bias_initializer': {'class_name': 'Zeros', 'config': {}},
'kernel_regularizer': None,
'bias_regularizer': None,
'activity_regularizer': None,
'kernel_constraint': None,
'bias_constraint': None}
```

###### For question 2, do the fields `units` and `activation` match this output?

```
{'name': 'dense_8',
'trainable': True,
'dtype': 'float32',
'units': 4,
'activation': 'sigmoid',
'use_bias': True,
'kernel_initializer': {'class_name': 'GlorotUniform',
'config': {'seed': None}},
'bias_initializer': {'class_name': 'Zeros', 'config': {}},
'kernel_regularizer': None,
'bias_regularizer': None,
'activity_regularizer': None,
'kernel_constraint': None,
'bias_constraint': None}
```

###### For question 3, do the fields `units` and `activation` match this output?

```
{'name': 'dense_9',
'trainable': True,
'dtype': 'float32',
'units': 1,
'activation': 'sigmoid',
'use_bias': True,
'kernel_initializer': {'class_name': 'GlorotUniform',
'config': {'seed': None}},
'bias_initializer': {'class_name': 'Zeros', 'config': {}},
'kernel_regularizer': None,
'bias_regularizer': None,
'activity_regularizer': None,
'kernel_constraint': None,
'bias_constraint': None}
```

---

---

#### Exercise 3: Architecture

###### For question 1, is code that creates the neural network the following?

```
model = keras.Sequential()
model.add(Dense(8, input_shape=(5,), activation= 'sigmoid'))
model.add(Dense(4, activation= 'sigmoid'))
model.add(Dense(1, activation= 'linear'))

```

The first two layers could use another activation function that sigmoid (eg: relu)

---

---

#### Exercise 4: Optimize

###### For question 1, does the output of `model.get_config()['layers']` match the fields `batch_input_shape`, `units` and `activation`?

```
[{'class_name': 'InputLayer',
  'config': {'batch_input_shape': (None, 30),
   'dtype': 'float32',
   'sparse': False,
   'ragged': False,
   'name': 'dense_134_input'}},
 {'class_name': 'Dense',
  'config': {'name': 'dense_134',
   'trainable': True,
   'batch_input_shape': (None, 30),
   'dtype': 'float32',
   'units': 10,
   'activation': 'sigmoid',
   'use_bias': True,
   'kernel_initializer': {'class_name': 'GlorotUniform',
    'config': {'seed': None}},
   'bias_initializer': {'class_name': 'Zeros', 'config': {}},
   'kernel_regularizer': None,
   'bias_regularizer': None,
   'activity_regularizer': None,
   'kernel_constraint': None,
   'bias_constraint': None}},
 {'class_name': 'Dense',
  'config': {'name': 'dense_135',
   'trainable': True,
   'dtype': 'float32',
   'units': 5,
   'activation': 'sigmoid',
   'use_bias': True,
   'kernel_initializer': {'class_name': 'GlorotUniform',
    'config': {'seed': None}},
   'bias_initializer': {'class_name': 'Zeros', 'config': {}},
   'kernel_regularizer': None,
   'bias_regularizer': None,
   'activity_regularizer': None,
   'kernel_constraint': None,
   'bias_constraint': None}},
 {'class_name': 'Dense',
  'config': {'name': 'dense_136',
   'trainable': True,
   'dtype': 'float32',
   'units': 1,
   'activation': 'sigmoid',
   'use_bias': True,
   'kernel_initializer': {'class_name': 'GlorotUniform',
    'config': {'seed': None}},
   'bias_initializer': {'class_name': 'Zeros', 'config': {}},
   'kernel_regularizer': None,
   'bias_regularizer': None,
   'activity_regularizer': None,
   'kernel_constraint': None,
   'bias_constraint': None}}]
```

You should notice that the neural network is struggling to learn. By luck the initialization of the weights might have led to an accuracy close of 90%. But when I trained the neural network, with `batch_size=300` on the data here is the output of the last epoch (50):

`Epoch 50/50 2/2 [==============================] - 0s 1ms/step - loss: 0.6559 - accuracy: 0.6274`

###### For question 2, is the the accuracy at epoch 50 higher than 95%?
