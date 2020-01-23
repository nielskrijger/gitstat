# gitstat

This tool generates a JSON logfile of one or more git repositories. This logfile is intended for [gitstat.com](https://gitstat.com).

## How to use?

1. Download the zip archive from the [releases](https://github.com/nielskrijger/gitstat/releases).
2. Extract the binary.
3. Run the following on linux/mac:

    ```
    $ cd <EXTRACTED_PATH>
    $ ./gitstat ../project-1 ../project-2
    ```
    
   Or on windows:
   
   ```
   $ dir <EXTRACTED_PATH>
   $ .\gitstat.exe ..\project-1 ..\project2
   ```
   
4. Upload the generated logfile to [gitstat.com](https://gitstat.com).

## Todo

See the Kanban board [here](https://github.com/nielskrijger/gitstat-web/projects/1).

## FAQ

__Is the data kept private?__

Yes. The [gitstat.com](https://gitstat.com) website runs within your browser and doesn't store the logfile remotely or even locally. As a result if you refresh the page you'll have to submit the logfile again.

The only data that is stored within your browser are the config settings.

__I miss feature X or found a bug__

If you found a bug or have a feature suggestion, please open an issue [here](https://github.com/nielskrijger/gitstat/issues). Issues/features related to the website should be posted in the [gitstat-web repository](https://github.com/nielskrijger/gitstat-web). 

__Can I contribute?__

Yes! Please open a PR. If you want to do significant work I'd recommend opening an issue first, share some thoughts before you invest a lot of your time.
