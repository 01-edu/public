## Select & style

### Instructions

Now that you created & identified properly the different sections of your being,
it's time to make it look more living-like! To achieve that, you're going to
style it with [CSS][0]. Create a CSS file, [link it][1] to your
`select-and-style.html`, and:

- target all the elements with the [universal selector][2] and style them with:
  - `margin` of `0`
  - `box-sizing` to `border-box`
  - `opacity` of `0.85`

- target the `body` tag and style it with a `height` of `100vh` so it takes the
  viewport height

- target all the `section` tags with the [type selector][3], and style it with:
  - `padding` of `20px`
  - `width` of `100%`
  - `height` of `calc(100% / 3)` _(one third of the `body` height)_

- target each following element with the [`id` selector][4], using the `id` you
  defined earlier for each section, and style it:
  - `face` with a "cyan" `background-color`
  - `upper-body` with a "blueviolet" `background-color`
  - `lower-body` with a "lightsalmon" `background-color`

To style an element, you systematically have to declare [rulesets][5], composed
of a property and a value. Here is an exemple of how to set the color of `div`
tags to `"red"`:

```css
div {
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
