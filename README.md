## Golang API using go-chi (https://github.com/go-chi/chi)

URL: localhost:8000/api/va/test

## Step 1: Return URL Parameter

Using the endpoint localhost:8000/api/v1/test/{mydata} the API will return {mydata}

### Notes

- In func main() create a new chi router
- Create a get request route that sends back whetever is passed into {mydata} at /api/v1/test/
- Start the server and listen on localhot:8000

## Step 2: Middleware

Create middleware that saves the URL parameter to context, then sends that value in the response

### Notes

- Modify r.Get to r.With({middlewareName}).Get
- In the middleware get the URL parameter value and add it to the context
- In the get route, get the context and then the value from the context
- Send back that value

## Step 3: Puzzle

Create another route for converting roman numerals to their integer values.
Example Route: (localhost:8000/api/v1/puzzle?numeral=V)

### Notes

- Create the route in main.go 
- This route uses a query parameter instead of a path parameter eg (?numeral={value})
- This route takes the value of the query and sets it to an integer value. Then loop through the characters and add the values together.
- Send back a message with the changed value
- Create main_test.go for testing
- Created test cases to check multiple values using uppercase and lowercase


## To test this API you can use Postman (https://www.postman.com)