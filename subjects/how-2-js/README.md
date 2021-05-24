## How 2 JS

### Instructions

Hello and welcome to the JS piscine, first you will have to learn
to execute javascript.

Being a special child, JS can run in different **runtime**, what you can
do with it greatly depend of your runtime.

Luckly you don't need to install anything for that since all you
need is a web browser.

> Main runtime for executing JS are: any web browser, NodeJS and Deno.

Let's make a hello world:

```bash
# first we create the javascript file
echo "console.log('Hello World')" > how-2-js.js

# To run JS in your browser you need to import it from an HTML page:
echo '<script type="module" src="how-2-js.js"></script>' > index.html

# Finally let's create a simple web server
&>/dev/null python3 -m http.server &

# Now open your browser at the specified port
xdg-open 'http://localhost:8000'
```

> `xdg-open` find your default application for the given argument
> on mac it's just `open` and it's `start` on windows

You can now open your web browser console (`ctrl`+`shift`+`i`)
and you should see your hello world.

> The console is a very handy place to test code and explore how the language
> works, don't be shy and play in it !

Great ! you are all set, if you want to re-execute your script, just refresh.

You now just have to create a repository named `{{ROOT}}`,
which will hold all your solutions for this piscine
and just add your 2 generated files to it, we will start slow for now... 🐢
