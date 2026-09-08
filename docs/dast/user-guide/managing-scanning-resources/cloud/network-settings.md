> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/cloud/network-settings

DAST

# Network and firewall settings for self-hosted scanning machines

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

 You need to configure your network and firewall to allow your self-hosted scanning machines to communicate with your Cloud instance of Burp Suite DAST:


-
 Allow your scanning machines to have outbound access to the **Dashboard IPs** listed on the [PortSwigger IP ranges](https://ip-ranges.portswigger.cloud/) page.

-
 Enable outbound access from the scanning machine to `*.oastify.com` on port `443`

#### Note

 These instructions only apply to Cloud instances of Burp Suite DAST. If you're looking for network and firewall settings for a self-hosted instance of Burp Suite DAST, see:


-  [Configuring your environment network and firewall settings - Standard](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/network-firewall-config)
-  [Configuring your environment network and firewall settings - Kubernetes](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/network-firewall-config)

#### Related pages

-  [System requirements](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/cloud/system-requirements)
-  [Setting up a self-hosted scanning machine for a Cloud instance](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/cloud/setup)
