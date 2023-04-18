#### General

##### Check the Repo content.

Files that must be inside the repository:

- Your documentation in the `README.md` file.
- Source code for the microservices and any scripts used for deployment.
- Configuration files for your Infrastructure as Code (IaC), containerization, and orchestration tools.

###### Are all the required files present?

##### Play the role of a stakeholder.

As part of the evaluation process, conduct a simulated real-world scenario where the students assume the role of a Cloud engineer and explain their solution to a team or stakeholder. Evaluate their understanding of the concepts and technologies used in the project, as well as their ability to communicate effectively and think critically about their solution.

During the roleplay, ask them the following questions:

- What is the cloud, and what are its benefits?

- Why should we deploy the solution in the cloud, instead of on-premises?

- How did you decide on the cloud provider for this project, and what factors did you consider when making that decision?

- Can you explain the architecture of your cloud-based microservices application and how the different components interact with each other?

- How did you manage the cost of your solution, and what strategies did you use to optimize it?

- How did you ensure that the application is secure, and what security best practices did you implement to protect your application?

- What monitoring and logging tools did you use, and how did they help you identify and troubleshoot issues with your application?

- Can you explain the auto-scaling policies you implemented and how they allow your application to handle varying workloads?

- How did you optimize the Docker images for each microservice, and how did it impact build times and image sizes?

- If you were to start this project again, what changes would you make to your approach or the technologies you used?

- How can your solution be extended or modified to handle future requirements, such as adding new microservices or migrating to a different cloud provider?

- What challenges did you encounter during the project, and how did you overcome them?

- How did you ensure that your documentation is clear and complete, and what steps did you take to make it easy for others to understand and maintain your solution?

###### Did the students have a good understanding of the concepts and technologies used in the project?

###### Did the students have the ability to communicate effectively and explain their decisions?

###### Are the students capable to think critically about their solution and consider alternative approaches?

##### Review the Architecture Design.

Review the student's architecture design, ensuring that it meets the project requirements:

1. `Scalability`: Ensure that the architecture can handle varying workloads and can scale up or down as needed.

2. `Availability`: Design the architecture to be fault-tolerant and maintain high availability, even during component failures.

3. `Security`: Incorporate security best practices into the architecture, such as encrypting data at rest and in transit, using private networks, and securing API endpoints.

4. `Cost-effectiveness`: Be mindful of the costs associated with the services and resources you select. Aim to design a cost-effective architecture without compromising performance, security, or scalability.

5. `Simplicity`: Keep the architecture as simple as possible, while still meeting the project requirements. Avoid overcomplicating the design with unnecessary components or services.

###### Did the student Architecture design meets the project requirements?

###### Did the students have the ability to design a cost-effective architecture that meets the project requirements?

###### Did the choice of services and architectural patterns align with best practices for scalability, availability, and security?

##### Check the student documentation in the `README.md` file.

###### Does the `README.md` file contain all the necessary information about the solution (prerequisites, setup, configuration, usage, ...)?

###### Is the documentation provided by the student clear and complete, including well-structured diagrams and thorough descriptions?

##### Verify the deployment.

###### Are all the microservices running as expected in the cloud environment, with no errors or connectivity issues?

###### Is the load balancing configured correctly, effectively distributing traffic across the services?

###### Are the microservices communicating with each other securely, using proper authentication and encryption methods?

##### Evaluate the infrastructure setup.

###### Are the `Infrastructure as Code (IaC)` tools, such as `Terraform` or `CloudFormation`, used effectively to provision and manage resources in the cloud environment?

###### Does the infrastructure setup follow the best practices for security and resource management?

##### Assess containerization and orchestration.

###### Are the Dockerfiles optimized for efficient container builds, and do they follow best practices?

###### Is the orchestration setup (e.g., Kubernetes manifests or AWS ECS task definitions) configured correctly and according to best practices?

##### Evaluate monitoring and logging.

###### Do monitoring and logging dashboards provide useful insights into the application performance and health?

##### Assess optimization efforts.

###### Are the auto-scaling policies configured correctly to handle varying workloads?

###### Does the application and resource allocation remain efficient under different load scenarios?

##### Check security best practices.

###### Has the student implemented security best practices, such as using HTTPS, securing API endpoints, and regularly scanning for vulnerabilities?

#### Bonus

##### Review Cloud Certification Preparation Efforts

Verify that the student has chosen a cloud platform certification to pursue and has made efforts to study the core concepts, services, and best practices for the chosen platform. Provide guidance and feedback on their preparation efforts, and offer suggestions for further learning and resources to help them succeed in their certification exam.

###### +Has the student chose a popular cloud platform certification to pursue, and is it relevant to their project or area of interest?

###### +Has the student demonstrated efforts to study and understand the core concepts, services, and best practices for the chosen platform?

###### +Is the student prepared and confident in taking the certification exam, and can they apply their knowledge to real-world scenarios?

###### +Are alerting mechanisms in place to notify the team of potential issues on time?

###### +Did the student add any optional bonus?

###### +Is this project an outstanding project?
