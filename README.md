# Laracasts downloader

This is a small utility to download all lessons from http://laracasts.com
written in [Go](http://golang.org/) adapted from
[LeonB's laracasts-dl](https://github.com/LeonB/laracasts-dl) with some updates
to and improvements.

Please note that you must have an account on laracasts to use this utility.

## Running the scraper

Download the appropriate pre-compiled package for your system and run it (coming soon):

``` bash
./laracasts-downloader myusername mypassword dest/
```
Run from source:

``` bash
go run laracasts-downloader.go myusername mypassword dest/
```

To build from source:

``` bash
go build laracasts-downloader.go
./laracasts-downloader myusername mypassword dest/
```

## Example

``` bash
$ laracasts-downloader myusername mypassword
2014/11/17 23:47:50 Logging in
2014/11/17 23:47:50 Building list of available lessons...
2014/11/17 23:47:51 Found 348 lessons
2014/11/17 23:47:51 Downloading lesson 1/348 (A Tour of the Laracasts Source)
50.36 MB / 50.66 MB [===================================] 99.41 % 3.22 MB/s 2014/11/17 23:48:08
Downloading lesson 2/348 (Important Breaking Change in 4.1.26)
12.38 MB / 12.44 MB [===================================] 99.52 % 3.25 MB/s 2014/11/17 23:48:12
Downloading lesson 3/348 (Maybe You Should Use SQLite)
20.53 MB / 20.90 MB [===================================] 98.21 % 3.30 MB/s 2014/11/17 23:48:20
Downloading lesson 4/348 (Enforcement, Entities, and Eloquent)
46.26 MB / 46.74 MB [===================================] 98.98 % 2.72 MB/s 2014/11/17 23:48:38
Downloading lesson 5/348 (PHP 5.6 in 10 Minutes)
15.40 MB / 15.93 MB [===================================] 96.64 % 3.20 MB/s 2014/11/17 23:48:43
Downloading lesson 6/348 (Entities vs. Value Objects)
40.88 MB / 41.39 MB [===================================] 98.76 % 3.24 MB/s 2014/11/17 23:48:57
Downloading lesson 7/348 (Supervise This)
9.98 MB / 10.15 MB [====================================] 98.30 % 3.32 MB/s 2014/11/17 23:49:01
Downloading lesson 8/348 (The Failed Job Interrogation)
28.19 MB / 28.48 MB [===================================] 98.99 % 3.20 MB/s 2014/11/17 23:49:11
Downloading lesson 9/348 (Beanstalkd Queues)
32.30 MB / 32.70 MB [===================================] 98.79 % 3.22 MB/s 2014/11/17 23:49:22
Downloading lesson 10/348 (How to use Eloquent Outside of Laravel)
16.47 MB / 16.80 MB [===================================] 98.06 % 3.16 MB/s 2014/11/17 23:49:28
Downloading lesson 11/348 (Form Macros for the Win)
50.83 MB / 83.88 MB [=================>-----------------] 60.60 % 2.98 MB/s 11s
```
