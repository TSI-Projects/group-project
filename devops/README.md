# Prerequisites

Please be sure you have the following utilities installed on your host machine:
* `python` - Required to run deployment scripts. You can download and install it from the [official website.](https://www.python.org/downloads/)
* `kubectl` - A command-line tool used to manage Kubernetes clusters (Docker Desktop, MiniKube, etc.).
* `helm` - Used to deploy the application. You can install Helm by following the instructions on the [official Helm website](https://helm.sh/docs/intro/install/).
* `docker` - Required to build and manage container images. You can download and install Docker Desktop from [Docker's official website.](https://docker.com)

# Dependencies

First, you’ll need to set up the required Kubernetes dependencies (`Nginx Ingress Controller` and `CloudNativePG`), navigate to the [`deploy-scripts`](./deploy-scripts/) directory and run the `install_dependencies.py` script by executing the following command:

```bash
python install_dependencies.py
```

# Server Docker Image Build

To create the server's Docker image, execute the `build_docker_image.py` located in the [`backend`](../backend) directory:

```bash
python build_docker_image.py
```

After the build completes, the image labeled `simple_go_server` will appear in the Docker Desktop GUI or listed in the terminal when running the `docker image ls` command.

# Local Deploy

## Deploy
Once the dependencies are installed and the server image is built, you can deploy and test the application locally. Navigate to the [`deploy-scripts`](./deploy-scripts/) directory and execute `deploy_local.py` script by running the following command:

```bash
python deploy_local.py
```

If deployment was successful you can now be able to access backend server by visiting this url [http://demo.localdev.me](http://demo.localdev.me).

## Destroy

If you want to destroy application, navigate to the [`deploy-scripts`](./deploy-scripts/) directory and execute the `destroy_local.py` script by running the following command:

```bash
python destroy_local.py
```