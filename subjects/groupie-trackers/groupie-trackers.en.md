## groupietrackers

### Objectives

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site.

- It will be given an [API](http://api_name.org/api), that consists in four parts:
  - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity and the date of their first album.

    ```json
     {
        "id": 1,
        "image": "http://localhost:8080/api/images/queen.jpeg",
        "name": "Queen",
        "members": [
            "Freddie Mercury",
            "Brian May",
            "John Daecon",
            "Roger Meddows-Taylor",
            "Mike Grose",
            "Barry Mitchell",
            "Doug Fogie"
        ],
        "creationDate": 1970,
        "firstAlbum": "13-07-1973",
        "locations": "http://localhost:8080/api/locations/1",
        "consertDates": "http://localhost:8080/api/dates/1",
        "relations": "http://localhost:8080/api/relation/1"
    }
    ```

  - The second one, `locations`, consists in their last and/or upcoming concert locations.

    ```json
    {
        "id": 1,
        "locations": [
            "north_carolina-usa",
            "georgia-usa",
            "los_angeles-usa",
            "saitama-japan",
            "osaka-japan",
            "nagoya-japan",
            "penrose-new_zealand",
            "dunedin-new_zealand"
        ]
    }
    ```

  - The third one, `dates`, consists in their last and/or upcoming concert dates.

    ```json
     {
        "id": 1,
        "dates": [
            "*23-08-2019",
            "*22-08-2019",
            "*20-08-2019",
            "*26-01-2020",
            "*28-01-2020",
            "*30-01-2019",
            "*07-02-2020",
            "*10-02-2020"
      ]
    }
    ```

  - And the last one, `relations`, does the link between all the other parts, `artists`, `dates` and `locations`.

    ```json
     {
        "id": 1,
        "datesLocations": {
            "dunedin-new_zealand": [
                "10-02-2020"
            ],
            "georgia-usa": [
                "22-08-2019"
            ],
            "los_angeles-usa": [
                "20-08-2019"
            ],
            "nagoya-japan": [
                "30-01-2019"
            ],
            "north_carolina-usa": [
                "23-08-2019"
            ],
            "osaka-japan": [
                "28-01-2020"
            ],
            "penrose-new_zealand": [
                "07-02-2020"
            ],
            "saitama-japan": [
                "26-01-2020"
            ]
        }
    }
    ```

- This project also focuses on the creation of events and on their visualization.
  - An event-driven system consists in a system that responds to some kind of action triggered by the client, time, or any other factor.

This project will help you learn about :

- Manipulation and storage of data.
- HTML.
- Event creation and display.
  - [Event-driven system](https://medium.com/omarelgabrys-blog/event-driven-systems-cdbe5a4b3d04).
- JSON files and format.

### Instructions

- The backend must be written in **Go**.
- The code must respect the [**good practices**](https://github.com/01-edu/public/blob/master/subjects/good-practices.en.md).
- It is recommended that the code should present a **test file**.
