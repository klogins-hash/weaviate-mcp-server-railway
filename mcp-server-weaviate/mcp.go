package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type MCPServer struct {
	server            *server.MCPServer
	weaviateConn      *WeaviateConnection
	defaultCollection string
}

func NewMCPServer() (*MCPServer, error) {
	conn, err := NewWeaviateConnection()
	if err != nil {
		return nil, err
	}
	s := &MCPServer{
		server: server.NewMCPServer(
			"Weaviate MCP Server",
			"0.1.0",
			server.WithToolCapabilities(true),
			server.WithPromptCapabilities(true),
			server.WithResourceCapabilities(true, true),
			server.WithRecovery(),
		),
		weaviateConn: conn,
		// TODO: configurable collection name
		defaultCollection: "DefaultCollection",
	}
	s.registerTools()
	return s, nil
}

func (s *MCPServer) Serve() {
	server.ServeStdio(s.server)
}

func (s *MCPServer) ServeHTTP(port string) {
	// Create HTTP handlers for Railway deployment
	http.HandleFunc("/health", s.healthHandler)
	http.HandleFunc("/tools", s.toolsHandler)
	http.HandleFunc("/call", s.callHandler)
	http.HandleFunc("/", s.rootHandler)

	log.Printf("Starting HTTP server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (s *MCPServer) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"service": "weaviate-mcp-server",
	})
}

func (s *MCPServer) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"name": "Weaviate MCP Server",
		"version": "0.1.0",
		"endpoints": []string{"/health", "/tools", "/call"},
		"tools": []string{"weaviate-insert-one", "weaviate-query"},
	})
}

func (s *MCPServer) toolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tools := []map[string]interface{}{
		{
			"name": "weaviate-insert-one",
			"description": "Insert an object into Weaviate",
			"parameters": map[string]interface{}{
				"collection": "Name of the target collection",
				"properties": "Object properties to insert (required)",
			},
		},
		{
			"name": "weaviate-query",
			"description": "Retrieve objects from Weaviate with hybrid search",
			"parameters": map[string]interface{}{
				"query": "Query data within Weaviate (required)",
				"targetProperties": "Properties to return with the query (required)",
			},
		},
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tools": tools,
	})
}

func (s *MCPServer) callHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Tool      string                 `json:"tool"`
		Arguments map[string]interface{} `json:"arguments"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Create MCP request format
	mcpReq := mcp.CallToolRequest{}
	mcpReq.Params.Name = req.Tool
	mcpReq.Params.Arguments = req.Arguments

	var result *mcp.CallToolResult
	var err error

	switch req.Tool {
	case "weaviate-insert-one":
		result, err = s.weaviateInsertOne(context.Background(), mcpReq)
	case "weaviate-query":
		result, err = s.weaviateQuery(context.Background(), mcpReq)
	default:
		http.Error(w, fmt.Sprintf("Unknown tool: %s", req.Tool), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (s *MCPServer) registerTools() {
	insertOne := mcp.NewTool(
		"weaviate-insert-one",
		mcp.WithString(
			"collection",
			mcp.Description("Name of the target collection"),
		),
		mcp.WithObject(
			"properties",
			mcp.Description("Object properties to insert"),
			mcp.Required(),
		),
	)
	query := mcp.NewTool(
		"weaviate-query",
		mcp.WithString(
			"query",
			mcp.Description("Query data within Weaviate"),
			mcp.Required(),
		),
		mcp.WithArray(
			"targetProperties",
			mcp.Description("Properties to return with the query"),
			mcp.Required(),
		),
	)

	s.server.AddTools(
		server.ServerTool{Tool: insertOne, Handler: s.weaviateInsertOne},
		server.ServerTool{Tool: query, Handler: s.weaviateQuery},
	)
}

func (s *MCPServer) weaviateInsertOne(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	targetCol := s.parseTargetCollection(req)
	props := req.Params.Arguments["properties"].(map[string]interface{})

	res, err := s.weaviateConn.InsertOne(context.Background(), targetCol, props)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("failed to insert object", err), nil
	}
	return mcp.NewToolResultText(res.ID.String()), nil
}

func (s *MCPServer) weaviateQuery(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	targetCol := s.parseTargetCollection(req)
	query := req.Params.Arguments["query"].(string)
	// TODO: how to enforce `Required` within the sdk so we don't have to validate here
	props := req.Params.Arguments["targetProperties"].([]interface{})
	var targetProps []string
	{
		for _, prop := range props {
			typed, ok := prop.(string)
			if !ok {
				return mcp.NewToolResultError("targetProperties must contain only strings"), nil
			}
			targetProps = append(targetProps, typed)
		}
	}
	res, err := s.weaviateConn.Query(context.Background(), targetCol, query, targetProps)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("failed to process query", err), nil
	}
	return mcp.NewToolResultText(res), nil
}

func (s *MCPServer) parseTargetCollection(req mcp.CallToolRequest) string {
	var (
		targetCol = s.defaultCollection
	)
	col, ok := req.Params.Arguments["collection"].(string)
	if ok {
		targetCol = col
	}
	return targetCol
}
