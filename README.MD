# EchoStorm

## Project Description
This is a sample project implemented with **Echo**, designed for developing and learning FastAPI. Users can take this project and expand it based on their needs.

## How to Run
Follow these steps to run the project:

1. Ensure that **Go** is installed on your system.
2. Navigate to the `cmd` directory.
3. Run the `main.go` file:

```sh
cd src/cmd
go run main.go
```

## API Routes
List of defined routes in this project:

### Products
- **POST `/products`** → Create a new product
- **GET `/products`** → Retrieve the list of products
- **GET `/products/:id`** → Get details of a specific product by ID
- **PUT `/products/:id`** → Update a product
- **DELETE `/products/:id`** → Delete a product

### Test and Other Routes
- **GET `/test`** → A simple test route
- **POST `/name`** → Process input data
- **POST `/bind`** → Check data binding capability

## Dependencies
- **Echo** as the web framework
- **Go Modules** for dependency management

## Additional Notes
- To customize the project, you can modify the API routes and handlers in `src/api/`.
- This project is a simple example and does not include advanced features like authentication or a database, but you can add them as needed.

