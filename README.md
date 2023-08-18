## Golang API using go-chi (https://github.com/go-chi/chi)

URL: localhost:8000/api/va/test

## Step 1: Return URL Parameter

Using the endpoint localhost:8000/api/v1/test/{mydata} the API will return {mydata}

### Notes

- In func main() create a new chi router
- Create a get request route that sends back whetever is passed into {mydata} at /api/v1/test/
- Start the server and listen on localhot:8000

## To test this API you can use Postman (https://www.postman.com)