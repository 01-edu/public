## wget

### Objectives

This project objective consists on recreating some functionalities of [`wget`](https://www.gnu.org/software/wget/manual/wget.html) using **Go**.

These functionalities will consist in:

- The normal usage of `wget`: downloading a file given an URL, example: `wget https://some_url.ogr/file.zip`
- Downloading a single file and saving it under a different name
- Downloading and saving the file in a specific directory
- Set the download speed, limiting the rate speed of a download
- Continuing interrupted downloads
- Downloading a file in background
- Downloading multiple files at the same time, by reading a file containing multiple download links asynchronously
- Main feature will be to download an entire website, [mirroring a website](https://en.wikipedia.org/wiki/Mirror_site)

### Introduction

Wget is a free utility for non-interactive download of files from the Web. It supports HTTP, HTTPS, and FTP protocols, as well as retrieval through HTTP proxies.

To see more about wget you can visit the manual by using the command `man wget`, or you can visit the website [here](https://www.gnu.org/software/wget/manual/wget.html).

#### Usage

Your program must have as arguments the link from where you want to download the file, for instance:

```console
student@student$ ./wget https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg
```

The program should be able to give feedback, displaying the:

- Time that the program started: it must have the following format **yyyy-mm-dd hh:mm:ss**
- Status of the request. For the program to proceed to the download, it must present a response to the request as status OK (`200 OK`) if not, it should say which status it got and finish the operation with an error warning
- Size of the content downloaded: the content length can be presented as raw (bytes) and rounded to Mb or Gb depending on the size of the file downloaded
- Name and path of the file that is about to be saved
- A progress bar, having the following:
  - A amount of `KiB` or `MiB` (depending on the download size) that was downloaded
  - A percentage of how much was downloaded
  - Time that remains to finish the download
- Time that the download finished respecting the previous format

It should look something like this

```console
student@student$ go run main.go https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg
start at 2017-10-14 03:46:06
sending request, awaiting response... status 200 OK
content size: 56370 [~0.06MB]
saving file to: ./EMtmPFLWkAA8CIS.jpg
 55.05 KiB / 55.05 KiB [================================================================================================================] 100.00% 1.24 MiB/s 0s

Downloaded [https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg]
finished at 2017-10-14 03:46:07
```

#### Flags

Your program should be able to handle different flags.

1. Download a file and save it under a different name by using the flag `-O` followed by the name you wish to save the file, example:

```console
student@student$ go run main.go -O=meme.jpg https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg
start at 2017-10-14 03:46:06
sending request, awaiting response... status 200 OK
content size: 56370 [~0.06MB]
saving file to: ./meme.jpg
 55.05 KiB / 55.05 KiB [================================================================================================================] 100.00% 1.24 MiB/s 0s

Downloaded [https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg]
finished at 2017-10-14 03:46:07
student@student$ ls -l
-rw-r--r-- 1 student student 56370 ago 13 16:59 meme.jpg
-rw-r--r-- 1 student student 11489 ago 13 10:28 main.go
```

---

2. It should also handle the path to where your file is going to be saved using the flag `-P` followed by the path to where you want to save the file, example:

```console
student@student$ go run main.go -P=~/Downloads/ -O=meme.jpg https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg
start at 2017-10-14 03:46:06
sending request, awaiting response... status 200 OK
content size: 56370 [~0.06MB]
saving file to: ~/Downloads/meme.jpg
 55.05 KiB / 55.05 KiB [================================================================================================================] 100.00% 1.24 MiB/s 0s

Downloaded [https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg]
finished at 2017-10-14 03:46:07
student@student$ ls -l ~/Downloads/meme.jpg
-rw-r--r-- 1 student student 56370 ago 13 16:59 /home/student/Downloads/meme.jpg
```

---

3. The program should handle speed limit. Basically the program can control the speed of the download by using the flag `--rate-limit`. If you download a huge file you can limit the speed of your download, preventing the program from using the full possible bandwidth of your connection, example:

```console
student@student$ go run main.go --rate-limit=400k https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg
```

This flag should accept different value types, example: k and M. So you can put the rate limit as `rate-limit=200k` or `rate-limit=2M`

---

4. Downloading different files should be possible. For this the program will receive the `-i` flag followed by a file name that will contain all links that are to be downloaded. Example:

```console
student@student$ ls
download.txt   main.go
student@student$ cat download.txt
http://ipv4.download.thinkbroadband.com/20MB.zip
http://ipv4.download.thinkbroadband.com/10MB.zip
student@student$ go run main -i=download.txt
content size: [10485760, 20971520]
finished 10MB.zip
finished 20MB.zip

Download finished:  [http://ipv4.download.thinkbroadband.com/20MB.zip http://ipv4.download.thinkbroadband.com/10MB.zip]

```

The Downloads should work asynchronously, it should download both files at the same time. You are free to display what you want for this option.

---

5. [**Mirror a website**](https://en.wikipedia.org/wiki/Mirror_site). This option should download the entire website being possible to use "part" of the website offline and for other useful [reasons](https://www.quora.com/How-exactly-does-Mirror-Site-works-and-how-it-is-done). For this you will have to download the website file system and save it into a folder that will have the domain name. Example: `http://www.example.com`, will be stored in a folder with the name `www.example.com` containing every file from the mirrored website. The flag should be `--mirror`.

To mirror a website you will have to implement the following `wget` flags so that the web mirror is complete (you do not need to do the literal flags, but just the theory behind it, so your flag `--mirror` need to behave like the following wget flags combined):

- [`--mirror`](https://www.gnu.org/software/wget/manual/wget.html): download recursive
- [`--convert-links`](https://www.gnu.org/software/wget/manual/wget.html): after the download is complete it will convert all links in the document to make them suitable for local viewing
- [`--page-requisites`](https://www.gnu.org/software/wget/manual/wget.html): downloads all files that are necessary to properly display a given HTML page
- [`--no-parent`](https://www.gnu.org/software/wget/manual/wget.html): this will not let the program ascend to the parent directory when retrieving

### Hint

You can take a look into the [html package](https://godoc.org/golang.org/x/net/html) for some help

---

This project will help you learn about:

- GNU Wget
- HTTP
- [FTP](https://en.wikipedia.org/wiki/File_Transfer_Protocol)
- Algorithms
- Mirror websites
- File system(fs)
