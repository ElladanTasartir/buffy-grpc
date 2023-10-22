# buffy-grpc
A gRPC study using the Buffy API

## Objective
Development of a simple gRPC API that encapsulates the [Buffy the Vampire Slayer and Angel series public API](https://github.com/Thatskat/btvs-angel-api)

## Setup
There are simple steps to get the project running. These are:
- Compile the proto files:
For this one you'll need [docker](https://www.docker.com) installed on your machine.
If you've already got Docker just run the Makefile command:
```sh

make compile-proto-go

```
After this a `gen` folder containing all the compiled protofiles will be created

- Build and Run:
You can build and run the project with the following command:
```sh

make run

```

This project was roughly based on [this article by Alexandre Miziara](https://medium.com/unicoidtech/escrevendo-um-micro-serviço-grpc-em-go-estrutura-definição-e-deploy-no-google-kubernetes-engine-82351e23a2f7)
