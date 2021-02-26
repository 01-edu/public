### Select & style

Now that you created & identified properly the different elements of your being, it's time to make it look more living-like! To achieve that, you're going to style it with [CSS](https://developer.mozilla.org/en-US/docs/Web/CSS).
Create a CSS file, [link it](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link#including_a_stylesheet) to your `index.html`, and :

- target all the elements with the [universal selector](https://developer.mozilla.org/en-US/docs/Web/CSS/Universal_selectors) and style them with an `opacity` of 0.8
- target all the `section` tags with the [type selector](https://developer.mozilla.org/en-US/docs/Web/CSS/Type_selectors), and style it with a `padding` of 20 pixels and a `margin` of 15 pixels
- target each following element with the [`id` selector](https://developer.mozilla.org/en-US/docs/Web/CSS/ID_selectors), using the `id` you defined earlier for each section, and style it:
  - `face` with a `width` & `height` of 80 pixels, and a "deeppink" `background-color`
  - `upper-body` with a `width` of 200 pixels, a `height` of 220 pixels, and a "blue" `background-color`
  - `lower-body` with a `width` of 100 pixels, a `height` of 220 pixels, and a "yellow" `background-color`
- target the `body` tag and style it with a `display` at "flex", a `flex-direction` at "column" and a `align-items` at "center" (this is to turn the `body` into a [`flex` container](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout/Basic_Concepts_of_Flexbox), so the elements will be centered in the page)

To style an element, you systematically have to declare [rulesets](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics#anatomy_of_a_css_ruleset), composed of a property and a value.
Here is an exemple of how to set the color of `div` tags to "red":

```div {
  color: red;
}
```

### Notions

- [`link` a CSS file](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link#including_a_stylesheet)
- [CSS basics](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics)
- [ruleset](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics#anatomy_of_a_css_ruleset)
- [List of different selectors](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics#different_types_of_selectors)
- [universal selector](https://developer.mozilla.org/en-US/docs/Web/CSS/Universal_selectors)
- [type selector](https://developer.mozilla.org/en-US/docs/Web/CSS/Type_selectors)
- [`id` selector](https://developer.mozilla.org/en-US/docs/Web/CSS/ID_selectors)
- [Flexbox layout](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout/Basic_Concepts_of_Flexbox), you can train on [Flexbox froggy](https://flexboxfroggy.com/)
