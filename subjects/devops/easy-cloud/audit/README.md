#### General

##### Check the Repo content:

Ensure the repository contains all required files:

- Comprehensive README.md file with detailed documentation of the architecture, design decisions, and component descriptions.
- All source code and scripts necessary for deploying the WordPress application.
- Terraform configuration files for the Infrastructure as Code (IaC) setup.
- A cost estimation file reflecting the budget and financial considerations for the project.

###### Are all required files present?

##### Play the role of a stakeholder:

Conduct a role-play session where students present their solution as Cloud Engineers. Evaluate their understanding of cloud concepts, their solution's architecture, and their ability to communicate effectively. Suggested questions:

- Explain the choice of cloud provider and the benefits of using managed services in this project.
- How does your architecture ensure high availability and scalability for the WordPress application?
- Describe how you've implemented security measures to protect your application and data.
- Discuss the auto-scaling strategy you've chosen and how it aligns with cost management principles.
- How have you managed the storage of dynamic and static content to ensure consistency and performance?
- What monitoring and alerting mechanisms have you put in place to maintain the health and performance of your application?
- Provide details on how you used Terraform for infrastructure provisioning and the benefits of using Infrastructure as Code (IaC).

###### Were the students able to answer all questions satisfactorily?

###### Did the students demonstrate a comprehensive understanding of the technologies and strategies used?

###### Could the students effectively communicate their design decisions and justify their architectural choices?

##### Review the Architecture Design:

Examine the student's architecture design to ensure it meets all project requirements:

1. `Scalability`: Does the architecture include an auto-scaling setup that efficiently handles varying loads?
2. `High Availability`: Are services deployed across multiple servers or data centers to ensure fault tolerance?
3. `Security`: Are appropriate security measures in place, including access controls and database security?
4. `Cost Management`: Is there evidence of cost optimization strategies and a clear understanding of the chosen cloud provider's pricing model?
5. `Simplicity and Compliance`: Has the student managed to design a straightforward solution without Kubernetes, as per the project's constraints?

###### Does the architecture align with the project's scalability, availability, security, and cost-effectiveness requirements?

##### Test WordPress Functionality:

- Navigate to the WordPress site URL.
- Upload a new media item to the WordPress site.
- Access various static assets (CSS, JavaScript, images).

###### Are these loading quickly, indicating effective CDN use?

###### Does the site load correctly from different locations or using different internet connections?

###### Is the item immediately accessible from different sessions, indicating successful NFS setup for dynamic content?

###### Is the WordPress application fully functional and serving content correctly?

##### High Availability and Auto-scaling Verification:

- Ask the student to demonstrate how WordPress instances are distributed across different physical locations or data centers.
- Simulate a server failure or take one instance offline.
- Show evidence of the auto-scaling configuration. Then, simulate or describe how the system scales up under high load and scales down when the load decreases.
- Perform a controlled load test to trigger auto-scaling. Monitor and record the scaling activity.

###### Does the site remain accessible, indicating successful failover and high availability?

###### Is the WordPress setup highly available with no downtime during server failures?

###### Does the auto-scaling mechanism work as expected, adjusting the number of instances according to load?

##### Load Balancing and Traffic Distribution:

- Explain or show the load balancer setup ensuring even traffic distribution.
- Ask the student to explain how is traffic evenly distributed across all servers.

###### Is the load balancing configured correctly, effectively distributing traffic across servers?

##### Alert and Monitoring Setup:

- Show the configuration for email alerts related to service problems or budget overruns.
- Trigger an Alert: (If possible), simulate a condition that would trigger an alert (e.g., stopping a service or exceeding a predefined budget threshold).

###### Are alerts properly set up and working as expected for system issues and budget overruns?

##### Security and Access Control:

- Ask the student to explain how are strict access controls implemented to secure NFS, object storage, and the database.
- Verify that the managed database is accessible only by WordPress servers and not exposed to the internet.
- Check the configuration of security groups or similar mechanisms to restrict traffic to necessary protocols and ports only.

###### Are security measures adequately implemented to protect infrastructure components?

##### Documentation Review:

Check the README.md file for completeness:

###### Does it include detailed setup and configuration instructions?

###### Are there clear, well-structured diagrams and thorough component descriptions?

###### Is there a clear explanation of design decisions and their rationale?

###### Is the documentation clear, complete, and easy to understand?

##### Infrastructure as Code (IaC) Evaluation:

- Request the student to demonstrate the use of Terraform commands (terraform plan, terraform apply) to verify the cloud environment setup.

###### Is Terraform used effectively to provision and manage cloud resources?

###### Do the Terraform configurations accurately reflect the architecture design and project requirements?

###### Does the IaC setup align with best practices and project requirements?

##### Cost Estimation Review:

- Examine the cost estimation file submitted by the student:

###### Does it provide a realistic estimate of the project's cost on the chosen cloud provider?

###### Are there considerations for scaling, data transfer, storage, and other potential costs?

###### Is the cost estimation detailed, realistic, and aligned with cloud provider pricing models?
