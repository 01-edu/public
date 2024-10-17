## Embedded Organs

> Brainpower Mode

### Context

Bravo! You've outlined the global shape of your robot friend, but now it's time to bring it to life on Earth by adding essential parts. Let's equip your robot with its vital organs! To do this, we're going to introduce you to the concept of nesting elements inside others. This will allow you to build a fully functional robot, one piece at a time.

> Don't be afraid of the density of things you are asked to do. Just divide the work up step by step.

### Instructions

So far, you just have a unique layer in your `<body>`: `face`, `upper-body`, and `lower-body` are all at the same level.
But as you know, on a face, there are two eyes, a nose, and a mouth - and inside that mouth, a tongue, etc.; any element can potentially be a container for other elements.

#### Task 1

Let's add new elements and wrap them in different layers; convert this list of organs into an HTML structure with the corresponding given tags!

```html
<section id="face">
  <div id="eyes">
    <p id="eye-left"></p>
    <p id="eye-right"></p>
  </div>
</section>

<section id="upper-body">
  <div id="arm-left"></div>
  <div id="torso"></div>
  <div id="arm-right"></div>
</section>

<section id="lower-body">
  <div id="leg-left"></div>
  <div id="leg-right"></div>
</section>
```

#### Task 2

Modify your CSS file to add rulesets to `section` tags: `display` at "flex" and `justify-content` at "center" (this is to turn the `section` tags into [`flex` containers](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout/Basic_Concepts_of_Flexbox), so the elements inside will be centered)

#### Task 3

Add the following CSS to your CSS file to see the freshly-added nested elements:

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

> From now on, you can customize the `background-color` of the three following background sections from the CSS code by choosing the colors that you think best suit your theme.  
> \#face  
> \#upper-body  
> \#lower-body  
> You don't know which color is configurable? [Internet will give you the list](https://letmegooglethat.com/?q=css+color+list)

### Notions

- [Anatomy of an HTML element](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started#anatomy_of_an_html_element)
- [Nesting HTML elements](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started#nesting_elements)
- [Flexbox layout](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout/Basic_Concepts_of_Flexbox), you can train on [Flexbox froggy](https://flexboxfroggy.com/)
