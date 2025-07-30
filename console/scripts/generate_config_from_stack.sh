#!/bin/bash

set -eux

if [ $# -ne 1 ]; then
  echo "Usage: $0 <STACK_NAME>"
  exit 1
fi

STACK_NAME=$1

# Determine region: first try aws configure get region, then fallback to AWS_REGION environment variable
REGION=$(aws configure get region || echo "")
if [ -z "${REGION}" ]; then
  if [ -z "${AWS_REGION}" ]; then
    echo "Error: AWS_REGION environment variable is not set and no region is configured in AWS CLI." >&2
    exit 1
  fi
  REGION=${AWS_REGION}
fi

# Function to resolve SSM parameters
resolve_ssm_parameter() {
  local param_value=$1
  if [[ "$param_value" == "{{resolve:ssm:"* ]]; then
    local ssm_param_name=$(echo "$param_value" | sed -e 's/^{{resolve:ssm://' -e 's/}}$//')
    echo "$(aws ssm get-parameter --region ${REGION} --name "${ssm_param_name}" --query Parameter.Value --output text)"
  else
    echo "${param_value}"
  fi
}

# Get UserPoolClientId from stack outputs
USER_POOL_CLIENT_ID=$(aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Outputs[?OutputKey=='UserPoolClientId'].OutputValue" \
  --output text)

if [ -z "${USER_POOL_CLIENT_ID}" ]; then
  echo "Error: UserPoolClientId not found in stack outputs for stack: ${STACK_NAME}" >&2
  exit 1
fi

# Get CallbackURL and LogoutURL from stack parameters
AUTH_REDIRECT_URL_RAW=$(aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Parameters[?ParameterKey=='CallbackURL'].ParameterValue" \
  --output text)

if [ -z "${AUTH_REDIRECT_URL_RAW}" ]; then
  echo "Error: CallbackURL not found in stack parameters for stack: ${STACK_NAME}" >&2
  exit 1
fi
AUTH_REDIRECT_URL=$(resolve_ssm_parameter "${AUTH_REDIRECT_URL_RAW}")

LOGOUT_REDIRECT_URL_RAW=$(aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Parameters[?ParameterKey=='LogoutURL'].ParameterValue" \
  --output text)

if [ -z "${LOGOUT_REDIRECT_URL_RAW}" ]; then
  echo "Error: LogoutURL not found in stack parameters for stack: ${STACK_NAME}" >&2
  exit 1
fi
LOGOUT_REDIRECT_URL=$(resolve_ssm_parameter "${LOGOUT_REDIRECT_URL_RAW}")

# Get UserPoolId from stack parameters and construct IDP URL
USER_POOL_ID_RAW=$(aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Parameters[?ParameterKey=='UserPoolId'].ParameterValue" \
  --output text)

if [ -z "${USER_POOL_ID_RAW}" ]; then
  echo "Error: UserPoolId not found in stack parameters for stack: ${STACK_NAME}" >&2
  exit 1
fi
USER_POOL_ID=$(resolve_ssm_parameter "${USER_POOL_ID_RAW}")

IDP="cognito-idp.${REGION}.amazonaws.com/${USER_POOL_ID}"

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
