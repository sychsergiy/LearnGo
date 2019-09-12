# Task
The popular web comic xkcd has a JSON interface. For example, a request to
https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd
that, using this index, prints the URL and transcript of each comic that matches a search term
provided on the command line.

# Start
to build cli tool(`xkcd` binary file output):
```bash
go run build .
```
to create download all comics to offline index:
```bash
./xkcd create-offline-index
```

to create search index from offline index:
```bash
./xkcd create-search-index
```

to search:
```bash
./xkcd search query
```