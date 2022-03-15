#!/bin/bash

# generate_clients() {
#   come back to this later
# }

generate_types() {
  echo "Generating types..."
  mkdir -p generated/api
  oapi-codegen -config config/types-config.yaml user-accounts-test.yaml  
  echo "Types generated successfully."
}

generate_server() {
  echo "Generating server..."
  mkdir -p generated/server
  oapi-codegen -config config/server-config.yaml user-accounts-test.yaml  
  echo "Server generated successfully." 
  }

generate_types
generate_server

exit 0