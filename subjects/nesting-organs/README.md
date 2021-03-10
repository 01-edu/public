## Nesting organs

### Instructions

Bravo! You displayed the global shape of your entity, but now it's time to populate each division ; let's add up some organs!
To do so, we're going to introduce you the concept of nesting elements inside others.

So far, you just have a unique layer in your `<body>`: `face`, `upper-body` & `lower-body` are all at the same level.
But as you know, on a face, there are 2 eyes, a nose, and a mouth - and inside that mouth, a tongue, etc. ; any element can potentially be a container for other elements.

Let's add new elements and wrap them in different layers ; convert this list of organs in a HTML structure with the corresponding given tags!

- face: `section` tag with `id` `face`
  - eyes: `div` tag with `id` `eyes`
    - eye left: `p` tag with `id` `eye-left`
    - eye right: `p` tag with `id` `eye-right`
- upper body: `section` tag with `id` `upper-body`
  - arm left: `div` tag with `id` `arm-left`
  - torso: `div` tag with `id` `torso`
  - arm right: `div` tag with `id` `arm-right`
- lower body: `section` tag with `id` `lower-body`
  - left leg: `div` tag with `id` `leg-left`
  - right leg: `div` tag with `id` `leg-right`

Modify your CSS file to add rulesets to `section` tags: `display` at "flex" and `justify-content` at "center" (this is to turn the `section` tags into [`flex` containers](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout/Basic_Concepts_of_Flexbox), so the elements inside will be centered)

Add also the following CSS to your CSS file to see the freshly-added nested elements:

```
div,
p {
  border: solid 1px black;
  padding: 10px;
  margin: 0;
  border-radius: 30px;
}

#face {
  align-items: center;
}

#eyes {
  display: flex;
  background-color: yellow;
  justify-content: space-between;
  align-items: center;
  border-radius: 50px;
  width: 200px;
}

#torso {
  width: 200px;
  background-color: violet;
}
```

### Code examples

Nest several elements:

```html
<div id="first-element">
  <span id="second-element"></span>
  <div id="third-element">
    <p id="fourth-element"></p>
  </div>
</div>
```

### Expected output

This is what you should see in the browser:
![](https://github.com/01-edu/public/raw/master/subjects/nesting-organs/nesting-organs.png)

### Notions

- [Anatomy of an HTML element](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started#anatomy_of_an_html_element)
- [Nesting HTML elements](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started#nesting_elements)
- [Flexbox layout](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout/Basic_Concepts_of_Flexbox), you can train on [Flexbox froggy](https://flexboxfroggy.com/)
