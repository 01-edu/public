## GraphQL

### Objectives

The objective of this project is to learn the [GraphQL](https://graphql.org/) query language, by creating your own profile page.

You'll use the GraphQL endpoint which is provided by the platform (`https://((DOMAIN))/api/graphql-engine/v1/graphql`). You'll be able to query your own data to populate your profile page.

So that you can access your data, you'll need to create a login page.

Your profile must display three pieces of information which you may choose. For example:

- Basic user identification
- XP amount
- level
- grades
- audits
- skills

Beside those sections it will have a mandatory section for the generation of statistic graphs.

### Instructions

You will have to create a profile UI where you can see your own school information. This information/data is present on the GraphQL endpoint, where you will have to query it.

The UI design is up to you. However, have in mind the principles of a [good UI](../good-practices/README.md).\
The UI will have a statistic section where you can generate graphs to see more about your journey and achievements on the school. This graphs must be done using [SVG](https://developer.mozilla.org/en-US/docs/Web/SVG). You will have to do at least **two different statistic graphs** for the data given.

Using SVG you can create several [types of graphs](https://www.tutorialspoint.com/svg/graph.htm) including interactive graphs and animated graph. It will be up to you to decide what type of graphs you are going to do.

Here are some possible combinations for the creation of the graphs:

- XP earned in a time period (progress over time)
- Levels over time
- XP earned by project
- Audit ratio
- Projects `PASS` and `FAIL` ratio
- Piscine (JS/Go) stats
  - `PASS` and `FAIL` ratio
  - Attempts for each exercise

Any other information you desire to display is welcome and will be noted.

### Login Page

You'll need a [JWT](https://jwt.io/introduction) to access the GraphQL API. A JWT can be obtained from the `signin` endpoint (`https://((DOMAIN))/api/auth/signin`).

You may make a `POST` request to the `signin` endpoint, and supply your credentials using `Basic` authentication, with base64 encoding.

Your login page must function with both:

- username:password
- email:password

If the credentials are invalid, an appropriate error message must be displayed.

You must provide a method to log out.

When making GraphQL queries, you'll supply the JWT using `Bearer` authentication.

### Hosting

Besides the creation of your own profile you will have to host it. There are several places where you can host your profile,\
for example: [github-pages](https://pages.github.com/), [netlify](https://www.netlify.com/) and so on. You are free to choose the hosting place.

Here are a selection of interesting tables and columns which are exposed via GraphQL:

- **User table**:

  This table will have information about the user

  | id  |   login |
  | --- | ------: |
  | 1   | person1 |
  | 2   | person2 |
  | 3   | person3 |

- **Transactions table**:

  This table will give you access to XP and audits ratio

  | id  | type | amount | objectId | userId |                        createdAt |                   path |
  | --- | :--: | -----: | -------: | -----: | -------------------------------: | ---------------------: |
  | 1   |  xp  |    234 |       42 |      1 | 2021-07-26T13:04:02.301092+00:00 | /madere/div-01/graphql |
  | 2   |  xp  |   1700 |        2 |      2 | 2021-07-26T13:04:02.301092+00:00 | /madere/div-01/graphql |
  | 3   |  xp  |    175 |       64 |      3 | 2021-07-26T13:04:02.301092+00:00 | /madere/div-01/graphql |

- **Progress table**:

  | id  | userId | objectId | grade |                        createdAt |                        updatedAt |                        path |
  | --- | :----: | -------: | ----: | -------------------------------: | -------------------------------: | --------------------------: |
  | 1   |   1    |     3001 |     1 | 2021-07-26T13:04:02.301092+00:00 | 2021-07-26T13:04:02.301092+00:00 | /madere/piscine-go/quest-01 |
  | 2   |   2    |      198 |     0 | 2021-07-26T13:04:02.301092+00:00 | 2021-07-26T13:04:02.301092+00:00 | /madere/piscine-go/quest-01 |
  | 3   |   3    |      177 |     1 | 2021-07-26T13:04:02.301092+00:00 | 2021-07-26T13:04:02.301092+00:00 | /madere/piscine-go/quest-01 |

- **Results table**:

  Both progress and result table will give you the student progression

  | id  | objectId | userId | grade | progressId | type |                        createdAt |                        updatedAt |                   path |
  | --- | -------: | -----: | ----: | ---------: | ---: | -------------------------------: | -------------------------------: | ---------------------: |
  | 1   |        3 |      1 |     0 |         58 |      | 2021-07-26T13:04:02.301092+00:00 | 2021-07-26T13:04:02.301092+00:00 | /madere/div-01/graphql |
  | 2   |       23 |      1 |     0 |         58 |      | 2021-07-26T13:04:02.301092+00:00 | 2021-07-26T13:04:02.301092+00:00 | /madere/div-01/graphql |
  | 3   |       41 |      6 |     1 |         58 |      | 2021-07-26T13:04:02.301092+00:00 | 2021-07-26T13:04:02.301092+00:00 | /madere/div-01/graphql |

- **Object table**:

  This table will give you information about all objects (exercises/projects)

  | id  | name |     type | childrenAttrs |
  | --- | ---: | -------: | ------------: |
  | 1   |    0 | exercise |          `{}` |
  | 2   |    0 |  project |          `{}` |
  | 3   |    1 | exercise |          `{}` |

Examples:

Lets take for instance the table `user` and try to query it:

```js
{
  query {
      user {
          id
      }
  }
}
```

This simple query will return an array with the ids of the users. Imagine if you wanted the `login`,
you could just add this attribute to the query like so:

```js
{
  query {
      user {
          id
          login
      }
  }
}
```

You can use `curl` query they API:

```console
curl --location 'https://((DOMAIN))/api/graphql-engine/v1/graphql' \
--header 'Authorization: Bearer <JWT>' \
--header 'Content-Type: application/json' \
--data '{"query":"{user {id login}}","variables":{}}'
```

Here is another example of a query using the table `user`:

```js
{
  query {
    user(where: { id: { _eq: 6 }}) {
      id
      login
    }
  }
}
```

**Note** that for this query the introduction of variables (arguments) **is required**,
so it will return just one user, the user that has the `id` equal to `6`.

You can see the result using `curl`:

```console
curl --location 'https://content.01-edu.org/api/graphql-engine/v1/graphql' \
--header 'Authorization: Bearer <JWT>' \
--header 'Content-Type: application/json' \
--data '{"query":"{user(where:{id:{_eq:6}}){id login}}","variables":{}}'
```

In GraphQL, the usage of arguments are specified in the schema. You can see the available query parameters by introspecting the API.

> If you're logged in to the platform, you may access [GraphiQL](<https://((DOMAIN))/graphiql/>) to more easily explore the schema.

Example of nesting, using the result and user table :

```js
{
  result {
    id
    user {
      id
      login
    }
  }
}
```

For this example we ask for the results `id` and `user`s that are associated to the `result`, requesting the users `login`s and `id`s.

**You must use all the types of querying present above** (_normal_, _nested_ and using _arguments_), do not forget that you can use the types together or separately.

This project will help you learn about:

- [GraphQL](https://graphql.org/)
- [GraphiQL](https://github.com/graphql/graphiql)
- [Hosting](https://en.wikipedia.org/wiki/Web_hosting_service)
- [JWT](https://jwt.io)
- Authentication
- Authorization
- Basics of human-computer interface
  - UI/UX
