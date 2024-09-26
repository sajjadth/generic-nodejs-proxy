# Generic Node.js Proxy

This project is a simple proxy server built with Node.js, Express, and `http-proxy-middleware`. It forwards incoming HTTP requests to a target URL, allowing you to access third-party APIs via the proxy.

## Features
- Forwards all incoming requests to the specified target URL
- Preserves query parameters and request payloads
- Configurable via environment variables
- Supports `GET`, `POST`, `PUT`, `PATCH`, and `DELETE` HTTP methods

## Prerequisites

- [Node.js](https://nodejs.org/) (v14+ recommended)
- [NPM](https://www.npmjs.com/) or [Yarn](https://yarnpkg.com/)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/sajjadth/generic-nodejs-proxy.git
   cd generic-nodejs-proxy
   ```

2. Install dependencies:
   ```bash
   npm install
   ```
   Or, if you prefer using yarn:
   ```bash
   yarn install
   ```

## Configuration

1. Create a `.env` file in the project root directory and add the target URL:
   ```
   TARGET_URL=https://your-target-api.com
   PORT=3000
   ```
   - Replace `https://your-target-api.com` with the URL you want to forward requests to.
   - Set `PORT` to the desired port number (default is 3000).

## Usage

### Run Locally
Start the server using:
```bash
npm start
```
Or with Yarn:
```bash
yarn start
```

The proxy server will run on `http://localhost:3000` by default and will forward all requests to your `TARGET_URL`.

You can host this proxy server on any platform that supports Node.js, such as a VPS, cloud service, or container orchestration platform like Docker or Kubernetes.

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

## Customization

You can modify `index.js` to add custom logic, such as:
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
