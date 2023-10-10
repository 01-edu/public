#### Functional
##### Download the project from its dedicated repository, run docker to initiate the web application and try to navigate it via the browser or using Postman.
###### Does the application work properly and you are able to navigate its core functionalities ?
##### Data and Design: Tell the auditee to show a graph representation of the database and check the following
###### Are nodes and relationships effectively representing movies, users, and ratings?
###### Are the relationships between nodes well represented and have the appropriate data attributes?
##### Movie Microservice
###### Does the Movie microservice effectively handle reading and writing movie data from the database?
###### Is it capable of generating accurate movie recommendations based on user ratings?
##### User Microservice
###### Does the User microservice handle user data operations (read and write) correctly?
###### Does it store user ratings for movies as intended?
##### Rating Microservice
###### Is the Rating microservice correctly handling data operations related to movie ratings?
##### Recommendation Microservice
###### Has Spring Data Neo4j ,or any other library, been employed to facilitate the interaction between Spring Boot and the Neo4j database?
###### Is Neo4j OGM used to map Java objects to Neo4j nodes and relationships?
###### Has one of the algorithms in Neo4j Graph Algorithms Library (content-based, collaborative filtering...) been employed to implement an efficient and accurate recommendation algorithm?
###### Does the auditee understand enough the Cypher Query Language used to express complex recommendation logic?
###### Is the Recommendation microservice generating accurate movie recommendations based on user ratings and related movies?
##### Search Functionality
###### Can users search for movies by various criteria, such as title, genre, release date, etc.?
##### Movie Details
###### Are movie details, including title, release date, genre, and rating, accurately displayed on the movie details page?
##### Rating Movies
###### Does the page for rating movies work as intended?
##### Viewing Recommendations
###### Can users view movie recommendations based on their ratings?
##### Saving Movies
###### Is there a feature allowing users to save movies to a watchlist?
##### Sharing Recommendations
###### Can users easily share movie recommendations with friends?
##### User-Friendly Interface
###### Is the Neo4flix web application's interface user-friendly and intuitive?
##### Web pages
###### Do essential web pages, including login, registration, home, movie details, movie rating, and recommendations, function as expected?
#### Security
##### Authentication and authorization
###### Is OAuth 2.0 or JWT properly implemented for user authentication?
###### Is two-factor authentication (2FA) implemented for enhanced security?
##### Password Security
###### Are strong password policies, including complexity requirements, enforced?
##### Data encryption
###### Is sensitive data protected, is HTTPS used and a valid SSL certificate is provided?
#### Testing
##### Test thouroughly the application against any possible flaws or edge cases, try as much tests as you can imagine.
###### Does the application appropriately handle errors coming from unexpected input or malicious use of its functionalities?
###### Is it functionning gracefully on stress tests?