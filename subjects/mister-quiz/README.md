## mister-quiz

### Objectives

Have you ever watched "Who Wants to Be a Millionaire?" and wonder how would you do if it was you? Well you are not going to play it this time yet. But after finishing this project, you will be able to.

You will have to continue to build a quiz game website that is already in development. The already made code is available [here](https://assets.01-edu.org/mister-quiz/mister_quiz.zip). There is missing some code that you will have to create in order to fullfil all the requests explained in the `Instructions` section.

The project is written in [PHP](https://www.php.net/) and it is using [Laravel](https://laravel.com/), one of the most (if not the most) known PHP web framework. So you have to install both PHP and Laravel on your machine (you can use which [method](https://laravel.com/docs/8.x/installation#getting-started-on-linux) you like the best). Get comfortable using dollar signs.

### Instructions

Firstly, users should be able to register and login. To register a user has to enter a username, an email, a password and a password confirmation. To login the user has to provide the email and the password. If the login credentials are not correct, an error message should appear.

In the Login page, it must be possible to navigate to the Registration page, and vice versa.

The user must have in its attributes the amount of XP, determined by the number of questions he or she answered right. Also the score for each category of questions should be stored in the format `x/y` where `x` are the correct answers and `y` are the total questions.

After logging in the user should be redirected to the homepage. The homepage should have at least three buttons present, four if they are
logged in:

- The Login button / Profile button => if the user is not logged in, the login button should be visible, otherwise the profile button must be visible.
- The Leaderboard button => shows a list of the best players.
- The Start Quiz button => button that when pressed starts a quiz.
- The Logout button => Is only shown if the user is logged in.

#### Quiz

When clicking the `Start Quiz` button on the homepage, if the user is not logged in, the user should be redirected to the Login page (instead of starting the quiz). In case the user is already logged in, the website must present the questions and the available answers for each question. The player can then choose an answer to each question and at the bottom of all the questions submit the quiz and get his results.

A question has:

- the question text
- a category. The available categories must be:
  - Art
  - History
  - Geography
  - Science
  - Sports
- XP

There are some conditions that you should have in account:

- After starting the quiz, if the user refreshes the page the same questions should pop up.
  - The same must happen if the user goes back to the homepage and start the quiz again.
- All questions should have been answered before being able to submit the quiz.
- At least one question of each category is required.

After submitting the quiz, the user should be presented to a `Results` page. This page must show:

- How many questions the user answered right and the number of total questions made
- How many questions the user answered right in each category

If a user gets a correct answer to a question, his XP is increased by the XP of the question he answered right. Also the score that the user has for the category of that question should be updated.

#### Profile

When the user visits his profile the following must be visible:

- his username
- his email
- his XP
- his rank (dependent on the amount of XP):
  - < 1500 XP => Quiz Aprentice
  - 1500 XP - 5000 XP => Average Quizer
  - 5000 XP - 10000 XP => Epic Quizer
  - \>= 10000 XP => Quiz Master
- the percentage of correct answers for each category
- number of correct answers for each category
- number of total questions answered for each category

Only the user should be able to see his own profile.

#### Leaderboard

The leaderboard page should display a table containing the top 10 players organized by XP amount. Each row of the table must have:

- username
- XP amount
- total correct answers

#### Laravel

If you already opened the project zip file provided, you might have been surprised with the amount of files and folders inside it. We know, it can be overwhelmeing.

Here are the main folders and files you will mess around with:

- `app/`
  - `Http/Controllers/` -> is the folder of every Controller in Laravel. A Controller is what controls the backend for a specific route.
  - `Models/` -> refers to the models/ classes that are used in the app.
- `database/migrations/` -> here are located the migrations that refer to the database.
- `public/` -> is the folder that applies CSS, JS, images, etc. to the website.
- `resources/views/` -> refers to the folder containing every view for the different pages of the website. You will notice that the files inside it have the extension `.blade.php`. This extension is used by Laravel to write HTML together with PHP. You can learn more about [Blade Templates](https://laravel.com/docs/5.1/blade) in the web.
- `routes/web.php` -> is the file responsible for the routes in the website.

Of course there are a lot more folders and files available that can help you achieve the idea that you have in mind for the website, but it is up to you to research them.

You can take advantage of the templates and design provided in the resource zip file provided, or you can choose to implement your own design ideas, as long as the project requirements are respected.

You may use some commands to, for example, create controllers, create migrations, apply migrations, etc. The options are endless. You will notice that in the root directory of the project there is a file `artisan`. Try running `php artisan` on the root folder of a Laravel project and you will see every command available for you.

#### Database

In order to apply the migrations in the `database/migrations/` directory, you will need to create a database. Shocker.

One of the ways to accomplish that is to work with [`XAMPP`](https://www.apachefriends.org/index.html). XAMPP is the most popular PHP development environment and unites useful tools in order to make web development easier. It pretty much uses Apache2 (an http web server host), MySQL (a database management service) and phpMyAdmin (a web MySQL administration app).

Once you have XAMPP and these three tools up and running, you can head over to [`localhost/phpmyadmin`](http://localhost/phpmyadmin/) and create a new database called mister_quiz. After that you apply the migrations and you are good to go on your quiz quest.

> [Here](https://assets.01-edu.org/mister-quiz/questions_and_answers.sql) is a sql file in which you can run in phpMyAdmin in order to have some data (questions and answers) for you to work with. You can also come up with your own questions and answers.

When first trying to run the project you will notice that you can in fact run it and go to the web page, but you will not be able to do much more. This is where you come in and save the day (or week depending on how much time you will take to solve this project).
