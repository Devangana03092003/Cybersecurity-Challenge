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

However, the code does not currently identify the correct certificate bundle or the supported cipher suites. To accomplish these objectives, you would need to modify the code to perform additional checks during the TLS handshake and potentially make additional connections to test different cipher suites.

Please note that identifying the correct certificate bundle typically involves checking the certificate chain provided by the server against a set of trusted root certificates. This is usually handled automatically by the Go standard library when InsecureSkipVerify is set to false, as it is in your code.

Identifying the supported cipher suites would involve making additional connections to the server with different cipher suites configured in the tls.Config struct and checking which ones are accepted. This could be a time-consuming operation and is typically not necessary unless you have specific security requirements.
