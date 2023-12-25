
#### Comprehension

##### Ask the student to describe the microservices architecture implemented.

###### Have you clearly defined the boundaries for each microservice based on business domains?

###### Do your microservices align directly with specific business functions?

###### Are your microservices designed to operate independently of one another?

###### Can each microservice be deployed, updated, and scaled without affecting others?

###### Is your architecture designed to support the independent scalability of each microservice?

###### Does your system maintain functionality even when one or more services fail?

###### Is there an API Gateway in your architecture to manage incoming requests?

###### Can you track and trace a request across multiple services easily?

##### Ask the student to explain one of the Ansible playbook

###### Did he/she clearly explain all the Ansible playbook

##### Discuss the CI/CD pipeline setup.

###### Are there unit tests for each functionality and are the tests running for each new PR?

###### Is the SonarQube report free from any error or warning that can break the CI/CD Process?

##### Detail the security measures implemented.

###### Were comprehensive security measures like SSL/TLS, secret management, and the principle of least privilege correctly implemented?

##### Ask the student to explain the database schema for PostgreSQL and Neo4j.

###### Did the data structure in PostgreSQL and Neo4j effectively support the application's requirements?

#### Functional

##### Verify the execution of Ansible playbooks.

###### Did the Ansible playbooks execute without errors and configure the environment as intended?

###### Were the playbooks able to handle re-running scenarios without causing disruptions or inconsistencies?

##### Verify Docker and Ansible setup.

###### Were Docker containers and Ansible playbooks set up correctly and functionally?

##### Test each microservice API.

###### Are all the microservices' APIs only accessible when logged in with an Admin profile?

##### Admin should be able to perform CRUD operations for users, travelers and payment methods. For each "entity" try to create, read, update and delete.

###### Is everything working as expected?

###### Are errors handled correctly?

##### Test Authentication and Authorization.

###### Was the authentication service robust and did the role-based access control function correctly?

##### Ask the student to Simulate load on microservices.

###### Did the microservices demonstrate effective load balancing and failover under heavy traffic?

##### Validate CI/CD pipeline and code quality.

###### Did the CI/CD pipeline function correctly for build, test, and deployment processes, and were code quality standards maintained?

##### Assess code review and best practices.

###### Is the code consistent and well-structured?

###### Are all pull requests following naming conventions such as (Camel Case, Pascal Case, ...), Consistency, clarity and descriptiveness?

###### Is the code consistence and well structured, also do they follow naming conventions in recent pull requests?

##### Check SonarQube logs in recent pull requests.

###### Is the log free of warnings about unsupported or deprecated libraries?

###### Are the security vulnerabilities found by SonarQube resolved in the pull requests?

#### Bonus

##### Documentation Quality

###### +Did the students provide clear documentation about the application and the database?

##### Kubernetes Incorporation

###### +Did the students Incorporate Kubernetes alongside Ansible to enhance service management, orchestration, and load-balancing capabilities?

##### Additional Bonuses

###### +Did the student add any valuable bonuses and it works fine without any error 