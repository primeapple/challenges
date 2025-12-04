#!/bin/bash

set -euo pipefail

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to convert string to camelCase (first word lowercase, rest capitalized)
to_camel_case() {
    local result=$(echo "$1" | sed -r 's/(^| |_)([a-z])/\U\2/g' | sed 's/ //g')
    # Make first character lowercase
    echo "$(echo "${result:0:1}" | tr '[:upper:]' '[:lower:]')${result:1}"
}

# Function to convert string to PascalCase (all words capitalized)
to_pascal_case() {
    echo "$1" | sed -r 's/(^| |_)([a-z])/\U\2/g' | sed 's/ //g'
}

# Get the Monday of the current week
get_monday_of_week() {
    # Get current day of week (1=Monday, 7=Sunday)
    day_of_week=$(date +%u)
    
    # Calculate days to subtract to get to Monday
    days_to_subtract=$((day_of_week - 1))
    
    # Get Monday's date
    date -d "-${days_to_subtract} days" +%Y_%m_%d
}

# Prompt for challenge name
echo -e "${BLUE}Enter the challenge name (e.g., 'sort html colors'):${NC}"
read -r challenge_name

if [ -z "$challenge_name" ]; then
    echo "Error: Challenge name cannot be empty"
    exit 1
fi

# Generate variables
MONDAY_DATE=$(get_monday_of_week)
CHALLENGE_DIR="${MONDAY_DATE}_${challenge_name// /_}"
PACKAGE_NAME=$(to_camel_case "$challenge_name")
FUNCTION_NAME=$(to_pascal_case "$challenge_name")

# Create directory
echo -e "${GREEN}Creating directory: ${CHALLENGE_DIR}${NC}"
mkdir -p "$CHALLENGE_DIR"

# Function to replace template variables
replace_vars() {
    local template_file=$1
    local output_file=$2
    
    sed -e "s/{{PACKAGE_NAME}}/${PACKAGE_NAME}/g" \
        -e "s/{{FUNCTION_NAME}}/${FUNCTION_NAME}/g" \
        -e "s/{{DATE}}/${MONDAY_DATE}/g" \
        "$template_file" > "$output_file"
}

# Generate files from templates
echo -e "${GREEN}Generating files from templates...${NC}"

replace_vars "templates/QUESTION.md.tmpl" "${CHALLENGE_DIR}/QUESTION.md"
replace_vars "templates/solution.go.tmpl" "${CHALLENGE_DIR}/solution.go"
replace_vars "templates/solution_test.go.tmpl" "${CHALLENGE_DIR}/solution_test.go"

# Summary
echo -e "${GREEN}✓ Challenge created successfully!${NC}"
echo ""
echo "Directory: ${CHALLENGE_DIR}/"
echo "Package: ${PACKAGE_NAME}"
echo "Function: ${FUNCTION_NAME}"
echo ""
echo -e "${BLUE}Next steps:${NC}"
echo "1. Edit ${CHALLENGE_DIR}/QUESTION.md with the challenge description"
echo "2. Implement your solution in ${CHALLENGE_DIR}/solution.go"
echo "3. Add tests in ${CHALLENGE_DIR}/solution_test.go"
echo "4. Run tests: cd ${CHALLENGE_DIR} && go test"
