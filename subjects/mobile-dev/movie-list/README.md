# Movie List

### Introduction

When developing a fullly functional app, you will most likely use some external files to display on your app. One of the most popular file types is JSON.  
Develop an app to work with JSON. You are given a file with information about different movies.
You must display top rated movies with descending rating on the first page of the app. By tapping on a movie, a new route with more detailed information regarding the movie must be displayed.  
Searching for a movie via search bar must be included. Searching must be done by movie name, Entered string must be completely within movie name, i.e. if "vatar" is searched, "Avatar" must be included in the response. Basically, search must work like SQL's **_ilike_** comparision.

### Objective

- Infinite scroll
- ListView
- Json Serialization
- Routes/Push
- AppBar actions
- Images
- Futures, async
- Provider

### Part 1

Create a ListView which will show the image, title and description of movies from JSON file (see **Assets** section);

Create a class with properties:

```
class Movie {
  final String genre;
  final String imdbRating;
  final String title;
  final String poster;

  Movie(this.genre, this.imdbRating, this.title, this.poster);
}
```

Add fromJson method to make a json serialization, see more on [here](https://flutter.dev/docs/development/data-and-backend/json).

Your ListView should use [FutureBuilder](https://api.flutter.dev/flutter/widgets/FutureBuilder-class.html) to wait for data from JSON file and then show it once it is loaded.

<center>
<img src="./resources/movieList.01.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>

### Part 2

Create a page with detailed information about the movie.

It should have an image of the movie, at least 5 parameters from the film's information, i.e. when it was filmed, main actors, anything else you might consider useful. Use scroll bar if the info doesn't fit in one page.

The appbar should have a name of the film and go back button.

Visit [documentaion](https://flutter.dev/docs/cookbook/navigation/named-routes) to see how to implement routing.

### **Assets**

[Movies](movies.json).

### **Bonus**

- Make [navigation transition](https://docs.flutter.dev/resources/platform-adaptations) when switching routes.
- Make [back navigation](https://docs.flutter.dev/resources/platform-adaptations#back-navigation).
- Animate [overscroll behaviour](https://docs.flutter.dev/resources/platform-adaptations#overscroll-behavior).
