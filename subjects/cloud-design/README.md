## Cloud-Design

<center>
<img 
    src="./resources/cloud-design.jpg?raw=true" style = "width: 600px
    !important; height: 600px !important;"/>
</center>

### Objective

The objective of this project is to challenge your understanding of DevOps and
cloud technologies by providing hands-on experience in deploying and managing a
microservices-based application on any major cloud platform (AWS, Azure, Google Cloud Platform, or others)
Your mission is to:

Set up and configure a cloud environment for deploying microservices on your chosen cloud provider. Deploy the
provided microservices application to your cloud environment. Implement
monitoring, logging, and scaling to ensure that the application runs
efficiently. Implement security measures, such as securing the databases and
making private resources accessible only from the Virtual Private Cloud
(VPC). Incorporate managed authentication for publicly accessible applications
using your cloud provider's identity services (Cognito, Azure AD, Cloud Identity, etc.). Optimize the application to handle
varying workloads and unexpected events.

### Hints

Before starting this project, you should know the following:

- Basic DevOps concepts and practices.
- Familiarity with containerization and orchestration tools, such as Docker and
  Kubernetes.
- Understanding of at least one major cloud platform.
- Familiarity with Infrastructure as Code (IaC) tools (Terraform, CloudFormation, ARM Templates, etc.).
- Knowledge of monitoring and logging tools, such as Prometheus, Grafana, and
  ELK stack.

> Any lack of understanding of the concepts of this project may affect the
> difficulty of future projects, take your time to understand all concepts.

> Be curious and never stop searching!

### Role play

To enhance the learning experience and assess your knowledge, a role play
question session will be included as part of the Cloud-Design Project. This
section will involve answering a series of questions in a simulated real-world
scenario where you assume the role of a Cloud engineer explaining your solution
to a team or stakeholder.

The goal of the role play question session is to:

- Assess your understanding of the concepts and technologies used in the
  project.
- Test your ability to communicate effectively and explain your decisions.
- Challenge you to think critically about your solution and consider
  alternative approaches.

Prepare for a role play question session where you will assume the role of a
Cloud engineer presenting your solution to your team or a stakeholder. You
should be ready to answer questions and provide explanations about your
decisions, architecture, and implementation.

### Architecture

Using your solutions from your previous projects `crud-master`,
`play-with-containers`, and `orchestrator`, you must design and deploy the
infrastructure on your chosen cloud platform respecting the project requirements, consisting of the
following components:

- `inventory-database container` is a PostgreSQL database server that contains
  your inventory database, it must be accessible via port `5432`.
- `billing-database container` is a PostgreSQL database server that contains
  your billing database, it must be accessible via port `5432`.
- `inventory-app container` is a server that contains your
  inventory-app code running and connected to the inventory database and
  accessible via port `8080`.
- `billing-app container` is a server that contains your billing-app
  code running and connected to the billing database and consuming the messages
  from the RabbitMQ queue, and it can be accessed via port `8080`.
- `RabbitMQ container` is a RabbitMQ server that contains the queue.
- `api-gateway-app container` is a server that contains your
  API Gateway code running and forwarding the requests to the other
  services, and it's accessible via port `3000`.

Design the architecture for your cloud-based microservices' application. You
are free to choose the services and architectural patterns that best suit your
needs, as long as they meet the project requirements and remain within a
reasonable cost range. Consider the following when designing your architecture:

1. `Scalability`: Ensure that your architecture can handle varying workloads
   and can scale up or down as needed. Cloud providers offer services like Auto Scaling
   that can be used to achieve this.

2. `Availability`: Design your architecture to be fault-tolerant and maintain
   high availability, even in the event of component failures.

3. `Security`: Incorporate security best practices into your architecture, such
   as encrypting data at rest and in transit, using private networks, and
   securing API endpoints. Also, ensure that the databases and private
   resources are accessible only from the AWS VPC and use AWS managed
   authentication for publicly accessible applications.

4. `Cost-effectiveness`: Be mindful of the costs associated with the services
   and resources you select. Aim to design a cost-effective architecture
   without compromising performance, security, or scalability.

5. `Simplicity`: Keep your architecture as simple as possible, while still
   meeting the project requirements. Avoid overcomplicating the design with
   unnecessary components or services.

### Cloud Provider Options

You may choose from any of the following cloud providers:

- **Amazon Web Services (AWS)**: Use services like EC2, ECS/EKS, RDS, VPC, Cognito, CloudWatch
- **Microsoft Azure**: Use services like Virtual Machines, Container Instances/AKS, SQL Database, Virtual Network, Azure AD, Azure Monitor
- **Google Cloud Platform (GCP)**: Use services like Compute Engine, GKE, Cloud SQL, VPC, Cloud Identity, Cloud Monitoring
- **Other providers**: DigitalOcean, Linode, Oracle Cloud, etc. (with instructor approval)

### Cost management:

1. `Understand the pricing model`: Familiarize yourself with the pricing model
   of the cloud provider and services you are using. Be aware of any free
   tiers, usage limits, and pay-as-you-go pricing structures.

2. `Monitor your usage`: Regularly check your cloud provider's billing
   dashboard to keep track of your usage and spending. Set up billing alerts to
   notify you when your spending exceeds a certain threshold.

3. `Clean up resources`: Remember to delete or stop any resources that you no
   longer need, such as virtual machines, storage services, and load balancers.
   This will help you avoid ongoing charges for idle resources.

4. `Optimize resource allocation`: Use the appropriate resource sizes for your
   needs and experiment with different configurations to find the most
   cost-effective solution. Consider using spot instances, reserved instances,
   or committed use contracts to save on costs, if applicable.

5. `Leverage cost management tools`: Many cloud providers offer cost management
   tools and services to help you optimize your spending. Use these tools to
   analyze your usage patterns and identify opportunities for cost savings.

> By being aware of your cloud usage and proactively managing your resources,
> you can avoid unexpected costs and make the most of your cloud environment.
> Remember that the responsibility for cost management lies with you, and it is
> crucial to stay vigilant and proactive throughout the project.

### Infrastructure as Code:

Provision the necessary resources for your cloud environment using Infrastructure as Code (IaC) tools. Choose the appropriate tool for your cloud provider:

- **Terraform**: Universal IaC tool (recommended for any provider)
- **AWS CloudFormation**: For AWS-specific deployments
- **Azure ARM Templates**: For Azure-specific deployments
- **Google Cloud Deployment Manager**: For GCP-specific deployments

This includes setting up compute instances, containers, networking components, and storage services.

### Containerize the microservices:

Use Docker to build container images for each microservice. Make sure to
optimize the Dockerfile for each service to reduce the image size and build
time.

### Deployment:

Deploy the containerized microservices on your chosen cloud platform using an orchestration tool like:

- **Kubernetes variants**: EKS (AWS), AKS (Azure), GKE (GCP), or self-managed
- **Cloud-native container services**: ECS (AWS), Container Instances (Azure), Cloud Run (GCP)

Ensure that the services are load-balanced and can communicate with each other securely.

> Use [this solution](https://github.com/01-edu/orchestrator) to kick-start
> your Kubernetes deployment.

### Monitoring and logging:

Set up monitoring and logging tools to track the performance and health of your
application. Use cloud-native tools or third-party solutions:

- **Cloud-native**: CloudWatch (AWS), Azure Monitor (Azure), Cloud Monitoring (GCP)
- **Third-party**: Prometheus/Grafana, ELK stack

### Optimization:

Implement auto-scaling policies to handle varying workloads and ensure high
availability. Test the application under different load scenarios and adjust
the resources accordingly.

### Security:

Implement security best practices such as using managed certificate services for
HTTPS, securing API endpoints with API gateways, regularly scanning for
vulnerabilities, and implementing managed authentication for
publicly accessible applications. Ensure
that the databases and private resources are secure and accessible only from
the cloud VPC/Virtual Network.

### Documentation

Create a `README.md` file that provides comprehensive documentation for your
architecture, which must include well-structured diagrams, thorough
descriptions of components, and an explanation of your design decisions,
presented in a clear and concise manner. Make sure it contains all the
necessary information about the solution (prerequisites, setup, configuration,
usage, ...). This file must be submitted as part of the solution for the
project.

### Bonus

If you complete the mandatory part successfully and you still have free time,
you can implement anything that you feel deserves to be a bonus, for example:

- Use your own `crud-master`, `play-with-containers`, and `orchestrator`
  solution instead of the provided ones.

- Use `Function as a Service (FaaS)` in your solution.

- Use `Content Delivery Network (CDN)` to optimize your solution.

- Implementing alert systems to ensure your application runs smoothly.

Challenge yourself!

### Submission and audit

Upon completing this project, you should submit the following:

- Your documentation in the `README.md` file.
- Source code for the microservices and any scripts used for deployment.
- Configuration files for your Infrastructure as Code (IaC), containerization,
  and orchestration tools.
