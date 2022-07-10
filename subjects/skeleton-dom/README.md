## Skeleton

### Welcome

Welcome to the `world-wide-what` quest!
In this new digital world you're gonna discover, it is possible to create beings with some lines of code. Yes, it is.
Like a modern Dr Frankenstein, you're about to dive into the `world-wide-what` and give birth to a new entity on the virtual space of your browser.
During this quest, step by step, you are going to shape it.

But before coding anything, make sure you checked the videos of the playlist; you need to have a server running on your computer.
For the whole quest, the principle is to iterate over your code: when you finish an exercise, copy-paste your code to use it in the next one.

### Resources

We provide you with some content to get started smoothly, check it out!

- Video [Basic set up of an HTML page](https://www.youtube.com/watch?v=QtKoO7tT-Gg&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=1)
- Video [Different HTML tags overview](https://www.youtube.com/watch?v=Al-Jzpib8VY&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=2)

Those videos are accompanying you step by step in each exercise, but if you want to check right away all the notions covered in the quest, you can watch the whole playlist [Web - HTML, CSS & DOM JS](https://www.youtube.com/playlist?list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF).

### Instructions

Ready? Let's code!

The first step to achieve in your quest is to conceive your being ; for that, 2 things have to be done:

- set some parameters (`<head>` tag)
- define the structure, or skeleton (`<body>` tag)

#### Parameters

To create any project, some things need to be declared in the HTML file - those are not visible elements in your page, but things cannot work without them.
Set your page with [`<head>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head) tags, and also put a [`<title>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title) to name the entity you're going to create.

#### Structure

Now you have to define the skeleton of your future creation.
Remember this as a general rule for any further project you will start: a good way of setting up a project is to think about all the elements needed, organize and divide them in blocks.

Let's define the first level of elements that your entity will be made of ; we're going to split it into 3 main chunks: the face, the upper body, and the lower body.

Inside the `<body>` tag of your HTML file, create 3 divisions using `<section>` tags, putting the following text content inside for each: `face`, `upper-body`, `lower-body`.

If you open you HTML file in the browser, you should see those 3 texts appear on the screen.

### Code examples

Create a `div` tag with `hello` as text content inside the `body`:

```html
<html>
  <body>
    <div>hello</div>
  </body>
</html>
```

#### How to submit

This exercise must be submited as a JS file, thankfully, javascript allows you to write `HTML` !

Here is how you would submit the `HTML` sample from above in the editor:

```js
document.documentElement.innerHTML = `
<body>
  <div>hello</div>
</body>
`
```

> You do only need to write what's inside the `<html>` tag excluding the `<html>` tag itself

### Notions

- [`html` tag](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html)
- [`head` tag](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head)
- [`title` tag](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title)
- [`body` tag](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body)
- [`section` tag](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section)
