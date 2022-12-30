# Bizz Card

### Introduction

### What is Flutter?

Flutter is Google’s UI toolkit for building beautiful, natively compiled applications for mobile, web, and desktop from a single codebase. (c) [https://flutter.dev/](https://flutter.dev/)

### How to start?

You can build Flutter apps using any text editor. It is recommended to use editors that have both Dart and Flutter plugins (Android Studio, IntelliJ, VS Code).

You should already have all installed on school computers, but if you want to set up your PC, then follow the official [installation guide](https://flutter.dev/docs/get-started/install).

Before we start take a look at Flutter's [official documentation](https://flutter.dev/docs)

Hint: You can see some Flutter samples [here](https://flutter.github.io/samples/#?type=cookbook) 👩🏽‍🍳

### Objective

### Making your first app

BizzCard is a simple application which shows a static card with your personal information.

This subject is divided into 3 parts. Overall objective is for you to learn:

- About and apply the essence of Flutter - widgets.
- Basic structure of a Flutter app.
- How to run Flutter app on physical device or Android/iOS emulator.

- Note: only standard Flutter packages and url_launcher (for bonus) are allowed

### First Part

In this part:

- Run Flutter generated counter app
- Understand structure of the Flutter app

To create your first Flutter application open Android Studio or similar IDE and follow the steps:

1. Open the IDE and select **Start a new Flutter project**.
2. Select **Flutter Application** as the project type. Then click **Next**.
3. Specify path to **Flutter SDK’s location**
   (select **Install SDK…** if the text field is blank).
4. Enter the project name (for example, `myapp`). Then click **Next**.
5. Click **Finish**.
6. Wait for Android Studio to install the SDK and create the project.

See the Run the app section [here](https://flutter.dev/docs/get-started/test-drive?tab=androidstudio#create-app) to run a sample app.

The starter point in Flutter app is in lib/main.dart. Change this file to change app's behavior.

### Second Part:

In this part:

- Build your first app

You should display static information about yourself, i.e. name, surname, age, phone number, email, and image.
Try to make it as it is shown in the example below:

<!-- <center> -->
<center>
<img src="./resources/bizzCard.01.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
<!-- </center> -->
</center>

Possible [diagram](https://flutter.dev/docs/development/ui/layout) of widget tree for the app.

#### Helpful keywords for research:

- AppBar
- Column to align text in a card
- Stack to place image on Container
- Padding to add paddings to Container
- Use DecorationImage inside BoxDecoration to style image

### **Bonus**

- Add qr code below the image, so anyone can scan it and get your full contact
