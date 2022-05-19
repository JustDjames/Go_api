# Go Api

This is a basic API written in Go. The plan is to connect it to a RDS instance and have it perform CRUD (Create, Read, Edit and Delete) on a table of users.

## plan for the API endpoints

the API should have the following endpoints that use the specified HTTP methods.

`/` (GET) - displays a help message with the a description of what the other endpoints do and what methods

`/users` (GET) - returns the full list of users. want it to be returned as json

`/user/<id>` (GET) - returns the details of a specific user. the RDS should be using a random id, instead of being incremental

`/user/<id>` (PUT) - updates the specified user. the new info needs to be in JSON

`/user/<id>` (DELETE) - delete the specified user

`/newuser` (POST) - add a new user to the users list. the used needs to be provided in JSON

## pulumi config

in the pulumi directory, create a file called `Pulumi.dev.yaml`. in that file you can specify your aws config and some sensitive values which are required for the stack like your ip address. below is a example of the file layout:

```yaml
config:
  aws:profile: default
  aws:region: eu-west-2
  Go_api:my_ip: <ip>
```

you will also need to add the database password to this file. You do this by running the following command:

```bash
pulumi config set --secret db_pass <password_value>
```

This will add the encrypted value to the `Pulumi.dev.yaml`. you can retrieve this value by running:

```bash
pulumi config get db_pass
```

