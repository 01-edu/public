## The calling

### Instructions

Congrats! You created the very first base for your being and you witnessed its
appearance on the digital world - your browser. But so far, it's a tiny seed of
the marvelous thing it could become ; be patient, there still is a bit of work.

First of all, instead of writing down what things are _(you're not writing down
on your hand the word 'hand', are you?)_ we're going to identify them
semantically with the very practical
[`id`](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/id)
attribute. \
This `id` has to be a **unique** identifier in your webpage, which will allow you
to target your element if needed. Your can compare it to your name & surname ; this
is what is identifying you to other people, and if someone's calling you by your
name, you answer.

So let's identify the 3 elements we have so far: in each section, remove the
text content from inside the tag and set it as the value of the `id` attribute
of the corresponding `section` text.

**Open the page in the browser:** \
you don't see _anything_? Don't freak out! \
Inspect the HTML that have been created with your
[browser inspector tool](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_are_browser_developer_tools),
& if you done it correctly, you should see inside the `body` the 3 sections with
the `id` attribute set in your HTML structure.

### Code examples

Set the `id` of a `div` tag to `"my-lil-div"`:

```html
<div id="my-lil-div"></div>
```

### Notions

- [`id` attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/id)
