> Source: https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/architecture-overview

DAST

# Kubernetes architecture overview

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 The following diagram shows the core components of Burp Suite DAST and the connections between them.


![Architecture for Burp Suite DAST](https://portswigger.net/burp/documentation/dast/images/enterprise-scan-04.jpg)

## DAST server

 The DAST server is the main application server. It coordinates between the other components.


## Web server

 The web server provides the interface to users either via the web UI or one of the APIs.


## Database

 The Kubernetes version of Burp Suite DAST requires you to connect your own external SQL database to store all the application data. For more information, see [System requirements for the external database](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/external-database-requirements).


## Scans and scanning resources

 For Kubernetes instances, your scanning resources automatically scale to cope with the number of concurrent scans that you need to run at any given time. These resources are then scaled back down again once they are no longer needed.


#### Read more
[Managing Kubernetes scanning resources](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/self-hosted/kubernetes)

#### Related pages
[Kubernetes system requirements](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/k8-system-requirements)

**Next step - **Kubernetes scanning resources overview  [CONTINUE](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/resource-overview)
