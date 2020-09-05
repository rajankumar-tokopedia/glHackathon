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
 ##### 1. Submit Hackathon Problem Solution 
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

 ##### 2.  Get Leader Board 
 ```shell script
curl --location --request GET 'http://localhost:7378/api/v1/event/1/leaders'
```

Response 
```json
{
    "header": {
        "process_time": "2.946233ms",
        "server_time": "Sun, 06 Sep 2020 02:51:56 IST",
        "status_code": 200,
        "message": ""
    },
    "data": [
        {
            "event_id": 1,
            "group_id": 2,
            "score": 154.97176,
            "rank": 1,
            "top_submissions": [
                {
                    "sub_id": 7,
                    "event_id": 1,
                    "problem_id": 1,
                    "language": "go",
                    "status": {
                        "label": "accepted",
                        "id": 1
                    },
                    "accuracy": 94.1104,
                    "created_at": "2020-09-05 15:38:07"
                },
                {
                    "sub_id": 10,
                    "event_id": 1,
                    "problem_id": 2,
                    "language": "c",
                    "status": {
                        "label": "accepted",
                        "id": 1
                    },
                    "accuracy": 60.861366,
                    "created_at": "2020-09-05 15:42:33"
                }
            ]
        },
        {
            "event_id": 1,
            "group_id": 0,
            "score": 90.01,
            "rank": 2,
            "top_submissions": [
                {
                    "sub_id": 1,
                    "event_id": 1,
                    "problem_id": 1,
                    "language": "js",
                    "status": {
                        "label": "accepted",
                        "id": 1
                    },
                    "accuracy": 90.01,
                    "created_at": "2020-09-05 11:40:05"
                }
            ]
        },
        {
            "event_id": 1,
            "group_id": 1,
            "score": 60.01,
            "rank": 3,
            "top_submissions": [
                {
                    "sub_id": 3,
                    "event_id": 1,
                    "problem_id": 1,
                    "language": "go",
                    "status": {
                        "label": "accepted",
                        "id": 1
                    },
                    "accuracy": 60.01,
                    "created_at": "2020-09-05 12:49:39"
                }
            ]
        }
    ]
}
```

#### 3. GET All Submissions by Group On Particular Event

```shell script
curl --location --request GET 'localhost:7378/api/v1/auth/event/1/submissions' \
--header 'Auth-Token: dhjfghefbkjerg555632bdsg3'
```

Response 
```json
{
    "header": {
        "process_time": "0.464677ms",
        "server_time": "Sun, 06 Sep 2020 02:53:17 IST",
        "status_code": 200,
        "message": ""
    },
    "data": [
        {
            "sub_id": 5,
            "event_id": 1,
            "group_id": 2,
            "problem_id": 1,
            "source_code": "{{AllTheCode}}",
            "language": "go",
            "status": {
                "label": "accepted",
                "id": 1
            },
            "accuracy": 60.01,
            "updated_at": "2020-09-05 12:49:51",
            "created_at": "2020-09-05 12:49:51"
        },
        {
            "sub_id": 6,
            "event_id": 1,
            "group_id": 2,
            "problem_id": 1,
            "source_code": "{{AllTheCode}}",
            "language": "go",
            "status": {
                "label": "accepted",
                "id": 1
            },
            "accuracy": 60.861366,
            "updated_at": "2020-09-05 15:38:06",
            "created_at": "2020-09-05 15:38:06"
        },
        {
            "sub_id": 7,
            "event_id": 1,
            "group_id": 2,
            "problem_id": 1,
            "source_code": "{{AllTheCode}}",
            "language": "go",
            "status": {
                "label": "accepted",
                "id": 1
            },
            "accuracy": 94.1104,
            "updated_at": "2020-09-05 15:38:07",
            "created_at": "2020-09-05 15:38:07"
        },
        {
            "sub_id": 8,
            "event_id": 1,
            "group_id": 2,
            "problem_id": 1,
            "source_code": "{{AllTheCode}}",
            "language": "js",
            "status": {
                "label": "accepted",
                "id": 1
            },
            "accuracy": 66.79145,
            "updated_at": "2020-09-05 15:38:11",
            "created_at": "2020-09-05 15:38:11"
        },
        {
            "sub_id": 9,
            "event_id": 1,
            "group_id": 2,
            "problem_id": 1,
            "source_code": "{{AllTheCode}}",
            "language": "c",
            "status": {
                "label": "accepted",
                "id": 1
            },
            "accuracy": 44.333706,
            "updated_at": "2020-09-05 15:38:15",
            "created_at": "2020-09-05 15:38:15"
        },
        {
            "sub_id": 10,
            "event_id": 1,
            "group_id": 2,
            "problem_id": 2,
            "source_code": "{{AllTheCode}}",
            "language": "c",
            "status": {
                "label": "accepted",
                "id": 1
            },
            "accuracy": 60.861366,
            "updated_at": "2020-09-05 15:42:33",
            "created_at": "2020-09-05 15:42:33"
        }
    ]
}
```

##Repo directory structure
```
- db (this folder contains sql db filtes)
- src
    - apperror (http error handller code0
    - coderunner (This package Contains code which help to exec code)
        - go (go executor implementation)
        - js (js executor implementation)
        - js (c executor implementation)
    - constants ( its contains contants used by applicatio)
    - handller ( http handller are implemanted here) 
    - helper ( application helpers) 
    - interface ( application component entity defins here) 
    - middleware ( http middleware define here like. Auth)
    - mocks (mock of application component entity which used in testing)
    - model (All db recated)
    - server (custom http server implemanted here which used by this application)
    - usecase (Independent code which used by handllers)
- vendor (third party lib )
server.go (main Entry Point Of Application)
server (executable stable binary of file)


```

Any Suggestion or complaint reach me out any of contact info.
Email : rajankumar549@gmail.com
Mob : +91-95401-52552

Author : Rajan Kumar