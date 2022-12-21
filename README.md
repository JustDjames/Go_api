# Go Api

This is a basic API written in Go. The plan is to connect it to a RDS instance and have it perform CRUD (Create, Read, Edit and Delete) on a table of users. This project also contains pulumi code to create the database and all other required resources in AWS.

## API

### plan for the API endpoints

the API should have the following endpoints that use the specified HTTP methods.

`/` (GET) - displays a help message with the a description of what the other endpoints do and what methods

`/users` (GET) - returns the full list of users. want it to be returned as json

`/user/<id>` (GET) - returns the details of a specific user. the RDS should be using a random id, instead of being incremental

`/user/<id>` (PUT) - updates the specified user. the new info needs to be in JSON

`/user/<id>` (DELETE) - delete the specified user

`/newuser` (POST) - add a new user to the users list. the used needs to be provided in JSON

### Arguments for API

below is a table with the arguments you can use to configure the API:

|   Argument  | Environment_var |                     Description                    | Default |
|:-----------:|:---------------:|:--------------------------------------------------:|:-------:|
| db_hostname |   DB_HOSTNAME   |            The hostname of the database            |    -    |
|   db_pass   |     DB_PASS     | The password for the user used connect to database |     -   |
|   db_user   |     DB_USER     |        The user used to connect to database        |    -    |
|   db_port   |     DB_PORT     |                  The database port                 |    -    |
|   db_name   |     DB_NAME     |              The name of the database              |    -    |
|   db_table  |     DB_TABLE    |     The name of the table used in the database     | db_name |
|   api_port  |     API_PORT    |         The port the api will listening on         |  "8080" |

## Pulumi

### pulumi config

In the pulumi directory, create a file called `Pulumi.dev.yaml`. in that file you can specify your aws config and some sensitive values which are required for the stack like your ip address. below is a example of the file layout:

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

### Creating Resources

when configured, run `pulumi up` in the `pulumi` directory and confirm that you want to create the resources. Once the resources have been created you can confirm that you are able to access the RDS database through a mysql docker container:

```bash
docker run -it --rm mysql mysql -h <database hostname> -D users -u root -P 3306 -p 
```

If the mysql container can connect, you will prompted you for the password for the user. Enter it and press enter and you should be connected to the database.

### Deleting Resources

When you are done, you can destroy the resources using `pulumi down`. Confirm that you want to destory the resources and pulumi will destroy them for you.
