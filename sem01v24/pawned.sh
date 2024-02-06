#!/bin/bash
candidate_password=$1
echo "Candidate passwordd: $candidate_password"

full_hash=$(echo -n $candidate_password | sha1sum | awk '{print substr($1, 0, 32)}')
prefix=$(echo $full_hash | awk '{print substr($1, 0, 5)}')
suffix=$(echo $full_hash | awk '{print substr($1, 6, 26)}')

if curl https://api.pwnedpasswords.com/range/$prefix | grep -i $suffix;
    then  echo "Candidate password is compromised";
    else echo "Candidate password is OK for use";
fi
