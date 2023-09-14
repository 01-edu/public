#### Functional

##### Download the project and start it according to the instructions provided by the student. Try accessing the API endpoints with Postman or similar API testing tool.

###### Does the application run and respond to API requests?

##### Check the implementation of CRUD operations for Users and Products.

###### Are all CRUD operations correctly implemented for Users and Products?

##### Try authenticating with a user. Then, try performing operations on Users and Products.

###### Does the authentication work correctly? Do operations respect the user roles?

##### Check the handling of exceptions. Try to create scenarios where the application could throw exceptions (like trying to update non-existing user).

###### Does the application handle the exceptions and return appropriate HTTP response codes and messages?

##### Check if the "GET Products" API is accessible without authentication.

###### Can you access the "GET Products" API without being authenticated?

#### Security

##### Check the implementation of the mentioned security measures (hashed passwords, input validation, protection of sensitive information, HTTPS).

###### Are all the mentioned security measures implemented?

#### Code Quality and Standards

##### Check the use of appropriate annotations in the data classes.

###### Are the annotations @Document, @Id, and @Field used correctly in the data classes?

##### Check the use of appropriate annotations in the controller classes.

###### Are the annotations @RestController, @RequestMapping, @GetMapping, @PostMapping, @PutMapping, and @DeleteMapping used correctly in the controller classes?

##### Check the use of the @Autowired annotation for dependency injection.

###### Is the @Autowired annotation used correctly for dependency injection?

##### Check the use of appropriate annotations for authentication and role-based access control.

###### Are the annotations @EnableWebSecurity, @EnableMethodSecurity, @PermitAll, @PostAuthorize and @PreAuthorize used correctly?

##### Check the use of validation annotations in the data classes.

###### Are validation annotations like @NotNull, @Size, @Email, etc., used to validate input data?

#### Bonus

##### Check if CORS policies are set appropriately.

###### +Are the CORS policies set appropriately?

##### Check if rate limiting is implemented to prevent brute force attacks.

###### +Is rate limiting implemented?

