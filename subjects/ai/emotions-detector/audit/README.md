#### Emotion detector

##### Preliminary

###### Does the structure of the project is equivalent to the one described in the subject `Delivery` section?

###### Does the README file summarize how to run the code and explain the global approach?

###### Does the environment contain all libraries used and their versions that are necessary to run the code?

###### Do the text files explain the chosen architectures?

#### CNN emotion classifier

###### Is the model trained only the training set?

###### Is the accuracy on the test set higher than 60%?

###### Do the learning curves prove that the model is not overfitting?

###### Has the training been stopped early enough to avoid the overfitting?

###### Does the screenshot show the usage of the `TensorBoard` to monitor the training?

###### Does the text document explain why the architecture was chosen, and what were the previous iterations?

###### Does the following command `python ./scripts/predict.py` run without any error and returns an accuracy greater than 60%?

```prompt
    python ./scripts/predict.py

    Accuracy on test set: 62%

```

#### Face detection on the video stream

###### Does the preprocessing pipeline take as input the webcam video stream of minimum 20 sec and save in a separate folder at least 20 preprocessed\* images?

###### Do all images contain a face?

###### Are all images reshaped and centered on the face?

###### Is the algorithm that detects the face imported via cv2?

###### Is the image converted to 48 x 48 grayscale pixels' image?

###### If there's an issue related to the webcam, does the code take as input a video recorded video stream?

###### Does the following command `python ./scripts/predict_live_stream.py` run without any error and return the following?

```prompt
    python ./scripts/predict_live_stream.py

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

The neural network trains by updating its weights given the training error. If an image is misclassified the neural network changes its weight to classify it correctly. The trick is to keep the neural network's weights unchanged and to modify the input pixels in order to force the neural network to predict the wanted class.
This part is validated if:

##### Choose an image from the database that gives more than 90% probability of `Happy`

###### Does the neural network modify the input pixels to predict Sad?

###### Can you recognize easily the chosen image? The modified image is SLIGHTLY changed. It means that you recognize very easily the original image.

Here are three resources that detail similar approaches:

- https://github.com/XC-Li/Facial_Expression_Recognition/tree/master/Code/RAFDB
- https://github.com/karansjc1/emotion-detection/tree/master/with%20flask
- https://www.kaggle.com/drbeanesp21/aliaj-final-facial-expression-recognition (simplified)
