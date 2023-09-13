## Movies recommendation engine with Neo4j

### Objectives
The aim of this project is to develop a movie recommendation engine using Neo4j, Spring Boot, Angular, microservices and Docker. The project will be divided into the following phases:

**Data modeling:** In this phase, the data model for the movie recommendation engine will be designed. This will involve creating nodes and relationships to represent the different entities in the system, such as movies, users, and ratings.

**Development of microservices:** In this phase, the microservices for the different components of the engine will be developed. These microservices will be responsible for tasks such as reading and writing data from the database, processing recommendations, and generating user interfaces.

**Front end development:** In this phase, the frontend user interface for the movie recommendation engine will be developed. This will involve using Angular to build a web application that allows users to search for movies, view recommendations, and rate movies.

**Deployment of the microservices:** In this phase, the microservices will be deployed to a Docker container. This will allow them to be easily deployed and scaled.

### Instructions
### 1. Microservices development
The microservices for the movie recommendation engine must be developed as follows:
-   **Movie microservice:** This microservice will be responsible for reading and writing data about movies from the database. It will also be responsible for generating recommendations for users based on the movies they have rated.
-   **User microservice:** This microservice will be responsible for reading and writing data about users from the database. It will also be responsible for storing the ratings that users have given to movies.
-   **Rating microservice:** This microservice will be responsible for reading and writing data about ratings from the database.
-   **Recommendation microservice:** This microservice will be responsible for generating recommendations for users based on the movies they have rated and the movies that are related to those movies.

CRUD operations must be developed for all of the microservices, also make sure to set up a secure authentication and authorization backend like `OAuth2` or `JWT`.

### 2. User interface
The web app should have a user-friendly interface that is easy to use. The following features should be included:
-   A login page where users can sign in to the application.
-   A registration page where users can create a new account.
-   A home page where users can see a list of movies.
-   A search bar where users can search for movies.
-   A page for each movie where users can see the movie's details, such as the title, release date, genre, and rating.
-   A page for rating movies.
-   A page for viewing recommendations.

**Functionnalities**
  The movie recommendation engine  must have the following features:
-   Users can search for movies by title, genre, release date, and other criteria.
-   Users can view recommendations for movies based on the movies they have rated.
-   Users can rate movies.
-   Users can filter recommendations by genre, release date, and other criteria.
-   Users can save movies to a watchlist.
-   Users can share recommendations with friends.

### 3. Security Measures

The web app should be secure and protect users' data. The following security measures should be implemented:
-   Use HTTPS to secure all communication between the client and the server.
-   Use `OAuth 2.0` or `JWT` for authentication.
-   Store user data in a secure database.
-   Use strong passwords and enforce password complexity requirements.
-   Implement two-factor authentication (2FA).

### 4. Testing
The web app should be tested to ensure that it is working properly. The following tests should be performed:
-   Functionality testing: Ensure that the web app is able to perform all of its intended functions.
-   Usability testing: Ensure that the web app is easy to use and navigate.
-   Security testing: Ensure that the web app is secure and protects users' data.

### Resources
[Build a Cypher recommendation engine](https://neo4j.com/developer/cypher/guide-build-a-recommendation-engine/)
[Neo4j Documentation](https://neo4j.com/docs/)