# <h1 align = "center"> Inventory And Orders Management </h1>

---
<p align = "center"> 
<a href = "https://go.dev">  
    <img alt = "Go" src="https://img.shields.io/badge/Go-1.19-blue.svg">
</a>
<a href = "https://docs.gofiber.io">  
    <img alt = "Go Fiber" src="https://img.shields.io/badge/Go Fiber-v2-green.svg">
</a>
<a href = "https://gorm.io/index.html">  
    <img alt = "Go Fiber" src="https://img.shields.io/badge/GORM-grey.svg">
</a>
</p>

---
<p align="left">

## Overview

This project is a CRUD (Create, Read, Update, Delete) application for managing inventory and orders . It provides a backend API built with  Go Fiber, uses Postgres as the database.

## Technologies Used

- **Framework:** Go Fiber
- **Language:** Golang
- **ORM Tool:** GORM
- **Database:** PostgreSQL
---

## Project Structure
```
Inventory_Order_Management/
├── database/
│   └── connect.go
├── controller/
│   └── userController.go
│   └── orderController.go
│   └── productController.go
├── routes/
│   └── userRoutes.go
│   └── orderRoutes.go
│   └── productRoutes.go
├── models/
│   └── user.go
│   └── product.go
│   └── order.go
├── go.mod
├── main.go
```
---

## Database

- The database package initializes the database connection and registers the models, orders and products. You can also auto migrate your models orders and products in database so whenever you change any model , you don't have to change tabels in database the auto migrate takes care of that. The global DB variable is initialized in connect.go.
---

## Contorller

- The controller package contains all the business logic. it is where all the responses goes back to the client and we it contains methods like CREATE, READ, UPDATE and DELETE these methods are called by routes package,
controller interacts with database to perform crud operations.
---

## Routes

- The routes package gets the objects of fiber(app) and it manages the API calls,the app contains Get, Put, Post, Delete methods associated with it which take end point and callback function as a parameter. 
---

## Model

- The model package contains entities, these are structs and this is the representaion of how the tables in the database will be created.
---

## Main.go 🚀

- The main.go is in the main package and it contains the main function which is the starting point of our program, it calls the connect.go to initilize the database and it creates fiber object and passes the app to routes.

<div style="border-bottom: 2px solid grey;"></div>


## API Documentation

1. User Endpoints:
```
- Post: /user/profile
- Get: /user/profiles
- Delete: /user/profile/:id
```

2. Product Endpoints:
```
- Post: /admin/product
- Get: /admin/products, /admin/product/:id
- Put: /admin/product/:id
- Delete: /user/profile/:id
```

3. Order Endpoints:
```
- Post: /user/order
- Get: /user/orders, /user/order/:id
- Delete: /user/order/:id
```


