# FizzBuzzServer

Dedicated web server to compute the Fizzbuzz series by calling the server with appropriate parameters.
You can also retrieve what was the most asked request to the server.

## Prerequisites

To work on this project, you should have Golang and Docker installed on your host

## Build

Once you pulled the project, you can enter it and create a docker image

    docker build . -t fizzbuzz:latest

## Run and use

Now you can launch a container to start the server

    docker run -i -t --publish 8080:8080 --name fizzbuzz --rm fizzbuzz:latest

## Endpoints

- GET /fizzbuzz/v1/produce

Accepts five URL parameters : three integers int1, int2 and limit, and two strings str1 and str2.
Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1,
all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

Execute :
    curl -X GET 'http://localhost:8080/fizzbuzz/v1/produce?int1=5&int2=4&limit=50&str1=Fizz&str2=Buzz'
or just put the URL in your browser:
    http://localhost:8080/fizzbuzz/v1/produce?int1=5&int2=4&limit=50&str1=Fizz&str2=Buzz

- GET /fizzbuzz/v1/stats

No parameters. It allows users to know what the most frequent request has been.
Returns the parameters corresponding to the most used request, as well as the number of hits for this request.
Execute :
    curl -X GET 'http://localhost:8080/fizzbuzz/v1/stats'
or just put the URL in your browser:
    http://localhost:8080/fizzbuzz/v1/stats


## Configure

You can modify the server configuration with environment variables. All the config tokens can be
overridden via environment variables. The variables' name follows a simple pattern:
`FIZZBUZZ_VALUE`.

The available variables are:
* FIZZBUZZ_BINDIP : Rest server IP address (default 0.0.0.0)
* FIZZBUZZ_BINDPORT : Rest server access port (default 8080)

For example:

    docker run -i -t -e FIZZBUZZ_RESTSERVER_BINDPORT='6060' --publish 6060:6060 --name fizzbuzz --rm fizzbuzz:latest

## Unit tests

There are some unit tests in the projects, they run each time you build the image.
You can also run them independently at project root :

    go test ./...
