# Laracasts downloader

This is a small utility to download all lessons from http://laracasts.com
written in [Go](http://golang.org/) adapted from
[LeonB's laracasts-dl](https://github.com/LeonB/laracasts-dl) with some updates
to and improvements.

Please note that you must have an account on laracasts to use this utility.

## Running the scraper

Download the appropriate pre-compiled package for your system and run it:

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
$ ./laracasts-downloader myusername mypassword
2014/11/18 23:56:03 Logging in...
2014/11/18 23:56:04 Building list of available lessons...
2014/11/18 23:56:13 Found 359 lessons
2014/11/18 23:56:13 Checking lesson 1/359 (limit your instance variables)
39.39 MB / 39.39 MB [==========================================================================] 100.00 % 2.98 MB/s 13s
2014/11/18 23:56:27 Checking lesson 2/359 (say hello to laravel homestead 20)
26.31 MB / 26.31 MB [===========================================================================] 100.00 % 2.90 MB/s 9s
2014/11/18 23:56:36 Checking lesson 3/359 (one level of indentation)
30.36 MB / 30.36 MB [==========================================================================] 100.00 % 2.81 MB/s 10s
2014/11/18 23:56:48 Checking lesson 4/359 (dont use else)
40.71 MB / 40.71 MB [==========================================================================] 100.00 % 2.52 MB/s 16s
2014/11/18 23:57:04 Checking lesson 5/359 (no abbreviations)
14.25 MB / 14.25 MB [===========================================================================] 100.00 % 2.72 MB/s 5s
//
...
//
2014/11/19 01:06:40 Checking lesson 353/359 (observables)
32.63 MB / 32.63 MB [==========================================================================] 100.00 % 2.34 MB/s 13s
2014/11/19 01:06:55 Checking lesson 354/359 (many to many relations)
35.42 MB / 35.42 MB [==========================================================================] 100.00 % 2.06 MB/s 17s
2014/11/19 01:07:13 Checking lesson 355/359 (easy asset management)
29.80 MB / 29.80 MB [==========================================================================] 100.00 % 2.24 MB/s 13s
2014/11/19 01:07:27 Checking lesson 356/359 (easy environments)
34.15 MB / 34.15 MB [==========================================================================] 100.00 % 2.29 MB/s 14s
2014/11/19 01:07:43 Checking lesson 357/359 (mailers)
51.92 MB / 51.92 MB [==========================================================================] 100.00 % 2.14 MB/s 24s
2014/11/19 01:08:09 Checking lesson 358/359 (atom feeds decoded)
47.16 MB / 47.16 MB [==========================================================================] 100.00 % 1.91 MB/s 24s
2014/11/19 01:08:35 Checking lesson 359/359 (from zero to deploy)
44.42 MB / 44.42 MB [==========================================================================] 100.00 % 2.24 MB/s 19s
2014/11/19 01:08:55 Finished in 1h12m52.8652351s.
```
