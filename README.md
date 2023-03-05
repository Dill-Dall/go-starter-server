Go Starter Server
===========================

This project is a simple but complete example of a GO server built using an API-first approach and oapi-codegen. The server implements a typical petstore case and is set up for Docker Compose, but can easily be modified to work with other container orchestration tools such as Minikube or K3s. The server implements a typical petstore case.

The project's file structure is organized and easy to navigate, with a clear separation between the server's code and the API code. The Makefile provides easy-to-use commands for building the project and generating the API code. Additionally, the README.md file provides clear and concise instructions for setting up and running the project.


Prerequisites
-------------

To run this project, you will need:

*   [Docker](https://www.docker.com/) installed and running.

Getting Started
---------------

To get started with this project, follow these steps:

1.  Run the server using Docker Compose: `docker-compose up`.
2.  Access the server at `http://localhost:3000`.


Building the Project
--------------------

To build the project, simply run `make build` in the project directory. This will build the binary in the current directory.

Generating the API Code
-----------------------

To generate the API code, run `make openapi-codegen`. This will generate the necessary Go files for the server based on the OpenAPI specification file.

Running the Project
-------------------

To run the project, use Docker Compose:

shCopy code

`docker-compose up`

This will start the server and Nginx. You can access the server at `http://localhost:3000`.

Notes
-----

*   The server uses a self-signed certificate for testing purposes. In a production environment, you should use a valid SSL certificate.
*   The project is set up for Docker Compose, but it can easily be modified to work with other container orchestration tools such as Minikube or K3s.
*   Testing testing ...