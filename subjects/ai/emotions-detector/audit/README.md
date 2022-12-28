#### Computer vision

##### Preliminary

```
project
│   README.md
│   environment.yml
│
└───data
│   │   train.csv
│   │   test.csv
│   │   xxx.csv
│
└───results
│   │
|   |───model (free format)
│   │   │   my_own_model.pkl
│   │   │   my_own_model_architecture.txt
│   │   │   tensorboard.png
│   │   │   learning_curves.png
│   │   │   pre_trained_model.pkl (optional)
│   │   │   pre_trained_model_architecture.txt (optional)
│   │
|   |───hack_cnn (free format)
│   │   │   hacked_image.png   (optional)
│   │   │   input_image.png
│   │
|   |───preprocessing_test
|   |   |   input_video.mp4  (free format)
│   │   │   image0.png  (free format)
│   │   │   image1.png
│   │   │   imagen.png
│   │   │   image20.png
|
|───scripts
│   │   train.py
│   │   predict.py
│   │   preprocess.py
│   │   predict_live_stream.py
│   │   hack_the_cnn.py

```

###### Does the structure of the project look as above?

###### Does the readme file summarize how to run the code and explain the global approach?

###### Does the environment contain all libraries used and their versions that are necessary to run the code?

###### Do the text files explain the chosen architectures?

#### CNN emotion classifier

###### Is the model trained only the training set?

###### Is the accuracy on the test set higher than 70%?

###### Do the learning curves prove that the model is not overfitting?

###### Has the training been stopped early enough to avoid the overfitting?

###### Does the screenshot show the usage of the tensorboard to monitor the training?

###### Does the text document explain why the architecture was chosen and what were the previous iterations?

###### Does the following command `python predict.py ` run without any error and returns an accuracy greater than 70%?

```prompt
    python predict.py

    Accuracy on test set: 72%

```

#### Face detection on the video stream

###### Does the preprocessing pipeline take as input the webcam video stream of minimum 20 sec and save in a separate folder at least 20 preprocessed\* images?

###### Do all images contain a face?

###### Are all images reshaped and centered on the face?

###### Is the algorithm that detects the face imported via cv2?

###### Is the image converted to 48 x 48 grayscale pixels' image?

###### If there's an issue related to the webcam, does the code takes as input a video recorded video stream?

###### Does the following command `predict_live_stream.py` run without any error and return the following?

```prompt
    python predict_live_stream.py

    Reading video stream ...

    Preprocessing ...
    11:11:11s : Happy , 73%

    Preprocessing ...
    11:11:12s : Happy , 93%

    Preprocessing ...
    11:11:13s : Surprise , 71%

    Preprocessing ...
    11:11:14s : Neutral , 82%

    ...

    Preprocessing ...
    11:13:29s : Happy , 63%
```

#### Hack the CNN - guidelines:

The neural network trains by updating its weights given the training error. If an image is misclassfied the neural network changes its weight to classify it correctly. The trick is to keep the neural network's weights unchanged and to modify the input pixels in order to force the neural network to predict the wanted class.
This part is validated if:

##### Choose an image from the database that gives more than 90% probability of `Happy`

###### Does the neural network modifies the input pixels to predict Sad?

###### Can you recognize easily the chosen image? The modified image is SLIGHTLY changed. It means that you recognise very easily the original image.

Here are three resources that detail similar approaches:

- https://github.com/XC-Li/Facial_Expression_Recognition/tree/master/Code/RAFDB
- https://github.com/karansjc1/emotion-detection/tree/master/with%20flask
- https://www.kaggle.com/drbeanesp21/aliaj-final-facial-expression-recognition (simplified)
