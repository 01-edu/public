## Class it!

> Brainpower Mode

### Context

Alright, your being is almost done; some elements still need a bit more shaping, and then we'll make it come to life!

### Resources

We provide you with some content to get started smoothly. Check it out!

- Video [CSS - Set & style with CSS class](https://www.youtube.com/watch?v=-U397k4VloU&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=6)

### Instructions

If you look at your page, you can observe that some elements come in pairs: the eyes, the arms, and the legs. They are the same organ, one on the left and one on the right; they have exactly the same shape. So for practicality and to avoid repeating the same style twice, we're not going to use their `id` to style them, but a [`class`](https://developer.mozilla.org/en-US/docs/Web/CSS/Class_selectors). Unlike an `id`, a `class` can be attributed to several different elements with common rulesets, so the style defined for that class will apply to all the HTML elements that have it.

#### Task 1

Create the following three classes, set them with the given rulesets, and attribute them to the corresponding HTML elements:

- Class `eye`:
  - `width` of 60 pixels
  - `height` of 60 pixels
  - `background-color` "red"
  - `border-radius` of 50%
  - Attributed to `eye-left` & `eye-right`
- Class `arm`:
  - `background-color` "aquamarine"
  - Attributed to `arm-left` & `arm-right`
- Class `leg`:
  - `background-color` "dodgerblue"
  - Attributed to `leg-left` & `leg-right`

Note that you can attribute several classes to the same element. Create the class `body-member`, which sets the `width` to 50 pixels and the `margin` to 30 pixels, and add it to the `class` attribute of these elements: `arm-left`, `arm-right`, `leg-left`, & `leg-right`.

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

> Need help? Ask your tablemate.

### Notions

- [CSS class](https://developer.mozilla.org/en-US/docs/Web/CSS/Class_selectors)