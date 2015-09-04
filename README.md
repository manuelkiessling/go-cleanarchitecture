go-cleanarchitecture
====================

An example Go application demonstrating The Clean Architecture.

[http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)

Install
-------

Create a new directory, where to host this project

    mkdir -p $GOPATH:src/github.com/manuelkiessling/

Check out the source

    cd $GOPATH:src/github.com/manuelkiessling/
    git clone https://github.com/manuelkiessling/go-cleanarchitecture

Setup the GOPATH to include this path

    cd go-cleanarchitecture
    export GOPATH=$GOPATH:`pwd`

Then build the project

    go get
    go build

Create the SQLite structure

    sqlite3 /var/tmp/production.sqlite < setup.sql

Run the server

    go-cleanarchitecture

Access the web endpoint at [http://localhost:8080/orders?userId=40&orderId=60](http://localhost:8080/orders?userId=40&orderId=60)

To run the tests, for each module, run

    cd src/infrastructure &&  go test
    cd src/interfaces && go test

Enjoy

