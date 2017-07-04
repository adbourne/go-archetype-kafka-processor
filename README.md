[![Build Status](https://travis-ci.org/adbourne/go-archetype-rest.svg?branch=master)](https://travis-ci.org/adbourne/go-archetype-rest)

# go-archetype-rest
An archetype project in Go making use of Dependency Injection in order to provide a simple REST server.
[gin-gonic](https://gin-gonic.github.io/gin/) is used to provide the REST server functionality.


## Why?
I really like Java + [Spring Boot](https://projects.spring.io/spring-boot/). It's battle tested,
projects are organised and [Dependency Injection](https://martinfowler.com/articles/injection.html) takes place almost
seamlessly. That being said, I really like Go; just the language alone is richly-featured without
the need of a framework, it produces much smaller binaries and it's new(ish) and shiny. So, as Go is pretty
non-prescriptive about what your project looks like, I've mirrored the structure of a typical Spring Boot project. Also,
as Go has the concept of dynamically implementing interfaces, Dependency Injection becomes a pretty easy feat, however,
to organise it I've borrowed Spring's concept of an 'Application Context'.

Is it the "Go way"? Maybe not. Valuable? Remains to be seen.

## Docker
The application can be packaged as a Docker container by using the command:

`make package `

The container is built on [Alpine Linux](https://alpinelinux.org/) and works out to be about 14MB. The container can be
run using the following command:

`docker run -it -p8080:8080 adbourne/go-archetype-rest`


## Development
This project makes use of [dep](https://github.com/golang/dep), which will eventually
be to be Go's [official dependency management tool](https://github.com/golang/go/wiki/PackageManagementTools).
All dependencies are already vendored, however to add new dependencies I recommend taking a look at [deps documentation](https://github.com/golang/dep).
