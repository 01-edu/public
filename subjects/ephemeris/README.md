## ephemeris

### Instructions

You are already feeling this change in yourself... You can know understand why people say that developers are lazy. You want to do more, but doing less. That's normal.
Let's try to enhance your own routine, to begin.
You want to save the world, to change it. But to do so, you have to know it. Every morning you do the same researches to stay aware and prepared for your new day. This project will help you saving this research time. We will automate it, in order for you to just get the informations you want by running a script. To do so, we will simulate navigation in a browser, get some information from it, and display it all gathered in the console.

> Your goal for the moment is to gain precious time, automating your own repetitive morning behaviour.
> This information can be found only in youself. It is very personnal, be creative. What seems relevant to you to begin a new day? This must be ***your own*** ephemeris, that you would like to use every day.
> Some examples we could see in other people minds: get the weather, a link to the last daily podcast that you like, a link to the last article of a media that you like, the last post of a bunch of people you follow, the last news article pushed from a social media, your horoscope, the tv program of the day, etc...

What is mandatory about the tests:
-> handle the opening and the closure of the browser (must be headless in last version)
-> go to a specific url
-> click on a specific element of the page which implies an event in this page (to like something, registrer to something, diplay something...)
-> navigate in a website (by clicking on a link)
-> navigate from a website to another
-> get the text content of an element
-> get the url of an element

In the end:
-> Your ephemeris must appear in your console.
-> The puppeteer tests must execute themselves without being seen.
-> The script has to finish and close properly.

#### Using puppeteer

In order to get your daily informations out of your favorite websites, you'll have to use `puppeteer` API.

First, create a new repository named `ephemeris` and run this command line in your terminal, in the repository:
```console
student@ubuntu:~/[[ROOT]]/ephemeris$ npm install puppeteer
```

Then, create a new file `ephemeris.mjs` in which you will use the puppeteer module:
```js
import puppeteer from 'puppeteer'

const browser = await puppeteer.launch({
  headless: false, // to be commented or removed when project is submitted
  devtools: true, // to be commented or removed when project is submitted
})
const [page] = await browser.pages()
await page.goto('https://www.google.com/')
// other actions...
await browser.close() // you can comment this during you tests
```

**Puppeteer helper - script model**

```javascript
if ((await page.$('button#myId')) !== null) {
  await page.click('button#myId')
}
await page.waitForSelector('input#myId')
await page.type(`input#myId`, 'myText')
await page.waitForSelector('div#myId')
await Promise.all([
  page.waitForNavigation(),
  page.click('div#myId'),
])
await page.waitForSelector('div#myId')
const myText = await page.$eval(
  'div#myId',
  (s) => s.textContent,
)
```

### Optional

1. Pimp your output. Here are some ideas or examples, be creative:
-> put some colors in your text
-> associate an emoji to the weather (< 10°: ❄️; > 10° && < 20°: ☀️; > 20°: 🔥)

2. Generate a `ephemeris.html` file from this ephemeris
-> Check the first raid! you can use the `style.css` and the `news` section of the `html` file.
-> Integrate your ephemeris in this first project.

3. Save a file every day in a `history` folder (with its date) and display the 10 latest files of the history in your HTML page.
