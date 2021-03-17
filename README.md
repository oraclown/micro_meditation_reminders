# Plan
- have a cronjob that runs every day at midnight to retrieve work intervals
- it then schedules a cronjob every 45th min of each working hour to run another function
- the function it schedules sends a text reminder to meditate and
- hits an endpoint of the api to update todays extra minutes meditated and
- the total extra minutes meditated using micro meditations
- which is displayed on a one-page web app that has a description of the project

# TODO
- Function that runs a cronjob each day at midnight which runs the next four functions (this will probs be in the main func of the webapp?).
- Function for retrieving all the day's calendar events.
- Function to get all work intervals from the day's cal events.
- Function that takes all the day's work intervals and gets all the times needed to send reminders.
- Function that takes all the times for reminders and schedules cron jobs at those times to run function that sends text reminders to meditate & updates webapp meditation stats.
- Function that runs the next two functions.
- Function that sends a text reminder to meditate.
- Function that hits api endpoint which updates webapp meditation stats.
- organize repo like snippetbox? ask others for common Go webapp organization
- make nicer readme w/ phone screen record gif
- 