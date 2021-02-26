### Class that

Alright, your being is almost done, some elements still need a bit more shaping and then we'll make it come to life!
If you look at your page, you can observe that some elements are repeated: the eyes, the arms & the legs. It is the same organ, one on the left and one on the right. They have exactly the same shape, so for practicity & to avoid to repeat twice the same style, we're not going to use their `id` to style them, but a [`class`](https://developer.mozilla.org/en-US/docs/Web/CSS/Class_selectors) ; contrary to an `id`, a `class` can be attributed to several different elements with common rulesets, and so the style defined for that class will apply to all the HTML elements that have it.

Create the 3 following classes, setting them with the given rulesets, & attribute them to the corresponding HTML elements:

- class `eye`:
  - `height` of 10 pixels
  - `background-color` "red"
  - `border-radius` of 50%
  - attributed to `eye-left` & `eye-right`
- class `arm`:
  - `background-color` "aquamarine"
  - attributed to `arm-left` & `arm-right`
- class `leg`:
  - `background-color` "blue"
  - attributed to `leg-left` & `leg-right`

You can attribute several classes to a same element ; create the class `width10`, which set the `width` to 10 pixels, and add it to the `class` attribute of all the repeated elements: `eye-left`, `eye-right`, `arm-left`, `arm-right`, `leg-left` & `leg-left`.

### Expected output

This is what you should see in the browser:

<!-- screenshot -->

### Notions

- [CSS class](https://developer.mozilla.org/en-US/docs/Web/CSS/Class_selectors)
