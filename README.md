# gochexec
Check if a path exists and is accessible before launching a command.


You can add this to for example a windows short cut to ensure network resources are available before launching a piece of software.


## Usage
gochexec.exe &lt;path&gt; &lt;command&gt; (arguments)

### CSV

If <path> is a path to a csv file then you can check multiple paths at once by simply having multiple lines like:
```csv
file,c:\path\to\file
file,c:\path\to\other\file
```
#### DNS

A csv file may contain dns lines like: `dns,<domain>`.
At this time it only uses the system resolver.

#### Socket

A csv file may contain socket lines like: `sock,<network>,<address>,<timeout>`

Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".

Address is a value relevant to your network type, for example `192.168.1.2:80`

Timeout is a value given in seconds.
