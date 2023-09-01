#### Initial Setup & Access

##### Download the project and initiate it using docker. Try accessing the web pages and the API endpoints using a web browser and tools like Postman.

###### Does the application run seamlessly, and are you able to interact with its functionalities?

#### User & Product CRUD Operations

##### Examine the CRUD operations for Users (both client and seller) and Products.

###### Are CRUD operations for Users and Products correctly implemented, and does each user role have the appropriate access levels?

#### Authentication & Role Validation

##### Sign up as a client and a seller. Afterwards, test functionalities specific to each role.

###### Does the authentication system work, and do the operations respect the user roles (seller vs client)?

#### Media Upload & Product Association

##### As a seller, attempt to upload media for a product and ensure it adheres to the given size and type constraints.

###### Can media be uploaded effectively, and are the constraints enforced? Are products correctly associated with these media?

#### Frontend Interaction

##### Explore the sign-in/up pages, seller product management page, product listing, and media upload page for products.

###### Are all the frontend pages functioning as expected, and is the user experience intuitive?

#### Security

##### Analyze for the specified security measures like hashed passwords, input validation, protection of sensitive data, and HTTPS usage.

###### Are the mentioned security practices appropriately enforced?

#### Code Quality and Standards

##### Evaluate the backend code for appropriate usage of Spring and other related annotations.

###### Are the Spring Boot, MongoDB, and other relevant annotations used correctly throughout the application?

#### Frontend Implementation

##### Investigate the frontend codebase, especially the Angular components, services, and modules.

###### Is the Angular code structured efficiently, and are components, services, and modules effectively utilized?

#### Error Handling & Edge Cases

##### Attempt actions that might cause errors, such as registering with an existing email, uploading invalid media formats, or exceeding the media size limit.

###### Does the application gracefully handle these errors, providing useful feedback to the user?