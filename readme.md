# Setup

1. git clone the project
2. run docker compose up command

# Config

1. Edit the docker-compose.yaml file if you use Windows or Mac with Apple Silicon.

Windows

```
image: "mysql:latest"
```

or

Mac

```
image: "arm64v8/mysql:latest"
```

# Todo
1. Proper token checking for each resource endpoint
2. Code should be in dockerhub instead of building it locally

# Test
1. You can test the endpoint by using app.http file that is included.
2. No unit test yet :(


# Build

Inside the web folder, you can refer the build step in web/Dockerfile.

# Credits
1. Some of the code is taken from https://github.com/notblessy/go-ingnerd
2. Facebook Login https://www.codershood.info/2020/04/16/facebook-login-in-golang-tutorial/