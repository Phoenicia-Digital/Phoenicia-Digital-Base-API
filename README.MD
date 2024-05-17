# Multiplexer Handling

The `server.go` file located in `~/source/server/server.go` contains the initialization function that stores all the functions via the multiplexer.

```go
multiplexer.Handle("<METHOD> <ROUTE>", PhoeniciaDigitalUtils.PhoeniciaDigitalHandler(<FUNCTION>))
```

- The `<METHOD>` tag stores the method if needed ('GET', 'DELETE', 'PUT', 'POST').
- The `<ROUTE>` tag stores the route where the function will be handled ('/', '/home', '/home/{id}').
- The `<FUNCTION>` tag stores the function itself, which will be handled on the specified route. This should be programmed in a folder inside the source folder `~/source/myfunctions/functions.go`.

**Keep in mind the tags should not contain `< >`:**
```go
multiplexer.Handle("GET /", PhoeniciaDigitalUtils.PhoeniciaDigitalHandler(MYFUNC))
```

## The `PhoeniciaDigitalUtils.PhoeniciaDigitalHandler()` Function

This function handles the logic of the function type `PhoeniciaDigitalHandler` which returns an error.

### Why Not Use the Base `http.HandleFunc`?

We don't use the base `http.HandleFunc` due to our error handling system in place:

- Our functions return an error by just returning a `ApiError` Struct.
- In case of an error, you can return the exact error to the frontend with the `ApiError` struct.
- Similarly, In case of no error you would return a `SendJSON` function will filled in parameters.

#### Middleware

Our middleware capabilities remain largely consistent with normal middleware functionality, with the following exception:

```
func Test(w http.ResponseWriter, r *http.Request) error {
    return PhoeniciaDigitalUtils.ApiError{Code: http.StatusAccepted, ErrorQuote: "This is the Test Function"}
}

func TestMiddleware(next PhoeniciaDigitalUtils.PhoeniciaDigitalHandler) PhoeniciaDigitalUtils.PhoeniciaDigitalHandler {
    return func(w http.ResponseWriter, r *http.Request) error {
        log.Println("This Passed Through The Middleware")
        in := 0
        if in == 0 {
            return next(w, r)
        } else {
            return PhoeniciaDigitalUtils.ApiError{Code: http.StatusBadGateway, ErrorQuote: "This Stoped At The Else Block"}
        }
    }
}
```

This middleware example demonstrates how middleware can be used to intercept and process requests before passing them on to the next handler in the chain.

Feel free to customize and extend this middleware pattern as needed for your application.

##### Instructions for Use:

- Define your handler functions in separate files within the myfunctions folder located inside the source folder.

- Implement the logic for each handler function, ensuring they adhere to the PhoeniciaDigitalHandler function signature.

- In your server.go file, initialize the multiplexer and register your handler functions using the Handle method, specifying the HTTP method, route, and corresponding function.

- Optionally, utilize middleware functions to intercept and process requests before they reach the designated handler function.

- Compile and run your server to start handling incoming HTTP requests.

By following these steps, you can effectively manage your HTTP request handling and middleware logic within your Go application.