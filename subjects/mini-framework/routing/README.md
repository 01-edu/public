## Routing example

The example consists in switching between a **'Home Page'** to a **'Active Page'** by clicking on a **'Click Me'** hyperlink.

You will need to two or three files depending if you include the script on the html files or not. All files should be in the same folder.

index.html:

```html
<!DOCTYPE html>
<html class="no-js">
  <head>
    <meta charset="utf-8" />
    <title></title>
  </head>
  <body>
    <h1>Home Page</h1>
    <a href="./active">Click here</a>
    <script src="./script.js"></script>
  </body>
</html>
```

active.html:

```html
<!DOCTYPE html>
<html class="no-js">
  <head>
    <meta charset="utf-8" />
    <title>Active</title>
  </head>
  <body>
    <h1>Active Page</h1>
    <a href="/">Click Me</a>
    <script src="script.js" async defer></script>
  </body>
</html>
```

script.js:

```js
const readTextFile = (file, allText = '') => {
  let rawFile = new XMLHttpRequest()
  rawFile.open('GET', file, false)
  rawFile.onreadystatechange = () => {
    if (rawFile.readyState === 4) {
      if (rawFile.status === 200 || rawFile.status == 0) {
        allText = rawFile.responseText
      }
    }
  }
  rawFile.send(null)
  return allText
}

document.body.innerHTML =
  location.pathname === '/'
    ? readTextFile('./index.html')
    : readTextFile('./active.html')
```
