# Bank Backend System

The Bank Backend provides APIs for managing bank accounts, recording balance changes, and facilitating money transfers.

## Features

- Create and manage bank accounts with owner’s name, balance, and currency.
- Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.
- Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.

## Setup local development

Follow these steps to install the required tools on Ubuntu:

### Docker

```bash
sudo apt update
sudo apt install docker.io
sudo systemctl enable --now docker
```

### TablePlus

Download TablePlus from the official website: [TablePlus Download](https://tableplus.com/download)

### Golang

```bash
sudo apt update
sudo apt install golang
```

### Migrate

```bash
sudo apt update
sudo apt install -y golang-migrate
```

### DB Docs

```bash
sudo npm install -g dbdocs
dbdocs login
```

### DBML CLI

```bash
sudo npm install -g @dbml/cli
dbml2sql --version
```

### Sqlc

```bash
sudo apt update
sudo apt install sqlc
```

### Gomock

```bash
go install github.com/golang/mock/mockgen@v1.6.0
```

## Setup infrastructure

Follow these steps to set up the necessary infrastructure for the Bank Backend System:

### Create the Bank Network

```bash
make network
```

### Start Postgres Container

```bash
make postgres
```

### Create the Simple_Bank Database

```bash
make createdb
```

### Run Database Migration Up

```bash
make migrateup
```

### Run Database Migration Down

```bash
make migratedown
```
## How to Generate Code

Follow these steps to generate code for the Bank Backend System:

### Generate Schema SQL File with DBML

```bash
make db_schema
```

### Generate SQL CRUD with sqlc

```bash
make sqlc
```

### Generate DB Mock with Gomock

```bash
make mock
```

### Create a New Database Migration

```bash
make new_migration name=<migration_name>
```
## How to Run

Follow these steps to run the Bank Backend System:

### Run Server

```bash
make server
```

### Run Tests

```bash
make test
```

## Deploy to Kubernetes Cluster

Follow these steps to deploy the Bank Backend System to a Kubernetes cluster:

### Install NGINX Ingress Controller

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.48.1/deploy/static/provider/aws/deploy.yaml
```

### Install Cert-Manager

```bash
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml
```
