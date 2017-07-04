FROM alpine:3.6
MAINTAINER Aaron Bourne <contact@aaronbourne.co.uk>

# Make gin-gonic run in production mode
ENV GIN_MODE "release"

# Specify the port the application is to run on
ENV REST_ARCHETYPE_PORT 8080

# Specify the random seed the application is to use
ENV REST_ARCHETYPE_RANDOM_SEED 1

# Create an app user to run the application as
RUN addgroup -S app && adduser -S -g app app

# Add the built application
ADD bin/app /

# Run as the app user
USER app

# Run the app
ENTRYPOINT ["/app"]