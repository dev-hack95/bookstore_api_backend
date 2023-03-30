# BOOK-STORE-API
====================================================================


## Authors

- [@dev-hack95](https://www.github.com/dev-hack95)

## Project Status
- Complete

## Table of Contents

  - [Tech Stack](#tech-stack)
  - [Flowchart](#Flowchart)

## Tech Stack
  - Golang, Docker, Mysql, Microservices

## Flowchart
```mermaid
flowchart LR;
    A(PKG) -----> B(config);
    A -----> C(controllers);
    A -----> D(Models);
    A -----> E(routes);
    A -----> F(utils);
    B -----> G(app.go);
    C -----> H(book-controller.go);
    D -----> I(book.go);
    E -----> J(bookstore-routes.go);
    F -----> K(utils.go);
```