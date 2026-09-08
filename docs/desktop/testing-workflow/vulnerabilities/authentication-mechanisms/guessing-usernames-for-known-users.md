> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/guessing-usernames-for-known-users

ProfessionalCommunity Edition

# Guessing usernames for known users

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

Burp Intruder has a built-in username generator that takes an input and produces a list of potential usernames using common patterns. For example, if you were to provide the input `Carlos Montoya` the generator would return `carlos.montoya`, `mcarlos`, and similar combinations.

This is useful in circumstances where you know details of a specific user (for example, their name or email address) but don't know their exact username. It provides a more targeted way of enumerating usernames than a generic name list.

#### Note

You can test this process out in the [Username enumeration via different responses](https://portswigger.net/web-security/authentication/password-based/lab-username-enumeration-via-different-responses) Web Security Academy lab.

## Steps

1. In Burp's HTTP history, identify a failure message for a username-based authentication mechanism.
1. In the message, highlight the username value, right-click, and select **Send to Intruder**.
1. Go to **Intruder**. Notice that Burp has automatically added the username as a payload position.
1. Make sure that **Sniper attack** is selected.
1. In the **Payloads** side panel, change the **Payload type** to **Username generator**.
1. Under **Payload configuration**, enter an input that you want to base the generated usernames on. You can add multiple inputs if required.
1. In the **Maximum payloads per item** field, enter the number of usernames you want Burp to generate. Burp generates this number of usernames for every input added.
1. Click ** Start attack**. Intruder generates its list of potential usernames and runs an attack testing each username in turn.
1. Analyze the attack results to check for interesting patterns, such as usernames that results in anomalous error messages, response times, or a different HTTP response code.

#### Related pages

- [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder)
- [Enumerating usernames with Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/enumerating-usernames)
