#!/bin/bash

YEAR=$1
DAY=$2

YEAR_DIR=problem/year"${YEAR}"
DAY_DIR="${YEAR_DIR}"/day"${DAY}"
INPUTS_DIR=./inputs/year"${YEAR}"/day"${DAY}"

mkdir -p "${YEAR_DIR}" "${DAY_DIR}" "${INPUTS_DIR}"
touch "${INPUTS_DIR}"/input
touch "${INPUTS_DIR}"/example

DAY_FILE="${DAY_DIR}"/day"${DAY}".go
if [[ ! -f "${DAY_FILE}" ]]; then
  cp template "${DAY_FILE}"
  sed -i "s/XX/${DAY}/g" "${DAY_FILE}"
fi
