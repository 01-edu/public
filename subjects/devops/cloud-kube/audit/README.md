#### General

##### Check the Repo content.

Files that must be inside the repository:

- Detailed documentation in the `README.md` file.
- Source code for the microservices and scripts required for deployment.
- Configuration files for the chosen cloud provider Infrastructure as Code (IaC), containerization, and orchestration tools.

###### Are all the required files present?

##### Play the role of a stakeholder.

Organize a simulated scenario where the students take on the role of Cloud engineers and explain their solution to a team or stakeholder. Evaluate their grasp of the concepts and technologies used in the project, their communication efficacy, and their critical thinking about their solution.
Suggested role play questions include:

- What is the cloud and its associated benefits?
- Why is deploying the solution in the cloud preferred over on-premises?
- How would you differentiate between public, private, and hybrid cloud?
- What drove your decision to select the specific cloud provider for this project, and what factors did you consider?
- Can you describe your microservices application's cloud-based architecture and the interaction between its components?
- How did you manage and optimize the cost of your cloud solution?
- What measures did you implement to ensure application security on the cloud, and what security best practices did you adhere to?
- What monitoring and logging tools did you utilize, and how did they assist in identifying and troubleshooting application issues?
- Can you describe the auto-scaling policies you implemented and how they help your application accommodate varying workloads?
- How did you optimize Docker images for each microservice, and how did it influence build times and image sizes?
- If you had to redo this project, what modifications would you make to your approach or the technologies you used?
- How can your solution be expanded or altered to cater to future requirements like adding new microservices or migrating to a different cloud provider?
- What challenges did you face during the project and how did you address them?
- How did you ensure your documentation's clarity and completeness, and what measures did you take to make it easily understandable and maintainable?

###### Were the students able to answer all the questions correctly?

###### Did the students demonstrate a thorough understanding of the concepts and technologies used in the project?

###### Were the students able to communicate effectively and justify their decisions?

###### Could the students critically evaluate their solution and consider alternative strategies?

##### Review the Architecture Design.

Review the student's architecture design, ensuring that it meets the project requirements:

1. `Scalability`: Does the architecture utilize services to manage varying workloads and scale as required?
2. `Availability`: Is the architecture designed to be fault-tolerant and maintain high availability, even during component failures?
3. `Security`: Does the architecture integrate security best practices, such as data encryption, use of VPC, and secure API endpoints with managed authentication?
4. `Cost-effectiveness`: Is the architecture designed to be cost-effective on the chosen cloud provider without compromising performance, security, or scalability?
5. `Simplicity`: Is the architecture straightforward and free of unnecessary complexity while still fulfilling project requirements?

###### Did the architecture design and choice of services align with all the project requirements above?

###### Were the students able to design a cost-effective architecture that meets the project requirements?

##### Check the student documentation in the `README.md` file.

###### Does the `README.md` file contain all the necessary information about the solution (prerequisites, setup, configuration, usage, ...)?

###### Is the documentation provided by the student clear and complete, including well-structured diagrams and thorough descriptions?

##### Verify the deployment. Ask the auditee **to show you**, the auditor, the use of the commands `kubectl`, the CLI tool for the chosen cloud provider and any other necessary with the right options to answer the following questions.

##### Ask the student to run the command `kubectl get pods -A`. Ensure the `KUBECONFIG` environment variable is correctly set to communicate with the K8s cluster deployed in the cloud.

###### Are all the microservices running as expected in the cloud environment?

##### Ask the student to run the command `kubectl get services`.

###### Are there enough services deployed to allow proper communication between the pods and the internet?

##### Ask the student to run the command `kubectl get pvc` or `kubectl get pv`.

###### Are there enough persistent volumes (or volumes claims) per each database defined in the application?

##### Test CRUD (create, read, update and delete) operations for the `inventory app` microservice.

###### Is everything working as expected?

##### Test the `billing app` microservice. Add an order to the queue and ask the student to show the updated database.

###### Is everything working as expected?

###### Is the load balancing configured correctly, effectively distributing traffic across the services?

###### Are the microservices communicating with each other securely, using proper authentication and encryption methods?

##### Evaluate the infrastructure setup. Ask the auditee **to show you**, the auditor, the use of the commands `terraform plan` and/or `terraform apply` to answer the following questions.

###### Is `Terraform` used effectively to provision and manage resources in the cloud environment?

###### Does the infrastructure setup follow the architecture design and the project requirements?

##### Assess containerization and orchestration. Ask the auditee **to show you**, the auditor, the use of the commands chose cloud provider CLI, `docker ps`, and/or `kubectl` or any other necessary with the right options to answer the following questions.

###### Are the Dockerfiles optimized for efficient container builds?

###### Is the orchestration setup (e.g., Kubernetes manifests) configured correctly?

##### Evaluate monitoring and logging.

###### Are monitoring and logging dashboards providing useful insights into the application performance and health?

##### Assess optimization efforts.

###### Are the auto-scaling policies configured correctly to handle varying workloads?

###### Does the application and resource allocation remain efficient under different load scenarios?

##### Check security best practices.

###### Has the student implemented security best practices, such as using HTTPS, securing API endpoints, and regularly scanning for vulnerabilities?

#### Bonus

###### +Did the student add any optional bonus?

###### +Is this project an outstanding project?
