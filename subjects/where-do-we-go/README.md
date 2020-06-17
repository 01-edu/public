## Where do we go?

### Instructions

Tired of staying home for too long, you decide to develop a page to index ideas for your next travel destinations, so that next time you'll ask yourself 'Where do we go?', you won't need to get lost for 3 hours!

Write the function `explore` which creates a page displaying the list of `places` provided in the data file below:

- sort the `places` from the Northest to the Southest
- display a fullscreen-size image for each place ; use the images hosted [here](/public/subjects/where-do-we-go/images)
- display a location indicator, displaying the `name` and the `coordinates` of the current place featured in the image, using the corresponding `color` as text color, which updates on scroll when another image is reached
- display a compass indicating the latitude direction ; North if the user is scrolling up, South if he's scrolling down
- when clicking on the page, open a link redirecting to the Google Maps' coordinates of the place currently displayed.

### Notions

- [Wheel event](https://developer.mozilla.org/en-US/docs/Web/API/Element/wheel_event): [`deltaY`](https://developer.mozilla.org/en-US/docs/Web/API/WheelEvent/deltaY)
- [`window`](https://developer.mozilla.org/en-US/docs/Web/API/Window): [`innerHeight`](https://developer.mozilla.org/en-US/docs/Web/API/Window/innerHeight), [`scrollY`](https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY), [`open()`](https://developer.mozilla.org/en-US/docs/Web/API/Window/open)
- Take a look at the [DMS coordinates system](https://en.wikipedia.org/wiki/Geographic_coordinate_system#Degrees:_a_measurement_of_angle)

### Provided files

- Check the HTML file [index.html](/public/subjects/where-do-we-go/index.html), which includes:

  - the JS script which will allow to run your code
  - some CSS pre-styled classes: feel free to use those as they are, or modify them

- Import `places` from the data file [data.js](/public/subjects/where-do-we-go/data.js)

### Expected result

You can see an example of the expected result [here](https://youtu.be/BLxNi1WH6_0)
