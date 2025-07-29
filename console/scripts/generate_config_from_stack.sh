#!/bin/bash

set -eu

if [ $# -ne 1 ]; then
  echo "Usage: $0 <STACK_NAME>"
  exit 1
fi

STACK_NAME=$1

# Get UserPoolClientId from stack outputs
USER_POOL_CLIENT_ID=$(aws cloudformation describe-stacks \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Outputs[?OutputKey=='UserPoolClientId'].OutputValue" \
  --output text)

if [ -z "${USER_POOL_CLIENT_ID}" ]; then
  echo "Error: UserPoolClientId not found in stack outputs for stack: ${STACK_NAME}" >&2
  exit 1
fi

# Get CallbackURL and LogoutURL from stack parameters
AUTH_REDIRECT_URL=$(aws cloudformation describe-stacks \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Parameters[?ParameterKey=='CallbackURL'].ParameterValue" \
  --output text)

if [ -z "${AUTH_REDIRECT_URL}" ]; then
  echo "Error: CallbackURL not found in stack parameters for stack: ${STACK_NAME}" >&2
  exit 1
fi

LOGOUT_REDIRECT_URL=$(aws cloudformation describe-stacks \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Parameters[?ParameterKey=='LogoutURL'].ParameterValue" \
  --output text)

if [ -z "${LOGOUT_REDIRECT_URL}" ]; then
  echo "Error: LogoutURL not found in stack parameters for stack: ${STACK_NAME}" >&2
  exit 1
fi

# Get UserPoolId from stack parameters and construct IDP URL
USER_POOL_ID=$(aws cloudformation describe-stacks \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Parameters[?ParameterKey=='UserPoolId'].ParameterValue" \
  --output text)

if [ -z "${USER_POOL_ID}" ]; then
  echo "Error: UserPoolId not found in stack parameters for stack: ${STACK_NAME}" >&2
  exit 1
fi

REGION=$(aws configure get region)
IDP="https://cognito-idp.${REGION}.amazonaws.com/${USER_POOL_ID}"

CONFIG_FILE="$(dirname "$0")"/../static/config.json

cat << EOF > "${CONFIG_FILE}"
{
  "UserPoolClientId": "${USER_POOL_CLIENT_ID}",
  "AuthRedirectURL": "${AUTH_REDIRECT_URL}",
  "LogoutRedirectURL": "${LOGOUT_REDIRECT_URL}",
  "IDP": "${IDP}"
}
EOF

echo "Generated ${CONFIG_FILE}"
