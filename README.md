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
- Each user creates 2 posts per a day;
- On average, each user:
    - Open feeds 10 times per a day;
    - Comment 20 posts per a month;
    - Every post at least have 2 comments;
    - Leaves 5 reactions per a day;
    - Every feed has 8 photos;
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

Post (write) = 10 000 000 * 2 / 86 400  ~= 232 RPS

Comments (write) = 10 000 000 * 20 / 30 / 86 400  ~= 78 RPS

Comments (read - need to open feeds) = 10 000 000 * 10 * 25 * 2 / 86 400  ~= 578 RPS

Reactions (write) = 10 000 000 * 5 / 86 400  ~= 115 RPS

Reactions (read) = 10 000 000 * 5 * 10 * 25 / 86 400  ~= 290 RPS

Subscribe (write) = 10 000 000 * 2 / 30 / 86 400  ~= 7 RPS

Open Feeds (read) = 10 000 000 * 10 / 86400 ~= 1157 RPS

## Traffic

### Create post

```
model
- post_id - 16 bytes
- user_id - 16 bytes
- point_id - 16 bytes
- descriptions - 150*4 = 600 bytes (UTF-8)
- photos - 8 * 2Mb = 16Mb
```


model size = 648 bytes 
media size = 16Mb 

Trafic:
```
232 RPS * 648 = 4536 bytes =~ 146 Kb/s
232 RPS * 16Mb =~ 4 Gb/s
```

### Open feeds

```
model
- post_id - 16 bytes
- user_id - 16 bytes
- point_id - 16 bytes
- descriptions - 150*4 = 600 bytes (UTF-8)
- photos - 8 * 2Mb = 16Mb
```

model size = 648 bytes
media size = 16Mb

Trafic:
```
1157 RPS * 648 * 25 = 18 743 400 bytes =~ 18 Gb/s
1157 RPS * 16Mb * 25 =~ 452 Gb/s
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
78 RPS * 152 = 1064 bytes =~ 12 Kb/s
```

### Read a comments

```
model
user_id = 16B
port_id = 16B
message = 30 * 4 = 120 bytes
```

model size = 152 bytes

Trafic:
```
578 RPS * 152 = 1064 bytes =~ 85 Kb/s
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

### Open reactions

```
model
user_id = 16B
port_id = 16B
reaction = 1B
```

size = 33 bytes

Trafic:
```
290 RPS * 33 = 3795 bytes =~ 9.5 Kb/s
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

## Resources

We will calculate resources for HDD and SSD. Let's take the average value from the table below:

| Type | Capacity  | IOps | Throughput  |
|------|-----------|-------|-------|
| HDD  | <= 32 Tb  | 100   | 100 Mb/s  |
| SSD(SATA)  | <= 100 Tb | 1000  | 1000 Mb/s  |
| SSD(nVME)  | <= 30 Tb | 10000  | 3 Gb/s  |


### Media

Dayly capacity (calculate all read operations) = `4 Gb/s * 2 * 86400 s = 691200 GB/day = 675 TB/day`

Yearly capacity = `675 TB/day * 365 = 240 PB/year`

#### HDD
```
Disks_for_capacity =  240 PB / 32 TB = 7680 
Disks_for_throughput = (4 Gb/s + 452 Gb/s) / 100 Mb/s = 467 
Disks_for_iops = (1157 + 232) / 100 = 14 
```

Disks = 7680

#### SSD(SATA)
```
Disks_for_capacity = 240 PB  / 100 TB = 1229 
Disks_for_throughput = (4 Gb/s + 452 Gb/s) / 1000 Mb/s = 3256 
Disks_for_iops = (1157 + 232) / 1000 = 2
```

Disks = 3256

#### SSD(nVME)
```
Disks_for_capacity = 240 PB / 30 TB = 7650 
Disks_for_throughput = (4 Gb/s + 452 Gb/s) / 3 Gb/s = 152 
Disks_for_iops = (1157 + 232) / 10000 = 1
```
Disks = 7650

#### Conclusion

Media is a specific type of data which need a object storage db to handle it efficiently (S3 or Ceph). It's not necessary to use SSDs for this task, but it would be better because it would allow us to handle more traffic and use lowest amount of disks.

- async master-slave replication with RF= 4
- `Hosts = disks / disks_per_host  = 3256 / RAID5 16 disks (as the most effective configuration) = 230`
- `Hosts_with_replication = hosts * replication_factor = 230 * 4 = 920`
- key-based sharding by location_id

### Posts

Dayly capacity (calculate all read operations) = `146 Kb/s * 2 * 86400 s = 24 GB/day`

Yearly capacity = `24 GB/day * 365 = 8,5 PB/year`

#### HDD
```
Disks_for_capacity =  8,5 PB / 32 TB = 272 
Disks_for_throughput = (18 Gb/s + 146 Kb/s) / 100 Mb/s = 184 
Disks_for_iops = (1157 + 232) / 100 = 14 
```

Disks = 272

#### SSD(SATA)
```
Disks_for_capacity = 8,5 PB  / 100 TB = 87 
Disks_for_throughput = (18 Gb/s + 146 Kb/s) / 1000 Mb/s = 18
Disks_for_iops = (1157 + 232) / 1000 = 2
```

Disks = 87

#### SSD(nVME)
```
Disks_for_capacity = 8,5 PB / 30 TB = 290 
Disks_for_throughput = (18 Gb/s + 146 Kb/s) / 3 Gb/s = 6
Disks_for_iops = (1157 + 232) / 10000 = 1
```
Disks = 290

#### Conclusion

Due to the peculiarities of the data structure, it can be stored in a relational database as PostgreSQL.

- async master-slave replication with RF=3
- `Hosts = disks / disks_per_host  = 87 / RAID5 16 disks = 8`
- `Hosts_with_replication = hosts * replication_factor = 8 * 3 = 24`
- key-based sharding by user_id

### Comments

Dayly capacity (calculate all read operations) = `85 Kb/s * 86400 s = 7 GB/day`

Yearly capacity = `7 GB/day * 365 = 2,5 PB/year`

#### HDD
```
Disks_for_capacity =  2,5 PB / 32 TB = 79 
Disks_for_throughput = (12 Gb/s + 85 Kb/s) / 100 Mb/s = 1 
Disks_for_iops = (78 + 578) / 100 = 7
```

Disks = 79

#### SSD(SATA)
```
Disks_for_capacity = 2,5 PB  / 100 TB = 26 
Disks_for_throughput = (12 Gb/s + 85 Kb/s) / 1000 Mb/s = 1
Disks_for_iops = (78 + 578) / 1000 = 1
```

Disks = 26

#### SSD(nVME)
```
Disks_for_capacity = 2,5 PB / 30 TB = 85
Disks_for_throughput = (12 Gb/s + 85 Kb/s) / 3 Gb/s = 1
Disks_for_iops = (78 + 578) / 10000 = 1
```
Disks = 85

#### Conclusion

Due to the peculiarities of the data structure, it can be stored in a relational database as PostgreSQL.

- async master-slave replication with RF= 2
- `Hosts = disks / disks_per_host  = 26 / RAID5 16 disks = 4`
- `Hosts_with_replication = hosts * replication_factor = 4 * 2 = 8`
- sharding seems redundant at the moment, but if it is needed it should be used key-based sharding by post_id

### Reactions

Dayly capacity (calculate all read operations) = `9.5 Kb/s * 86400 s = 801 MB/day`

Yearly capacity = `801 MB/day * 365 = 285 GB/year`

#### HDD
```
Disks_for_capacity =  285 GB / 32 TB = 1 
Disks_for_throughput = (9.5 Kb/s + 3 Kb/s) / 100 Mb/s = 1 
Disks_for_iops = (290 + 115) / 100 = 4
```

Disks = 4

#### SSD(SATA)
```
Disks_for_capacity = 285 GB  / 100 TB = 26 
Disks_for_throughput = (9.5 Kb/s + 3 Kb/s) / 1000 Mb/s = 1
Disks_for_iops = (290 + 115) / 1000 = 1
```

Disks = 26

#### SSD(nVME)
```
Disks_for_capacity = 285 GB  / 30 TB = 85
Disks_for_throughput = (9.5 Kb/s + 3 Kb/s)  / 3 Gb/s = 1
Disks_for_iops = (290 + 115) / 10000 = 1
```
Disks = 85

#### Conclusion

Due to the peculiarities of the data structure, it can be stored in a relational database as PostgreSQL.

- async master-slave replication with RF=2
- `Hosts = disks / disks_per_host  = 4 / 2 = 2`
- `Hosts_with_replication = hosts * replication_factor = 2 * 2 = 4`
- sharding seems redundant at the moment, but if it is needed it should be used key-based sharding by post_id

### Subscriptions


Dayly capacity (calculate all read operations) = `0.2 Kb/s * 86400 s = 16 MB/day`

Yearly capacity = `16 MB/day * 365 = 6 GB/year`

#### HDD
```
Disks_for_capacity =  6 GB/ 32 TB = 1 
Disks_for_throughput = 0.2 Kb/s / 100 Mb/s = 1 
Disks_for_iops = 7 / 100 = 1
```

Disks = 1

#### SSD(SATA)
```
Disks_for_capacity =  6 GB / 100 TB = 1
Disks_for_throughput = 0.2 Kb/s / 1000 Mb/s = 1
Disks_for_iops = 7/ 1000 = 1
```

Disks = 1

#### SSD(nVME)
```
Disks_for_capacity =  6 GB / 30 TB = 1
Disks_for_throughput = 0.2 Kb/s / 3 Gb/s = 1
Disks_for_iops = 7 / 10000 = 1
```
Disks = 1

#### Conclusion

Due to the peculiarities of the data structure, it can be stored in a relational database as PostgreSQL.

- async master-slave replication with RF=2
- `Hosts = disks / disks_per_host  = 1 / 2 = 1`
- `Hosts_with_replication = hosts * replication_factor = 1 * 2 = 2`
- ыharding seems redundant at the moment, but if it is needed, should be used key-based sharding by user_id
