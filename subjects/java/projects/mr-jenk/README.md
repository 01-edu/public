## MR-Jenk

### Objectives

Within this module, you will set up a Continuous Integration (CI) and Continuous Deployment (CD) pipeline using Jenkins. This pipeline will automatically build, test, and deploy your e-commerce microservices project.

### Instructions

#### 1. Setting Up Jenkins

- Download, install, and configure Jenkins.
  - **Hint**: Use Jenkins official documentation or Docker to set up Jenkins.
- Set up build agents if necessary.

#### 2. Create a CI/CD Pipeline for Your E-commerce Platform

- Create a Jenkins Job that fetches the source code from your Git repository (e.g., GitHub).
- Set up build triggers to initiate a build whenever there's a new commit.
  
#### 3. Automated Testing

- Integrate automated tests within your pipeline.
  - **Hint**: Use JUnit for backend testing and Jasmine/Karma for Angular frontend testing.
- Ensure that the pipeline fails when a test fails.

#### 4. Deployment

- Automatically deploy your application to a server or platform of your choice after successful builds. Consider platforms like AWS, Heroku, or a local server.
- Implement a rollback strategy in case a deployment fails.

#### 5. Notifications

- Set up email or Slack notifications to inform team members of build status, whether it's a success or a failure.

### Bonus

- **Parameterized Builds**: Allow certain parameters to be customizable for each build run. For example, choose different deployment environments.
- **Distributed Builds**: Use multiple build agents to carry out builds in parallel or to build for different platforms or environments.

### Testing

Your CI/CD setup will be assessed on:

- Successful and automated fetching of the latest code changes.
- Effective implementation of automated tests and correct response to their outcomes.
- Proper deployment strategies, ensuring new versions are smoothly transitioned into live environments.
- The immediacy and accuracy of build and deployment notifications.
- (Bonus) The correct and innovative use of parameterized and distributed builds.

### Resources
[Jenkins Official Documentation](https://www.jenkins.io/doc/)
[JUnit Documentation](https://junit.org/junit5/docs/current/user-guide/)
[Jasmine/Karma Testing for Angular](https://angular.io/guide/testing)
