## Where do we go?

### Instructions

Tired of staying home for too long, you decide to develop a page to index ideas for your next travel destinations, so that next time you'll ask yourself 'Where do we go?', you won't need to get lost for 3 hours!

Write the function `explore` which creates a page displaying the list of `places` provided in the data file below:

- sort the `places` from the Northest to the Southest
- display a fullscreen-size `<section>` for each place ; use the pics hosted in the `./where-do-we-go_images` folder (find the download link below) to set the `background` attribute with the corresponding image URL. The URL has to be formatted like so: `./where-do-we-go_images/name-of-the-place.jpg`
- display a location indicator as a `<a>` tag in the middle of the screen which:
  - has the class `location`
  - displays as text strings separated by `\n`, the `name` and the `coordinates` of the current place featured in the image
  - using the corresponding `color` as text color
  - updates the `name`, `coordinates` and `color` on scroll, when the top of the next image reaches the middle of the screen height
  - has the `href` attribute set to open **a new tab** redirecting to a Google Maps' URL with the coordinates of the place currently displayed
- display a compass as a `div` tag indicating the latitude direction which:
  - has the class `direction`
  - displays `N` for North if the user is scrolling up
  - displays `S` for South if he's scrolling down

### Notions

- [Scroll event](https://developer.mozilla.org/en-US/docs/Web/API/Element/scroll_event)
- [`window`](https://developer.mozilla.org/en-US/docs/Web/API/Window): [`innerHeight`](https://developer.mozilla.org/en-US/docs/Web/API/Window/innerHeight), [`scrollY`](https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY), [`pageYOffset`](https://developer.mozilla.org/en-US/docs/Web/API/Window/pageYOffset)
- Take a look at the [DMS coordinates system](https://en.wikipedia.org/wiki/Decimal_degrees)

### Files

You only need to create & submit the JS file `where-do-we-go.js` ; we're providing you the following files to download (click right and save link) & test locally:

- the HTML file [where-do-we-go.html](./where-do-we-go.html) to open in the browser, which includes:

  - the JS script which will allow to run your code
  - some CSS pre-styled classes: feel free to use those as they are, or modify them

- the data file [where-do-we-go.data.js](./where-do-we-go.data.js) from which you can import `places`

- the images to use, in this [compressed folder](https://assets.01-edu.org/where-do-we-go_images.zip)

### Expected result

You can see an example of the expected result [here](https://youtu.be/BLxNi1WH6_0)
