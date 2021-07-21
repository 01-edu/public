# Examples of graphql (queries)

> Note: that all queries present where can be ran in your own school. For this you must login with an admin role and access the following route <https://((DOMAIN))/graphiql>. Depending on our role the tables permissions may differ.

## Simple queries to get info

- The following query gives the basic information of a give user :

```graphql
query getUserInfo($name: String!) {
  user(where: {login: {_eq: $name}}) {
    profile
    login
    attrs
    auditRatio
    campus
    createdAt
    totalDown
    totalUp
    updatedAt
  }
}
```

query variable:

```graphql
{"name": "Joao"}
```

---

- The following query gives the record of a given user, the output would be :
  - author of the record
  - the time that the record was created and will end
  - message/reason for the band

 to see band students

```graphql
query bandUsers($userName: String!) {
  record(where: {user: {login: {_eq: $userName}}}) {
    authorLogin
    banEndAt
    createdAt
    message
  }
}
```

query variable:

```graphql
{"userName": "Joao"}
```

---

- The following query gives a list of groups given the project `name`. The output will be the `captainLogin` and all the members from that group, including the `userLogin`.

```graphql
query getGroupInfo($object: String!) {
  group(where: {object: {name: {_eq: $object}}}) {
    captainLogin
    members {
      userLogin
    }
  }
}
```

query variable:

```graphql
{"object": "ascii-art"}
```

---

- The following query get all events in a given campus, the output for this query will be the:
  - object name
  - path of the event
  - parent event (this will be the parent event of the given child event)

```graphql
query eventsByCampus($campus: String!) {
  event(where: {campus: {_eq: $campus}}) {
    object {
      name
    }
    path
    parent {
      path
      object {
        name
      }
    }
  }
}
```

query variable:

```graphql
{ "campus": "madere" }
```

---

- The following query gets the information of users that are associated to an event. For the query to work it should be given two arguments : `campus` and the `object`. The output will be the user name/login, audit ration and the event info.

```graphql
query usersEvent($campus: String!, $object: String!) {
  event_user(where: {event: {object: {name: {_eq: $object}}, _and: {campus: {_eq: $campus}}}}) {
    userAuditRatio
    userLogin
    event {
      path
      object {
        name
      }
    }
  }
}
```

query variable:

```graphql
{"campus": "pyc", "object": "Final Exam"}
```

- If we wanted to filter the users that where registered to a type of object in the event, we would just need to filter the object by `type` instead of `name`. Should look something like this:

```graphql
query usersEvent($campus: String!, $objectType: String!) {
  event_user(where: {event: {object: {type: {_eq: $objectType}}, _and: {campus: {_eq: $campus}}}}) {
    userAuditRatio
    userLogin
    event {
      path
      object {
        name
      }
    }
  }
}
```

query variable:

```graphql
{"campus": "pyc", "objectType": "raid"}
```

---

- The following query gets the current `XP` of a given user, the output should be the `amount` that this use should have

```graphql
query xp($userName: String!) {
  xp(where: {user: {login: {_eq: $userName}}}) {
    user {
      login
    }
    amount
  }
}
```

query variable:

```graphql
{"userName": "Joao"}
```

- If we wanted to filter all `XP` gain from a user from an given object, it would look something like this:

```graphql
query eventXp($userName: String!, $objectName: String!) {
  xp_by_path(where: {user: {login: {_eq: $userName}}, _and: {object: {name: {_eq: $objectName}}}}) {
    path
    user {
      login
    }
    event {
      path
    }
    object {
      name
    }
    amount
  }
}
```

query variable:

```graphql
{"userName": "Joao", "objectName": "ascii-art"}
```

This query will return the amount of `XP` gain from each of the events.

---

- The following query gives information relevant to the onboarding games.

```graphql
query getGameInfo($name: String!) {
  toad_result_view(where: {user: {login: {_eq: $name}}}) {
    games
  }
}
```

query variables:

```graphql
{"name": "Qazaqbala"}
```
