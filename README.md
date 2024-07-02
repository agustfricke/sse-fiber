# SSE Fiber

This project demonstrates a simple implementation of Server-Sent Events (SSE) using the Fiber web framework in Go. 
SSE allows servers to push real-time updates to clients over a single, long-lived HTTP connection.

## Getting Started

To run this project, make sure you have Go installed on your machine.

1. Clone the repository:
   ```bash
   git clone https://github.com/agustfricke/sse-fiber
   cd sse-fiber
   ```

2. Install dependencies (assuming you have Go modules enabled):
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

4. Open `http://localhost:3000` in your web browser to see the SSE in action.

## Project Structure

- **`main.go`**: Contains the main Go application code using the Fiber web framework.
  
- **`static/index.html`**: Simple HTML file that connects to the SSE endpoint (`/sse/validation`) and logs incoming messages to the browser console.

## How It Works

### Server (`main.go`)

- **Fiber Setup**: Initializes a Fiber web server (`app`).
  
- **SSE Handler (`/sse/validation`)**:
  - Sets appropriate headers (`Content-Type: text/event-stream`, `Cache-Control: no-cache`, `Connection: keep-alive`) to support SSE.
  - Creates a new `Client` when a request is received, which represents a connected client.
  - Continuously sends `DashBoard` updates to the connected clients every second.

- **Client Struct**: Represents a client connected to the SSE endpoint. It maintains a channel (`events`) to send `DashBoard` updates.

- **DashBoard Struct**: Simple structure representing some data (`User` in this case) that is periodically sent to clients.

### Client (`static/index.html`)

- **HTML Setup**: Basic HTML structure with a `<script>` tag.

- **JavaScript**:
  - Defines an `onload` event handler that initializes an `EventSource` object.
  - Connects to the server's SSE endpoint (`/sse/validation`).
  - Logs incoming messages (`onmessage` event) from the server to the browser console.

## Notes

- SSE is suitable for applications where the server needs to push updates to clients in real-time without clients having to repeatedly poll the server.
- Fiber is a lightweight web framework for Go, known for its speed and minimalism.
