#!/bin/bash

# Load the .env file that will could contain the session cookie
if [ -f "../.env" ]; then
	source ../.env
else
	echo "Error: not .env file found to load session cookie"
	exit 1
fi

if [ -z "$SESSION_COOKIE" ]; then
	echo "SESSION_COOKIE is not set in the .env file! Exiting..."
	exit 1
fi

normalize_day() {
	local day=${1#0}
	echo "$day"
}

# Fetch the input for the given year and day
function getInput() {
	if [ $# -lt 2 ]; then
		echo "Error: Not enought args for $0"
		exit 1
	fi

	day=$(normalize_day ${2})

	# The day dir should be forced to dd format
	OUTPUT_DIR="${1}/$(printf "%02d" "${day}")"
	FILENAME="$OUTPUT_DIR/input.txt"

	# If we already have the input downloaded, no need to redownload it
	if [ -f $FILENAME ]; then
		echo "Input file '$FILENAME' already exists"
		return
	fi

	# Download the input for the specific year of the specific day into a temp file
	echo "Downloading input for year $1 day $2..."
	HTTP_STATUS=$(curl -s -w "%{http_code}" -b "session=$SESSION_COOKIE" "https://adventofcode.com/${1}/day/${day}/input" -o "tmp.txt")

	# If there was a 404 error remove the temp file and continue
	if [ $? -eq 0 ]; then
		if [ "$HTTP_STATUS" -eq 404 ]; then
			echo "Day ${day} not found (404). Skipping this day."
			rm -f "tmp.txt"
			return 1
		fi
		mkdir -p "$OUTPUT_DIR"
		mv "tmp.txt" "${FILENAME}"
		echo "Day ${day} input saved to ${FILENAME}"
		return 0
	else
		echo "Failed to download input for day ${day}"
		rm -f "tmp.txt"
		return 1
	fi

}

# If no args are provided then we'll download all the inputs
if [ $# -lt 2 ]; then
	echo "Year and/or Day not specified. Downloading everything"
	for YEAR in {2015..2024..1}; do
		for DAY in {1..25..1}; do
			getInput ${YEAR} ${DAY}
			code=$?
			if [ ${code} -ne 0 ]; then
				exit ${code}
			fi
		done
	done
else
	getInput ${1} ${2}
	code=$?
	if [ ${code} -ne 0 ]; then
		exit ${code}
	fi
fi
