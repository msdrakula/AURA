> Source: https://portswigger.net/burp/documentation/dast/user-guide/reference/scanning-machines

DAST

# Scanning machines

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 A scanning machine is a virtual or physical machine that runs [scans](https://portswigger.net/burp/documentation/dast/user-guide/working-with-scans). You can run scans on the same machine as the web server and DAST server to begin with. When you follow the installation wizard, this is the default option that is selected.


![Bundled scanning machine](https://portswigger.net/burp/documentation/dast/images/enterprise-scan-05.jpg)

 Alternatively, you can set up as many [external machines](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/additional-scanning-machines) as you want and use them as dedicated scanning machines. In this way, you can spread your scans across multiple scanning machines to avoid overloading a single machine. You can [assign a maximum number of concurrent scans](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/assigning-scan-limits) that can be run on each machine. Please refer to the [system requirements](https://portswigger.net/burp/documentation/dast/setup/self-hosted/system-req-overview) section for more information about [how many machines](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/single-v-multi-machine) you might need to deploy.


![Setup with two external scanning machines](https://portswigger.net/burp/documentation/dast/images/enterprise-scan-03.jpg)

 Note that if you decide to use external scanning machines, you need to [authorize them](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/additional-scanning-machines#authorizing-a-new-scanning-machine) in Burp Suite DAST before you can use them to run scans.


## Scanning pools

 For standard instances (as opposed to Kubernetes instances), scanning machines exist in one of several scanning pools. These pools are used to manage resources for different kinds of scans, or for scanning different sorts of sites. To assign the machine to a different pool, see [Managing scanning pools](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/scanning-pools).


## CI-driven scans

 You can run CI-driven scans in your CI/CD pipeline. These run in temporary scanning containers that are created in a container platform.


 For more information, see [Integrating CI-driven scans](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans).


#### Related pages

- [Adding additional scanning machines](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/additional-scanning-machines).
- [Assigning scan limits](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/assigning-scan-limits).
- [Managing scanning pools](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/scanning-pools).
