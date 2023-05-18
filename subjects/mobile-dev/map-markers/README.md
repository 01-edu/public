## map-markers

In this project, your team will create an app that uses the `Google Maps API`. The app must allow users to save and display their favorite places on a map.

### Instructions

The app should have the following functionality:

A `TabBar` with three screens:

- A `Maps` screen like [google maps](https://codelabs.developers.google.com/codelabs/google-maps-in-flutter#0) where the user can see their favorite places marked on the map with an `info view`.

- A `Favorite Places` screen which will have a list of the user's favorite places.

- An `InfoPage` screen that displays the names and email addresses of the app's authors, along with a general description of the app.

**On the Map screen:**

- The user's favorite places should be displayed as `markers` on the map (with an info view) and saved so that they persist after the app is closed.
- When a `marker` is tapped, a dialog window should open displaying the place's title/name.
- A button should be provided to navigate to the user's current location.
- A search bar needs to be included in order to allow the user to search for addresses and places, you can use the `Google Places API`. When a suggested address is tapped, the map should navigate to that place.

<center>
<img src="./resources/map.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
<img src="./resources/favorite-in-map.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
<img src="./resources/search.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>

**On the Favorite Places screen:**

- The user's list of favorite places should be displayed and the user should be able to delete them.
<center>
<img src="./resources/favorites.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
<img src="./resources/delete.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>

Here when the user slides it to the right it deletes the favorite place.

</center>

**On the Info screen:**

- The names and email addresses of the app's developers should be displayed, along with the year of development and any other relevant information.

<center>
<img src="./resources/infopage.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>

### Hints

> Here is how to get your [Google Maps API key](https://blog.logrocket.com/google-maps-flutter/) and use it.

> Note: Don't forget to add the following to your `AndroidManifest.xml` file as well as your API key:

```xml
<uses-permission android:name="android.permission.ACCESS_COARSE_LOCATION"/>
<uses-permission android:name="android.permission.ACCESS_FINE_LOCATION"/>
```

> Note: Don't forget to add the following to your `info.plist` file as well as your API key in `AppDelegate.swift`:

```xml
<key>NSLocationWhenInUseUsageDescription</key>
<key>NSLocationAlwaysUsageDescription</key>
```

These entries are required in order to use the device's location services.
