# gochexec
Check if a path exists and is accessible before launching a command.


You can add this to for example a windows short cut to ensure network resources are available before launching a piece of software.


## Usage
gochexec.exe <path> <command> (arguments)

### CSV

If <path> is a path to a csv file then you can check multiple paths at once by simply having multiple lines like:
```csv
file,c:\path\to\file
file,c:\path\to\other\file
```


## TODO

Support more checks, maybe check for printer resources, ping a host, or make sure webserver is up and not giving 5xx errors.
