# Emotions detection with Deep Learning

Cameras are everywhere. Videos and images have become one of the most interesting data sets for artificial intelligence.
Image processing is a quite board research area, not just filtering, compression, and enhancement. Besides, we are even interested in the question, “what is in images?”, i.e., content analysis of visual inputs, which is part of the main task of computer vision. The study of computer vision could make possible such tasks as 3D reconstruction of scenes, motion capturing, and object recognition, which are crucial for even higher-level intelligence such as
image and video understanding, and motion understanding.
For this 2 months project we will focus on two tasks:

- emotion classfication
- face tracking

With the computing power exponentially increasing the computer vision field has been developping exponentially. This is a key element because the computer power allows to use more easily a type of neural networks very powerful on images: CNN's (Convolutional Neural Networks). Before the CNN's were democratized, the algorithms used relied a lot on human analysis to extract features which obviously time consuming and not reliable. If you're interested in the "old school methodology" this article explains it: towardsdatascience.com/classifying-facial-emotions-via-machine-learning-5aac111932d3.
The history behind this field is fascinating ! Here is a short summary of its history https://kapernikov.com/basic-introduction-to-computer-vision/

### Project goal and suggested timeline

The goal of the project is to implement a **system that detects the emotion on a face from a webcam video stream**. To achieve this exciting task you'll have to understand how to:

- deal with images in Python
- detect a face in an image
- train a CNN to detect the emotion on a face

That is why I suggest to start the project with a preliminary step. The goal of this step is to understand how CNNs work and how to classify images. This preliminary step should take approximately **two weeks**.

Then starts the emotion detection in a webcam video stream step that will last until the end of the project !

The two steps are detailed below.

### Preliminary:

- Take this lesson. This course is a reference for many reasons and one of them is the creator: **Andrew Ng**. He explains the basics of CNNs but also some more advanced topics as transfer learning, siamese networks etc ... I suggest to focus on Week 1 and 2 and to spend less time on Week 3 and 4. Don't worry the time scoping of such MOOCs are conservative ;-). Here is the link: https://www.coursera.org/learn/convolutional-neural-networks . You can attend the lessons for free !

- Participate to this challenge: https://www.kaggle.com/c/digit-recognizer/code . The MNIST dataset is a reference in computer vision. Researchers use it as a benchmark to compare their models. Start first with a logistic regression to understand how to handle images in Python. And then train your first CNN on this data set.

### Face emotions classification

Emotion detection is one of the most researched topics in the modern-day machine learning arena. The ability to accurately detect and identify an emotion opens up numerous doors for Advanced Human Computer Interaction. The aim of this project is to detect up to seven distinct facial emotions in real time. This project runs on top of a Convolutional Neural Network (CNN) that is built with the help of Keras whose backend is TensorFlow in Python. The facial emotions that can be detected and classified by this system are Happy, Sad, Angry, Surprise, Fear, Disgust and Neutral.

Your goal is to implement a program that takes as input a video stream that contains a person's face and that predicts the emotion of the person.

**Step 1**: **Fit the emotion classifier**

- Train a CNN on the dataset `train.csv`. Here is an example of architecture you can implement: https://www.quora.com/What-is-the-VGG-neural-network . **The CNN has to perform more than 70% on the test set**. You will see that the CNNs take a lot of time to train. You don't want to overfit the neural network. I strongly suggest to use early stopping, callbacks and to monitor the training using the tensorboard.

You have to save the trained model in `my_own_model.pkl` and to explain the chosen architecture in `my_own_model_architecture.txt`. Use `model.summary())` to print the architecture. It is also expected that you explains the iterations and how you end up choosing your final architecture. Save a screenshot of the tensorboard while the model's training in `tensorboard.png` and save a plot with the learning curves showing the model training and stopping BEFORE the model starts overfitting in `learning_curves.png`.

- Optional: Use a pre-trained CNN to improve the accuracy. You will find some huge CNN's architecture that perform well. The issue is that it is expensive to train them from scratch. You'll need a lot of GPUs, memory and time. **Pre-trained CNNs** solve partially this issue because they are already trained on a dataset and perform well on some use cases. However, building a CNN from scratch is required, as mentioned, this step is optional and doesn't replace the first one. Similarly, save the model and explain the chosen architecture.

**Step 2**: **Classify emotions from a video stream**

- Use the video stream outputted by your computer's webcam and preprocess it to make it compatible with the CNN you trained. One of the preprocessing steps is: face detection. As you may have seen the training samples are imaged centered on a face. To do so, I suggest to use a pre-trained model to detect faces. OpenCV for image processing tasks where we identify a face from a live webcam feed which is then processed and fed into the trained neural network for emotion detection. The preprocessing pipeline will be corrected with a functional test in `preprocessing_test`:

  - **Input**: Video stream of 20 sec with a face on it
  - **Output**: 20 (or 21) images cropped and centered on the face with 48 x 48 grayscale pixels

- Predict at least one emotion per second from the video stream. The minimum requirement is printing in the prompt the predicted emotion with its associated probability. If there's any problem related to the webcam use as input the a recorded video stream.

For that step, I suggest again to use **OpenCV** as much as possible. This link shows how to work with a video stream with OpenCV. OpenCV documentation may become deprecated in the futur. However, OpenCV will always provide tools to work with video streams, so search on the internet for OpenCV documentation and more specifically "opencv video streams". https://docs.opencv.org/4.x/dd/d43/tutorial_py_video_display.html

- Optional: **(very cool)** Hack the CNN. Take a picture for which the prediction of your CNN is **Happy**. Now, hack the CNN: using the same image **SLIGHTLY** modified make the CNN predict **Sad**. https://medium.com/@ageitgey/machine-learning-is-fun-part-8-how-to-intentionally-trick-neural-networks-b55da32b7196

### Deliverable

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

- Run **predict.py** expected output:

```prompt
python predict.py

Accuracy on test set: 72%

```

- Run **predict_live_stream.py** expected output:

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

### Useful ressources:

- https://machinelearningmastery.com/what-is-computer-vision/

- Use a pre-trained CNN: https://arxiv.org/pdf/1812.06387.pdf

- Hack the CNN https://medium.com/@ageitgey/machine-learning-is-fun-part-8-how-to-intentionally-trick-neural-networks-b55da32b7196

- http://ice.dlut.edu.cn/valse2018/ppt/WeihongDeng_VALSE2018.pdf

- https://arxiv.org/pdf/1812.06387.pdf

### Files needed for this project

[link](https://assets.01-edu.org/ai-branch/project3/challenges-in-representation-learning-facial-expression-recognition-challenge.zip)
