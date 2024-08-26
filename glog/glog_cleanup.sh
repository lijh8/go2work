#!/bin/bash

# Delete the oldest log files
# Keep most latest log files at specified count

# $1 - Level of log file (e.g., INFO, WARNING, ERROR)
# $2 - Maximum count of log files to keep
glog_cleanup() {
    local log_level=$1
    local max_count=$2

    if [[ -z "$log_level" ]]; then
        echo "Error: Log level not specified."
        return
    fi

    if ! [[ "$max_count" =~ ^[0-9]+$ ]] || (( max_count < 0 )); then
        echo "Error: max_count must be a non-negative integer."
        return
    fi

    if (( max_count < 5 )); then
        max_count=5
    fi

    log_files=($(ls -1t ${LOG_DIR}/*.log.${log_level}.*))

    if [ ${#log_files[@]} -gt $max_count ]; then
        num_to_delete=$((${#log_files[@]} - $max_count))

        # Delete the oldest log files!!!
        for (( i=0; i<num_to_delete; i++ )); do
            echo "delete: ${log_files[$((max_count + i))]}"
            rm "${log_files[$((max_count + i))]}"
        done
    fi
}

# $ sh glog_cleanup.sh

# path to the log files
LOG_DIR="./logs"

MAX_COUNT_KEEP=5

glog_cleanup "INFO" $MAX_COUNT_KEEP
glog_cleanup "WARNING" $MAX_COUNT_KEEP
glog_cleanup "ERROR" $MAX_COUNT_KEEP
