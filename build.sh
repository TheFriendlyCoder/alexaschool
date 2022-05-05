#!/bin/sh -ex
GOARCH=amd64 GOOS=linux go build -o alexaschool
[ -f alexaschool.zip ] && rm alexaschool.zip
zip alexaschool.zip alexaschool
aws lambda update-function-code --function-name GoSchool --zip-file fileb://alexaschool.zip
ask deploy
ask smapi get-skill-status -s amzn1.ask.skill.5cb6ba49-8028-4dc8-87ac-ba0e3ba100a3