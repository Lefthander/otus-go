#!/bin/bash

# Build folder with 5 files

mkdir -p ./envtest
for i in {0..5}
do
   echo value"$i" > envtest/"File$(printf "%03d" "$i").txt"
done

# Run the test

./goenv ./envtest env


# Clean up

rm -r ./envtest
