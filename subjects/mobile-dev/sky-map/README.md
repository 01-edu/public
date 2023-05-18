## Sky map

A sky map app is a mobile application that provides users with an interactive map of the night sky. It utilizes the device's sensors, such as the GPS, accelerometer, and magnetometer, to determine the user's location and orientation. With this information, the app can display a real-time representation of the stars, planets, constellations, and other celestial objects visible from the user's location.

### Instructions

To manage states within the app, you should use either the BLoC or Provider pattern.

To complete this exercise, you will need to do the following:

- Utilize the `GPS`, `accelerometer`, and `magnetometer` sensors to determine the `location` and `position` of the device.
- Implement `real-time updates` by updating an image at least **_10 times per second_**.
- Send a request to a public API to retrieve information about celestial objects or use a valid file.
- Display the celestial objects on a black canvas within the app.
- Allow users to tap on an object to view a short description of it.
- Use either the `BLoC` or `Provider` pattern to manage states within the app.
- The objects to be displayed in the app include:
  - All planets of the solar system
  - The sun
  - The moon
  - At least three constellations of your choice
- When moving the phone around, the celestial objects shown on the display should change

Remember to follow best practices for coding and app development, and be sure to document your code and any decisions made during the development process. Don't forget to consider user experience and design when creating the app, as a visually appealing and intuitive interface will be important for its success.

Possible examples of the app:

<center>
<img src="./resources/skyMap.01.jpg?raw=true" style = "width: 500px !important; height: 300px !important;"/>
</center>

<center>
<img src="./resources/skyMap.02.png?raw=true" style = "width: 500px !important; height: 300px !important;"/>
</center>
