## Movie List

### Introduction

When developing a fully functional app, you will most likely use some external files to display on your app. One of the most popular file types is `JSON`.

In this exercise, you will develop a fully functional app that uses `JSON` files to display information about movies. You will be given a [movies](movies.json) file with information about different movies.Your app needs to have the following features:

- Display the top rated movies with descending rating on the home page of the app.
- By tapping on a movie, a new route with more detailed information regarding the movie must be displayed.
- You need to include a `search bar` to search for movies by name. The search function must return movies whose name includes the entered string. For example, if "vatar" is searched, "Avatar" must be included in the response. The search function must work like **SQL's ilike** comparison.

### ListView

Create a `ListView` which will show the `image`, `title`, and `genre` of movies from the given [movies](movies.json) file. To create the `ListView`, you need to create a `Movie class` with the following properties:

```flutter
class Movie {
  final String genre;
  final String imdbRating;
  final String title;
  final String poster;

  Movie(this.genre, this.imdbRating, this.title, this.poster);
}
```

You also need to add a `fromJson method` to make a [JSON serialization](https://flutter.dev/docs/development/data-and-backend/json).

Once you have the `Movie class`, use [FutureBuilder](https://api.flutter.dev/flutter/widgets/FutureBuilder-class.html) to wait for data from the `JSON` file and then show it once it is loaded.

Here's an example of what your `ListView` should look like:

<center>
<img src="./resources/movieList.01.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>

### Movie page

Create a page with detailed information about the movie.

The page should have an `image` of the movie and at least five parameters from the movie's information, i.e. the year it was filmed, main actors, and anything else you might consider useful. If the information doesn't fit on one page, use a scroll bar to show all the information.

The `appbar` in this page should have the name of the movie and a back button.

> Note: widgets should not be [overflowed](https://blog.flutterflow.io/preventing-layout-overflows/)!

### Bonus

- Make [navigation transitions](https://docs.flutter.dev/resources/platform-adaptations) when switching routes.
- Make [back navigation](https://docs.flutter.dev/resources/platform-adaptations#back-navigation).
- Animate [overscroll behaviour](https://docs.flutter.dev/resources/platform-adaptations#overscroll-behavior).
- Use a database to store movie data instead of a JSON file for better performance and scalability.

### Notions

- [Navigate with named routes](https://flutter.dev/docs/cookbook/navigation/named-routes)
- [Infinite scroll](https://pub.dev/packages/infinite_scroll_pagination)
- [ListView](https://api.flutter.dev/flutter/widgets/ListView-class.html)
- [Json Serialization](https://docs.flutter.dev/development/data-and-backend/json#serialization)
- [Routes/Push](https://docs.flutter.dev/development/ui/navigation)
- [AppBar actions](https://docs.flutter.dev/development/ui/widgets/material#AppBar)
- [Futures, async](https://dart.dev/codelabs/async-await)
