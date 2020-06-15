## Where do we go?

### Instructions

Tired of staying home for too long, you decide to develop a page to index ideas for your next travel destinations, so that next time you'll ask yourself 'Where do we go?', you won't need to get lost for 3 hours!

Create a page which displays the list of `places` provided in the data file below:

- sort the `places` from the Northest to the Southest
- display a fullscreen-size image for each place ; use the images hosted here: `https://github.com/MarieMalarme/dom-js/tree/master/assets/images`, also available as Github Pages here `https://mariemalarme.github.io/dom-js/assets/images/locationName.jpg`
- display a location indicator, displaying the `name` and the `coordinates` of the current place featured in the image, using the corresponding `color` as text color, which updates on scroll when another image is reached
- display a compass indicating the latitude direction ; North if the user is scrolling up, South if he's scrolling down
- when clicking on the page, open a link redirecting to the Google Maps' coordinates of the place currently displayed.

### Notions

- [Wheel event](https://developer.mozilla.org/en-US/docs/Web/API/Element/wheel_event): [`deltaY`](https://developer.mozilla.org/en-US/docs/Web/API/WheelEvent/deltaY)
- [`window`](https://developer.mozilla.org/en-US/docs/Web/API/Window): [`innerHeight`](https://developer.mozilla.org/en-US/docs/Web/API/Window/innerHeight), [`scrollY`](https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY), [`open()`](https://developer.mozilla.org/en-US/docs/Web/API/Window/open)

### Provided files

- Import the `places` from the data file: [https://mariemalarme.github.io/dom-js/assets/data/where-do-we-go.js](https://mariemalarme.github.io/dom-js/assets/data/where-do-we-go.js)

### Expected result

You can see an example of the expected result [here](https://youtu.be/BLxNi1WH6_0)
