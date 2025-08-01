#!/bin/bash
set -e

openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -subj "/CN=guardian-ca" -days 365 -out ca.crt
