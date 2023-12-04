#### General

##### Check the Repo content:

Files that must be inside the repository:

- CI/CD pipeline configuration files, scripts, and any other required artifacts.
- An Ansible playbook and used scripts for deploying and configuring a GitLab instance.
- A well-documented README file that explains the pipeline design, the tools used, and how to set up and use the pipeline.

###### Are all the required files present?

##### Play the role of a stakeholder:

As part of the evaluation process, conduct a simulated real-world scenario where the students assume the role of a DevOps engineer and explain their solution to a team or stakeholder. Evaluate their understanding of the concepts and technologies used in the project, as well as their ability to communicate effectively and think critically about their solution.

During the roleplay, ask them the following questions:

- Can you explain the concept of DevOps and its benefits for the software development lifecycle?

- How do DevOps principles help improve collaboration between development and operations teams?

- What are some common DevOps practices, and how did you incorporate them into your project?

- How does automation play a key role in the DevOps process, and what tools did you use to automate different stages of your project?

- Can you discuss the role of continuous integration and continuous deployment (CI/CD) in a DevOps workflow, and how it helps improve the quality and speed of software delivery?

- Can you explain the importance of infrastructure as code (IaC) in a DevOps environment, and how it helps maintain consistency and reproducibility in your project?

- How do DevOps practices help improve the security of an application, and what steps did you take to integrate security into your development and deployment processes?

- What challenges did you face when implementing DevOps practices in your project, and how did you overcome them?

- How can DevOps practices help optimize resource usage and reduce costs in a cloud-based environment?

- Can you explain the purpose and benefits of using GitLab and GitLab Runners in your project, and how they improve the development and deployment processes?

- What are the advantages of using Ansible for automation in your project, and how did it help you streamline the deployment of GitLab and GitLab Runners?

- Can you explain the concept of Infrastructure as Code (IaC) and how you implemented it using Terraform in your project?

- What is the purpose of using continuous integration and continuous deployment (CI/CD) pipelines, and how did it help you automate the building, testing, and deployment of your application?

- How did you ensure the security of the application throughout the pipeline stages?

- Can you explain the continuous integration (CI) pipeline you've implemented for each repository?

- Can you explain the continuous deployment (CD) pipeline you've implemented for each repository?

###### Do all of the students have a good understanding of the concepts and technologies used in the project?

###### Do all of the students have the ability to communicate effectively and explain their decisions?

###### Are all of the students capable of thinking critically about their solution and considering alternative approaches?

##### Review the GitLab and Runners Deployment. Ask the auditee **to show you**, the auditor, the use of the commands `ansible-playbook --list-tasks`, and/or `systemctl status` or any other necessary with the right options to answer the following questions.

###### Was the GitLab instance deployed and configured successfully using Ansible?

###### Are the GitLab Runners integrated with the existing pipeline and executing tasks as expected for all repositories?

##### Review the Infrastructure Pipeline:

###### Did the student deploy the infrastructure of the `cloud-design` project and the source code of `crud-master` project for two environments (staging, prod) on a cloud platform (e.g., AWS, Azure, or Google Cloud) using `Terraform`?

###### Are the two environments similar in design, resources and services used?

###### Does the student's infrastructure configuration exist in an independent repository with a configured pipeline?

###### Are the "Init", "Validate", "Plan", "Apply to Staging", "Approval", and "Apply to production environment" stages implemented correctly in the infrastructure pipeline?

##### Review the CI Pipeline:

- `Build`: Compile and package the application.
- `Test`: Run unit and integration tests to ensure code quality and functionality.
- `Scan`: Analyze the source code and dependencies for security vulnerabilities and coding issues. Consider using tools such as `SonarQube`, `Snyk`, or `WhiteSource`.
- `Containerization`: Package the applications into Docker images using a Dockerfile, and push the images to a container registry (e.g., Docker Hub, Google Container Registry, or AWS ECR).

###### Are the Build, Test, Scan, and Containerization stages implemented correctly in the CI pipeline for each repository?

##### Review the CD Pipeline:

- `Deploy to Staging`: Deploy the application to a `staging environment` for further testing and validation.
- `Approval`: Require manual approval to proceed with deployment to the `production environment`. This step should involve stakeholders and ensure the application is ready for production.
- `Deploy to Production`: Deploy the application to the `production environment`, ensuring zero downtime and a smooth rollout.

###### Are the "Deploy to Staging", "Approval", and "Deploy to Production" stages implemented correctly in the CD pipeline for each repository?

##### Review the functionality of pipelines. Ask the auditee **to show you**, the auditor, that the pipelines are functional by running one or several tests of their choosing.

###### Are the pipelines working properly and updating the application and infrastructure after each modification in each repository?

##### Check whether the students have effectively implemented the following cybersecurity guidelines:

`Restrict triggers to protected branches`: Ensure that the pipelines are triggered only on protected branches, preventing unauthorized users from deploying or tampering with the application. Check that access control measures are in place to minimize risk.

`Separate credentials from code`: Confirm that the students have not stored credentials in application code or infrastructure files. Look for the use of secure methods like secret management tools or environment variables to prevent exposure or unauthorized access.

`Apply the least privilege principle`: Assess if the students have limited user and service access to the minimum required level. This reduces potential damage in case of breaches or compromised credentials.

`Update dependencies and tools regularly`: Check if the students have a process for keeping dependencies and pipeline tools updated. Verify if they have automated updates and monitored for security advisories and patches to minimize security vulnerabilities.

###### Are triggers restricted to protected branches, ensuring unauthorized users cannot deploy or tamper with the application?

###### Have the students separated credentials from code, using secure methods like secret management tools or environment variables?

###### Did the students apply the least privilege principle to limit user and service access to the minimum required level?

###### Do the students have a process for updating dependencies and tools regularly, automating updates, and monitoring for security advisories and patches?

##### Review the Documentation:

###### Does the `README.md` file contain all the necessary information about the solution (prerequisites, setup, configuration, usage, ...)?

###### Is the documentation provided by the student clear and complete, including well-structured diagrams and thorough descriptions?

#### Bonus

###### +Did the student implemented any feature or anything that you would consider a bonus?

###### +Is this project an outstanding project?