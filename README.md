# Thorr

![Thorr logo](assets/images/Thorr.gif)

Thorr is a simple, yet powerful, tool for doing integration test and load testing of HTTP APIs. It is written in Go and is designed to be easy to use and extend.

## Installation

### Download the binary

```bash
curl -L
```

### Build from source

```bash
git clone https://github.com/HungOnBlog/thorr.git
cd thorr
go build -o thorr .
```

## Usage

### Define a test

All the test of Thorr are defined in a JSON file and will be put in "thorr" folder. By default, Thorr will look for all files named "*.json" in the "thorr" folder. You can also specify a different file (or folder) name by using the `-f` flag.

```json
# Example of Thorr test configuration file
# thorr/user.json
{
  "name": "user",
  "description": "Test the user service",
  "baseURL": "http://localhost:8080",
  "tests": [
    {
      "name": "Get user",
      "description": "Get a user by id",
      "method": "GET",
      "path": "/user/:id",
      "params": {
        "id": "1"
      },
      "expect": {
        "status": 200,
        "body": {
          "id": 1,
          "name": "Hung"
        }
      }
    }
  ]
}
```
