> Source: https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-http-proxy-server

DAST

# Configuring an HTTP proxy server

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If your organization does not allow you to connect to the public internet directly, you can configure a network proxy that the DAST server can use to reach external domains, such as `portswigger.net`.


 To activate your license and perform automatic software updates, the DAST server needs access to `portswigger.net` on port 443. For the best experience give the server permanent access, rather than just for the initial installation.


1.
 From the settings menu , select **Network**.

1.
 Scroll down to **HTTP proxy server** and select **Use an HTTP proxy server**.

1.
 Enter the **Host** and **Port** for your proxy server.

1.
 If your proxy server requires a login:


  -
 Select **Authenticated**.

  -
 Enter a valid username and password.


1.
 To use the proxy server for connecting to an SMTP server, select **Use proxy to connect to email server**.

1.

To bypass the proxy for specific hosts, enter them in **No proxy for**. Click the  icon to add additional entries. This is useful for internal hosts that are not reachable through your proxy.

You can add the following types of entry:

  - An exact hostname, for example `internal.example.com`.
  - A wildcard that matches all subdomains of a domain, for example `*.example.com`. This matches subdomains such as `api.example.com` and `a.b.example.com`, but not `example.com`. You can only use the `*` wildcard at the start of an entry, immediately followed by a dot.
  - An IP address, for example `192.168.0.10` or `2001:db8::1`. Wildcards and CIDR ranges are not supported for IP addresses.

#### Note

 Each entry is matched against the host exactly as your connection requests it: the hostname or IP address in the URL. Burp Suite DAST does not perform DNS resolution for matching, so enter each host in the same form your connections use. For example, excluding a host by IP address does not bypass the proxy for requests made to that host by its hostname.


#### Note

 You can only use an unauthenticated proxy to connect to an SMTP server. For more information, see [configure a connection to your SMTP server](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server).


#### Related pages

[Planning to deploy Burp Suite DAST](https://portswigger.net/burp/documentation/dast/setup)
