# Blog aggregator using go

To run this project you need to have a go installation and a postgresql server in your machine. When.
If you have go install you can do `go install .` to have the blog aggregator cli in your enviroment. 

You will notice that you need a .gatorconfig.json in your home directory it should look something like this:

{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":"victor"}

The current_user_name will be modify when login in.

## Available commands

* login (username): modify the .gatorconfig.json and will set up a user session.
* register (username): create a user in the database and will log in
* reset: deletes every entry from the user (!! danger !!)
* users: displays every user in the database
* agg (timer): updates the feeds and posts database in a timer
* addfeed (name) (url): creates a feed globally, and follows it for the current user
* feeds: displays all the feeds in the database
* follow (url): follows an existing feed
* following: displays the user's followed feeds
* unfollow (url): unfollows a feed for the current user
* browse (limit): displays the latest posts with an optional limit (default: 2)
