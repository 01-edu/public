#### General

##### Check the Repo content.

###### Are all files related to the CloudPress project, including Terraform configuration files, Ansible playbooks, and any additional scripts, included in the repository for example (Helper Scripts, Configuration Scripts, Utility Scripts, Deployment Control Scripts)?

#### AWS Infrastructure Setup with Terraform

###### Were EC2 instances provisioned using Terraform?

###### Were security groups, storage, and networking configurations set up correctly using Terraform?

###### Was the EC2 Instance securely configured after setup (SSH, security, etc.)?

#### Nginx Installation and Configuration using Ansible

###### Was Nginx successfully installed on the EC2 instance using Ansible? Run `nginx -v` on the EC2 instance to validate the installed Nginx version, ensuring the successful installation via Ansible.

###### Was Nginx configured to serve web content appropriately?

###### Is Nginx service running? Use `systemctl status nginx`.

###### Were firewall settings adjusted to allow HTTP/HTTPS traffic?

#### MariaDB Installation and Configuration using Ansible

###### Was MariaDB installed on the EC2 Instance using Ansible? Run` mariadb -V` on the EC2 instance to confirm the MariaDB version, verifying the installation through Ansible.

###### Is MariaDB service active? Confirm the status of the MariaDB service by using the command `systemctl status mariadb` to ensure its active state, indicating proper functionality post-deployment.

###### Was a MariaDB database created specifically for WordPress? Confirm by running the command `SHOW DATABASES` to verify the existence of the WordPress database.

###### Were the user rights and privileges for the WordPress user correctly configured?

#### PHP Installation and Configuration using Ansible

###### Was PHP installed on the EC2 Instance using Ansible? Run `php -v` on the EC2 instance to verify the PHP version, ensuring the installation was successfully carried out by Ansible.

###### Is PHP-FPM running? Run `systemctl status php-fpm`.

###### Were PHP settings appropriately configured for the WordPress site?

###### Was the PHP service successfully started?

#### WordPress Installation and Configuration using Ansible

###### Was WordPress downloaded and set up on the EC2 Instance using Ansible? Execute `wp core version` on the EC2 Instance provisioned by Ansible to confirm the successful download and setup of WordPress.

###### Was WordPress configured to use the MariaDB database?

###### Is the database connectivity confirmed?

###### Was the WordPress service successfully started?

###### Using curl or browser: Is WordPress accessible?

###### Can users access the WordPress site without encountering errors or downtime?

###### Was the environment configuration secure and effective?

Consider the following examples of secure and effective environment configuration:

- **Sensitive Information Handling:** Ensure sensitive information like passwords, API keys, or access credentials are not hardcoded in configuration files. Utilize specialized tools such as AWS Secrets Manager, HashiCorp Vault, or equivalent solutions to securely manage and access sensitive data.

- **Scalability Measures:** Check for implemented strategies to handle traffic spikes or increased loads, such as auto-scaling mechanisms or load balancing.

- **Fault Tolerance:** Validate the setup's ability to manage various failure scenarios, maintaining high availability of the WordPress site.

###### Are sensitive information such as passwords and access keys securely stored and managed (e.g., using AWS Secrets Manager or similar tools)?

###### Were proper encryption methods implemented for data in transit and at rest?

#### Documentation

###### Does the README.md file contain an architecture overview?

###### Does the documentation cover the deployment process comprehensively?

###### Are additional configurations, tips, and potential pitfalls included in the documentation?

#### Technical Verification

##### Change some Terraform settings as instructed below:

- Modify the count of EC2 instances. For instance, if you initially provisioned two instances, change the count to three or reduce it to one.
- Alter inbound or outbound rules in the security groups, like opening or closing specific ports.
- Other changes that you consider relevant.

###### Did altering Terraform settings reflect changes in the deployed infrastructure?

###### Is WordPress fully functional? (Adding a page, accessing the website, etc.)

##### Nginx Configuration Validation

###### Were changes in Nginx configurations regarding traffic limitation or exposed ports effective in controlling or limiting access?
