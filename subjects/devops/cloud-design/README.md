## Cloud-Design

![Cloud-Design](pictures/cloud-design.jpg)

### Objective

This subject aims to challenge your understanding of DevOps and cloud technologies by providing hands-on experience in deploying and managing a microservices-based application on a cloud platform. You will use your `crud-master` project source code, your mission is to:

- Set up and configure a cloud environment for deploying microservices.
- Deploy the provided microservices application to the cloud environment.
- Implement monitoring, logging, and scaling to ensure the application runs efficiently.
- Optimize the application to handle varying workloads and unexpected events.

### Tips

Before starting this project, you should know the following:

- Basic DevOps concepts and practices.
- Familiarity with containerization and orchestration tools, such as Docker and Kubernetes.
- Understanding of cloud platforms, such as AWS, Azure, or GCP.
- Familiarity with Infrastructure as Code (IaC) tools, such as Terraform or CloudFormation.
- Knowledge of monitoring and logging tools, such as Prometheus, Grafana, and ELK stack.

> Any lack of understanding of the concepts of this project may affect the difficulty of future projects, take your time to understand all concepts.

> Be curious and never stop searching!

### Cloud Certification Preparation (optional)

Choose a popular cloud platform certification to pursue, based on the platform you will for this project or your area of interest. Study the core concepts, services, and best practices for the chosen platform, and use the provided resources to help prepare for the certification exam.

#### Popular Fundamental Cloud Platform Certifications:

1. `AWS Certified Cloud Practitioner`: This certification covers the fundamentals of AWS cloud services, architecture, and cost management. It is intended for individuals who want to validate their understanding of the AWS platform.

- [AWS Certified Cloud Practitioner Exam Guide](https://aws.amazon.com/fr/certification/certified-cloud-practitioner/)
- [AWS Certified Cloud Practitioner Learning Path](https://aws.amazon.com/fr/training/learn-about/cloud-practitioner/)
- [AWS Well-Architected Framework](https://aws.amazon.com/architecture/well-architected/)

2. `Microsoft Azure Fundamentals (AZ-900)`: This certification covers the basics of Microsoft Azure cloud services, architecture, and security. It is designed for individuals who want to demonstrate their understanding of the Azure platform.

- [Microsoft Azure Fundamentals (AZ-900) Exam Guide](https://docs.microsoft.com/en-us/learn/certifications/exams/az-900)
- [Microsoft Azure Fundamentals (AZ-900) Learning Path](https://docs.microsoft.com/en-us/learn/paths/azure-fundamentals/)
- [Microsoft Azure Well-Architected Framework](https://docs.microsoft.com/en-us/azure/architecture/framework/)

3. `Google Cloud Platform Associate Cloud Engineer`: This certification covers the fundamentals of the Google Cloud Platform (GCP) services, architecture, and security. It is intended for individuals who want to validate their understanding of the GCP platform.

- [Google Cloud Associate Cloud Engineer Exam Guide](https://cloud.google.com/certification/cloud-engineer)
- [Google Cloud Associate Cloud Engineer Learning Path](https://cloud.google.com/training/cloud-engineer)
- [Google Cloud Architecture Framework](https://cloud.google.com/architecture/framework)

### Roleplay

To further enhance the learning experience and assess your knowledge, a roleplay question session will be included as part of the Cloud-Design Project. This section will involve your answering a series of questions in a simulated real-world scenario, where you will assume the role of a Cloud engineer explaining your solution to a team or a stakeholder.
The goal of the roleplay question session is to:

- Assess your understanding of the concepts and technologies used in the project.
- Test your ability to communicate effectively and explain your decisions.
- Challenge you to think critically about your solution and consider alternative approaches.

Prepare for a roleplay question session where you will assume the role of a Cloud engineer presenting your solution to your team or a stakeholder. You should be ready to answer questions and provide explanations about your decisions, architecture, and implementation.

### Design the Architecture

Design the architecture for your cloud-based microservices application. You are free to choose the services and architectural patterns that best suit your needs, as long as they meet the project requirements and remain within a reasonable cost range. Consider the following when designing your architecture:

1. `Scalability`: Ensure that your architecture can handle varying workloads and can scale up or down as needed.

2. `Availability`: Design your architecture to be fault-tolerant and maintain high availability, even in the event of component failures.

3. `Security`: Incorporate security best practices into your architecture, such as encrypting data at rest and in transit, using private networks, and securing API endpoints.

4. `Cost-effectiveness`: Be mindful of the costs associated with the services and resources you select. Aim to design a cost-effective architecture without compromising performance, security, or scalability.

5. `Simplicity`: Keep your architecture as simple as possible, while still meeting the project requirements. Avoid overcomplicating the design with unnecessary components or services.

### Instructions

#### Choose a cloud platform:

Select a cloud provider (AWS, Azure, GCP, or any other provider of your choice) and create an account.

> While working on this project, it is essential to be aware of the potential costs associated with using cloud resources. Each cloud provider offers various services with different pricing models, and it can be easy to incur unexpected charges if you don't carefully manage your resources.

#### Cost management:

1. `Understand the pricing model`: Familiarize yourself with the pricing model of the cloud provider and services you are using. Be aware of any free tiers, usage limits, and pay-as-you-go pricing structures.

2. `Monitor your usage`: Regularly check your cloud provider's billing dashboard to keep track of your usage and spending. Set up billing alerts to notify you when your spending exceeds a certain threshold.

3. `Clean up resources`: Remember to delete or stop any resources that you no longer need, such as virtual machines, storage services, and load balancers. This will help you avoid ongoing charges for idle resources.

4. `Optimize resource allocation`: Use the appropriate resource sizes for your needs and experiment with different configurations to find the most cost-effective solution. Consider using spot instances, reserved instances, or committed use contracts to save on costs, if applicable.

5. `Leverage cost management tools`: Many cloud providers offer cost management tools and services to help you optimize your spending. Use these tools to analyze your usage patterns and identify opportunities for cost savings.

> By being aware of your cloud usage and proactively managing your resources, you can avoid unexpected costs and make the most of your cloud environment. Remember that the responsibility for cost management lies with you, and it is crucial to stay vigilant and proactive throughout the project.

#### Set up the infrastructure:

Provision the necessary resources for your cloud environment using IaC tools like `Terraform` or `CloudFormation`. This includes setting up virtual machines, networking components, and storage services.

#### Containerize the microservices:

Use Docker to build container images for each microservice. Make sure to optimize the Dockerfile for each service to reduce the image size and build time.

> You can use your `play-with-containers` project solution.

#### Deploy the application:

Deploy the containerized microservices on your cloud platform using an orchestration tool like Kubernetes or AWS ECS. Ensure that the services are load-balanced and can communicate with each other securely.

> You can use your `orchestrator` project solution.

#### Implement monitoring and logging:

Set up monitoring and logging tools to track the performance and health of your application. Use tools like Prometheus, Grafana, and ELK stack to visualize metrics and logs.

#### Optimize the application:

Implement auto-scaling policies to handle varying workloads and ensure high availability. Test the application under different load scenarios and adjust the resources accordingly.

#### Secure the application:

Implement security best practices, such as using HTTPS, securing API endpoints, and regularly scanning for vulnerabilities.

### Documentation

In a `README.md` file create clear and concise documentation for your architecture, including diagrams, descriptions of the components, and the rationale behind your design choices.

### Bonus

If you complete the mandatory part successfully and you still have free time, you can implement anything that you feel deserves to be a bonus, for example:

- Use Function as a Service in your solution.

- Use CDN to optimize your solution.

Challenge yourself!

### Submission and audit

Upon completing this project, you should submit the following:

- Your documentation in the `README.md` file.
- Source code for the microservices and any scripts used for deployment.
- Configuration files for your IaC, containerization, and orchestration tools.
