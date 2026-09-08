> Source: https://portswigger.net/burp/documentation/collaborator/uses

DASTProfessional

# Typical uses for Burp Collaborator

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Burp Collaborator can be used to detect a wide range of out-of-band vulnerabilities. Some common use cases are described below:


- [External service interaction](https://portswigger.net/burp/documentation/collaborator/uses#external-service-interaction).
- [Out-of-band resource load](https://portswigger.net/burp/documentation/collaborator/uses#out-of-band-resource-load).
- [Blind SQL injection](https://portswigger.net/burp/documentation/collaborator/uses#blind-sql-injection).
- [Blind cross-site scripting](https://portswigger.net/burp/documentation/collaborator/uses#blind-cross-site-scripting).

## External service interaction

 Burp Collaborator can induce and detect a typical external service interaction as follows:


1. When the application receives a payload from Burp Collaborator, it performs a DNS lookup on the payload URL, then performs an HTTP request.
1. The Collaborator server receives the DNS lookup and HTTP request.
1. Burp polls the Collaborator server for payload interactions.
1. Burp reports the external service interaction, including the full interaction messages.

![Detecting external service interaction](https://portswigger.net/burp/documentation/images/collaborator/collaborator2.png)

 You can detect some types of service vulnerabilities by analyzing the details of the service interaction with a collaborating instance of that service. For example, mail header injection can be detected in this way.


#### Note

 The ability to induce an external service interaction doesn't always indicate a vulnerability. You need to determine whether this is intended behavior. For more information about external service interaction vulnerabilities, see the
 **External service interaction** issues in [Vulnerabilities detected by Burp Scanner](https://portswigger.net/burp/documentation/scanner/vulnerabilities-list).


## Out-of-band resource load

 This is when an application can be induced to load content from an external source and include it in its own response.


 To detect this vulnerability, the Collaborator server returns specific data in its responses to the application's interactions. Burp then analyzes the application's in-band response for the data.


![Detecting out-of-band resource load](https://portswigger.net/burp/documentation/images/collaborator/collaborator3.png)

#### Related pages

 For more information, see the **Out-of-band resource load** issue in [Vulnerabilities detected by Burp Scanner](https://portswigger.net/burp/documentation/scanner/vulnerabilities-list).


## Blind SQL injection

 This is when an application is vulnerable to SQL injection, but its HTTP responses don't contain the results of the relevant SQL query or the details of any database errors.


 To detect this vulnerability, Burp sends injection-based payloads that trigger an external interaction when the injection is successful.


 The following example uses an Oracle-specific API to trigger an interaction when injected into a SQL statement.


![Detecting blind SQL injection](https://portswigger.net/burp/documentation/images/collaborator/collaborator4.png)

#### Related pages

 For more information, see the Web Security Academy documentation on [Blind SQL injection](https://portswigger.net/web-security/sql-injection/blind).


## Blind cross-site scripting

 The Collaborator server can notify Burp of deferred interactions that occur asynchronously when the relevant in-band payload is submitted to the target. This enables the detection of various stored vulnerabilities, such as second-order SQL injection and blind XSS.


 In the example below:


1. Burp sends a stored XSS payload that triggers a Collaborator interaction if it is rendered to a user.
1. An admin user views the payload later. Their browser performs the interaction.
1. Burp polls the Collaborator server. Burp receives details of the interaction and reports the stored XSS vulnerability.

![Detecting blind cross-site scripting](https://portswigger.net/burp/documentation/images/collaborator/collaborator5.png)
