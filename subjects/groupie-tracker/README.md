## groupie-tracker

### Objectives

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

- It will be given an [API](https://groupietrackers.herokuapp.com/api), that consists in four parts:

  - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

  - The second one, `locations`, consists in their last and/or upcoming concert locations.

  - The third one, `dates`, consists in their last and/or upcoming concert dates.

  - And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

- It should be used at least the `artists` and the `relation` parts given in the API.

- Given all this you should build a user friendly website where you can display the bands info through several data visualizations (examples : blocks, cards, tables, list, pages, graphics, etc). It is up to you to decide which info you will present and how you will display it.

- This project also focuses on the creation of events and on their visualization.

  - An event consists in a system that responds to some kind of action triggered by the client, time, or any other factor.

This project will help you learn about :

- Manipulation and storage of data.
- [JSON](https://www.json.org/json-en.html) files and format.
- HTML.
- Event creation and display.
  - [Events](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Building_blocks/Events).

### Instructions

- The backend must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommended that the code should present a **test file**.

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Usage

- You can see an example of a RESTful API [here](https://rickandmortyapi.com/)
