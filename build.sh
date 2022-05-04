#!/bin/sh
GOARCH=amd64 GOOS=linux go build -o goschool
rm goschool.zip
zip goschool.zip goschool
aws lambda update-function-code --function-name GoSchool --zip-file fileb://goschool.zip
ask smapi get-skill-status -s amzn1.ask.skill.5cb6ba49-8028-4dc8-87ac-ba0e3ba100a3