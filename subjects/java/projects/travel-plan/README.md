## Travel-Plan

### Objectives

The primary goal of this project is to construct a robust and scalable environment for the Travel Management System and to develop a comprehensive Admin Dashboard that serves as the control center for travel management operations.

### Instructions

This is the first part of the last project and for this part the students will be asked to build the working enviroment and impliment the admin part to ease the last part for the project.

#### 1. Environment Configuration

##### Configure a Microservices Architecture:
-    Follow industry best practices for scalability and high availability.
-    Implement multiple replicas of each microservice for load balancing and failover mechanisms.

##### Database Installation and Configuration:
-    Set up PostgreSQL and Neo4j databases.
-    Ensure databases are containerized to facilitate scalability and replication.

##### Continuous Integration/Continuous Delivery (CI/CD) Pipeline:
-   Employ Jenkins for CI/CD and unit testing.
-   Employ SonarQube for automating code quality checks

##### Use of Docker and Ansible:
-    Docker to support automated provisioning and deployment of microservices.
-    Ansible for creating playbooks to deploy all system elements consistently and scalably.

##### Outcome:
-    Achieve an automated, scalable infrastructure.
-    Prepare the groundwork for Admin Dashboard development and additional features of the Travel Management System.

#### 2. Development and Design

- Craft an Admin Dashboard enabling administrators to oversee users, travels, and payment gateways.
- the admin should be able to add, edit or delete any user, travel, payment method and all the related part correctly (database cascading!!!) 
- Ensure each travel entry includes a destination or multiple destinations, dates, duration, activities, accommodation, and transportation details.
- for the payment methods you can search about it to know what you need to support the most you have to support at least strip and paypal.
- Develop an authentication and authorization service to safeguard access and operations to the Admin Dashboard.
- Design a responsive and intuitive UI for the Admin Dashboard that shows well in diffirent screen sizes, ensuring browser compatibility at least for Mozilla and Chrome.
- Utilize any beneficial packages to enhance development efficiency, with the expectation to justify package selections during project reviews.
- every feature should have it's unit tests to ensure the the new modifications doesn't break the code. 


#### 3. Best Practices & Ecosystem Familiarity

- Adopt a collaborative development workflow using pull requests (PRs) for introducing changes and features.
- Perform thorough code reviews on each PR to maintain code integrity, security adherence, and best practice compliance.
- Integrate a CI/CD pipeline via Jenkins for seamless build, test, and deployment workflows of PRs.
- Ensure all PRs undergo a rigorous review process, securing approval before merging into the main codebase.

#### 4. Security Measures

- Implement SSL/TLS encryption for all data in transit.
- Ensure databases and services are accessible only within the internal network or via secure, authenticated endpoints.
- Use secret management tools like HashiCorp Vault to handle sensitive information like API keys and database credentials.
- Apply the principle of least privilege across all levels of the system, particularly in role-based access controls within the Admin Dashboard.
- Regularly update all components to patch known vulnerabilities and ensure compliance with security best practices.

### Bonus

- Compose clear and detailed documentation for the application and database schema.
- Incorporate Kubernetes alongside Ansible to enhance service management, orchestration, and load balancing capabilities.
- impliment integration tests and E2E tests one of them or both of them. 
- any other bonus that add a real value to the project.

