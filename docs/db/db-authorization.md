# DB authorization

Every request to Hasura executes against a set of session variables. Normally there are some main variables for the authorization context:

- `X-Hasura-User-Id`: this variable usually denotes the user executing the request.

- `X-Hasura-Role`: This variable denotes the role with which the user is executing the current request. Hasura has a built-in notion of a role, and will explicitly look for this variable to infer the role.

- `X-Hasura-campuses`: this variable contains the campus that the user currently is on.

In our database we have several roles which are simple arbitrary names.
Each role can be given a set of permissions and actions (`select`, `insert`, `update`, `delete`). That will execute against each table of the database.

Example:

To give the user with a certain role a permission to make a request, we must set a permission rule that would look something like this:

```json
{ "user_id": { "_eq": "X-Hasura-User-Id" } }
```

This is the same as saying : if the value of the `user_id` column equals the value of the session variable `X-Hasura-User-Id`, it allows this request to execute on the current row and get information on that user id.

## Roles

These are the roles presented in the database:

- `anonymous`: this role allows non logged in users to query tables (only using the `select` action):

  - `users` : columns `id`, `login`
  - `object` : columns `id`, `childAttr`, `campus`, `name`, `type`
  - `result` : columns `groupId`, `objectId`, `progressId`, `userId`, `grade`, `campus`
  - `transactions` : columns `objectId`, `userId`, `amount`, `type`
  - `progress` : columns `isDone`, `objectId`, `userId`, `id`, `grade`, `campus`

---

- `user` : this role allows the following queries:

  - `selects` action:

    - without permission rules:

      - `event_user`
      - `group`
      - `group_user`
      - `match`
      - `registration_user`
      - `event`
      - `object`
      - `object_child`
      - `object_status`
      - `object_type`
      - `registration`
      - `role`
      - `transaction_type`
      - _`event_user_view`_
      - _`object_children_view`_
      - _`registration_user_view`_
      - _`user_public_view`_

    - with permission rules:
      - `audit`, the user only can query this table if the user id from the `X-Hasura-User-Id` variable is equal to one of the `members`, or the `auditorId`.
      - `result`, the user only can query this table if the `X-Hasura-User-Id` is equal to the `userId` or one of the members from their group.
      - `transaction`, this table can be queried by users if the `X-Hasura-User-Id` is equal to the `userId`.
      - `record`, the same applies to this table and the tables below.
      - `progress`
      - `user_role`
      - `user`, the same applies to this table but with the `id`.
      - _`user_role_view`_
      - _`audit_private`_
      - _`progress_by_path_view`_
      - _`progress_view`_
      - _`xp`_
      - _`xp_by_event`_
      - _`xp_by_path`_

---

- `campus-admin` : this role allows users to query every table, but with the variables `X-Hasura-campus` checked (campus check). This means that users with this role will only be able to query information from their own campus. Example: a user in campus `madere` can only query the content associated to that campus. The following tables can be queried:

  - `select` action:
    - without permission rules:
      - `group_status`
      - `object_status`
      - `object_type`
      - `result_type`
      - `role`
      - `transaction_type`
      - `user_role`
      - _`user_public_view`_
      - _`user_role_view`_
    - with permission rules:
      - `event_user`
      - `audit`
      - `group`
      - `group_user`
      - `match`
      - `object`
      - `event`
      - `progress`
      - `record`
      - `registration`
      - `registration_user`
      - `result`
      - `transaction`
      - `user`
      - _`progress_by_path_view`_
      - _`audit_private`_
      - _`event_user_view`_
      - _`event_with_results_ready_view`_
      - _`progress_view`_
      - _`registration_user_view`_
      - _`registration_with_event_ready_view`_
      - _`toad_result_view`_
      - _`xp_by_event`_
      - _`xp_by_path`_

---

- `campus-admin-read-only` : this role allows users to query almost all tables (only using the `select` action). But with the same permission rule in the `campus-admin`. The following tables can be queried:

  - without permission rules:
    - `group_status`
    - `object_status`
    - `object_type`
    - `result_type`
    - `role`
    - `transaction_type`
    - `user_role`
    - _`user_public_view`_
    - _`user_role_view`_
  - with permission rules:
    - `event_user`
    - `audit`
    - `group`
    - `group_user`
    - `match`
    - `object`
    - `event`
    - `progress`
    - `record`
    - `registration`
    - `registration_user`
    - `result`
    - `transaction`
    - `user`
    - _`progress_by_path_view`_
    - _`audit_private`_
    - _`event_user_view`_
    - _`event_with_results_ready_view`_
    - _`progress_view`_
    - _`registration_user_view`_
    - _`registration_with_event_ready_view`_
    - _`toad_result_view`_
    - _`xp_by_event`_
    - _`xp_by_path`_

---

- `admin-read-only` : this role allows users to query all tables only using the `select` action.

---

- `admin` : this role allows users to query using any action in any table on the database.

> You can see more about each role by going to the [graphiql](https://((DOMAIN))/graphiql) in the docs section. Note that you must be logged in with the user role you desire to see. For the role `anonymous` you do not need to be logged in. If you want to see the possible tables that can be queried by an admin, you must login with an admin, and so on...

A role is given by default to every user, if an user has more roles, the highest would be taken by default when login.

This table can describe the permissions for each table of the database and each user role.

✅ : with permission rules\
🟩 : without permission rules\
❌ : not allowed\
S : select\
U : update\
I : insert\
D : delete

| tables                             | `anonymous`                     | `user`                          | `campus-admin`                  | `campus-admin-read-only`        | `admin`                         | `admin-read-only`               |
| :--------------------------------- | :------------------------------ | :------------------------------ | :------------------------------ | :------------------------------ | :------------------------------ | ------------------------------- |
|                                    | S&emsp; U&emsp; I&emsp; D&emsp; | S&emsp; U&emsp; I&emsp; D&emsp; | S&emsp; U&emsp; I&emsp; D&emsp; | S&emsp; U&emsp; I&emsp; D&emsp; | S&emsp; U&emsp; I&emsp; D&emsp; | S&emsp; U&emsp; I&emsp; D&emsp; |
| event_user                         | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| audit                              | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| group                              | ❌ ❌ ❌ ❌                     | 🟩 ✅ ✅ ✅                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| group_user                         | ❌ ❌ ❌ ❌                     | 🟩 ✅ ✅ ✅                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| match                              | ❌ ❌ ❌ ❌                     | 🟩 ✅ ✅ ✅                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| progress                           | 🟩 ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| record                             | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| registration_user                  | ❌ ❌ ❌ ❌                     | 🟩 ❌ ✅ ✅                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| result                             | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| transaction                        | 🟩 ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| user                               | 🟩 ❌ ❌ ❌                     | ✅ 🟩 ❌ ❌                     | ✅ ✅ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| audit_expired_view                 | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| audit_private                      | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| discordToken                       | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| event                              | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| event_user_view                    | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| object                             | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| event_with_results_ready_view      | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| group_status                       | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| job                                | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| match_to_remove_view               | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ✅ ✅ ✅                     | ❌ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| object_child                       | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ❌ ✅ ✅ ✅                     | ❌ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| object_children_view               | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| object_status                      | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| object_type                        | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| progress_by_path_view              | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| progress_view                      | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| registration                       | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| registration_user_view             | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| registration_with_event_ready_view | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| result_type                        | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| role                               | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| toad_result_view                   | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| token                              | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| token_status                       | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| transaction_type                   | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| user_public_view                   | ❌ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| user_role                          | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| user_roles_view                    | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | 🟩 ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| xp                                 | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ❌ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| xp_by_event                        | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
| xp_by_path                         | ❌ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ❌ ❌ ❌                     | ✅ ✅ ✅ ✅                     | ✅ ❌ ❌ ❌                     |
