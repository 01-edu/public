## CRUD Master

### Instructions

APIs are a very common and convenient way to deploy services in a modular way.
In this exercise we will create a simple API infrastructure, having an API gateway connected with two other APIs.
Those two APIs will in turn get data from two distinct databases.
The communication between APIs will be done by using HTTP and message queuing systems.
All those services will be in turn encapsulated in different virtual machines.

#### General overview

We will setup an e-commerce system, where one API (`inventory`) will take track of the orders and another one (`billing`) will process the payments.

The API gateway will communicate in HTTP with `inventory` and using RabbitMQ for `billing`.

#### API 1: Inventory

#### API 2: Billing

#### The API Gateway

