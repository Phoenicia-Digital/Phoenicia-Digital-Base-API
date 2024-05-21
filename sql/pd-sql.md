Integrating PostgreSQL and Golang with Dynamic Queries and .sql Files

This document explains how to integrate PostgreSQL and Golang for database interactions using base queries stored in separate .sql files. You'll learn how to incorporate dynamic values extracted from JSON requests received by your Golang API.

##  THIS CAN SIMPLY BE DONE USING THE 
```GO 
PhoeniciaDigitalDatabase.Postgres.ReadSQL(filename)
```
# Where you replace filename with the filename inside ./sql folder ex: myQuery | NO NEED FOR .sql at the end of file name

Scenario:

    You have a Golang API that receives a JSON request containing user data (e.g., ID).
    You want to use a base query stored in a .sql file to fetch data from a PostgreSQL database based on the received user ID.

Steps:

    Define the Base Query in a .sql File:

1) **Create a file named `myQuery.sql` (or any preferred name) and store your base SQL query inside it. Here's an example:**

    ```SQL
    SELECT * FROM users WHERE id = $1;
    ```
This query selects all columns from the users table where the id column matches the provided value (represented by $1). You can use positional or named placeholders here.

2) **Golang API Code:**

Create a Golang file (e.g., main.go) to handle the API logic. Here's a breakdown of the essential steps:

    Define a Struct for JSON Data:

    ``` GO
    type UserRequest struct {
        ID int `json:"id"`
    }
    ```
This struct represents the expected data format of the user ID received in the JSON request.

Parse the JSON Request:

    ```GO
    var request UserRequest
    err := json.Unmarshal(requestBody, &request)
    if err != nil {
        log.Fatal("Error unmarshalling JSON request:", err)
    }
    ```
This code snippet uses the `encoding/json` package to unmarshal the JSON request body into the UserRequest struct.

3) **Connect to the PostgreSQL Database:**
    ```GO
    db, err := sql.Open("postgres", "connection_string")
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }
    defer db.Close() // Close the connection when the function exits
    ```
This code establishes a connection to your PostgreSQL database using the `database/sql` package. Replace `"connection_string"` with your actual connection details.

4) **Read the Base Query from the File:**
    ```GO
    query, err := ioutil.ReadFile("myQuery.sql")
    if err != nil {
        log.Fatal("Error reading base query:", err)
    }
    ```
This block demonstrates how to read the base query directly from the `.sql` file using the `ioutil.ReadFile` function (from io/ioutil).

5) **Execute the Query with Dynamic Value:**
    ```GO
    rows, err := db.Query(query, request.ID)
    if err != nil {
        log.Fatal("Error executing query:", err)
    }
    defer rows.Close() // Close the rows object
    ```
The db.Query function executes the query with the `request.ID` value replacing the `$1` placeholder.