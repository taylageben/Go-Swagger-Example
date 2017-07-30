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

    See the [DB setup](./examples/database/store/db.go) file as an example

3.  Initialize the DB

    Initialize your database in [the main method](./examples/database/cmd/to-do-demo-server/main.go):
 
    ```bash
    //This code has been manually added
    var dataStore store.DataStore
    dataStore.Init()
    ```

    **NB!** Once you have customized Main, any regeneration of the API **MUST** include the `--exclude-main` flag:
    
    ```bash
    swagger generate server --target ../examples/database --name  --spec ../swagger.yml --exclude-main
    ```
    
4.  Add your db functionality
    
    See the code in [the DB handler](./examples/database/restapi/dbHandler/todo.go) and [the configure file](./examples/database/restapi/configure_to_do_demo.go)