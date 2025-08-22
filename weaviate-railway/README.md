# Weaviate Instance for Railway

This directory contains the configuration to deploy a Weaviate vector database instance on Railway.

## Deployment Steps

1. **Create New Railway Service**:
   - Go to Railway dashboard
   - Create new project from GitHub repo
   - Select this `weaviate-railway` folder as the root

2. **Railway will automatically**:
   - Build using the Dockerfile
   - Set environment variables from railway.toml
   - Deploy on a public URL

3. **Get Your Weaviate URL**:
   - After deployment, Railway will provide a URL like: `https://your-app.railway.app`
   - Your Weaviate endpoint will be: `https://your-app.railway.app`

## Environment Variables

The following are automatically configured via `railway.toml`:

- `QUERY_DEFAULTS_LIMIT=25`
- `AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true`
- `PERSISTENCE_DATA_PATH=/var/lib/weaviate`
- `DEFAULT_VECTORIZER_MODULE=none`
- `ENABLE_MODULES=""`
- `CLUSTER_HOSTNAME=node1`
- `ORIGIN=*`

## Health Check

Railway will monitor the health endpoint at `/v1/.well-known/ready`

## Usage with MCP Server

Once deployed, update your MCP server's `.env` file:

```bash
WEAVIATE_HOST=your-weaviate-app.railway.app
WEAVIATE_SCHEME=https
# No API key needed with anonymous access enabled
```

## API Endpoints

Your Railway-hosted Weaviate will be available at:

- GraphQL: `https://your-app.railway.app/v1/graphql`
- REST: `https://your-app.railway.app/v1/objects`
- Health: `https://your-app.railway.app/v1/.well-known/ready`
