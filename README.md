go-cleanarchitecture
====================

An example Go application demonstrating The Clean Architecture.

[http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)


Install
-------

Check out the source

    git clone https://github.com/manuelkiessling/go-cleanarchitecture && cd go-cleanarchitecture

Download modules

    go mod download

Then build the project

    go build

Create the SQLite structure

    sqlite3 /var/tmp/production.sqlite < setup.sql

Run the server

    ./go-cleanarchitecture

Access the web endpoint at [http://localhost:8080/orders?userId=40&orderId=60](http://localhost:8080/orders?userId=40&orderId=60)

To run the tests, for each module, run

    go test ./...

Enjoy.


License
-------
The MIT License (MIT)

Copyright (c) 2012 Manuel Kiessling

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
