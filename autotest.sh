#!/bin/bash


for file in ./tests/*; do
    output=$(./main < $file)
    ans=$(cat "./anses/$(basename $file).ans")
    if [ "$output" = "$ans" ]; then
        echo "Passed $file"
    else
        echo "Not passed $file"
    fi    
done    