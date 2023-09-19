# Artifact Management System with Nexus

## Project Description
Welcome to the **Nexus** project! This hands-on project is designed to provide you with a comprehensive learning experience in utilizing Nexus by Sonatype as a central repository for storing, managing, and deploying software artifacts. Through a series of tasks and requirements, you'll explore various aspects of artifact management, version control, Docker integration, continuous integration, and security considerations.

## Project Requirements
In this project, You will have to be working with, at least, java 11 and a version of Maven that is compatible with it.

1. **Setup Nexus Repository Manager:**
   - Obtain the most recent release of Nexus Repository Manager, and proceed to install and set it up on either a local or remote server. It is essential to configure Nexus to operate under a "nexus" user rather than the "root" user.
   - Create repositories and set them up for different artifact types such as JARs, WARs, and Docker images.

2. **Sample Web Application:**
   - Clone buy-02 project repository, configure the web application as Maven project.

3. **Artifact Publishing:**
   - Configure Maven to publish the built artifacts to the appropriate repositories in Nexus.

4. **Dependency Management:**
   - Utilize Nexus as a proxy for fetching external dependencies required for the web application.
   - Configure the project to resolve dependencies from Nexus repositories.

5. **Versioning:**
   - Implement versioning for the web application and its artifacts using Nexus capabilities.
   - Demonstrate how to retrieve and manage different versions of the artifacts.

6. **Docker Integration:**
   - Set up a Docker repository in Nexus and publish the Docker image to the repository.

7. **Continuous Integration (CI):**
   - Configure the pipeline, in addition to builds and tests, add artifact publishing whenever changes are pushed to the repository.

8. **Documentation:**
   - Create clear and detailed documentation that explains the project setup, configuration steps, and usage instructions.
   - Include screenshots and examples to guide users through the artifact management process.

9. **Bonus: Nexus Security and Access Control:**
   - Explore Nexus security features such as user authentication, role-based access control, and repository-level permissions.
   - Configure security settings to restrict access to certain artifacts or repositories.

## Evaluation
Your project will be evaluated based on the completion of all required tasks. Refer to the provided evaluation audit document for the criteria and checklist used by the evaluator.

Enjoy your journey into the world of Nexus and artifact management!
