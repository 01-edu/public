## Select then style

> Brainpower mode

### Context

To bring your robot friend to life, it's not enough to just build its skeleton with HTML. Your robot needs personality and style, and that's where CSS (Cascading Style Sheets) comes in. CSS is like the clothing and accessories that give your robot its unique look and feel. It helps you decorate and arrange the elements on your page, making everything look just right. Ready to start styling? Let’s dive in and give your robot buddy the perfect look!

### Resources

We provide you with some content to get started smoothly, check it out!

- Video [Link a CSS stylesheet to your HTML file](https://www.youtube.com/watch?v=e7G-KhaqTjs&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=3)
- Video [CSS - Style with type selectors](https://www.youtube.com/watch?v=q0ur7YWBzhs&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=4)
- Video [HTML/CSS - Set & style with ID selector](https://www.youtube.com/watch?v=3b3MiY-MR-Y&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=5)

### Instructions

Now that you have created and properly identified the different sections of your being, it's time to make it look more lifelike! To achieve that, you're going to style it with [CSS][0]. Create a CSS file, [link it][1] to your `select-and-style.html`, and:

- Target all the elements with the [universal selector][2] and style them with:

  - `margin` of `0`
  - `box-sizing` set to `border-box`
  - `opacity` of `0.85`

- Target the `body` tag and style it with a `height` of `100vh` so it takes the viewport height.

- Target all the `section` tags with the [type selector][3], and style them with:

  - `padding` of `20px`
  - `width` of `100%`
  - `height` of `calc(100% / 3)` _(one third of the `body` height)_

- Target each of the following elements with the [`id` selector][4], using the `id` you defined earlier for each section, and style them:
  - `face` with a "cyan" `background-color`
  - `upper-body` with a "blueviolet" `background-color`
  - `lower-body` with a "lightsalmon" `background-color`

### Code examples

To style an element, you systematically have to declare [rulesets][5], composed of a property and a value.

Set the color of `div` tags to `"red"`:

```css
div {
  color: red;
}
```

Set the `background-color` of the HTML element with the `id` `"block-1"`:

```css
#block-1 {
  color: red;
}
```

### Expected output

This is what you should see in the browser: ![screenshot][8]

### Notions

- [`link` a CSS file][1]
- [CSS basics][7]
- [ruleset][5]
- [List of different selectors][6]
- [universal selector][2]
- [type selector][3]
- [`id` selector][4]

[0]: https://developer.mozilla.org/en-US/docs/Web/CSS
[1]: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link#including_a_stylesheet
[2]: https://developer.mozilla.org/en-US/docs/Web/CSS/Universal_selectors
[3]: https://developer.mozilla.org/en-US/docs/Web/CSS/Type_selectors
[4]: https://developer.mozilla.org/en-US/docs/Web/CSS/ID_selectors
[5]: https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics#anatomy_of_a_css_ruleset
[6]: https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics#different_types_of_selectors
[7]: https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics
[8]: https://github.com/01-edu/public/blob/master/subjects/select-and-style/select-and-style.png?raw=true
