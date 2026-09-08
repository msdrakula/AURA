> Source: https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/single-v-multi-machine

DAST

# Single vs. multi-machine architecture

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 The number of machines needed to run Burp Suite DAST very much depends on the scale of your intended usage.


## Single machine architecture

 You can run all of the components on a single machine, including the embedded database. The embedded database is designed for trials and evaluations of Burp Suite DAST. It is not intended for production use.
 For production use, we recommend that you connect to an external database and use a multi-machine architecture. The diagram below shows a single-machine architecture:


![A single-machine installation of Burp Suite DAST](https://portswigger.net/burp/documentation/dast/images/enterprise-scan-05.jpg)

## Multi-machine architecture

 Alternatively, you can run scans on several different machines and you can use your own external database for storage. This lets you scale the number of concurrent scans that you could potentially run
 to be indefinitely large and utilize any existing database infrastructure that you have. The diagram below shows a multiple-machine architecture, with an external database and separate scanning machines:


![A multiple-machine installation of Burp Suite DAST](https://portswigger.net/burp/documentation/dast/images/enterprise-scan-03.jpg)

 Note that the DAST server and the web server are always deployed on a single machine.


## Requirements

 For specific details about the system requirements for both of these options, see [System requirements for standard
 instances](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/standard-system-requirements).


## Network and firewall settings

To ensure that Burp Suite DAST can work correctly, you need to configure your network to allow the various components to communicate with each other and your target applications.
 The network requirements vary depending on whether you intend to use a single-machine or multi-machine architecture.
 To learn more, see [Configuring your network and firewall settings](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/network-firewall-config).

**Next step - **Configuring your network and firewall  [CONTINUE](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/network-firewall-config)
