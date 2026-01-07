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
- Each user creates 2 posts per a month;
- On average, each user:
    - Open feeds 10 times per a day;
    - Comment 2 posts per a month;
    - Leaves 5 reactions per a day;
    - Every feed has 2 photos;
    - Comment length 30 simbols;
    - Description length 150 simbols;
    - User subscribe to anouther user 2 times per month;
    - Photo size 2 Mb;
- Feeds view by default have 25 posts;
- Commonwealth of Independent States;
- Seasonality
    - In may, the number of posts increases by 12 percent;
    - In April, the number of posts falls by 7 percent;
- Time limits:
    - Creating post no more than 2 second;
    - Viewing a post no more than 1 second;
    - Rating a post no more than 1 second;
    - Leaving a comment no more than 2 seconds;
    - Feed loading no more than 2 seconds;
    - Searching no more than 2 second;
    - Following / Unfollowing user no more 1 second;

# Basic calculations

## RPS 

Post (write) = 10 000 000 * 2 / 30 / 86 400  ~= 7 RPS

Comments (write) = 10 000 000 * 2 / 30 / 86 400  ~= 7 RPS

Reactions (write) = 10 000 000 * 5 / 86 400  ~= 115 RPS

Subscribe = 10 000 000 * 2 / 30 / 86 400  ~= 7 RPS

Open Feeds (read) = 10 000 000 * 10 / 86400 ~= 1157 RPS

## Traffic

### Create post

```
model
- post_id - 16 bytes
- user_id - 16 bytes
- point_id - 16 bytes
- descriptions - 150*4 = 600 bytes (UTF-8)
- photos - 2 * 2Mb = 4Mb
```


model size = 648 bytes 
media size = 4Mb 

Trafic:
```
7 RPS * 648 = 4536 bytes =~ 4 Kb/s
7 RPS * 4Mb =~ 28 Mb/s
```

### Open one specific post

```
model
- post_id - 16 bytes
- user_id - 16 bytes
- point_id - 16 bytes
- descriptions - 150*4 = 600 bytes (UTF-8)
- photos - 2 * 2Mb = 4Mb
```


model size = 648 bytes
media size = 4Mb

Trafic:
```
1157 RPS * 648 = 749 736 bytes =~ 749 Kb/s
1157 RPS * 4Mb =~ 4 Gb/s
```

### Open feeds

```
model
- post_id - 16 bytes
- user_id - 16 bytes
- point_id - 16 bytes
- descriptions - 150*4 = 600 bytes (UTF-8)
- photos - 2 * 2Mb = 4Mb
```

model size = 648 bytes
media size = 4Mb

Trafic:
```
1157 RPS * 648 * 25 = 18 743 400 bytes =~ 18 Gb/s
1157 RPS * 4Mb * 25 =~ 318 Gb/s
```

### Leave a comment

```
model
user_id = 16B
port_id = 16B
message = 30 * 4 = 120 bytes
```

model size = 152 bytes

Trafic:
```
7 RPS * 152 = 1064 bytes =~ 1 Kb/s
```

### Leave a reaction

```
model
user_id = 16B
port_id = 16B
reaction = 1B
```

size = 33 bytes

Trafic:
```
115 RPS * 33 = 3795 bytes =~ 3 Kb/s
```

### Subscribe

```
model
user_id = 16B
author_id = 16B
```

size = 32 bytes

Trafic:
```
7 RPS * 33 = 231 bytes =~ 0,2 Kb/s
```
