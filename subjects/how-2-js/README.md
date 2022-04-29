## How 2 JS

### Instructions

Welcome to the JS piscine 👋.

First you will have to learn to execute JavaScript.

JavaScript can run in different **runtime** environments. What you can do with JavaScript will greatly depend on the runtime. Even different web browsers (Chrome, Firefox, Safari etc) count as different runtime environments.

We'll start by running JavaScript in your browser, so you don't need to install anything to get started.

Some common JavaScript runtime environments:
- Web browsers
- [Node.js](https://nodejs.org/)
- [Deno](https://deno.land/)

Let's make a hello world. First we create the JavaScript file

```sh
$ echo "console.log('Hello World')" > how-2-js.js
```
To run JavaScript in your browser, you can import it into a HTML page:
```sh
$ echo '<script type="module" src="how-2-js.js"></script>' > index.html
```

Let's create a simple web server:
```sh
$ &>/dev/null python3 -m http.server &
```

Now open your browser at the specified port. You'll use an appropriate command for your system:

- Linux: `xdg-open`
- macOS: `open`
- Windows: `start`
```sh
xdg-open 'http://localhost:8000'
```


You can now open your web browser console. From Google Chrome, press **Command+Option+J** (Mac) or **Control+Shift+J** (Windows, Linux, ChromeOS), and you should see your hello world.

The console is a handy place to test your code.

You are all set. If you want to re-execute your script, just refresh.

Create a repository named `((ROOT))` which will hold all your solutions for this piscine. Add your 2 generated files to it.

We'll start slow for now... 🐢
