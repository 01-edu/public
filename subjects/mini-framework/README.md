## mini-framework

Now that you have already used a framework of your choice, you must now implement some features on a framework of your own. That's right, you are going to create a framework.

Be aware that a framework is different from a library. When you call a method from a library, you are in control. But with a framework, the control is inverted: the framework calls you.

### Objectives

Your framework should implement:

- Abstracting the DOM
- Routing System
- State Management
- Event Handling

You will also need to make a [todoMVC](http://todomvc.com/) project using your framework.

### Instructions

You must create documentation for your framework, so that users (auditors) are able to understand and know how to use your framework without experiencing any awkwardness.

Your framework will be tested by using it, like you previously have used one, in the social network project. So the user has to be presented to a folder structure that allows him to run the app from the root of that folder structure. The user testing your framework will have to implement some simple code in order to test the features described below.

> You are not allowed to use any framework/library like `React`, `Angular`, `Vue` and similar to create your own framework.

#### Documentation

By documentation we mean, the explaining of how does the framework works and how to work with it, for example: how to create a div, how to add an event to a button, etc. A new user of your framework, after reading the documentation has to be able to use it without too much guessing work.

So for this you will have to create a [markdown](https://www.markdownguide.org/getting-started/) file, in which will have to contain:

- Explanation on the features of your framework
- Code examples and explanations on how to:
  - Create an element
  - Create an event
  - Nest elements
  - Add attributes to an element
- Explanation on why things work the way they work

#### Abstracting the DOM

You will have to implement a way to handle the DOM. The DOM can be seen as a big object, like in the example below:

```html
<html>
  <div class="nameSubm">
    <input type="text" placeholder="Insert Name" />
    <input type="submit" placeholder="Submit" />
  </div>
</html>
```

The HTML above can be written as:

```json
{
  "tag": "html",
  "attrs": {},
  "children": [
    {
      "tag": "div",
      "attrs": {
        "class": "nameSubm"
      },
      "children": [
        {
          "tag": "input",
          "attrs": {
            "type": "text",
            "placeholder": "Insert Name"
          }
        },
        {
          "tag": "input",
          "attrs": {
            "type": "submit",
            "placeholder": "Submit"
          }
        }
      ]
    }
  ]
}
```

With this in mind you can manipulate the DOM more easily in JS. And that is what you will do using a method. Here are some methods you can use:

- [Virtual DOM](https://bitsofco.de/understanding-the-virtual-dom/) - Using a second DOM with the wanted changes, to compare with the real DOM and change just what is needed
- [Data Binding](https://docs.microsoft.com/en-us/dotnet/desktop-wpf/data/data-binding-overview?redirectedfrom=MSDN) - binds together two data sources and keeps them synchronized
- [Templating](https://medium.com/@BuildMySite1/javascript-templating-what-is-templating-7ff49d97db6b) - refers to the client side data binding method implemented with the JavaScript language.

There are a lot of ways to achieve this. Above are just some examples, what matters is that the DOM must respond to certain actions of the user.

You have to take into account the events, children and attributes of each element of the DOM.

---

#### Routing System

Routing in this case refers to the synchronization of the state of the app with the URL. In other words you will have to develop a simple way to change the URL through actions of the user that will also change the state (explained in the next topic).

---

#### State Management

The state of an app can be seen as the outcome of all the actions that the user has taken since the page loaded. In other words, if a user clicks on a button to execute an action, the state should then change.\
What you will need to do is to implement a way to handle this state. Remember that multiple pages may need to interact with the same state, so you need to find a way to let the state be reachable at every time.

---

#### Event Handling

You will also have to implement a way to handle the events triggered by the user, like: scrolling, clicking, certain keybindings, etc.... Note that this new way of handling events must be different from the `addEventListener()` method that already exists.

---

#### TodoMVC

A todoMVC project consists of creating a [webpage](https://todomvc.com/examples/react/dist/) (this example is written in React) with the same elements present in the example, so we advise you to test it around first. You have to make your todoMVC project, a functional copy of the examples given in the links above, but using your framework.\
Be aware that every thing that we can't visually see has to be present too (IDs, classes, etc.).

This project will help you learn about:

- Web Development
  - JS
  - HTML
  - CSS
- Frameworks
- Documentation
- DOM
- Routing
- State of an Application
- Event Handling
