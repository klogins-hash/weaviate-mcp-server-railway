# Weaviate MCP Server Project

This project contains the official Weaviate MCP (Model Context Protocol) server implementation, optimized for Railway deployment.

## Project Structure

```text
server doesn't know v3/
├── mcp-server-weaviate/          # Official Weaviate MCP server
│   ├── client/
│   │   ├── mcp-server            # Built server binary
│   │   └── client.go             # Test client
│   ├── .env                      # Environment configuration
│   ├── Dockerfile                # Railway deployment
│   ├── railway.toml              # Railway configuration
│   ├── start-server.sh           # Server startup script
│   └── ...                       # Other project files
└── README.md                     # This file
```

## Prerequisites

- Go 1.25+ (installed via Homebrew)
- Weaviate instance (local, cloud, or Weaviate Cloud Services)

## Configuration

The server is configured via environment variables in `.env`:

```bash
# Railway will automatically set PORT
PORT=8675

# Weaviate connection settings
WEAVIATE_HOST=localhost:8080
WEAVIATE_SCHEME=http

# For Weaviate Cloud Services (recommended for Railway)
# WEAVIATE_HOST=your-cluster.weaviate.network
# WEAVIATE_SCHEME=https
# WEAVIATE_API_KEY=your_api_key_here
```

## Railway Deployment

### Quick Deploy

1. **Connect to Railway**: Link your GitHub repository to Railway
2. **Set Environment Variables**: Configure Weaviate connection in Railway dashboard
3. **Deploy**: Railway will automatically build and deploy using the Dockerfile

### Environment Variables for Railway

Set these in your Railway project dashboard:

```bash
WEAVIATE_HOST=your-cluster.weaviate.network
WEAVIATE_SCHEME=https
WEAVIATE_API_KEY=your_weaviate_api_key
```

### API Endpoints

Once deployed, your Railway app will expose:

- `GET /` - Server information and available endpoints
- `GET /health` - Health check endpoint
- `GET /tools` - Available MCP tools
- `POST /call` - Execute MCP tools

## Local Development

### 1. Start Weaviate (if running locally)

```bash
docker run -p 8080:8080 -e QUERY_DEFAULTS_LIMIT=25 -e AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true -e PERSISTENCE_DATA_PATH='/var/lib/weaviate' -e DEFAULT_VECTORIZER_MODULE='none' -e ENABLE_MODULES='' -e CLUSTER_HOSTNAME='node1' semitechnologies/weaviate:latest
```

### 2. Build and Start the MCP Server

```bash
cd "mcp-server-weaviate"
make build
./start-server.sh
```

The server will be available at `http://localhost:8675`

## Available Tools

- **weaviate-insert-one**: Insert an object into Weaviate
- **weaviate-query**: Retrieve objects from Weaviate with hybrid search

## API Usage Examples

### Insert Object

```bash
curl -X POST http://your-app.railway.app/call \
  -H "Content-Type: application/json" \
  -d '{
    "tool": "weaviate-insert-one",
    "arguments": {
      "collection": "MyCollection",
      "properties": {
        "title": "Example Document",
        "content": "This is example content"
      }
    }
  }'
```

### Query Objects

```bash
curl -X POST http://your-app.railway.app/call \
  -H "Content-Type: application/json" \
  -d '{
    "tool": "weaviate-query",
    "arguments": {
      "query": "example content",
      "targetProperties": ["title", "content"]
    }
  }'
```

## Development Commands

- **Build**: `make build`
- **Test client**: `make run-client` (requires server to be running)
- **Docker build**: `docker build -t weaviate-mcp-server .`
- **Docker run**: `docker run -p 8675:8675 --env-file .env weaviate-mcp-server`

## Notes

- The server supports both stdio (for local MCP clients) and HTTP (for Railway deployment)
- Railway deployment uses HTTP endpoints for better cloud compatibility
- For production, use Weaviate Cloud Services for better reliability
- Health check endpoint at `/health` is configured for Railway monitoring
