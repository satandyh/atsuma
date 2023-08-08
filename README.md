# **AT**tack **SU**rface **MA**nagement

Just another one system for get knowledge about how to attack your lovely org.

Atsuma expands bash pipeline approach by adding data insert and store in sqlite3 database. Why do not use pure bash with functions and scripts instead? Because I curse myself when try to use files instead of database:
- you can't add metadata to file
- files have no any relationship between each other

## Requirements

- bash - because we reuse it inside
- sqlite - only if you need to find some stored data by yourself
- all other apps that you want to start thru atsuma

## Usage

Atsuma works like server software - always in memory. But anyway it's stateless app. So you can start two-three and even more instances at same time. It's something like automatization software. So pipelines there require more options than usual. Here are them:
- command
- in
- out
- trigger
- alert
- metadata
- lifetime

# TODO for search in database need use simple sql states - need to create simple syntax for SQL SELECT statements

# search and metadata opt
When task try to get some data from database it always will look the last result, and only after process it's state (if it need).

# session control and owners
each instance create session id and give counter of liveness. if counter of liveness more than x sec than session marked as non active. after y*x sec that none active session dropped. this work do every 3 active app when they try to update their live state!
