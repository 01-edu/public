## clonernews

### Objectives

Technology nowadays continue to evolve. As a programmer or developer you must be updated with this exponential evolution of technology.

This is where tech news shines, just like Hacker News where you can see all about new technology, jobs and much more, but some websites do not perform very well or their are not so appealing.

So your objective for this raid is to create an UI for [`HackerNewsAPI`](https://github.com/HackerNews/API).

You must handle at least:

- Posts, this includes :
  - [stories](https://github.com/HackerNews/API#ask-show-and-job-stories)
  - [jobs](https://github.com/HackerNews/API#ask-show-and-job-stories)
  - [polls](https://github.com/HackerNews/API#items)
- [comments](https://github.com/HackerNews/API#items), each comment must have the proper post parent.

Post and comments must be ordered by newest post to oldest.
You must not load all posts at once. You must be able to load posts when ever the users need to see more posts. This can be done with the help of [event](https://developer.mozilla.org/en-US/docs/Web/Events).

[Live Data](https://github.com/HackerNews/API#live-data) : is one of the features from this Hacker News API, you can handle requests so that the news you provide are updated.

- You must have a section that present the newest information. You will have to notify the user at least every 5 seconds, whenever the live data is updated. In other words, after every 5 seconds if a change was made in the live data, you have to notify the user.

Currently this API does not present [rate limit](https://en.wikipedia.org/wiki/Rate_limiting). But that does not mean you should abuse/overload the API!!!

Best ways you can avoid rate limiting :

- optimize your code to eliminate any unnecessary requests
- usage of throttling/debouncing function to regulates the amount of requests

### Optional

You can handle sub-comments for each stories, jobs and polls this meaning nested comments
