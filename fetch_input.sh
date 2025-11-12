#!/usr/bin/env bash

set -uo pipefail

readonly SUCCESS=0
readonly WARN=1
readonly ERROR=2
readonly PANIC=3

AOC_SESSION_COOKIE="${AOC_SESSION_COOKIE-}"

function info() {
	echo "[INFO]: $1"
}

function warn() {
	echo "[WARN]: $1" >&2
}

function error() {
	echo "[ERROR]: $1" >&2
}

function panic() {
	echo "[PANIC]: $1" >&2
	exit $PANIC
}

# Load the .env if the env variable AOC_SESSION_COOKIE isn't set
function load_session_cookie() {
	if [ -f ".env" ]; then
		source ".env"
	else
		panic "no .env file found to load session cookie"
	fi

	if [ -z "${AOC_SESSION_COOKIE}" ]; then
		panic "AOC_SESSION_COOKIE environment variable is not set or in the root .env file"
	fi
}

# Downloads the input for a specific year and day
# Args: <year> <day>
# Returns:
#   0 (SUCCESS) - Successfully downloaded
#   1 (WARN)    - Nothing was done but no error
#   2 (ERROR)   - Day not found
#   3 (PANIC)   - Critical failure (network error, etc.)
function get_input() {
	if [ $# -ne 2 ]; then
		panic "get_input requires exactly 2 arguments"
	fi

	local year="$1"
	local day="$2"

	# The day dir should be forced to dd format
	local output_dir="inputs/${year}/$(printf "%02d" "${day}")"
	local filename="$output_dir/input.txt"

	if [ -f "${filename}" ]; then
		warn "Input file '$filename' already exists"
		return $WARN
	fi

	# Download the input for the specific year of the specific day into a temp file
	info "Downloading input for year $year day $day"

	local url="https://adventofcode.com/${year}/day/${day}/input"
	info "executing http get request to $url"

	local temp_file=$(mktemp)
	local http_status=$(
		curl \
			--request GET \
			--silent \
			--write-out "%{http_code}" \
			--cookie "session=${AOC_SESSION_COOKIE}" \
			"$url" \
			--output "${temp_file}"
	)

	local curl_exit=$?
	if [ ${curl_exit} -ne 0 ]; then
		rm -f "${temp_file}"
		panic "Failed to download input for day ${day} (curl exit code: ${curl_exit})"
	fi

	case ${http_status} in
	200)
		mkdir -p "${output_dir}"
		mv "${temp_file}" "${filename}"
		info "✓ Day ${day} input saved to ${filename}"
		return $SUCCESS
		;;
	400)
		rm -f "${temp_file}"
		# Show first 20 chars of cookie for debugging
		error "Cookie starts with: ${AOC_SESSION_COOKIE:0:20}..."
		panic "Bad request (400) for day ${day}. Check session cookie!"
		;;
	404)
		# This day isn't ready yet
		rm -f "${temp_file}"
		return ${ERROR}
		;;
	*)
		rm -f "${temp_file}"
		panic "Unexpected HTTP status ${http_status} for day ${day}"
		;;
	esac

}

# Call get_input using generated year and day args
# If $ERROR is returned then it'll stop trying to download and complete the process
function download_all() {
	info "Year and/or Day not specified. Downloading all inputs..."

	# From 2015 to 2024 there are 25 puzzles
	for year in {2015..2024}; do
		for day in {1..25}; do
			get_input $year $day
			local code=$?
			if [ ${code} -gt $WARN ]; then
				return $code
			fi
		done
	done

	# After 2025 there are 12 puzzles
	local year=2025
	while [ $? -eq 0 ]; do
		for day in {1..12}; do
			get_input ${year} ${day}
			local code=$?
			if [ ${code} -gt $WARN ]; then
				return ${code}
			fi
		done
		((year++))
	done
}

function main() {
	if [ ! -d ".git" ]; then
		panic "fetch script must be executed in the root directory of the project"
	fi

	if [ -z "${AOC_SESSION_COOKIE:-}" ]; then
		load_session_cookie
	fi

	# If no args are provided then we'll download all the inputs
	if [ $# -lt 2 ]; then
		download_all
		local code=$?
		if [ $code -gt $WARN ]; then
			info "All available puzzle inputs downloaded"
			return $SUCCESS
		fi

		return $code
	else
		get_input "$1" "$2"
		local code=$?
		if [ ${code} -eq $ERROR ]; then
			error "Day not found. Puzzle may not be released yet."
			exit ${code}
		fi
	fi
}

main $@
