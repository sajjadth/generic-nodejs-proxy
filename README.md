# Generic Proxy

This project is a simple proxy server, initially built with Node.js and now migrated to Go for improved performance and lower memory usage. It forwards incoming HTTP requests to a target URL, allowing you to access third-party APIs via the proxy.

## Features

- Forwards all incoming requests to the specified target URL
- Preserves query parameters and request payloads
- Configurable via environment variables
- Supports `GET`, `POST`, `PUT`, `PATCH`, and `DELETE` HTTP methods

## Prerequisites

- [Go](https://golang.org/) (v1.22+ recommended)

## Configuration

1.Create a .env file in the project root directory and add the target URL:

```
TARGET_URL=https://your-target-api.com
PORT=3000
```

- Replace `https://your-target-api.com` with the URL you want to forward requests to.
- Set `PORT` to the desired port number (default is 3000).

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository_url>
   cd <repository_directory>
   ```
2. Create a .env file:
   - Copy the provided environment variables template and add your specific values.
3. Install dependencies:
   ```
   go mod tidy
   ```
4. Run the application:
   ```
   go run main.go
   ```

## Examples

### Forwarding a GET Request

If you make a request to your proxy URL:

```
GET http://localhost:3000/api/users?name=John
```

The proxy will forward it to:

```
GET https://your-target-api.com/api/users?name=John
```

### Forwarding a POST Request with Payload

Request to Proxy:

```
POST http://localhost:3000/api/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com"
}
```

The proxy forwards it to:

```
POST https://your-target-api.com/api/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com"
}
```

## Migration from Node.js to Go

This project was initially developed in Node.js and has since been migrated to Go for improved performance and lower memory usage.

If you're interested in the Node.js version, you can check out the `v1.0-node` tag.

## Customization

You can modify `main.go` to add custom logic, such as:

- Changing request headers
- Adding authentication
- Logging requests

## Troubleshooting

- Ensure your `TARGET_URL` is correct in the `.env` file.
- Make sure the target API is reachable from your server.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
