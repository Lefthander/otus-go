## Lesson 6

Develop a tool similar to Unix 'dd' tool

Input parameters:

- from   : Source file
- to     : Destination file
- offset : Offset in the source file
- limit  : limit of bytes to read from the source file

Negative limit does not accepted. In case of limit is greater then size of source file the source file will be copied completely.
In case of offset is negative the read pointer will be set regarding end of file. 
Usage: 
```shell
$ gocopy --from <source file> --to <destination file> [ --offset <offset in source file>] [ --limit <amount of bytes to copy>]
```
In case of optional parameters limit and offset are ommited the all from source will be copied to destination file.