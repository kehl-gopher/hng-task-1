Certainly! Here's a more comprehensive README for your **Number Classification API** project:

---

# Number Classification API

The **Number Classification API** is a simple web service that classifies a number based on various mathematical properties. The API identifies whether a given number is prime, perfect, or Armstrong, as well as whether it is odd or even. Additionally, it provides a fun fact about the number fetched from the Numbers API.

## Table of Contents
1. [Features](#features)
2. [Technology Stack](#technology-stack)
3. [API Endpoint](#api-endpoint)
   - [Request](#request)
   - [Response](#response)
4. [Functionality](#functionality)
5. [Installation](#installation)
6. [Usage](#usage)
7. [Error Handling](#error-handling)
8. [Deployment](#deployment)
9. [Testing](#testing)
10. [Contributing](#contributing)
11. [License](#license)

## Features
- **Prime Number Check**: Determines if the given number is prime.
- **Perfect Number Check**: Determines if the given number is perfect.
- **Armstrong Number Check**: Identifies if the number is an Armstrong number.
- **Odd/Even Classification**: Classifies the number as odd or even.
- **Fun Fact**: Fetches an interesting fun fact about the number from the Numbers API.

## Technology Stack
- **Programming Language**: Go (Golang)
- **API Framework**: `net/http` for handling HTTP requests
- **Fun Fact API**: [Numbers API](http://numbersapi.com)
- **Deployment**: Hosted on **[Your Hosting Platform]** (e.g., Heroku, Render, Koyeb)
- **Version Control**: GitHub (public repository)

## API Endpoint

### `GET /api/classify-number?number={number}`

#### Query Parameters:
- **number** (required): The number to be classified (integer)

#### Request Example:
```bash
GET http://your-domain.com/api/classify-number?number=371
```

### Response Formats:

#### 200 OK Response:
When the request is valid and successfully processed:
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

- `number`: The input number.
- `is_prime`: Boolean value indicating if the number is prime.
- `is_perfect`: Boolean value indicating if the number is perfect.
- `properties`: An array of properties related to the number (e.g., Armstrong, Odd, Even).
- `digit_sum`: The sum of the digits of the number.
- `fun_fact`: A fun fact about the number, fetched from the Numbers API.

#### 400 Bad Request Response:
When the request has an invalid parameter or is not an integer:
```json
{
    "number": "alphabet",
    "error": true
}
```
- `number`: The invalid input.
- `error`: A boolean flag indicating that the request was invalid.

## Functionality

The **Number Classification API** processes a number and determines:
1. **Prime Check**: Determines if a number is prime by checking divisibility from 2 up to the square root of the number.
2. **Perfect Number Check**: Determines if a number is perfect (a number that is equal to the sum of its proper divisors).
3. **Armstrong Number Check**: Checks if a number is Armstrong (a number equal to the sum of its digits raised to the power of the number of digits).
4. **Odd/Even Classification**: Classifies the number as odd or even.
5. **Fun Fact**: Fetches a fun fact about the number from the Numbers API.

The API uses a combination of these checks and returns a comprehensive response in JSON format.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/kehl-gopher/hng-task-1.git
   ```

2. **Install Go (if not installed)**:
   Follow the [Go installation guide](https://golang.org/doc/install) to install Go on your local machine.

3. **Navigate to the project directory**:
   ```bash
   cd number-classification-api
   ```

4. **Run the application**:
   ```bash
   go run main.go
   ```

5. The API will start and be accessible at `http://localhost:8080`.

## Usage

Once the server is running, you can interact with the API by sending GET requests to the `/api/classify-number` endpoint.

### Example:

- **Request**:
  ```bash
  GET https://curly-joelle-hng-ef7ab010.koyeb.app/api/classify-number?number=371
  ```

- **Response**:
  ```json
  {
      "number": 371,
      "is_prime": false,
      "is_perfect": false,
      "properties": ["armstrong", "odd"],
      "digit_sum": 11,
      "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
  }
  ```

## Error Handling

The API handles errors gracefully. If an invalid number or non-integer value is provided in the query string, the API will respond with a **400 Bad Request** status and an appropriate error message.

For example:
- If a string is provided instead of a number:
  ```json
  {
      "number": "alphabet",
      "error": true
  }
  ```

## Deployment

The API is deployed and can be accessed publicly.

To deploy the API:
1. Choose a cloud platform  Koyeb
2. Follow the deployment instructions for the chosen platform.
3. Ensure the application is publicly accessible and can handle GET requests at the `/api/classify-number` endpoint.

## Testing

To test the API:
1. Run the API locally or use the hosted version.
2. Send GET requests to the `/api/classify-number` endpoint with various valid and invalid numbers.
3. Verify that the responses are in the correct JSON format.
4. Check that the API correctly classifies numbers as prime, perfect, Armstrong, odd, or even.

## Contributing

We welcome contributions to improve the **Number Classification API**. If you have suggestions for new features or fixes, feel free to:

1. Fork the repository.
2. Make changes in your forked repository.
3. Create a pull request to submit your changes.
