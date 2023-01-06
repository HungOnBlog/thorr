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
// Example of a test file for user APIs
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

## Road map

- [x] Support GET, POST, PUT, PATCH, DELETE methods
- [ ] Add more assertions type (e.g. body contains, body not contains, body match regex, ...)
- [ ] Add more test types (e.g. load test, ...)
- [ ] Add more reporters (e.g. HTML, ...)
- [ ] Add more test file formats (e.g. YAML, ...)
- [ ] Add dynamic path (e.g. /user/:id) with customer resource loader (e.g. database, file, aws s3, ...)
- [ ] Build Github Action for Thorr
- [ ] Web UI for Thorr
- [ ] Helm chart for Thorr
