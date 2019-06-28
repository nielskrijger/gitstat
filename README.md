# gitstat

Outputs a log of git commits & stats in json. This json file is intended for analytic tools for further processing.

## How to use?

TODO

## Why this lib?

There are plenty of git statistics programs out there, but for all of them I found myself doing a lot of post-processing.

This post-processing consisted of two steps usually:

- Ignore unwanted changes (e.g. ignore certain filetypes; extreme commits due to linter change; ...)
- Make diagrams, overviews, etc

Being a web developer