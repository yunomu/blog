#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -euo pipefail

# Set default stack name if not provided
STACK_NAME=${1:-wagahai-blog}

# Get CloudFront distribution ID from CloudFormation stack
echo "Searching for CloudFront distribution ID in stack '${STACK_NAME}'..."
DISTRIBUTION_ID=$(aws cloudformation describe-stack-resources --stack-name "${STACK_NAME}" | jq -r '.StackResources[] | select(.LogicalResourceId=="Distribution") | .PhysicalResourceId')

# Check if the distribution ID was found
if [ -z "${DISTRIBUTION_ID}" ]; then
    echo "Error: Could not find a CloudFront distribution in stack '${STACK_NAME}'."
    exit 1
fi

echo "Found CloudFront distribution ID '${DISTRIBUTION_ID}'."

# Create an invalidation
echo "Creating invalidation for path '/*'..."
INVALIDATION_ID=$(aws cloudfront create-invalidation --distribution-id "${DISTRIBUTION_ID}" --paths "/*" | jq -r '.Invalidation.Id')

echo "Invalidation created with ID: ${INVALIDATION_ID}"
echo "Waiting for invalidation to complete..."

# Wait for the invalidation to complete
aws cloudfront wait invalidation-completed --distribution-id "${DISTRIBUTION_ID}" --id "${INVALIDATION_ID}"

echo "Invalidation completed successfully."
