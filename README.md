## A RESTful web service API with Go and the Gin Web Framework

Gin simplifies many coding tasks associated with building web applications, including web services.

In this project, you'll use Gin to route requests, retrieve request details, and marshal JSON for responses.

### Design API endpoints

You'll build an API that provides access to a store selling vintage recordings on vinyl. So you'll need to provide endpoints through which a client can get and add albums for users.

/albums
- GET - Get a list of all albums, returned as JSON
- POST - Add a new album from request data sent as JSON

/albums/:id
- GET - Get an album by its ID, returning the album data as JSON

### Data

To keep things simple for this project, you'll store data in memory. A more typical API would interact with a database.

### Todo

- Use a database to store the data
- Use docker
