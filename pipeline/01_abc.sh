#!/bin/bash 

echo "some output"
echo "another line"

>&2 echo "this goes to stderr"
>&2 echo "another error"
