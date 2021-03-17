# Plan
- have a cronjob that runs every day at midnight to retrieve work intervals
- it then schedules a cronjob every 45th min of each working hour to run another function
- the function it schedules sends a text reminder to meditate and
- hits an endpoint of the api to update todays extra minutes meditated and
- the total extra minutes meditated using micro meditations
- which is displayed on a one-page web app that has a description of the project