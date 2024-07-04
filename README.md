# fase-4-hf-payment

## Summary

* [Requirements](#Requirements)
* [Installation](#Installation)
* [Tests](#Tests)
* [Others Microservices](#Microservices)
* [Documentation](#Documentation)



### Requirements

On Windows:
```
- Windows Subsystem for Linux (WSL);
- GNU Make v4.0 or later (on WSL);
- Docker v20.10.7 or later (on WSL);
- Docker-compose v1.25 or later (on WSL).
```


On Linux/MacOS:
```
- GNU Make v4.0 or later;
- Docker v20.10.7 or later;
- Docker-compose v1.25 or later.
```

### Installation

1. Rename the file `.env.example` to `.env` and setup your environment variables;
2. Run the command below:
```bash
$ make run-terraform
```

### Tests:

Execute the command below:
```bash
$ make run-bdd
```
### Microservices: 

* [Orch](https://github.com/FIAP-SA-Hermes-Foods/fase-4-hf-orch)
* [Voucher](https://github.com/FIAP-SA-Hermes-Foods/fase-4-hf-voucher)
* [Product](https://github.com/FIAP-SA-Hermes-Foods/fase-4-hf-product)
* [Order](https://github.com/FIAP-SA-Hermes-Foods/fase-4-hf-order)
* [Client](https://github.com/FIAP-SA-Hermes-Foods/fase-4-hf-client)

### Documentation

* [Domain Story Telling](https://github.com/FIAP-SA-Hermes-Foods/fiap-hf-storytelling)
* [Ubiquitious Language](https://github.com/FIAP-SA-Hermes-Foods/fiap-hf-ubiquitious-language)
* [Context Map](https://github.com/FIAP-SA-Hermes-Foods/fiap-hf-context-map)
* [Event Storming](https://github.com/FIAP-SA-Hermes-Foods/fiap-hf-event-storming)
* [Project structure](https://github.com/FIAP-SA-Hermes-Foods/fiap-hf-src/tree/main/docs/project_structure.md)
* [Postman Collection](https://github.com/FIAP-SA-Hermes-Foods/fiap-hf-src/blob/main/infrastructure/postman_collection/hermes-foods.postman_collection.json)
* [Contribution Guide](https://github.com/FIAP-SA-Hermes-Foods/fiap-hf-src/tree/main/docs/contribution.md)
