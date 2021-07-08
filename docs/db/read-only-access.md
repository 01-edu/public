# Read Only token

## instructions

1. **How to get read-only role?**, for applications to get read-only access they can contact an `admin` user so that he/she can add this role to that external application.

2. **How to create and associate application with read-only access?**, for this you must create an application on `gitea` then you can create your own application token. This token will grant access to your account using the `gitea` API. To create this token you must go to **user/settings/application** then **Generate New Token**, you can use this link <https://((DOMAIN))/user/settings/applications>. It should display a token.

3. **How to get the read only token?**, to get this token you must send a request to the authentication service with the application token. The authentication service can be accessed/reached by sending the request to : `https://((DOMAIN))/api/auth/apptoken?token=${appToken}`.
This route will validate the application token and build a new **JWT** that allows you to query the information needed.

The following example should help :

```js
const APPTOKEN = '<app token>' // put your application token here
const res = await fetch(
  `https://dev.01-edu.org/api/auth/apptoken?token=${APPTOKEN}`
)
const { token } = await res.json()

// you should take care of error cases

console.log('JWT : ', token)
```

4. **What can you query?**, querying information is dependent on the users role. A role is given by default to every user, in this case it should be a `read_only` role. You can see more about possible information you can query [here](db-authorization.md)

5. **How to query the information?**, to query the information you must send a request with the query you desire to use. You can see more information and instructions about this on the project [graphql](../../subjects/graphql/README.md).

The following example should help :

```js
const jwt = '<jwt>' // put your jwt here
const query = '{user(where:{id:{_eq:1}}){login}}' // what you want to query
const res = await fetch(
  'https://((DOMAIN))/api/graphql-engine/v1/graphql',
  {
    method: 'POST',
    headers: { Authorization: `Bearer ${token}` },
    body: JSON.stringify({ query }),
  }
)
const { data } = await res.json()

// you should take care of error cases...

console.log(data)
```

## Observations

Because of the nature of JWT you should renew the token often, normally this token will have a life spam of one day.
To refresh the tokens you need to do the following:

- Send a request to the authentication service with the `JWT`. The authentication service can be accessed/reached by sending the request to : `https://((DOMAIN))/api/auth/refresh?token=${jwt}`. This route will create a new token and expire the current token.

The following example should help :

```js
const JWT = '<jwt>' // put your jwt here
const res = await fetch(
  `https://dev.01-edu.org/api/auth/refresh?token=${JWT}`
)
const { token } = await res.json()

// you should take care of error cases...

console.log('JWT : ', token)
```
