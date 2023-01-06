# Thorr

![Thorr logo](assets/images/Thorr.gif)

Thorr is a simple, yet powerful, tool for doing integration test and load testing of HTTP APIs. It is written in Go and is designed to be easy to use and extend.

## Key concepts

- **[Test Suit](#test-suit)** A test suit is a collection of tests. It is a JSON file that contains the test cases.
- **[Test](#test-test-case)** A test is a single test case. It contains the request and the assertions.
- **[Assertion](#assertions)**: An assertion is a condition that you want to check. It can be the status code, the response body, the response headers, ...

## Test suit

Test suit is a collection of test cases. A test suit will be a JSON file that contains the test cases with the following format:

```json
{
  "name": "Test suit name",
  "description": "Test suit description",
  "default": {
    "base_url": "http://localhost:8080",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "tests": [],
}
```

## Test (Test case)

Test is a specific case (request) you want to test. It will be defined in the test suit file with the following format:

```json
{
  "name": "Test name",
  "description": "Test description",
  "request": {
    "method": "POST",
    "path": "/",
    "headers": {
      "apikey": "apikey"
    },
    "queries": {
      "key": "value"
    },
    "body": {
      "key": "value",
      "key2": 1
    }
  },
  "assertions": [
    {
      "on": "status_code",
      "check": "exact",
      "expected": 200
    },
    {
      "on": "body",
      "check": "type",
      "expected": {
        "key": "string",
        "key2": "string"
      }
    },
    {
      "on": "headers",
      "check": "exist",
      "expected": {
        "apikey": true,
        "Content-type": true
      }
    }
  ]
}
```

## Assertions

Assertions are the conditions that you want to check. It can be the status code, the response body, the response headers, ...

### Assertion `on`

- **status_code**: Check the status code of the response
- **body**: Check the body of the response
- **headers**: Check the headers of the response

### Assertion `check`

- **exact**: Check the value is exactly the same as the expected value & type
- **type**: Check the type of the value is the same as the expected type
- **exist**: Check the value exists in the response

For assertion Thorr will automatically flatten nested object so you can also assert for the nested object by using dot notation (e.g. `body.key`).

#### Data type

For assertion `type` Thorr will automatically flatten nested object so you can also assert for the nested object by using dot notation (e.g. `body.key`).

##### Simple

- **string**: Check the value is a string
- **int**: Check the value is an integer
- **double**: Check the value is a double
- **boolean**: Check the value is a boolean
- **date**: Check the value is a valid date (e.g. 2020-01-01)
- **datetime**: Check the value is a valid datetime (e.g. 2020-01-01T00:00:00Z)
- **time**: Check the value is a valid time (e.g. 00:00:00)

##### Advanced

String advanced type assertion:

- **string::email**: Check the value is a valid email
- **string::length::min-max**: Check the length of the string is between min and max (e.g. `string::length::1-10`)
- **string::uuid**: Check the value is a valid UUID
- **string::url**: Check the value is a valid URL
- **string::base64**: Check the value is a valid base64 string

Int advanced type assertion:

- **int::range::min-max**: Check the value is between min and max (e.g. `int::range::1-10`)
- **int::(lt|lte|gt|gte)::value**: Check the value is less than, less than or equal, greater than, greater than or equal to the value (e.g. `int::lt::10`)

Double advanced type assertion:

- **double::range::min-max**: Check the value is between min and max (e.g. `double::range::1-10`)
- **double::(lt|lte|gt|gte)::value**: Check the value is less than, less than or equal, greater than, greater than or equal to the value (e.g. `double::lt::10`)

Date advanced type assertion:

- **date::(lt|lte|gt|gte)::value**: Check the value is less than, less than or equal, greater than, greater than or equal to the value (e.g. `date::lt::2020-01-01`)
- **date::(before|after)::value**: Check the value is before or after the value (e.g. `date::before::2020-01-01`)
- **date::range::min to max**: Check the value is between min and max (e.g. `date::range::2020-01-01 to 2020-01-10`)
- **date::format::format**: Check the value is in the format (e.g. `date::format::YYYY-MM-DD`)
- **date::format::utc**: Check the value is in UTC format (e.g. `date::format::utc`)

Time advanced type assertion:

- **time::(lt|lte|gt|gte)::value**: Check the value is less than, less than or equal, greater than, greater than or equal to the value (e.g. `time::lt::00:00:00`)
- **time::(before|after)::value**: Check the value is before or after the value (e.g. `time::before::00:00:00`)
- **time::range::min to max**: Check the value is between min and max (e.g. `time::range::00:00:00 to 00:00:10`)
- **time::format::format**: Check the value is in the format (e.g. `time::format::HH:mm:ss`)

### Assertion `expected`

The expected value is the value that you want to check. It can be a string, a number, a boolean, an object, ...

```json
{
  "assertions": [
    {
      "on": "status_code",
      "check": "exact",
      "expected": 200
    },
    {
      "on": "body",
      "check": "type",
      "expected": {
        "key": "string",
        "key2": "string",
        "key": {
          "key3": "string"
        },
      }
    },
    {
      "on": "headers",
      "check": "exist",
      "expected": {
        "apikey": true,
        "Content-type": true
      }
    }
  ]
}
```

## Example of a test suit definition

```json
{
  "name": "Test suit template",
  "description": "This is a template for test suit. Tests in a in a suit are executed in sequence. If one test fails, it will be skipped and the next test will be executed.",
  "type": "integration",
  "base_url": "https://postman-echo.com",
  "default_header": {
    "Content-Type": "application/json"
  },
  "tests": [
    {
      "name": "Test name",
      "description": "Test description",
      "request": {
        "method": "POST",
        "path": "/",
        "headers": {
          "apikey": "apikey"
        },
        "queries": {
          "key": "value"
        },
        "body": {
          "key": "value",
          "key2": 1
        }
      },
      "assertions": [
        {
          "on": "status_code",
          "check": "exact",
          "expected": 200
        },
        {
          "on": "body",
          "check": "type",
          "expected": {
            "key": "string",
            "key2": "string"
          }
        },
        {
          "on": "headers",
          "check": "exist",
          "expected": {
            "apikey": true,
            "Content-type": true
          }
        }
      ]
    }
  ]
}
```

## Road map

- [ ] Implement the integration test
- [ ] Implement the load test
- [ ] Support GET, POST, PUT, PATCH, DELETE methods
- [ ] Add more assertions type (e.g. body contains, body not contains, body match regex, ...)
- [ ] Add more test types (e.g. load test, ...)
- [ ] Add more reporters (e.g. HTML, ...)
- [ ] Add more test file formats (e.g. YAML, ...)
- [ ] Add dynamic path (e.g. /user/:id) with customer resource loader (e.g. database, file, aws s3, ...)
- [ ] Build Github Action for Thorr
- [ ] Web UI for Thorr
- [ ] Helm chart for Thorr

## Contributing

I would love to see your ideas for Thorr. Please feel free to open an issue or a pull request. 🥰🥰

You should follow the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) convention when writing commit messages.

## License

Thorr is licensed under the MIT License. See [LICENSE](LICENSE) for the full license text.
