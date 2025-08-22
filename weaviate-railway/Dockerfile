# Weaviate Docker configuration optimized for Railway
# Use the official Weaviate Docker image
FROM semitechnologies/weaviate:1.24.1

# Set environment variables for Railway
ENV AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true \
    PERSISTENCE_DATA_PATH=/var/lib/weaviate \
    QUERY_DEFAULTS_LIMIT=25 \
    DEFAULT_VECTORIZER_MODULE=none \
    CLUSTER_HOSTNAME=weaviate \
    AUTOSCHEMA_ENABLED=true

# Expose port (Railway will override with PORT env var)
EXPOSE 8080

# Start Weaviate with Railway's PORT
CMD ["sh", "-c", "/bin/weaviate --host 0.0.0.0 --port ${PORT:-8080} --scheme http"]
