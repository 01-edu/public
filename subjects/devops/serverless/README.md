## Serverless Payments Reminder

Serverless Payments Reminder is a basic Slack Bot that reminds companies / users to pay their bills. The reminder gets triggered by AWS EventBridge (also known as CloudWatch Events). The bot itself is hosted on a simple AWS Lambda function.

### Requirements for the task

Create a simple Slack bot using:

- [AWS](https://aws.amazon.com/)
- [Serverless Framework](https://serverless.com/)
- [Cloudformation](https://aws.amazon.com/cloudformation/)
- [Terraform](https://www.terraform.io/)

### Task

The task is separated into multiple different levels. The levels are ordered by complexity. You can choose the level that you feel most comfortable with.

## Level 0: Basic Lambda Function hosted via Serverless

**1. Create Slack Bot**

Write a slack bot, that sends a reminder to slack channel with the following message:

```sh
Dear Board Members! This is a reminder to make the payment for the licenses of the software. The due date is 07.XX.YY, where XX is the month and YY is the year. The amount to be paid is $ZZZ.VV. Please make the payment as soon as possible to <IBAN_NUMBER>. Thank you!
```

The environment variables which must be available in the Lambda function are:

- `SLACK_WEBHOOK_URL` - The Slack Webhook URL
- `AMOUNT` - The amount to be paid
- `IBAN_NUMBER` - The IBAN number

The due date must be calculated based on the current date for the next month.

Before step two, you need to test that Slack Bot is functioning properly by running an application locally. It should send a message to the Slack channel.

**2. Write tests for the application**

Write integration tests for the application to ensure that each part of the application is functioning properly.

**3. Use [Serverless Framework](https://github.com/serverless/serverless) to host the lambda function**

Serverless framework is an automation tool used to deploy serverless applications which can help with event-driven architecture deployments.

Use the serverless framework to host the lambda function combined with AWS Eventbridge.

**Refs**

- [AWS EventBridge](https://www.serverless.com/framework/docs/providers/aws/events/event-bridge)
- [AWS Lambda Function](https://www.serverless.com/framework/docs/providers/aws/guide/functions).

AWS Eventbridge should trigger lambda function every month on the 1st day of the month at 10:00 AM.

**Note** Before hosting, you need to ensure that the bot is written in a way that it can be hosted on AWS Lambda. You can look at the example [here](https://github.com/KostLinux/aws-incident-manager-notifier/blob/56d52e90f8a14e689e7d2a1c7ee44590de5af2f5/main.go#L158).

## Level 1: Basic Lambda Function automation via Cloudformation

**1. Repeat the 1st and 2nd step from Level 0**

Create a slack bot, that sends a reminder to slack channel. Write integration tests for the application to ensure that each part of the application is functioning properly.

**Note** Don't forget to add Lambda handler for the application

**2. Use [Cloudformation](https://aws.amazon.com/cloudformation/) to host the lambda function**

Cloudformation is Infrastructure As Code (IaC) tool used to automate the provisioning of AWS resources inside the AWS. It allows you to use a simple yaml syntax to define the resources you want to create.

**2.1 Hosting the Lambda function**

Write cloudformation template with parameters (variables) `SlackWebhookUrl`, `Amount`, `IbanNumber`. The template should automatically provision the Slackbot based on architecture mentioned in [Level 0](#level-0-basic-lambda-function-hosted-via-serverless) step 3.

## Level 2: Basic Lambda Function automation via Basic Terraform with Local State\*\*

Terraform is an Infrastructure As Code (IaC) tool used to automate the provisioning of resources inside different SaaS solutions and Cloud providers. It uses HCL syntax, which is easy maintainable. The main problem of terraform is that you need to know, how everything works under the hood (e.g. IAM roles, policies, serverless platforms etc).

**1. Repeat the 1st and 2nd step from Level 0**

Create a slack bot, that sends a reminder to slack channel. Write integration tests for the application to ensure that each part of the application is functioning properly.

**Note** Don't forget to add Lambda handler for the application

**2. Automate the hosting via Terraform**

**2.1 Write an IAM role with the least privilege principle for the lambda function and EventBridge.**

**2.2 Automatically pack up the lambda function into a zip file**

**2.3 Build the lambda function based on the zip file**

**2.4 Build AWS Eventbridge, which will trigger the lambda function**

**Note** Lambda function should be triggered every month on the 1st day of the month at 10:00 AM.

At this level you can use local state for terraform, no modules needed to write.

## Level 3: Basic Lambda Function automation via Terraform Module with Remote S3 State (advanced)\*\*

In a DevOps world, mostly we use remote state for terraform. The main reason is that we can share the state with other team members and we can easily manage the state of the infrastructure. Although it is a bit more complex to set up, it allows you to avoid issues when someone even removes the state file, cause S3 has versioning turned on.

This level is more advanced, but more near to the real-world scenario. As a DevOps Engineer / SRE mostly all the terraform code is written in modules (or using community pre-built modules). Modules help us to maintain the codebase and reuse the code in different projects, so you don't need to write the same code again and again.

**1. Repeat the steps from [Level 2](#level-2-basic-lambda-function-automation-via-basic-terraform-with-local-state)**

Repeat all the steps from Level 2, but now you need to write a terraform module for the lambda function and EventBridge. Which means that instead of writing values directly into main.tf file, everything should be written using variables. For example:

```tf
resource "aws_lambda_function" "slack_bot" {
  function_name = var.function_name
  role = var.role
  handler = var.handler
  runtime = var.runtime
  timeout = var.timeout
}
```

We're using variables for each value, so we can reuse the module in different projects.

**2. Use remote state for terraform**

In this step you need to configure IAM User with least privilege principle for terraform to access S3 bucket, provision resources in Eventbridge and Lambda function.

**Note** We would recommend using [Cloudformation](https://aws.amazon.com/cloudformation/) to create an IAM User with the least privilege principle + S3 Bucket for the remote state. Cloudformation can be applied via AWS CLI which means that everything persists to be AS CODE.

**3. Use the S3 bucket for the remote state and module to provision the resources**

Write your backend configuration to backend.tf file:

```tf
terraform {
  backend "s3" {
        bucket = ""
        key = ""
        region = ""
        encrypt = ""
    }
}
```

Use the module written at step 1 to provision the resources

```tf
module "slack_bot" {
  source = "./modules/slack_bot"
  function_name = "slack_bot"
  role = "arn:aws:iam::123456789012:role/lambda-role"
  handler = "main.handler"
  runtime = "nodejs14.x"
  timeout = 10
}
```

## Helpful commands

- `serverless deploy` - Deploy the serverless application
- `aws cloudformation deploy --template-file template.yaml --stack-name my-stack` - Deploy the cloudformation stack
- `terraform init` - Initialize the terraform project and remote state.
- `terraform plan` - Plan the resources you're going to provision. It will show you the changes that will be made.
- `terraform apply` - Apply the terraform project only if you're sure that everything is correct during the plan.
