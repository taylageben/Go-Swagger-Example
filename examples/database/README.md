# Steps to Adding Database Functionality

1. Adding 'db' tag to struct

    To add a db tag to the struct, add the `x-go-custom-tag` to each property of a definition:
    
    ```yaml
    properties:
          name:
            type: string
            x-go-custom-tag: db:"name"
          completed:
            type:
              - boolean
              - 'null'
            x-go-custom-tag: db:"completed"
    ```

2. Add database connection and initialization functions

    See the [DB setup](./store/db.go) file as an example

3.  Initialize the DB

    Initialize your database in [the main method](./cmd/to-do-demo-server/main.go):
 
    ```go
    //This code has been manually added
    var dataStore store.DataStore
    dataStore.Init()
    ```

    **NB!** Once you have customized Main, any regeneration of the API **MUST** include the `--exclude-main` flag:
    
    ```bash
    swagger generate server --target .. --name  --spec ../swagger.yml --exclude-main
    ```
    
4.  Add your db functionality
    
    See the code in [the DB handler](./restapi/dbHandler/todo.go) and [the configure file](./restapi/configure_to_do_demo.go)
    
# Example Responses

## Post Todo

```bash
$ curl -H "Content-Type: application/json" -X POST -d '{"name":"test","completed":false}' http://localhost:8080/todos?apikey=123


{"name":"test","completed_at":"0001-01-01T00:00:00.000Z","created_at":"2017-07-30T21:56:55.192+02:00","id":1,"updated_at":"2017-07-30T21:56:55.192+02:00"}
```

## Get Todo

```bash
$ curl http://localhost:8080/todos/1


{"name":"test","completed_at":"0001-01-01T00:00:00.000Z","created_at":"2017-07-30T21:56:55.192+02:00","id":1,"updated_at":"2017-07-30T21:56:55.192+02:00"}
```