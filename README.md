# Hackathon Sample Application in Golang
GL Hackathon Problem Statement

GL wants to host time-bound coding hackathons. These will be available for registration on the GL website and interested folks can signup, create their group, invite folks to join their group.

During the Hackathon, the problem statement along with some sample test cases will be made available to participants. There would be public test cases and hidden test cases. 

Each participating group will be able to upload their code any number of times during the hackathon. 

Based on the number and complexity of tests which the solution passes groups will be ranked into a real-time leaderboard. This leaderboard will be visible to all participants.

At the end of the hackathon top N groups will receive awards which they can claim.



###Requirment Docs
https://docs.google.com/document/d/1cE4F2ixCKZAYTEDF2T-gm2ZnMccFJEyd-bUAT3iJ1BM/edit#


###API Enpoints 
| Submit Hackathon Problem Solution to system
```shell script
curl --location --request POST 'localhost:7378/api/v1/auth/problem/submission' \
--header 'Authorization: Bearrer <<AuthToken>>' \
--header 'Auth-Token: dhjfghefbkjerg555632bdsg3' \
--header 'Content-Type: application/json' \
--data-raw '{
    "event_id":1,
    "problem_id":2,
    "source_code":"{{AllTheCode}}",
    "lang":"c"
}'
```
Response 
```json
{
    "header": {
        "process_time": "6.887708",
        "server_time": "1599320553",
        "status_code": 200,
        "message": ""
    },
    "data": {
        "status": "prosecessing",
        "sub_id": 10
    }
}
```

#### Get Leader Board 
