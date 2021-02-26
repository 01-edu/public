### Nesting organs

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

Add the following CSS to your `styles.css` file to see the freshly-added nested elements:

```
section {
  display: flex;
  justify-content: space-between;
}

div,
p {
  border: solid 1px black;
  padding: 5px;
  margin: 0;
}

#face {
  flex-direction: column;
  align-items: center;
  border-radius: 30px;
}

#eyes {
  width: 60px;
  display: flex;
  justify-content: space-between;
  background-color: cyan;
}

#torso {
  width: 100px;
  background-color: violet;
}
```

### Expected output

This is what you should see in the browser:

<!-- screenshot -->

### Notions

- [Anatomy of an HTML element](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started#anatomy_of_an_html_element)
- [Nesting HTML elements](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started#nesting_elements)
