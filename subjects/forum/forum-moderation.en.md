## forum-moderation

### Objectives

You must follow the same [principles](https://public.01-edu.org/subjects/forum/forum.en) as the first subject.

The `forum-moderation` will be based on a moderation system. It must present a moderator that, depending on the access level of a user or the forum set-up, approves posted messages before they become publicly visible.

- The filtering can be done depending on the categories of the post being sorted by irrelevant, obscene, illegal or insulting.

For this optional you should take into account all types of users that can exist in a forum and their levels.

You should implement at least 4 types of users :

#### Guests

- These are unregistered-users that can neither post, comment, like or dislike a post. They only have the permission to **see** those posts, comments, likes or dislikes.

#### Users

- These are the users that will be able to create, comment, like or dislike posts.

#### Moderators

- Moderators, as explained above, are users that have a granted access to special functions :
  - They should be able to monitor the content in the forum by deleting or reporting post to the admin
- To create a moderator the user should request an admin for that role

#### Administrators

- Users that manage the technical details required for running the forum. This user must be able to :
  - Promote or demote a normal user to, or from a moderator user.
  - Receive reports from moderators. If the admin receives a report from a moderator, he can respond to that report
  - Delete posts and comments
  - Manage the categories, by being able to creating and deleting them.

This project will help you learn about :

- Moderation System
- User access levels

### Allowed packages

- All [standard go](https://golang.org/pkg/) packages are allowed.

### Instructions

- You must handle website errors, HTTPS status.
- You must handle all sort of technical errors.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommended that the code should present a **test file**.
