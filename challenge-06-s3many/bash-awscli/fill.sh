#!/bin/bash
# Call with bucketname in environment variable
# BUCKET
for i in 0 1 2 3 4 5 6 7 8 9
do
    for k in 0 1 2 3 4 5 6 7 8 9
    do
        aws s3 cp readme.md s3://${BUCKET}/test-bash-awscli-${i}-${k}.md
    done
done
