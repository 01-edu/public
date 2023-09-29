#### Functional

##### Download the project and trigger a Jenkins build. Observe if the pipeline runs as expected.

###### Does the pipeline initiate and run successfully from start to finish?

##### Trigger some intentional build errors and observe Jenkins' response.

###### Does Jenkins respond appropriately to build errors?

##### Examine the automated testing step.

###### Are tests run automatically during the pipeline execution? Does the pipeline halt on test failure?

##### Make a minor change in the source code, commit, and push. Observe if the pipeline is triggered automatically.

###### Does a new commit and push automatically trigger the Jenkins pipeline?

##### Check the deployment process.

###### Is the application deployed automatically after a successful build? Is there a rollback strategy in place?

#### Security

##### Examine the permissions on the Jenkins dashboard.

###### Are permissions set appropriately to prevent unauthorized access or changes?

##### Review how sensitive data (like API keys, passwords) is managed in Jenkins.

###### Is sensitive data secured using Jenkins secrets or environment variables?

#### Code Quality and Standards

##### Examine the Jenkinsfile or the build configuration.

###### Is the code/script well-organized and understandable? Are there any best practices being ignored?

##### Look into the test report formats and outputs.

###### Are test reports clear, comprehensive, and stored for future reference?

##### Check for notifications setup.

###### Are notifications triggered on build and deployment events? Are they informative?

#### Bonus

##### Examine if parameterized builds are implemented.

###### +Are there options for customizing the build run with different parameters?

##### Examine the distributed builds (if implemented).

###### +Are multiple agents utilized effectively for distributed builds?