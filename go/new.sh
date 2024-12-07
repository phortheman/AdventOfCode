#!/bin/bash

# Creates a main.go and main_test.go file using the templates
# Make sure do to a sed command to replace the input puzzle name to be the relative location
# Also could embed the puzzle input. Make it technically more potable

if [ "$#" -ne 2 ]; then
	echo "Usage: $0 <year> <day>"
	exit 1
fi

year=$1
day=$2

# Validate year format
if [[ $year =~ ^[0-9]{2}$ ]]; then
	year="20$year"
elif [[ $year =~ ^[0-9]{4}$ ]]; then
	if [ "$year" -lt 2000 ] || [ "$year" -gt 2099 ]; then
		echo "Error: <year> must be between 2000 and 2100: ${year}"
		exit 1
	fi
else
	echo "Error: invalid <year> '${year}'. Must be either yy or yyyy format. And obviously make sense for AOC."
	exit 1
fi

# Check if day is greater than 25
if [ "$day" -gt 25 ]; then
	echo "Error: <day> cannot be greater than 25."
	exit 1
fi

# Force day to be full dd format
day=$(printf "%02d" "$day")

directory="./${year}/${day}"

if [ -d ${directory} ]; then
	echo "Directory already exists: ${directory}"
	exit 1
fi

# Get the input for the day. If a non zero returns then stop processing
../inputs/fetch.sh ${year} ${day}
code=$?
if [ $code -ne 0 ]; then
	echo "Error downloading input. Exiting..."
	exit $code
fi

# Make the directory
mkdir -p ${directory}

# Copy the templated base files to the solution directory
cp template.go ${directory}
cp test_template.go ${directory}

# Update the copied files to set the relative input path
sed -i -e "s/<YEAR>/${year}/g" -e "s/<DAY>/${day}/g" "${directory}/template.go"

echo "Setup for year ${year}, day ${day} completed"
echo "Files created:"
echo "	- ${directory}/main.go"
echo "	- ${directory}/main_test.go"
