# insights-service in Golang

Gateway for Insights services

### Installation

### Requirements

1. [Go](https://golang.org/doc/install) ( = v1.14 )

1) Clone the repository to your local machine and install required dependecies

> ```shell
> git clone git@github.com:medwing/insights-service.git
> cd insights-service
> ```

### Create Database
```shell
make createdb
```

### Run Migrations
```shell
make migrateup
```

### Generate schema
> ```shell
> go run github.com/99designs/gqlgen generate
> ```

### To start the server

> ```shell
> go run server.go
> ```

### Development Scripts
```shell
  go help ## Get go commands
  go fmt  ## Format files
  go test  ## Run tests
  go get package_name  ## To add new project Dependency
```

### Git-workflow

**Terminal**

1. Create feature branch upon development branch
2. Commit your changes
3. Check if the development branch has changes. If yes rebase.
4. Push your branch to the remote (Github)

**Github**

1. Create merge-request(PR) (Github)
2. Request a team-member to review your PR (Github)

**Example of the git-workflow**

```shell
git checkout -b my-feature ## create my-feature branch
git add . ## stage changes
git commit  ## commit changes
git checkout development ## checkout the branch upon you created my-feature branch
git pull ## pull changes
git checkout my-feature
git rebase development # rebase if there were changes to master-branch
git push origin -u my-feature # push to the remote
```