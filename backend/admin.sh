#!/bin/bash

# Function to prompt user and get input
prompt_input() {
  read -p "$1" input
  echo "$input"
}

if [ -f ".env" ]; then
  # Export all variables from the .env file into the script's environment
  export $(grep -v '^#' .env | xargs)
else
  echo "Error: .env file not found in the current directory."
  exit 1
fi

# Ensure DATABASE_URL is set
if [ -z "$DATABASE_URL" ]; then
  echo "Error: DATABASE_URL environment variable is not set."
  exit 1
fi

# Get user hash
user_hash=$(prompt_input "Enter the user hash: ")

# Confirm valid hash input
if [[ ! $user_hash =~ ^[a-f0-9]{64}$ ]]; then
  echo "Error: Invalid user hash. It must be a 64-character SHA256 hash."
  exit 1
fi

# Ask for action
action=$(prompt_input "Do you want to add or remove beta tester rights? (add/remove): ")

# Determine SQL command
case $action in
  add)
    is_beta_tester="TRUE"
    ;;
  remove)
    is_beta_tester="FALSE"
    ;;
  *)
    echo "Error: Invalid action. Use 'add' or 'remove'."
    exit 1
    ;;
esac

# Execute SQL command
sql_command="UPDATE users SET is_beta_tester = $is_beta_tester WHERE user_hash = '$user_hash';"

# Run SQL command via psql
PGPASSWORD=$(echo $DATABASE_URL | sed -n 's/.*:.*:\(.*\)@.*/\1/p') \
psql "$DATABASE_URL" -c "$sql_command"

if [ $? -eq 0 ]; then
  echo "Successfully updated beta tester status."
else
  echo "Error updating beta tester status."
  exit 1
fi
