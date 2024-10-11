## Get them all

### Instructions

You've been attributed the task to find the main architect of the Tower of Pisa before he achieves his plans, avoiding us nowadays all those lame pictures of people pretending to stop it from falling.

You arrive at the architects' chamber to find him, but all you have in front of you is a bunch of unknown people.
Step by step, with the limited details you have, gather information and figure out by elimination who he is.

Launch the provided HTML file in the browser to begin your investigation.<br/>
On top of the webpage, each of the four buttons fires a function:

- Write the body of the `getArchitects` function, which returns an array containing 2 arrays of HTML elements:

  - the first array contains the architects, all corresponding to a `<a>` tag
  - the second array contains all the non-architects people

- Write the body of the `getClassical` function, which returns an array containing 2 arrays of HTML elements:

  - the first array contains the architects belonging to the `classical` class
  - the second array contains the non-classical architects

- Write the body of the `getActive` function, which returns an array containing 2 arrays of HTML elements:

  - the first array contains the classical architects who are `active` in their class
  - the second array contains the non-active classical architects

- Write the body of the `getBonannoPisano` function, which returns an array containing:
  - the HTML element of the architect you're looking for, whose `id` is `BonannoPisano`
  - an array which contains all the remaining HTML elements of active classical architects

> From now on, don't forget to [**export**](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/export) all the expected functions, so that they can be imported to be tested<br/> > `export const getArchitects = () => {...}`

### Notions

- [HTML Element](https://developer.mozilla.org/en-US/docs/Web/API/Element)
- [`getElementsByTagName()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName)
- [`getElementsByClassName()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName)
- [`getElementById()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById)
- [`querySelectorAll()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll) / [`querySelector()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector)
- ...and bit of CSS that could help with the [`:not` pseudo class](https://developer.mozilla.org/en-US/docs/Web/CSS/:not)

### Files

You only need to create & submit the JS file `get-them-all.js` ; we're providing you the following files to download (click right and save link) & test locally:

- the HTML file [get-them-all.html](./get-them-all.html) to open in the browser, which includes:

  - the JS script running some code, and which will also allow to run yours
  - some CSS pre-styled classes: feel free to use those as they are, or modify them
  - the import of the data

- the data file [get-them-all.data.js](./get-them-all.data.js) used to generate content in the HTML

- feel free to use the CSS file [get-them-all.data.css](./get-them-all.data.css) as it is or you can also modify it.
