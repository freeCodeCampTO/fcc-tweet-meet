go build && \
zip fcc-tweet-meet.zip fcc-tweet-meet && \
rm fcc-tweet-meet && \
aws lambda update-function-code --function-name fcc-tweet-meet --zip-file fileb://fcc-tweet-meet.zip && \
rm fcc-tweet-meet.zip