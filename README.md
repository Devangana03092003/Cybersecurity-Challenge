# Cybersecurity-Challenge


1. *Connection Profiles*: The code uses the httptrace package from the Go standard library to trace the events within HTTP client requests. It captures the following timings:
   - DNS lookup time: From the start of the DNS lookup to when it's done.
   - TCP connection time: From the start of the TCP connection to when it's established.
   - TLS handshake time: From the start of the TLS handshake to when it's done.
   - Server processing time: From the end of the TLS handshake to when the first byte of the response is received.
   - Content transfer time: From when the first byte of the response is received to when the request is fully completed.

2. *TLS Information*: The code captures the following TLS details during the TLS handshake:
   - Cipher Suite: The cipher suite used in the TLS connection.
   - TLS Version: The version of the TLS protocol used in the connection.
   - Server Name: The server name used in the SNI extension.
   - Certificate Details: The issuer, subject, and validity period of each certificate in the server's certificate chain.
