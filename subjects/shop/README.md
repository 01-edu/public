### Shop

### Objectives

In this project, a website will be given to you. This website is still in the creation process, so not all features are available. The purpose of the project is for you to implement some of those missing features.

You should also know that the website was created using Rails, a web application development framework written in Ruby. You can learn more about this framework in its own [website](https://guides.rubyonrails.org/getting_started.html).

You will have to create an e-Commerce shop application with user authentication where you can sell any type of products.
As mentioned, some of the process is already done. So first of all, let's see what you will have to implement.

### Instructions

In this website the goal is for you to fix it so that it will be possible to create users. Additionally, those users will be able to create ads to sell products in the website.

To create an ad you will have to go to the `Sell` button and create the ad. You will be able to give a title to your ad, a price, a model and a description. You also must select a brand, a color and a condition for your product. Moreover, you will be able to choose images from your folder in app/assets/images.

First you will have to implement a controller called `registrations_controller` that can be found incomplete in app/controllers. You will need [Devise](https://github.com/heartcombo/devise) which is a flexible authentication solution for Rails based on Warden.

- You will have to create two definitions in this controller:

The first one is `sign_up_params` where you will have to allow the necessary params to sign up, which are:

        - name
        - email
        - password
        - password_confirmation

The second one is`account_update_params` where you will also need to allow the necessary params to edit or update an account, which are:

        - name
        - email
        - password
        - password_confirmation
        - current_password

- You will also need to complete the `products_helper.rb` located in app/helpers. This helper will be responsible for showing who is selling the products and who is able to edit and delete the ads (in this case the creator of the ad). So you will have to find a way to show the name of the seller in each product.

- You have to find a way so that, only the user who creates an ad, can edit or also permanently delete it.

- All e-Commerce apps need some sort of Cart. So you will have to create a shopping cart for your application.

To create the cart you will have to define a new model `Cart`.
The cart must work like any other e-commerce shopping cart.

- You must be able to add one or more items to the cart.
- A total must be shown, which is the sum of all the products in the cart. This total must be updated every time you add or remove an item from the cart.
- The cart must have an `Empty Cart` button, which clears the cart and takes you back to the home page.
- It must be possible to remove items individually from the cart as well, without emptying the whole cart.
- The message "Added to your cart" and "Removed from your cart" must be shown when you add or remove something from the cart. These messages must disappear from the screen after a short time.
- You must add a shopping cart icon in the home page, which shows the number of items in the cart, which updates each time you add a new product.

- You will have to create a `concern` in models/concerns/ called `current_cart.rb`. A `concern` is similar to a helper which you can include in controllers throughout the app. This `concern` must allow you to add stuff to the cart when you are not signed in, and then when you sign in to an account, the cart will have the products that you added before you signed in.

#### Ruby on Rails

Ruby is a programming language similar to Python and Perl. It is dynamically typed, interpreted, and can be modified at runtime (such as adding new methods to classes). It has many shortcuts that makes it very clean, methods are rarely over 10 lines. It has good RegEx support and works well for shell scripting.

Rails is a gem, or a Ruby library. Rails helps make web applications, providing classes for saving to the database, handling URLs and displaying html (along with a webserver, maintenance tasks, and much more).

This project was created using:

- Ruby 3.0.0
- Rails 6.1.3

If you decide to install other versions you will need to update your gems according to the version you are using. If you go with these versions, you already have the necessary gems in your gemfile.

Now, that you know what the project is about, we will explain to you the file system architecture.

When starting a new Rails project (with the command `rails new name_of_project`).
This will create a Rails application in a directory with the given name and then you will need to install the gem dependencies using the command `bundle install`.

You can learn more about the purpose of each directory that is created [here](https://www.tutorialspoint.com/ruby-on-rails/rails-directory-structure.htm).

Here are some commands that you may need and that were used in the project:

- `bundle install`
- `bundle update`
- `rails` -> It will give you a list of commands you can use.
- `rails s` -> runs a server in the port 3000 by default.
- `rails g scaffold Product brand:string model:string description:text condition:string finish:string title:string price:decimal --no-stylesheets --no-javascripts` -> Product will be our main model within the app. Scaffolding in Ruby on Rails refers to the auto-generation of a set of a model, views and a controller usually used for a single database table.
- `rails g controller store index` -> Creates a controller named store that will have the Index path to go to if you need.
- `rails generate uploader Image` -> This create an uploader and should give you a file in "app/uploaders/Image_uploader.rb"
- `rails g migration add_image_to_products image:string` -> This will create the migration.
- `rails db:migrate` -> Checks which missing migrations still need to be applied.

You can get the code which is already done [here](https://assets.01-edu.org/shop/shop.zip). You will notice that there is some code missing. Your task is to complete it so that the website works as explained above.

You should do this steps after unzipping:

- bundle install
- bundle update
- rails db:migrate
- rails db:seed
- rails s

### Bonus

As bonus for this project there are a couple things you could do:

- You can add a category to the product so that the user can chose between cars, clothes, computers and many more. Or you can add more fields to describe the products, for example add more brands.
- You can create a button `add to cart` so that it is possible to add a product to the cart without opening the ads.
- You can find a way to remove items from the cart one by one.
- You can implement your own display and design.
- You can add a payment method.
