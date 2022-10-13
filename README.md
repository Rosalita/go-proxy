# go-proxy

Learning about proxies.

A forward proxy sits in front of clients and masks them.
Users -> proxy -> internet -> server
A forward proxy is the most common form of a proxy server.
Common use for a forward proxy is to pass requests from an isolated private network
to the internet through a firewall.

A reverse proxy sits in front of backend servers and ensures that no client
ever communicates with them directly.
Users -> internet -> proxy -> server
Common uses for reverse proxy include load balancing, it distributes client requests
across a group of servers. Web acceleration, can compress inbound and outbound data, as
well as caching to speed up the flow of traffic between clients and servers.
Revers proxy can perform additional tasks such as SSL encryption to take load off webservers
and boost their performance. Reverse proxies can intercept requests for security, protecting 
the identity of the servers and being an additional defence against attacks.

TCP is transport control protocol and allows applications and devices to exchange messages 
over a network. It is designed to send packets across the internet and ensure successful
delivery of data and messages over networks. The IP in TCP/IP stands for internet protocol

As a connection based protocol, a connection is established and maintained until applications and 
devices have finished exchanging data. It determines how the original message should be broken into
packets.

TCP/IP uses a threeway handshake to establish a connection between a device and a server which 
ensures TCP socket connections can be transferred in both directions concurrently.

Headers

X-Forwarded-For (XFF) http request header is a standard header for identifying the originating 
IP address of a client connecting to a webserver through a proxy server.
It contains a list of IP addresses, the left most address being the client and each subsequent
address is a proxy with the right most IP address being the most recent proxy. 
Security concerns of using this header are that this header is only trustworthy if a trusted
reverse proxy is between the client and the server. If the server is directly connectable 
from the internet, even if it is also behind a reverse proxy, no part of the X-Forwarded-For 
IP list can considered trustworthy or safe.
See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For

User-Agent is a header that lets servers identify the application, os, vendor, and/or version
that the client request originated from. User agent is software that acts on behalf of a user.

TE header is a request header that specifies the transfer encodings that the user agent is 
willing to accept. If it's set to trailers, the response header can include additional 
fields at the end of chunked messages.