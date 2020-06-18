## Get them all

### Instructions

You've been attributed the task to find the main architect of the Tower of Pisa before he achieves his plans, avoiding us nowadays all those lame pictures of people pretending to stop it from falling.

Launch the provided HTML file in the browser to begin your investigation.
You arrive at the architects' chamber to find him, but all you have in front of you is a bunch of unknown people.
Step by step, with the little information you have, gather information and figure out by elimination who he is.

On top of the webpage, each of the four buttons fires a function which has to return an array containing 2 arrays of HTML elements (except for the last function): the targetted people, and the others eliminated at that step - the ones previously eliminated mustn't be included.

- Write the body of the `getArchitects` function, which targets the architects, all corresponding to a `<a>` tag, and eliminates all the non-architects people.

- Write the body of the `getClassical` function, which targets the architects belonging to the `classical` class, and eliminates the non-classical architects.

- Write the body of the `getActive` function, which targets the classical architects who are `active` in their class, and eliminates the non-active classical architects.

- Write the body of the `getBonannoPisano` function, which targets the architect you're looking for, whose `id` is `BonannoPisano` (return only this one as a single element, not an array), and eliminates all the remaining active classical architects.

> From now on, don't forget to [**export**](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/export) all the expected functions, so that they can be imported to be tested

### Notions

- [HTML Element](https://developer.mozilla.org/en-US/docs/Web/API/Element)
- [`getElementsByTagName()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName)
- [`getElementsByClassName()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName)
- [`getElementById()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById)
- [`querySelectorAll()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll) / [`querySelector()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector)
- ...and bit of CSS that could help with the [`:not` pseudo class](https://developer.mozilla.org/en-US/docs/Web/CSS/:not)

### Provided files

- Check the HTML file [index.html](/public/subjects/get-them-all/index.html), which includes:

  - the JS script running some code, and which will also allow to run yours
  - some data used to generate content
  - some CSS pre-styled classes: feel free to use those as they are, or modify them
