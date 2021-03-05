## Carbon copy

### Video

Check out [this little video](https://youtu.be/ym-Qiio2ktI)!

### Instructions

Today is a big day: you're going to make your own webpage. Like a boss, yes.
If that one already sounds scary to you, don't worry, we've got your back ; you'll be provided some assets to complete this task.

The goal of this raid is to make a carbon copy of this webpage template:
![](https://github.com/01-edu/public/blob/master/subjects/carbon-copy/page-template.jpg)
Feel free to replace (or not, if you're lazy) the text contents by details about you / your company / your organisation / your dog / anything you want to talk about. Just try to approximatively match the same amount of content, so it still looks visually similar to the placeholder version.

The raid is divided in 3 phases:

- pure HTML structure
- custom CSS style
- JS interactions

### Phase 1: HTML only _(mandatory)_

Create & write the HTML file `index.html` to build the structure of the page.
For this phase, we provide you a CSS file (`styles.css`) that you just have to [link](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link) to your `index.html`, meaning you won't need to style anything. However, you can do it from scratch and write yourself the whole CSS if you feel up to it.

NB: Each item of the navbar menu - in the top right corner - has to bring to the corresponding section on click (clicking on "About" scrolls to the "About" section in the page); take a look at the [`href` anchor use](https://www.w3.org/TR/html401/struct/links.html#h-12.2.3).

Here is a [wireframe](https://en.wikipedia.org/wiki/Website_wireframe) of the webpage, showing the HTML tags you have to use:
![](https://raw.githubusercontent.com/01-edu/public/master/subjects/carbon-copy/page-wireframe.jpg)

### Phase 2: custom CSS _(mandatory)_

First of all, let's customize the color atmosphere of the webpage to your own taste: go to the CSS file `styles.css`, & replace the current blue & yellow with 2 new colors of your choice.

Now that the page is mainly built, you have to populate the "Dashboard" section with 3 new elements.
![](https://raw.githubusercontent.com/01-edu/public/master/subjects/carbon-copy/dashboard-template.jpg)

Those 3 cards have to display respectively one information with:

- a title
- a subtitle
- a text paragraph

For this phase, you'll have to make the whole HTML & CSS by yourself.

### Phase 3: JS interactions _(mandatory & optional)_

If you made it until here pretty fast, now the fun will begin! You're going to add a bunch of JS interactions to make elements appear / disappear / change in the HTML & CSS by linking a JS [script](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script) to your `index.html`.

- Change the order of the pictures when clicking on the pictures' section, [toggling](https://css-tricks.com/snippets/javascript/the-classlist-api/) the pre-defined class `row-reverse` from the CSS file `styles.css`.

![](https://raw.githubusercontent.com/01-edu/public/master/subjects/carbon-copy/images-order.gif)

- Option 1: In the Contact section, when clicking on the "Introduce yourself" button, get the text typed in the `input` and display it in the middle of the following sentence: "Nice to meet you _[put here the input data]_ 👋! Thanks for introducing yourself." Also, the `<p>`, `<input>` & `<button>` elements have to disappear after the button has been clicked.

![](https://raw.githubusercontent.com/01-edu/public/master/subjects/carbon-copy/contact-input.gif)

- Option 2: When clicking on a card, open a modal window that will show the whole article ; the modal will be closed either when clicking on a "Close" button, or when the "Escape" key is pressed.

![](https://raw.githubusercontent.com/01-edu/public/master/subjects/carbon-copy/modale.gif)

- Option 3: In the modal article, create a widget that allows to change the text alignment ; on click on `left` or `center` buttons, the layout changes to the chosen justification, and the selected option's `font-weight` becomes `bold` whereas the other becomes `light`.

![](https://raw.githubusercontent.com/01-edu/public/master/subjects/carbon-copy/text-alignment.gif)

- Warrior option: set the `header` text content with a random quote every time the page is loaded, and then every 10 seconds. You can use this marvelous [Chuck Norris API](https://api.chucknorris.io/) to [fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch) his most inspiring sayings and display them in your own page!

![](https://raw.githubusercontent.com/01-edu/public/master/subjects/carbon-copy/fetch-quote.gif)

### Notions

- [HTML tags](https://developer.mozilla.org/en-US/docs/Web/HTML/Element)
- [CSS basics](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics)
- [Css flexbox layout](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout/Basic_Concepts_of_Flexbox) & [Flexbox froggy](https://flexboxfroggy.com/) are always useful
- [Link a JS script](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script)
- [`classList` / `toggle`](https://css-tricks.com/snippets/javascript/the-classlist-api/)
- [`getElementById`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById) or [`querySelector`](https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector)
- [`textContent`](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent)
- [`style`](https://developer.mozilla.org/en-US/docs/Web/API/ElementCSSInlineStyle/style)
- [`addEventListener`](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener): [`click` event](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event) / [`keydown` event](https://developer.mozilla.org/en-US/docs/Web/API/Element/keydown_event) / [`load` event](https://developer.mozilla.org/en-US/docs/Web/API/Window/load_event)
- [`fetch`](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch)
- [`setInterval`](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/setInterval)

### Files

Download [`carbon-copy.zip`](https://assets.01-edu.org/carbon-copy) to have at your disposal the following files:

- the CSS file `styles.css` containing the pre-styled elements
- the `images` folder containing the images to display in the webpage
- the `assets` folder containing the templates & wireframes images of the webpage
