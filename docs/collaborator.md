> Source: https://portswigger.net/burp/documentation/collaborator

DASTProfessional

# Burp Collaborator

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Burp Collaborator is a network service that enables you to detect invisible vulnerabilities. These are vulnerabilities that don't:


- Trigger error messages.
- Cause differences in application output.
- Cause detectable time delays.

 Burp Collaborator uses its own server to identify invisible vulnerabilities, as part of [Out-of-band Application Security Testing (OAST)](https://portswigger.net/burp/application-security-testing/oast). The general process is as follows:


-
 Burp sends Collaborator payloads in a request to the target application. These are subdomains of the Collaborator server's domain. When certain vulnerabilities occur, the target application may use the injected payload to interact with the Collaborator server.

-
 Burp polls the server, to see whether interactions have occurred.


![Burp Collaborator OAST](https://portswigger.net/burp/documentation/images/collaborator/collaborator1.png)

 Burp Collaborator is used in both Burp Suite Professional and Burp Suite DAST:


- Burp Scanner automates the Collaborator process as part of various scan checks. Scanner reports on issues identified in this process.
- Some extensions and custom scan checks use automated Collaborator functionality.

 In Burp Suite Professional, you can use Burp Collaborator to [manually test for vulnerabilities](https://portswigger.net/burp/documentation/desktop/tools/collaborator).


#### Related pages

- [Typical uses for Burp Collaborator](https://portswigger.net/burp/documentation/collaborator/uses).
- [Burp Collaborator server](https://portswigger.net/burp/documentation/collaborator/server).
- [Deploying a private Burp Collaborator server](https://portswigger.net/burp/documentation/collaborator/server/private).
- [Using Collaborator in checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating/writing-guide#using-collaborator-in-checks).
