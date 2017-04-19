Tweetpost
=========

Tweetpost can chunk and breakdown tweets intelligently. 

Running
-------
```Bash
$ go get github.com/vogxn/tweetpost
$ cd $GOPATH/src/github.com/vogxn/tweetpost
$ cd server/
$ go run main.go
```

Open (http://localhost:8080/post)[http://localhost:8080/post] on your browser
and start splitting tweets

Screenshot
---------
![tweetpost](https://github.com/vogxn/tweetpost/blob/master/server/static/favicons/tweetpost.png "Tweetpost screenshot")

### Note
This app is a demonstration of golang used to build web applications. It should
serve as a simple guide to layout your MVC apps. Most of the server side code
is in the server directory.

Features
--------
- Simple to use user-interface
- Easy editing before posting
- Selective and Bulk Posting
- Tweet sequence numbers

Coming soon
-----------
- Scheduled posts support
- URL shortening using bit.ly, goo.gl and youtu.be?
- Chain tweet reading support
- User history?


