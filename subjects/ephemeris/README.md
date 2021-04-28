## ephemeris

### Video

Check out [this little video](https://youtu.be/hnDVnKajqqU)!

### Instructions

You are already feeling this change in yourself... You can now understand why people say that developers are lazy. You want to do more, but doing less. That's normal.
To begin, let's try to enhance your own routine. Every morning you do the same researches to stay aware and be prepared for your new day. We're going to save this precious time.
To do so, we're going to create a script which automates this researches by simulating them in a browser, get the informations you want, and display them all gathered in the console.

> We are talking here about your own interests. This can be found only in yourself. What's relevant to you to begin a new day? It is very personnal, be creative!
> Some examples we could see in other people minds: get the weather, the last podcast/article of a media that you like, the last post of a someone/something you follow, your horoscope, the tv program of the day, etc.

About the browser control, you must:

- handle the opening and the closure of the browser
- navigate to a specific url
- click on an element which triggers an event in the page (display something, register to something, like something...)
- navigate in a website
- get the text content of an element
- get the url of an element

In the end:

- The whole browser control must execute itself without being seen.
- Your `ephemeris` must appear in your console.
- The script has to finish and close properly.

#### Using puppeteer

In order to get your daily informations out of your favorite websites, you'll have to use the `puppeteer` API.

First, create a new repository named `ephemeris` and run this command line in your terminal, in the repository:

```console
$ npm install puppeteer
```

Then, create a new file `ephemeris.mjs` in which you will use the puppeteer module:

```js
import puppeteer from 'puppeteer'

const browser = await puppeteer.launch({
  headless: false, // to be commented or removed when project is submitted
  devtools: true, // to be commented or removed when project is submitted
})
const [page] = await browser.pages()
// actions...
await browser.close() // you can comment this during you tests
process.exit(0) // terminate the process as succeeded (you can comment this during you tests)
```

**Puppeteer helper - script model**

```javascript
// navigate to a specific url
await page.goto('https://www.google.com/')

// wait an element in the page
await page.waitForSelector('input#myId')

// find an element in the page:
const button = await page.$('button#myId')
// check the element was found
if (button) {
  // use the element
  await button.click()
} else {
  // do something else: for example treat the exception as an error
  console.error('Button not found !')
  // and terminate the process as failed (code 1)
  process.exit(1)
}

// navigate in a website
await Promise.all([page.waitForNavigation(), page.click('div#myId')])

// get the textContent of an element
const myText = await page.$eval('div#myId', (elem) => elem.textContent)

// type something in an input
await page.type(`input#myId`, 'some text')
```

### Optional

1. Pimp your output. Here are some ideas or examples, be **creative**:

   - Put some colors in your text
   - Associate an emoji to the weather `(< 10°: ❄️; > 10° && < 20°: ☀️; > 20°: 🔥)`
   - Add randoms messages to say hello
   - Etc.

2. Generate a `ephemeris.html` file from this ephemeris

   - Check the first raid! you can use the `style.css` and the `cards` section of the `html` file as an inspiration...

3. Save a file every day in a `/history` folder (with its date) and display the 10 latest files of the history in an HTML page.

4. Integrate your `ephemeris` in the result of your first `carbon-copy` raid, on the `cards` section.
   - Integration must be dynamic: _each day, your new ephemeris should replace the old one._

### Notions

- [HTML Elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Element)
- [HTML Attributes](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes)
- [Puppeteer](https://pptr.dev/)
- [Node process: `exit`](https://nodejs.org/api/process.html#process_process_exit_code)
- [Node file system: `readdir`](https://nodejs.org/api/fs.html#fs_fspromises_readdir_path_options)
- [Node file system: `readFile`](https://nodejs.org/api/fs.html#fs_fspromises_readfile_path_options)
- [Node file system: `writeFile`](https://nodejs.org/api/fs.html#fs_fspromises_writefile_file_data_options)
- [Node path: `resolve`](https://nodejs.org/api/path.html#path_path_resolve_paths)
- [`getElementById`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById))
- [`textContent`](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent)
- [`http-server`](https://www.npmjs.com/package/http-server) or [`serve`](https://www.npmjs.com/package/serve)
