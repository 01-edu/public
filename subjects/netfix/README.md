## Netfix

### Objectives

Imagine there was a place in which you could ask someone to fix your bidet and at the same time ask for a service to paint the walls of your house. And while you are at it, home cleaning would be quite handy. Well, in this project you will help create a website in which all of this and more will be possible.

This website is still in the creation process, so not all features are available. You are going to implement some of the missing features.

You should also know that the website is being created in Django, a Python framework for web development. You can learn more about this framework in its own [website](https://www.djangoproject.com/). Yes, you will be messing around with Python, so be aware of the indentation. Also, you will need to know that this project was build in the version `v3.1.14` of Django, so if you have problems to run this project, please check your current version of Django and make the necessary changes in order to get the project to run.

### Instructions

The website should contain two types of users:

- Company: that can create new services
- Customer: that can request existing services

Both users should be able to register and login. In order to login, users should enter the email and the password to do so. However, to register, each type of user has to provide different information:

- Customer:
  - email
  - password
  - password confirmation
  - username
  - date of birth
- Company:
  - email
  - password
  - password confirmation
  - username
  - field of work

> Each user should have a unique email and username. So while registering a user should be alerted if the email and/or username has already been registered.

Every user should have his own profile page in which their information should be displayed (apart from the password, obviously). And, in case the user is a customer, all previous requested services have to be displayed as well. In case of a company profile, along with all their information, should also be present the services it provides.

The `field of work` attribute in the companies will restrict the kind of services it can provide. The `field of work` attribute should only accept these values:

- Air Conditioner
- All in One
- Carpentry
- Electricity
- Gardening
- Home Machines
- Housekeeping
- Interior Design
- Locks
- Painting
- Plumbing
- Water Heaters

A `Carpentry` company can only create carpentry services as a `Housekeeping` company can only provide housekeeping services. However, the `All in One` companies can create any kind of services. But what does a service include? A service must have:

- name
- description
- field (can have the same categories of fields as the companies, except you can not have an All in One service)
- price per hour
- date it was created (this attribute should automatically take a value when creating a service)

The website should have a page displaying the most requested services. There should also be a page showing every service in creation order (last created first) and a page for every service category, which displays the services available for that category.

Every service should have its own page, in which should be displayed the above information and also the company name that provides this service. A user should be able to check every service from that company by clicking on the name displayed and seeing the company profile.

Once a service is created by a company, every user has access to it. Only customers can request a service by providing the `address` where the service is required and the `service time` (in hours) needed for the service to be completed. Once a customer requires a service, it gets added to the list of previously requested services. In this list, for each service requested, the customer can see the service name, service field, calculated service cost, the date in which the service was requested and the company who provided the service.

You can check out this [video](https://youtu.be/GyRo3CUWQzE) to see what is expected from your project.

#### Django

Now, that you know what the project is about, we will explain you the file system architecture. A Django project is organized by a main folder of the project and different apps.

When starting a new Django project (with the command `django-admin startproject <// name of project \\>`), the main folder (that has the same name of the project) and a file called `manage.py` get created. This file is used to run various commands. The most common ones are:

- `python3 manage.py runserver` -> runs a server in the port 8000 by default, and serves your project to it.
- `python3 manage.py startapp <// name of app \\>` -> creates an app, in other words, creates a directory with most of the needed files for an app.
- `python3 manage.py makemigrations` and `python3 manage.py migrate` -> these two commands are always hand in hand with each other. We use them to declare any changes made to the Django database (usually made in the `models.py` file of each app).
- `python3 manage.py createsuperuser` -> creates a super user (admin) that can access and change information in the Django database (accessible in the url `localhost:8000/admin`).
- `python3 manage.py` -> to get a list of all the commands available.

> You may note that we use the command `python3`. You could simply use the command `python` but you must have the 3.0 version or higher installed.

Each app usually is linked to a different aspect of the project. For example, in this project, there are three apps:

- services: handles the essential features related to services (service creation, services display, service request ...)
- users: handles the essential features related to the users (user registration, user profile, user login ...)
- main: handles the information that is common in all the project (home page, navigation bar, logout page ...)

By default, when starting an app, Django creates several files that are very useful to the development of a website. The ones in which we spend more time in are:

- `models.py` -> where you can add objects and define their attributes to the Django database
- `views.py` -> where you can manipulate the backend of a web page on the project

In addition to these two files, as a rule of thumb other files are manually created:

- `urls.py` -> here you can specify the urls of your project and associate them to a view (from the `views.py` file), in order to display a web page. For example, the url `profile/` is associated to the view `views.ProfileView`)
- `forms.py` -> here you can create your forms to use in your `views.py` file. For example, you can create a default login form to use in the `LoginView` (present in the `views.py` file).

Usually we also create a folder where we can store templates (in HTML). This folder usually follows the following structure:

```sh
netfix/                     # project directory
  ⌊ users/                  # users app directory
    ⌊ __pycache__/
    ⌊ migrations/
    ⌊ templates/            # users template directory
      ⌊ users/
        ⌊ login.html
        ⌊ profile.html
        ⌊ register.html
        ⌊   ...
    ⌊   ...
    ⌊ urls.py
    ⌊ views.py
```

You will also use what it is known as `Django templates`, which is used inside of HTML files. It is a way of coding inside of HTML, allowing to go through objects passed in (for loop) or even make conditional verifications (if/else statement), amongst other cool stuff. To learn more about it, you can check the [Django templates documentation](https://docs.djangoproject.com/en/3.1/topics/templates/).

You can get the code which is already done [here](https://assets.01-edu.org/netfix/netfix.zip). You may see that there are some code missing, both in HTML and Python files. Your task is to complete it so that the website works as explained above.

A css file is already provided (in the path `netfix/static/css/style.css`), but do feel free to change it up and come up with your own design for your site. This also means that you can change the HTML already provided.

> When first trying to run the project with the `python3 manage.py runserver` command, you will notice that you will get an error. You must start there, by defining the customer model.

### Bonus

As bonus for this project there are a couple things you could do:

- add a rating system, where customers could rate the services they have requested
- add pages to the service list
- feel free to implement your own bonus
