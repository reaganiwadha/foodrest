# foodrest ![GCI Badge](https://img.shields.io/badge/Google%20Code%20In-JBoss%20Community-red?style=flatr&labelColor=fdb900&link=https://codein.withgoogle.com/organizations/jboss-community/)

A simple golang REST server with CRUD functionality for storing food information

## Configuring Port
The port is configured with environment variables. You can also configure the port by using the ```.env``` file

## Starting the server
```
go run main.go
```

## Endpoints
### CreateFoodEndpoint
This endpoint is a post endpoint to create a new food.

```bash
POST http://server/foods
```

Body

```json
{
    "name": "Sandwich",
    "price" : 3.5
}
```
Returns
```
HTTP STATUS 200
```

### AllFoodsEndpoints
This endpoint will return all the foods that is stored in the server.

```bash
GET http://server/foods
```

Returns
```json
[{"ID":0,"Name":"Pasta","Price":5},{"ID":1,"Name":"Burger","Price":3}]
```

### GetFoodByIDEndpoint
This endpoint will return a specific food based on its ID.
```
GET http://server/foods/{id}
```

```
GET http://server/foods/0
```


Returns
```json
{
    "ID": 0,
    "Name": "Pasta",
    "Price": 5
}
```

### UpdateFoodByIDEndpoint
This endpoint is for updating a specific food based on its ID.

```
PUT http://server/foods/{id}
```

```
PUT http://server/foods/0
```

Body
```json
{
    "Name": "Pasta",
    "Price": 6
}
```
Returns
```
HTTP Status 200
```

### DeleteFoodByIDEndpoint
This endpoint will delete a specific food based on its ID.

```
DELETE http://server/foods/{id}
```

```
DELETE http://server/foods/0
```
Returns
```
HTTP Status 200
```