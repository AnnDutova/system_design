# SystemDesign of social network

Homework for course of System Design. A web service for a social network for travelers demonstrates the candidate's knowledge of creating highly scalable systems.

## Functional requirements
- Publishing travel posts. Travel posts can include:
    - photos;
    - small descriptions;
    - geo-map;
- Add to other travelers' posts:
    - rating;
    - comments;
- Subscribing / Unsubscribing to other travelers' blogs;
- Search popular places;
- Search posts with popular places;
- View other travelers' feeds and the user's feed based on subscriptions in reverse chronological order;

## Non-functional requirements
- DAU = 10_000_000
- Availability 99,95%
- Posts always stored;
- Each user creates 1 posts per a day;
- On average, each user:
    - Open feeds 10 times a day;
    - Comment 2 posts a day;
    - Leaves 5 reactions a day;
- Feeds view by default have 25 posts;
- Commonwealth of Independent States;
- No seasonality;

# Basic calculations
RPS (write):
```
DAU = 10 000 000
Each user creates 2 comments and leaves 5 reaction for posts
RPS = 10 000 000 * 2 * 5 / 86 400  ~= 115 
```
RPS (read):
```
DAU = 10 000 000
Each user open feeds 10 times
Maximum RPS = 10 000 000 * 10 / 86400 ~= 1157
```
Post model:
- post_id - 8 bytes
- user_id - 8 bytes
- point_id - 8 bytes
- descriptions - 100 bytes
- photo_id - 8 bytes

size = 150 bytes 


Trafic (write):
```
115 * 150 = 17250 bytes =~ 17Kb
```
Trafic (read):
```
1157 * 150 = 173550 bytes =~ 17Mb
```

