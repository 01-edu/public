## Class that!

### Resources

We provide you with some content to get started smoothly, check it out!

- Video [CSS - Set & style with CSS class](https://www.youtube.com/watch?v=-U397k4VloU&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=6)

### Instructions

Alright, your being is almost done, some elements still need a bit more shaping and then we'll make it come to life!
If you look at your page, you can observe that some elements come by pair: the eyes, the arms & the legs. It is the same organ, one on the left and one on the right ; they have exactly the same shape, so for practicity & to avoid to repeat twice the same style, we're not going to use their `id` to style them, but a [`class`](https://developer.mozilla.org/en-US/docs/Web/CSS/Class_selectors) ; contrary to an `id`, a `class` can be attributed to several different elements with common rulesets, and so the style defined for that class will apply to all the HTML elements that have it.

Create the 3 following classes, setting them with the given rulesets, & attribute them to the corresponding HTML elements:

- class `eye`:
  - `width` of 60 pixels
  - `height` of 60 pixels
  - `background-color` "red"
  - `border-radius` of 50%
  - attributed to `eye-left` & `eye-right`
- class `arm`:
  - `background-color` "aquamarine"
  - attributed to `arm-left` & `arm-right`
- class `leg`:
  - `background-color` "dodgerblue"
  - attributed to `leg-left` & `leg-right`

Note that you can attribute several classes to a same element ; create the class `body-member`, which set the `width` to 50 pixels and the `margin` to 30 pixels, and add it to the `class` attribute of those elements: `arm-left`, `arm-right`, `leg-left` & `leg-right`.

### Code examples

Declare a class `my-first-class` and style it with a `color` to `"blue"` and a `background-color` to `"pink"`:

```css
.my-first-class {
  color: blue;
  background-color: pink;
}
```

Apply classes to HTML elements:

```html
<div class="my-first-class"></div>
<div class="another-class"></div>
<div class="my-first-class another-class"></div>
```

### Expected output

This is what you should see in the browser:
![](https://github.com/01-edu/public/raw/master/subjects/class-that/class-that.png)

### Notions

- [CSS class](https://developer.mozilla.org/en-US/docs/Web/CSS/Class_selectors)
