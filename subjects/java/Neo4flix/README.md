## Movies recommendation engine with Neo4j

### Objectives
The aim of this project is to develop a movie recommendation engine using Neo4j, Spring Boot, Angular, microservices and Docker. The project will be divided into the following phases:

**Data modeling:** In this phase, the data model for the movie recommendation engine will be designed. This will involve creating nodes and relationships to represent the different entities in the system, such as movies, users, and ratings.

**Development of microservices:** In this phase, the microservices for the different components of the engine will be developed. These microservices will be responsible for tasks such as reading and writing data from the database, processing recommendations, and generating user interfaces.

**Front end development:** In this phase, the frontend user interface for the movie recommendation engine will be developed. This will involve using Angular to build a web application that allows users to search for movies, view recommendations, and rate movies.

**Deployment of the microservices:** In this phase, the microservices will be deployed to a Docker container. This will allow them to be easily deployed and scaled.

### Instructions

#### 1. Features and requirements
  The movie recommendation engine  must have the following features:
-   Users can search for movies by title, genre, release date, and other criteria.
-   Users can view recommendations for movies based on the movies they have rated.
-   Users can rate movies.
-   Users can filter recommendations by genre, release date, and other criteria.
-   Users can save movies to a watchlist.
-   Users can share recommendations with friends.

The microservices for the movie recommendation engine must be developed as follows:

-   **Movie microservice:** This microservice will be responsible for reading and writing data about movies from the database. It will also be responsible for generating recommendations for users based on the movies they have rated.
-   **User microservice:** This microservice will be responsible for reading and writing data about users from the database. It will also be responsible for storing the ratings that users have given to movies.
-   **Rating microservice:** This microservice will be responsible for reading and writing data about ratings from the database.
-   **Recommendation microservice:** This microservice will be responsible for generating recommendations for users based on the movies they have rated and the movies that are related to those movies.
#### 2. Testing (WIP)
