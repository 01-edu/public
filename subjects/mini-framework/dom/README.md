## DOM example

The example consists in a red box moving by clicking on a button.

```html
<!DOCTYPE html>
<html class="no-js">
  <head>
    <meta charset="utf-8" />
    <title>SSS</title>
  </head>
  <body>
    <div id="app">
      <div
        id="box"
        style="
          width: 100px;
          height: 100px;
          background: red;
          transition: 0.3s;
          transform: translateX(0px);
        "
      ></div>
      <button id="button" class>Click Me</button>
    </div>
  </body>
  <script>
    document.getElementById('button').addEventListener('click', (e) => {
      let box = document.getElementById('box')
      let actualTranslate = Number(box.style.transform.match(/\d+/g)[0]) + 50
      box.style.transform = `translateX(${actualTranslate}px)`
      console.log(box.style.transform)
    })
  </script>
</html>
```
