## Lesson 7

Develop a tool which runs a specific command with arguments under specific environment.
The environment variables are located in the folder and looks like a file, where:
 - Name of file - represents an environment variable name
 - File content - represents a value of environment variable

Input parameters:

- path   : path to folder where files are located
- cmd    : Command with argument to run.


Usage: 
```shell
$ goenvdir  <path to env files> <cmd args...>
```

Test:

There are only shell test available. In order to run them please use script below.

```shell
$ ./goenv_test.sh
```
