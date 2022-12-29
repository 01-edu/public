## map-markers

### Introduction

Create an app where you can save the list of your favorite places and show them in the app.

In this raid, your team should implement an app which uses google maps.

### Objective

Implement following functionality:

- Tabbar app with 3 screens:
  - Google maps screen [example](https://codelabs.developers.google.com/codelabs/google-maps-in-flutter#0)
  - List of your favorite places.
  - Info about your app. Names of the authors, their email adresses, and general description of what your application does.

On map there must be:

- Your favorite places as markers (with infoView) (places must be saved after reopenig the app)
- When some place is tapped, open dialog window with title and description textfields
- A button, which will navigate to current position of device
- Search bar to search for addresses, places, use google places API. When suggested address is tapped on, navigate map to that place

On FavouritesList Screen:

- List of favorite places, which you can delete

On InfoPage Screen:

- Your names as developers, year, etc

- Note: don't forget about info.plist:

```
NSLocationWhenInUseUsageDescription
NSLocationAlwaysUsageDescription
```
