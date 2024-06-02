package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"time"
)

func main() {
	urls := []string{
		"https://vitbhopal.ac.in/",
	}

	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			return
		}

		var connect, dns, tlsHandshake, firstByte time.Time

		trace := &httptrace.ClientTrace{
			DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
			DNSDone:  func(ddi httptrace.DNSDoneInfo) { fmt.Printf("DNS Done: %v\n", time.Since(dns)) },
			ConnectStart: func(network, addr string) {
				connect = time.Now()
			},
			ConnectDone: func(network, addr string, err error) {
				if err != nil {
					fmt.Printf("Error establishing connection: %v\n", err)
					return
				}
				fmt.Printf("TCP Connection Time: %v\n", time.Since(connect))
			},
			TLSHandshakeStart: func() { tlsHandshake = time.Now() },
			TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
				if err != nil {
					fmt.Printf("Error during TLS handshake: %v\n", err)
					return
				}
				fmt.Printf("TLS Handshake Time: %v\n", time.Since(tlsHandshake))
				fmt.Printf("Cipher Suite: %v\n", tls.CipherSuiteName(cs.CipherSuite))
				fmt.Printf("TLS Version: %v\n", cs.Version)
				fmt.Printf("Server Name: %v\n", cs.ServerName)
				for _, cert := range cs.PeerCertificates {
					fmt.Printf("Certificate Issuer: %v\n", cert.Issuer)
					fmt.Printf("Certificate Subject: %v\n", cert.Subject)
					fmt.Printf("Certificate Validity: From %v to %v\n", cert.NotBefore, cert.NotAfter)
				}
			},
			GotFirstResponseByte: func() {
				firstByte = time.Now()
				fmt.Printf("Server Processing Time: %v\n", time.Since(tlsHandshake))
			},
		}

		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
		client := &http.Client{
			Timeout: time.Second * 60, // Increase the timeout to 60 seconds
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: false,
				},
				MaxIdleConnsPerHost: 0, // Disable connection pooling
			},
		}
		startRequest := time.Now()
		resp, err := client.Do(req)
		endRequest := time.Now()

		if err != nil {
			fmt.Printf("Request error: %v\n", err)
			return
		}

		fmt.Printf("Total Time for %v: %v\n", url, endRequest.Sub(startRequest))
		fmt.Printf("Content Transfer Time: %v\n\n", endRequest.Sub(firstByte))
		resp.Body.Close()
	}
}