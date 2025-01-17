# GMicro 

<a href="https://github.com/superdecimal/gmicro/actions">
  <img src="https://github.com/superdecimal/gmicro/workflows/CI/badge.svg?style=flat" alt="Actions Status">
</a>
<a href="https://github.com/superdecimal/gmicro/commits/master">
  <img src="https://img.shields.io/github/last-commit/superdecimal/gmicro.svg?style=flat&logo=github&logoColor=white"
alt="GitHub last commit">
</a>
<a href="https://github.com/superdecimal/gmicro/issues">
  <img src="https://img.shields.io/github/issues-raw/superdecimal/gmicro.svg?style=flat&logo=github&logoColor=white"
alt="GitHub issues">
</a>
<a href="https://github.com/superdecimal/gmicro/pulls">
  <img src="https://img.shields.io/github/issues-pr-raw/superdecimal/gmicro.svg?style=flat&logo=github&logoColor=white" alt="GitHub pull requests">
</a>
<a href="https://github.com/superdecimal/gmicro/blob/master/LICENSE">
  <img src="https://img.shields.io/github/license/superdecimal/gmicro.svg?style=flat" alt="License Status">
</a>


**GMicro is a calculator service built with Go and gRPC**

**Requirements**
* Go 1.14+ with modules enabled
* Make
* Docker
* Helm v3
* Minikube

## ToC
* [Usage](docs/usage.md)
* [12 Factor](docs/12f.md)
* [Deployment](docs/deployment.md)
* [Further Work](docs/further_work.md)


## Repo Structure
This is an opinionated non-standard repo structure.

* `.github/workflows` - ci files
* `deploy` - deployment files
* `docs` - documentation files
* `pkg`  - code to be shared by multiple services
* `pkg/proto` - generated code from protobuf
* `proto` - protobuf files
* `services` - code for our services

## Makefile 

### Targets

Install dev environment tools
```
make dev-env
```
Build all Dockerfiles
```
make build-all -B
```
Generate mocks
```
make generate-mocks
```
Generate code from proto files
```
make generate-proto
```