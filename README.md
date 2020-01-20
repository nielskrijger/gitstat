# gitstat

Generates a logfile of git commits and statistics. This file is intended for use on [gitstat.com](https://gitstat.com).

## How to use?

1. Download the zip archive from the [releases](https://github.com/nielskrijger/gitstat/releases).
2. Extract the binary.
3. Run the following on linux/mac:

    ```
    $ dir <EXTRACTED_PATH>
    $ ./gitstat ../project-1 ../project-2 <etc>
    ```
    
   Or on windows:
   
   ```
   $ dir <EXTRACTED_PATH>
   $ .\gitstat.exe ..\project-1 ..\project2 <etc>
   ```
   
4. Upload the generated file to [gitstat.com](https://gitstat.com).

## Why?

There are plenty of git statistics programs out there but usually I have to do quite a bit of pre- and post-processing of the data to get the information I'm looking for. This processing usually entailed:

- Eliminating user aliases
- Eliminating unwanted changes, for example:
    - limit to certain file types;
    - exclude certain directories;
    - remove specific large commits (usually caused by an accident or a lint auto formatter);
    - ignore merge commits;
    - ...
- Making diagrams, tables, overviews, etc

The goal of [gitstat.com](https://gitstat.com) is to make such git analysis easier.

However...! Right now the website is very limited feature-wise, likely other git statistics tools will serve your needs better.

## Todo

- [x] support multiple code repositories
- [ ] add homebrew
- [ ] add change stats for file renames

## FAQ

__I miss feature X or found a bug__

If you found a bug or feature suggestions, please open a an issue [here](https://github.com/nielskrijger/gitstat/issues). Issues/features related to the website should be posted in the [gitstat-web repository](https://github.com/nielskrijger/gitstat-web). 

__Can I contribute?__

Yes! Please open a PR. If you want to do significant work I'd recommend opening an issue first, share some thoughts before you invest a lot of your time.
