## Favorite Images

Nowadays, it is very difficult to imagine a popular app that does not work with storing, accessing images. This is a practical exercise designed to help you practice it in a mobile app.

This app must allows you to create a personal gallery where you can upload images from your phone library or take new pictures and immediately add them to your collection.

To build this app, you have as allowed packages the [image_picker](https://pub.dev/packages/image_picker), which enables you to access the device's camera and gallery.

### Instructions

### Image Picker

In the first part of the exercise, you will create an `appbar` with an **IconButton**.
When the user clicks on the button, they should see two options:

- Open the camera.
- Access the gallery.

> Note: Your camera must work both in `iOS` and `Android`.

You should use the `image_picker` package to implement this functionality. Additionally, you should show a `"No images selected"` message if no images have been loaded.

> Note: Don't forget to add the special keys to the Info.plist file under the iOS folder to access the camera and gallery, see the [example](https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/AboutInformationPropertyListFiles.html).

<center>
<img src="./resources/imageLibrary.01.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>

### Image Gallery

In the second part of the exercise, you will add the following functionalities:

- Add the selected image or taken picture to a `GridList` to display all the images in your collection.
- Implement the ability for the user to `tap on an image` and view the entire picture, providing a better user experience.
- Add the ability for the user to `zoom in and out` of the image, allowing them to see the image details more clearly.

<center>
<img src="./resources/imageLibrary.02.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
<img src="./resources/imageLibrary.03.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>

### Bonus

Here are three possible bonus features you could add to this project:

- `Image editing features`, you could add basic image editing functionalities such as cropping, rotating, or adding filters to the images before they are added to the gallery.

- `Cloud storage` to make sure that the user's images are safe, you could implement cloud storage so that the images are stored on a remote server rather than just on the user's device.

- `Sharing options`, to improve the user experience, you could add sharing options so that users can easily share their favorite images on social media platforms or through email or messaging.
