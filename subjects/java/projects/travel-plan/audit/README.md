
#### Comprehension

##### Ask the student to describe the microservices architecture implemented.

###### Did the architecture effectively split functionalities into different microservices based on key factors like scalability and domain separation?

##### Ask the student to explain one of the Ansible playbook

###### Did he/she clearly explain all the Ansible playbook

##### Discuss the CI/CD pipeline setup.

###### Is there unit tests for each functionality and are the test run for each PR?

###### Is the SonarQube report free from any error or warning that can break the CI/CD Process?

##### Detail the security measures implemented.

###### Were comprehensive security measures like SSL/TLS, secret management, and the principle of least privilege correctly implemented?

##### Ask the student to explain the database schema for PostgreSQL and Neo4j.

###### Did the data structure in PostgreSQL and Neo4j effectively support the application's requirements?

#### Functional

##### Verify the execution of Ansible playbooks.

###### Did the Ansible playbooks execute without errors and configure the environment as intended?

###### Were the playbooks able to handle re-running scenarios without causing disruptions or inconsistencies?

##### Test each microservice API.

###### Are all microservice APIs role-based and don't work with an admin role-based user authenticated except the login API?

###### Are all microservice APIs functioning correctly and handling errors effectively?

##### Verify Admin Dashboard functionality.

###### Did the Admin Dashboard correctly manage users, travels, and payment gateways?

##### Test Authentication and Authorization.

###### Was the authentication service robust and did the role-based access control function correctly?

##### Ask the student to Simulate load on microservices.

###### Did the microservices demonstrate effective load balancing and failover under heavy traffic?

##### Verify Docker and Ansible setup.

###### Were Docker containers and Ansible playbooks set up correctly and functionally?

##### Validate CI/CD pipeline and code quality.

###### Did the CI/CD pipeline function correctly for build, test, and deployment processes, and were code quality standards maintained?

##### Assess code review and best practices.

###### Is the code consistence and well structured, also do they follow naming conventions in recent pull requests?

###### Are all the used libraries supported and not deprecated and not having any deprecated methods in use, in recent pull requests?

###### Are the security vulnerabilities found by SonarQube resolved in the pull requests?

#### Bonus

##### Documentation Quality

###### +Did the students provide clear documentation about the application and the database?

##### Kubernetes Incorporation

###### +Did the students Incorporate Kubernetes alongside Ansible to enhance service management, orchestration, and load-balancing capabilities?

##### Additional Bonuses

###### +Did the student add any valuable bonuses and it works fine without any error 