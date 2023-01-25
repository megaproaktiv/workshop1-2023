#!/bin/bash
# Call with bucketname in environment variable
# BUCKET
for i in 0 1 2 3 4 5 6 7 8 9
do
    for k in 0 1 2 3 4 5 6 7 8 9
    do
        ./go-s3cp --from readme.md --to s3://${BUCKET}/test-bash-go-s3cp-${i}-${k}.md
    done
done



