#### Functional

##### Try to run the following command `"./wget https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg"`

###### Did the program download the file `"EMtmPFLWkAA8CIS.jpg"`?

##### Try to run the following command with a link at your choice `"./wget <https://link_of_your_choice.com>"`

###### Did the program download the expected file?

##### Try to run the following command `"./wget https://golang.org/dl/go1.15.linux-amd64.tar.gz"`

###### Did the program download the file `"go1.15.linux-amd64.tar.gz"`?

###### Did the program displayed the start time?

###### Did the start time and the end time respected the format? (yyyy-mm-dd hh:mm:ss)

###### Did the program displayed the status of the response? (200 OK)

###### Did the Program displayed the content length of the download?

###### Is the content length displayed as raw (bytes) and rounded (Mb or Gb)?

###### Did the program displayed the name and path of the file that was saved?

##### Try to download a big file, for example: `"./wget http://ipv4.download.thinkbroadband.com/100MB.zip"`

###### Did the program download the expected file?

###### While downloading, did the progress bar show the amount that is being downloaded? (KiB or MiB)

###### While downloading, did the progress bar show the percentage that is being downloaded?

###### While downloading, did the progress bar show the time that remains to finish the download?

###### While downloading, did the progress bar progressed smoothly (kept up with the time that the download took to finish)?

##### Try to run the following command, `"./wget -O=test_20MB.zip http://ipv4.download.thinkbroadband.com/20MB.zip"`

###### Did the program downloaded the file with the name `"test_20MB.zip"`?

##### Try to run the following command, `"./wget -O=test_20MB.zip -P=~/Downloads/ http://ipv4.download.thinkbroadband.com/20MB.zip"`

###### Can you see the expected file in the "~/Downloads/" folder?

##### Try to run the following command, `"./wget --rate-limit=300k http://ipv4.download.thinkbroadband.com/20MB.zip"`

###### Was the download speed always lower than 300KB/s?

##### Try to run the following command, `"./wget --rate-limit=700k http://ipv4.download.thinkbroadband.com/20MB.zip"`

###### Was the download speed always lower than 700KB/s?

##### Try to run the following command, `"./wget --rate-limit=2M http://ipv4.download.thinkbroadband.com/20MB.zip"`

###### Was the download speed always lower than 2MB/s?

##### Try to create a text file with the name `"downloads.txt"` and save into it the links below. Then run the command `"./wget -i=downloads.txt"`

```
https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg
http://ipv4.download.thinkbroadband.com/20MB.zip
http://ipv4.download.thinkbroadband.com/10MB.zip
```

###### Did the program download all the files from the downloads.txt file? (EMtmPFLWkAA8CIS.jpg, 20MB.zip, 10MB.zip)

###### Did the downloads occurred in an asynchronous way? (tip: look to the download order)

#### Mirror

##### Try to run the following command `"./wget --mirror http://corndog.io/"`, then try to open the `"index.html"` with a browser

###### Is the site working?

##### Try to run the following command `"./wget --mirror https://oct82.com/"`, then try to open the `"index.html"` with a browser

###### Is the site working?

##### Try to run the following command `"./wget --mirror --reject=gif https://oct82.com/"`, then try to open the `"index.html"` with a browser

###### Did the program download the site without the GIFs?

##### Try to run the following command `"./wget --mirror https://trypap.com/"`, then use the command `"ls"` to see the file system of the created folder.

```
css  img  index.html
```

###### Does the created folder has the same fs as above?

##### Try to run the following command `"./wget --mirror -X=/img https://trypap.com/"`, then use the command `"ls"` to see the file system of the created folder.

```
css  index.html
```

###### Does the created folder has the files above?

##### Try to run the following command `"./wget --mirror https://theuselessweb.com/"`

###### Is the site working?

##### Try to run the following command to mirror a website at your choice `"./wget --mirror <https://link_of_your_choice.com>"`

###### Did the program mirror the website?

#### Bonus

###### +Does the project runs quickly and effectively? (Favoring recursive, no unnecessary data requests, etc)

###### +Does the code obey the [good practices](https://public.01-edu.org/subjects/good-practices/README.md)?
