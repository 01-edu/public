## groupie-tracker

### Objectives

Groupie Trackers consists of receiving a given API and manipulating the data contained in it, in order to create a site, displaying the information.

- It will be given an [API](https://groupietrackers.herokuapp.com/api), that consists in four parts:

  - `artists`: Contains information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

  - `locations`: Contains their last and/or upcoming concert locations.

  - `dates`: Contains their last and/or upcoming concert dates.

  - `relation`: Links the data of all the other parts, `artists`, `dates` and `locations`.

- Given all this you should build a user friendly website where you can display the bands info through several data visualizations (examples : blocks, cards, tables, list, pages, graphics, etc). It is up to you to decide how you will display it.

- This project also focuses on the creation of events/actions and on their visualization.

  - The event/action you need to implement is a client call to the server (client-server). We can say it is a feature of your choice that needs to trigger an action. This action must communicate with the server in order to receive information, ([request-response])(https://en.wikipedia.org/wiki/Request%E2%80%93response)
  - An event consists in a system that responds to some kind of action triggered by the client, time, or any other factor.

### Instructions

- The backend must be written in **Go**.
- The site and server cannot crash at any time.
- All pages must work correctly and you must take care of any errors.
- The code must respect the [**good practices**](../good-practices/README.md).
- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

### Allowed packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed.

### Usage

- You can see an example of a RESTful API [here](https://rickandmortyapi.com/)

This project will help you learn about :

- Manipulation and storage of data.
- [JSON](https://www.json.org/json-en.html) files and format.
- HTML.
- Event creation and display.
- [Client-server](https://developer.mozilla.org/en-US/docs/Learn/Server-side/First_steps/Client-Server_overview).
