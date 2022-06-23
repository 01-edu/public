## clonernews

### Objectives
Technology is a rapidly evolving sector. As a programmer, it is not always easy to keep up to date with every new advancement.

Tech news (like Hacker News) is a really great way to keep up to date with the exponential evolution of technology, as well as tech jobs and much more. Some websites do not offer very appealing propositions to consume their media.

So your objective for this raid is to create an UI for the [HackerNews API](https://github.com/HackerNews/API).

You must handle at least:
- Posts: including:
  - [stories](https://github.com/HackerNews/API#ask-show-and-job-stories)
  - [jobs](https://github.com/HackerNews/API#ask-show-and-job-stories)
  - [polls](https://github.com/HackerNews/API#items)
- [comments](https://github.com/HackerNews/API#items): each comment must have the proper post parent.

Post and comments must be ordered by newest post to oldest.
You must not load all posts at once, so that you only load posts once the users need to see more posts. This can be done with the help of [events](https://developer.mozilla.org/en-US/docs/Web/Events).

The point of the project is to keep users updated, so we'll need to inform our users of changes to the data using [Live Data](https://github.com/HackerNews/API#live-data). Create a section that presents the newest information. You'll need to notify the user at least every 5 seconds, whenever the live data is updated.

Currently this API does not present a [rate limit](https://en.wikipedia.org/wiki/Rate_limiting), but that does not mean that you should abuse or overload the API.

Best ways you can avoid rate limiting:
- optimize your code to eliminate any unnecessary requests
- usage of a throttling/debouncing function to regulate the number of requests,

### Optional

You can handle sub-comments for stories, jobs and polls, by implementing nested comments.
