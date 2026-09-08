> Source: https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/managing-certificates

DAST

# Managing certificates for outbound connections

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can add or remove certificates so that Burp Suite DAST knows which external systems to trust. For example, you can add certificates for email or Jira servers, or to allow updates from PortSwigger.


 You can also upload certificates for an internal certificate authority, or individually self-signed certificates.


#### Note

 Only upload certificates from trusted sources.


 Certificates for inbound traffic are managed separately, refer to [Enabling TLS](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/configure-web-server#enabling-tls).


 To manage your certificates for outbound connections:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the settings menu , select **Network**.

1.
 Scroll down to **Manage certificates**.

1.
 To add a certificate, click **Upload certificate**.

1.
 To remove a certificate, click .
