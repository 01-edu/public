## GraphQL

### Objectives

The objective of this project is to learn graphql query language by creating profiles. It will be provided,\
by the platform, a graphql endpoint that is connected to the database. So you can query this endpoint to obtain the information you desire.\
**Note** that for security reasons some tables are private and some are public, you will only be provided with certain content.

Your profile must have at least:

- Basic user identification
- XP amount
- level
- grades
- skills
- statistics

### Instructions

> Here is a small introduction and advantages of graphQL, if you want to read more about graphQL you can visit there [site](https://graphql.org/)

GraphQL is a query language for any API endpoint and runtime. The syntax lets you specify what data you want to receive from that API.\
You must have in mind that this language is used for API endpoints and not database. Unlike SQL,\
graphql query does not go to your database directly but to your graphql API endpoint which can connect to a database or any other data source.

You are already familiar with REST, since REST is a robust methodology of creating APIs and not elastic/scalable.\
It can be at the same time painful, because it requires individual creation of each API, example: `v1/user/item/{id}`, `v1/post/item/{id}` and so on.

The main feature of graphql compared to REST is that it lets you ask for specific information. And even better then that is the nesting feature.

Lets take for instance the social network project, if the server had a Graphql API that was connected to the database we could query this API using graphql language:

```js
query {
    user {
        name
    }
}
```

This simple query will return an array with the name of the users. Imagine if you wanted the `date_of_birth`,\
you could just add this attribute to the query, example:

```js
query {
    user {
        name
        dateOfBirth
    }
}
```

Example for the nesting fetching:

```js
query {
    user(id: "13") {
        name
        dateOfBirth
        followers {
            id
            name
        }
    }
}
```

For this example we ask for the user followers attribute, requesting the followers names and ids.\
The examples above are simple examples. **Note** that for this query is required the introduction of variables (arguments)\
 so it will return just one user, the user that as the id equal to `13`.

In graphql the usage of arguments can be specified in the schema of the API endpoint. Here is the _docs_ for the graphql endpoint you are going to use, _https://[[DOMAIN]]/public/subjects/grapqhl_

---

### **profile**

You will have to create a profile system where you can see all the info of a student. The information/data is present\
 on the graphQL endpoint, where you will have to query it.

The display of the information is up to you to design, but it must include:

- Basic user identification, for example **githubLogin**
- **XP** amount
- **level**
- **grades**
- **skills**
- **statistics**

Any other information you desire to display is welcome and will be noted

Beside this information, you will have to create a search engine which returns a selection of students profiles, based on the students:

- Name
- XP
- level
- skills

For instance you can search for a student by their `githubLogin` or filter the students by the amount of `XP` or which `level`/`skill`.

---

### Usage

> Here is the list of tables that you are allowed to query, for more information you can check out the _docs_ `https://[[DOMAIN]]/public/subjects/grapqhl`:

- **User table**:

  This table will have information about the user

  | id  | githubLogin |
  | --- | ----------: |
  | 1   |     person1 |
  | 2   |     person2 |
  | 3   |     person3 |

- **Transactions table**:

  This table will give you access to XP and audits ratio

  | id  | type | amount | userId |                    attrs |                        createdAt |
  | --- | :--: | -----: | -----: | -----------------------: | -------------------------------: |
  | 1   |  xp  |    234 |      1 | `{"objectId": 3001, ...` | 2019-03-14T12:02:23.168726+00:00 |
  | 2   |  xp  |   1700 |      2 | `{"objectId": 3001, ...` | 2019-03-14T12:02:23.168726+00:00 |
  | 3   |  xp  |    175 |      3 | `{"objectId": 3001, ...` | 2019-03-14T12:02:23.168726+00:00 |

- **Progress table**:

  | id  | userId |                    attrs | bestResultId | objectId |
  | --- | :----: | -----------------------: | -----------: | -------: |
  | 1   |   1    |                     `{}` |           61 |     3001 |
  | 2   |   2    | `{"name": "memory", ...` |         NULL |      198 |
  | 3   |   3    | `{"name": "memory", ...` |        14319 |      177 |

- **Results table**:

  Both progress and result table will give you the student progression

  | id  |             createdAt |             updatedAt | grade | progressId |                        attrs |
  | --- | --------------------: | --------------------: | ----: | ---------: | ---------------------------: |
  | 1   | 2019-07-06T13:52:5... | 2019-07-06T13:52:5... |     0 |         58 |                         `{}` |
  | 2   | 2019-07-06T13:52:5... | 2019-07-06T13:52:5... |     0 |         58 | `{"errors": {"steps": ...}}` |
  | 3   | 2019-07-06T13:52:5... | 2019-07-06T13:52:5... |     1 |         58 |                         `{}` |

- **Object table**:

  This table will give you information about all objects (exercises/projects)

  | id  | name |     type | status |                     attrs | childrenAttrs |             createdAt |             updatedAt |
  | --- | ---: | -------: | -----: | ------------------------: | ------------: | --------------------: | --------------------: |
  | 1   |    0 | exercise |  draft | `{"language": "dom", ...` |          `{}` | 2019-03-14T12:02:2... | 2019-03-14T12:02:2... |
  | 2   |    0 |  project | online |  `{"language": "go", ...` |          `{}` | 2019-03-14T12:02:2... | 2019-03-14T12:02:2... |
  | 3   |    1 | exercise | online |  `{"language": "js", ...` |          `{}` | 2019-03-14T12:02:2... | 2019-03-14T12:02:2... |

This project will help you learn about:

- [Graphql](https://graphql.org/) language
- Custom search operations (`include`/`exclude`/`fuzzy`)
- Basics of human-computer interface
