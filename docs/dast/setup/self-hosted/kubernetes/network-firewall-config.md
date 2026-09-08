> Source: https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/network-firewall-config

DAST

# Configuring your environment network and firewall settings (Kubernetes)

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 To ensure that Burp Suite DAST is able to function correctly, you may need to configure your firewall to allow the various components to communicate with each other and the public web. We support IPv4 and IPv6.


#### Warning

 For security reasons, make sure that your Kubernetes cluster can only reach systems that you intend to scan. Failure to do so may result in unintended user access to internal functionality.


## Configuring your instance

 Configure the connections as follows:


-

 Allow your users and API clients to access the web server on the configured port.


#### Note

 You can't change the web server port on a Kubernetes instance as your external port should be configured as part of your ingress solution.

-
 To activate your license and perform automatic software updates, allow the DAST server to access `portswigger.net` on port 443. If necessary, configure a network proxy to reach the public web.

-
 Allow your Kubernetes cluster to access the websites that you want to scan on the relevant ports.

-
 Allow the Kubernetes cluster to have access to the database service on the configured host and port.


![Simplified Kubernetes network diagram](https://portswigger.net/burp/documentation/dast/images/k8-network-diagram.jpg)

**Next step - **Kubernetes system requirements  [CONTINUE](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/k8-system-requirements)
