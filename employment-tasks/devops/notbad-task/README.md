# Senior DevOps Engineer - NotBad Header based API

Look at the following tasks and estimate how much time you will spend on them.

## Preconditions

### Technical & Knowledge
You need at least:

* Experience with AWS stack
* Experience with CI/CD
* Experience with Bash scripts
* Experience in at least one programming language (Java, Python, PHP, Perl, etc.)
* A text editor of your choice

## The tasks
1) We have a Terraform securitygroups.tf file. Every time Terraform runs, it says the security group in that file will be updated in place. Find a way to prevent this.  

2) Get the keycloak-test.tar.gz archive. What we can improve?

**2.1 Pull the archive from S3 bucket via `wget`**

```bash
wget https://lookinlabs-contribution-data.s3.eu-west-1.amazonaws.com/keycloak-test.tar.gz
```

**2.2 Find what can be improved in `keycloak-test.tar.gz`**

    
3) Provide infrastructure and create CI/CD with a web app that will listen to 8089 port and return "ReallyNotBad" string when POST request contains header "NotBad" with value "true", eg. `curl -X POST -H "NotBad: true" https://someurl:8089/` should return "ReallyNotBad".
Use any technology you want to deploy the application to AWS. It can be Ansible, Terraform, etc. or a combination of some of them.
Hint: https://aws.amazon.com/free/