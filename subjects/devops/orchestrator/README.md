## Orchestrator

![Orchestrator](pictures/Orchestrator.jpg)

### Objectives

In this project, You will deploy a microservices' architecture on Kubernetes,
you will gain experience with key technologies and concepts such as Kubernetes
architecture, deployments, services, ingresses, and API gateways. Additionally,
this project will provide you with an opportunity to practice DevOps skills
such as containerization, continuous integration, and deployment (CI/CD), and
infrastructure as code (_IaC_) using Kubernetes manifests. By completing this
project, you will have a solid understanding of microservices architecture and
the tools and techniques used to deploy and manage such systems using
Kubernetes.

### Tips

- Spend time on the theory before rushing into the practice.
- Read the official documentation.
- You must understand the K8s components.

> Any lack of understanding of the concepts of this project may affect the
> difficulty of future projects, take your time to understand all concepts.

> Be curious and never stop searching!

### Architecture

![Architecture](pictures/Architecture.png)

You have to deploy this microservices' architecture in a K3s cluster
consisting of the following components:

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
  API gateway code running and forwarding the requests to the other
  services, and it's accessible via port `3000`.

<!--TODO: add link to solution-->

> Use the Dockerfiles provided [here](...)

### The cluster

By using K3s in Vagrant you must create two virtual machines:

1. `Master`: the master in the K3s cluster.

2. `Agent`: an agent in the K3s cluster.

You must install `kubectl` on your machine to manage your cluster.

The nodes must be connected and available!

```console
$> kubectl get nodes -A NAME
STATUS   ROLES    AGE    VERSION
<master-node>   Ready    <none>   XdXh   vX
<agent1-node>   Ready    <none>   XdXh   vX
$>
```

You must provide a `orchestrator.sh` script that run and create and manage the
infrastructure:

```console
$> ./orchestrator.sh create
cluster created
$> ./orchestrator.sh start
cluster started
$> ./orchestrator.sh stop cluster stopped $>
```

### Docker Hub

You will need to push the Docker images for each component to Docker Hub.

> You will use it in your Kubernetes manifests.

![Docker Hub example](pictures/dockerhub-example.jpg)

### Manifests

You should create a YAML Manifest that describes each component or resource of
your deployment.

### Secrets

You must store your passwords and credentials as a K8s secrets.

> It's forbidden to put your passwords and credentials in the YAML manifests,
> except the secret manifests!

### Applications deployment instructions

The following applications must be deployed as a deployment, and they
must be scaled horizontally automatically, depending on CPU consumption.

- `api-gateway`: max replication: 3 min replication: 1 CPU percent trigger: 60%

- `inventory-app`: max replication: 3 min replication: 1 CPU percent trigger:
  60%

The `billing-app` must be deployed as _StatefulSet_.

### Databases

Your databases must be deployed as _StatefulSet_ in your K3s cluster, and you
must create volumes that enable containers to move across infrastructure
without losing the data.

### Documentation

You must push a `README.md` file containing full documentation of your solution
(prerequisites, configuration, setup, usage, ...).

### Bonus

If you complete the mandatory part successfully, and you still have free time,
you can implement anything that you feel deserves to be a bonus, for example:

- Use the `Dockerfile` you have defined in your solution for
  `play-with-containers`

- Deploy a Kubernetes Dashboard to monitor the cluster

- Deploy a dashboard for applications logs

- Kubernetes in cloud ?!

Challenge yourself!

### Submission and audit

You must submit the `README.md` file and all files used to create and delete
and manage your infrastructure: Vagrantfile, Dockerfiles, Manifests, ...

```console
.
├── Manifests
│   └── [...]
├── Scripts
│   └── [...]
├── Dockerfiles
│   └── [...]
└── Vagrantfile
```

If you decide to use a different structure for your project remember you should
be able to explain and justify your decision during the audit.

> In the audit you will be asked different questions about the concepts and the
> practice of this project, prepare yourself!

#### What's next?

In order to develop your knowledge and career as a DevOps engineer, we highly
recommend you to learn and practice more about Kubernetes and even get a
certification for Kubernetes.

[https://kubernetes.io/training/](https://kubernetes.io/training/)
