## Where do we go?

### Instructions

Tired of staying home for too long, you decide to develop a page to index ideas for your next travel destinations, so that next time you'll ask yourself 'Where do we go?', you won't need to get lost for 3 hours!

Write the function `explore` which creates a page displaying the list of `places` provided in the data file below:

- sort the `places` from the Northest to the Southest
- display a fullscreen-size `<section>` for each place ; use the pics hosted in the `./images` folder to set the `background` attribute with the corresponding image URL. The URL has to be formatted like so: `./images/name-of-the-place.jpg` ; you can see [here](https://github.com/01-edu/public/tree/master/subjects/where-do-we-go/images) the list of images that the folder contains
- display a location indicator as a `<a>` tag in the middle of the screen which:
  - has the class `location`
  - displays as text strings separated by `\n`, the `name` and the `coordinates` of the current place featured in the image
  - using the corresponding `color` as text color
  - updates the `name`, `coordinates` and `color` on scroll, when the top of the next image reaches the middle of the screen height
  - has the `href` attribute set to open **a new tab** redirecting to a Google Maps' URL with the coordinates of the place currently displayed
- display a compass indicating the latitude direction ; `N` for North if the user is scrolling up, `S` for South if he's scrolling down

### Notions

- [Scroll event](https://developer.mozilla.org/en-US/docs/Web/API/Element/scroll_event)
- [`window`](https://developer.mozilla.org/en-US/docs/Web/API/Window): [`innerHeight`](https://developer.mozilla.org/en-US/docs/Web/API/Window/innerHeight), [`scrollY`](https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY), [`pageYOffset`](https://developer.mozilla.org/en-US/docs/Web/API/Window/pageYOffset)
- Take a look at the [DMS coordinates system](https://en.wikipedia.org/wiki/Decimal_degrees)

### Provided files

You have at your disposal the following files that you can download to test locally:

- the HTML file [index.html](/public/subjects/where-do-we-go/index.html) to open in the browser, which includes:

  - the JS script which will allow to run your code
  - some CSS pre-styled classes: feel free to use those as they are, or modify them

- the data file [data.js](/public/subjects/where-do-we-go/data.js) from which you can import `places`

### Expected result

You can see an example of the expected result [here](https://youtu.be/BLxNi1WH6_0)
