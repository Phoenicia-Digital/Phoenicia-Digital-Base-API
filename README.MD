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

- Our functions return an error with the help of the `PhoeniciaDigitalUtils.SendJSON` function.
- In case of an error, you can return the exact error to the frontend with the `ApiError` struct.
- Similarly, in case of success with no data, you can return the success code and quote with the `ApiSuccess` struct.
