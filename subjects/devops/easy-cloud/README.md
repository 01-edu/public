## Easy-Cloud

<center>
<img 
    src="./resources/easy-cloud.png?raw=true" style="width: 600px !important; height: 600px !important;"/>
</center>

### Objective

This project is designed to immerse you in the world of DevOps and Cloud engineering, focusing on deploying and managing a scalable application using a cloud provider from your choice, your mission is:

- Establishing and configuring a public cloud environment suitable for a scalable web application.
- Utilizing managed web service or managed container service for deploying containerized microservices deployment.
- Implementing effective monitoring, logging, auto-scaling, and security measures to ensure the application's smooth and efficient operation.
- Emphasizing the use of managed services to reduce the operational overhead on the startup.

### Hints

Before embarking on this project, ensure you are equipped with:

- An understanding of basic DevOps practices and principles.
- Familiarity with the public cloud concepts, managed public cloud services, and related networking/security services.
- Proficiency in Terraform for infrastructure as code.
- Knowledge of containerization technologies, specifically Docker.
- Insights into monitoring and logging practices and public cloud services.

> Any lack of understanding of the concepts of this project may affect the difficulty of future projects, take your time to understand all concepts.
> Be curious and never stop searching!

### Role Play

The project includes a role-play session to simulate a real-world scenario where you, as a Cloud Engineer, present and justify your infrastructure design and decisions to stakeholders or your team. This session aims to:

- Evaluate your grasp of the technologies and strategies employed in your project.
- Test your communication skills and ability to articulate your choices.
- Encourage critical thinking and the consideration of alternative solutions.

Be prepared to defend your architecture, explaining the rationale behind each decision and how it aligns with the project's objectives.

### Architecture

You are a Cloud engineer and you have received a mission from a small company with limited resources and limited cloud/DevOps knowledge, they want to create an infrastructure with the following:

- Setup WordPress on at least two servers to ensure high availability. These servers should be configured identically to serve the WordPress application.
- Implement an auto-scaling mechanism that automatically scales the number of servers up during high load periods and scales down during low load periods. This ensures that your infrastructure can handle traffic spikes efficiently without manual intervention.
- Configure a load balancer to distribute incoming traffic evenly across all active servers. This ensures that no single server bears too much load, improving the responsiveness of your WordPress site.
- Utilize a network file system (NFS) for storing dynamic WordPress content, such as media uploads, plugins, and themes. This setup allows all servers to access and serve the same dynamic content, ensuring consistency across different user sessions.
- Store static content in an object storage service, designed for high durability and availability. This storage should serve as static website assets like CSS, JavaScript, and image files.
- Integrate a Content Delivery Network (CDN) to cache and deliver static content from locations close to your users. This reduces load times and bandwidth costs while improving user experience.
- Use a managed database service for WordPress. This database should be accessible only by your WordPress servers and not exposed to the public internet, enhancing security.
- Ensure that the database service offers high-availability features, such as automatic backups, failover mechanisms, and scalability options.
- Design your infrastructure to be fault-tolerant by deploying servers in different physical locations or data centers, if possible. This minimizes the impact of hardware failures and network issues.
- Implement health checks and failover strategies to automatically reroute traffic away from failed instances and maintain high availability.
- Monitor and adjust your infrastructure according to actual needs. Use the auto-scaling feature to align operational costs with traffic patterns, ensuring you only pay for the resources you need.
- Apply strict access controls to ensure only authorized access to your infrastructure components. This includes securing your network file system, object storage, and database.
- Use the LTS version for WordPress applications, plugins, and themes to mitigate vulnerabilities.
- Use security groups or similar network security mechanisms to restrict inbound and outbound traffic to only necessary ports and protocols.
- Email alerts in the event of a problem with one of the services or a skip of the budget, you can define the budget by yourself.

> You are not allowed to use Kubernetes services or install it in the servers, the clients don't want to use it.

### Cost Management

Effective cost management is crucial. Embrace the pricing model, monitor your usage, optimize resource allocation, and leverage your cloud provider cost management tools to control your project's financial footprint. Regularly review and adjust your resources to ensure you achieve the best balance between performance, scalability, and cost.

> Remember that the responsibility for cost management lies with you.

You should provide a cost estimation with the submitted files, you can use a cloud service for that.

### Infrastructure as Code

Provision the necessary resources for your Cloud environment using Terraform as an Infrastructure as Code (IaC) tool. This includes setting up Cloud computing instances, containers, networking components, and storage services (e.g. AWS S3).

### Documentation

Create a `README.md` file that provides comprehensive documentation for your architecture, which must include well-structured diagrams, thorough descriptions of components, and an explanation of your design decisions, presented clearly and concisely. Make sure it contains all the necessary information about the solution (prerequisites, setup, configuration, usage, ...). This file must be submitted as part of the solution for the project.

### Submission and audit

Upon completing this project, you should submit the following:

- Your documentation in the `README.md` file.
- Configuration files for your Infrastructure as Code (IaC), containerization, and orchestration tools.
- Your cost estimation file.
