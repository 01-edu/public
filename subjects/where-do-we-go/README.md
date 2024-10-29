## Where do we go?

### Instructions

Where will you go on your next holiday?

Let's make a page to index your options, so that next time you ask yourself that question, you'll be ready with some ideas.

Create a function named `explore`, which creates a page displaying the list of `places` provided in the data file below.

- Sort the `places` from north to south, so that the northern-most place is first.
- Display a fullscreen-size `<section>` for each place. Use the pics hosted in the `./where-do-we-go_images` folder below. Set the `background` attribute with the corresponding image URL. The URL has to be formatted like so: `./where-do-we-go_images/name-of-the-place.jpg`.
- Display a location indicator as an `<a>` tag in the middle of the screen. It should:
  - have the class `location`
  - display the `name` and `coordinates` of the current place, as text strings separated by `\n`
  - set the text color as `color`
  - update the `name`, `coordinates` and `color` on scroll, at the point when the next image reaches the middle of the screen height
  - make the `href` attribute open **a new tab** redirecting to a Google Maps URL with the coordinates of the place currently displayed
- Display a compass as a `div` tag, indicating the latitude direction which:
  - has the class `direction`
  - displays `N` for North if the user is scrolling up
  - displays `S` for South if he's scrolling down

### Files

You only need to create & submit the JS file `where-do-we-go.js`; we're providing you the following files to download and test locally:

- the HTML file [where-do-we-go.html](./where-do-we-go.html) to open in the browser, which includes:

  - the JS script which will enable you to run your code

- the data file [where-do-we-go.data.js](./where-do-we-go.data.js) from which you can import `places`

- feel free to use the CSS file [where-do-we-go.data.css](./where-do-we-go.data.css) as it is or you can also modify it.

- you can get the images to be used in this [compressed folder](https://assets.01-edu.org/where-do-we-go_images.zip) or you can get them in the where-do-we-go_images folder from the public URL. Example: `https://public.01-edu.org/subjects/where-do-we-go/where-do-we-go_images/arlit.jpg
### Expected result

You can see an example of the expected result [here](https://youtu.be/BLxNi1WH6_0)

### Notions

- [Scroll event](https://developer.mozilla.org/en-US/docs/Web/API/Element/scroll_event)
- [window](https://developer.mozilla.org/en-US/docs/Web/API/Window)
- [innerHeight](https://developer.mozilla.org/en-US/docs/Web/API/Window/innerHeight)
- [scrollY](https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY)
- [pageYOffset](https://developer.mozilla.org/en-US/docs/Web/API/Window/pageYOffset)
- Take a look at the [DMS coordinates system](https://en.wikipedia.org/wiki/Decimal_degrees)
