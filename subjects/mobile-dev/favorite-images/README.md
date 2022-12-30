# Favorite Images

### Introduction

Nowadays, it is very difficult to imagine a popular app that does not work with storing, accessing images.  
In this app you will get a chance to practise it.  
You should create your gallery, where you can load images from phone library or take a picture and immediately upload it.

Allowed package: [image_picker](https://pub.dev/packages/image_picker).

### First Part

- Create an appbar with **IconButton**.
- When you click on a button you should see 2 options: open camera or a gallery.
- See the usage of [image_picker](https://pub.dev/packages/image_picker) package.
- Don't forget to add special keys to Info.plist file under ios folder to access the camera and gallery, see [example](https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/AboutInformationPropertyListFiles.html).
- Show "No images selected" message if no images are loaded.

<center>
<img src="./resources/imageLibrary.01.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>

</center>

### Second Part

- When user is done picking image or taking picture it should be added to GridList
- When an image is tapped, user should see the whole image
- Add an ability to zoom in, zoom out the image, see [example](https://api.flutter.dev/flutter/widgets/InteractiveViewer-class.html).

<center>
<img src="./resources/imageLibrary.02.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>

<img src="./resources/imageLibrary.03.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>
