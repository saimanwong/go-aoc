#!/bin/bash

YEAR=$1
DAY=$2

YEAR_DIR=problem/year"${YEAR}"
DAY_DIR="${YEAR_DIR}"/day"${DAY}"
INPUTS_DIR="${DAY_DIR}"/inputs

mkdir -p "${YEAR_DIR}" "${DAY_DIR}" "${INPUTS_DIR}"
touch "${INPUTS_DIR}"/input

DAY_FILE="${DAY_DIR}"/day"${DAY}".go
if [[ ! -f "${DAY_FILE}" ]]; then
  cp template "${DAY_FILE}"
  sed -i "s/XX/${DAY}/g" "${DAY_FILE}"
fi
