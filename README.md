## Golang API using go-chi (https://github.com/go-chi/chi)

Base URL: localhost:8000/api/v1
End points: 
/test
/puzzle
/pokemon
/word

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

## Step 4: Implement Databases and aditional API calls

New endpoints: 

/api/v1/pokemon 
This endpoint uses the PokeAPI (https://pokeapi.co/). If you use this endpoint as a POST it will GET a pokemon from the pokeAPI and store it in BoltDB database. If you use this route as a GET it will respond with all the pokemon in the BoltDB database.

/api/v1/word
This endpoint uses SQLite. If you use this endpoint as POST and parse JSON into the body such as "word": "{your word}" it will store that word in the SQL database. If you use this endpoint as a GET it will send back all the words stored in the database.

### Notes

- Install BoltDB (github.com/boltdb/bolt)
- Create BoltDB file in main func
- Make POST route that takes a random pokemon from the PokeAPI and store it in BoltDB
- Make GET route that returns all the pokemon stored in the BoltDB
- Install SQLite (github.com/mattn/go-sqlite3)
- Create POST endpoint for word
- Post a word that is parsed into the body as JSON into SQL database
- Create GET endpoint for word
- Return all words in SQL database

## To test this API you can use Postman (https://www.postman.com)

Example Post request:

![Imgur](https://i.imgur.com/IH15S0h.png)
