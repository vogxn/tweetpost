Tweetpost
=========

Tweetpost can chunk and breakdown tweets intelligently. 

Running
-------
```Bash
$ go get github.com/vogxn/tweetpost
```

```Bash
$ cd vogxn/tweetpost
$ cd server/
$ go run main.go
```

Open [http://localhost:8080/post](http://localhost:8080/post) on your browser
and start splitting tweets

Screenshot
---------
![tweetpost](https://github.com/vogxn/tweetpost/blob/master/server/static/favicons/tweetpost.png "Tweetpost screenshot")

### Note

This app is a demonstration of golang used to build web applications. It should
serve as a simple guide to layout your MVC apps. Most of the server side code
is in the `server` directory.

### OAuth

The OAuth 3-legged authorization in this repository does not work since
Twitter's API does not accept CORS headers. This is because Twitter is still on
the OAuth 1.0 standard while most implementations have moved to the OAuth 2.0
standard

Features
--------
- Simple to use user-interface
- Easy editing before posting
- Selective and Bulk Posting
- Tweet sequence numbers

Coming soon
-----------
- OAuth with Twitter
- Post the actual tweets
- Scheduled posts support
- URL shortening using bit.ly, goo.gl and youtu.be?
- Chain tweet reading support
- User history?


