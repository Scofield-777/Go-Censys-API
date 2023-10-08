# Go-Censys-API
# Go Script for Censys API Interaction

This Go script provides a simple and convenient way to interact with the Censys API, allowing you to perform various operations like certificate search, IPv4 search, and more. This README provides a step-by-step guide on how to use the script effectively.

## Table of Contents

- [Installation](#installation)
- [Authentication](#authentication)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

-Before using the script, ensure that you have Go installed on your system. You can download and install Go from the [official website](https://golang.org/dl/).

-To install the necessary dependencies for this script, you can use the following `go get` command:

  ```sh
  go get github.com/Scofield-777/Go-Censys-API
```

## Authentication
-Set your Censys API credentials as environment variables. Replace YOUR_API_ID and YOUR_API_SECRET with your actual Censys API credentials.
```sh
  export CENSYS_API_ID=YOUR_API_ID
  export CENSYS_API_SECRET=YOUR_API_SECRET
```
## Usage
-main.go provides an example of how to interact with the Censys API securely using environment variables. It performs a certificate search based on your query and retrieves relevant data.
```sh
  go run main.go SQL
```
