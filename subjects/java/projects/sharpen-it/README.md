## Sharpen-It

### Objectives

In this project, your task is to finalize the provided environment for your project, applying coding best practices to enhance clarity and cleanliness.

### Instructions

You have to download this [Environment](./env.zip), then you will find two folders, one for Jenkins and another for SonarQube. Your goal is to finish the given infrastructure and make the environment functional and compliant with the below requirements.

#### Jenkins

- Build and run the Jenkins Docker image.

- Install missing plugins.

- Set up credentials to access your repository.

- Fix the provided pipeline to build and test the back-end and front-end

- Complete and rectify the configuration for deployment.

- Finalize the configuration for the rollback strategy.

- Configure your repository to trigger the CI/CD pipeline for every approved pull request (PR).

- Configure Jenkins to automatically merge a new PR if the build is successful during deployment.

> 💡 Feel free to use GitHub or GitLab for environment configuration.

#### SonarQube

- Run the provided Docker Compose for SonarQube.

- Complete the configuration to integrate SonarQube with your repository.

- Ensure SonarQube runs regularly for continuous monitoring of code quality and security.

- Implement a code review and approval process to prevent Jenkins CI/CD pipeline execution in case of reported issues.

> 💡 Use the provided `sonar.properties` file or configure SonarQube through the web interface.

### Testing

Your CI/CD setup will be assessed based on:

- Successful and automated fetching of the latest code changes.

- Effective implementation of automated tests and proper responses to their outcomes.

- Sound deployment strategies, ensuring smooth transitions of new versions into live environments.

- Immediacy and accuracy of build and deployment.

### Resources

- [Jenkins Official Documentation](https://www.jenkins.io/doc/)
- [JUnit Documentation](https://junit.org/junit5/docs/current/user-guide/)
- [Jasmine/Karma Testing for Angular](https://angular.io/guide/testing)
- [SonarQube Official Documentation](https://docs.sonarqube.org/latest/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
