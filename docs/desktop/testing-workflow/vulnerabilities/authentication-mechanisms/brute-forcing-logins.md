> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/brute-forcing-logins

ProfessionalCommunity Edition

# Brute-forcing logins with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Although it's far more efficient to first [enumerate a valid username](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/enumerating-usernames) and then attempt to guess the matching password, this may not always be possible. Using Burp Intruder, you can attempt to brute-force both usernames and passwords in a single attack.


#### Note

 The example below is simplified to demonstrate how to use the relevant features of Burp Suite. To run this kind of attack on real websites, you usually need to also bypass defenses such as rate limiting. For some ideas on how to do this, see the [Authentication](https://portswigger.net/web-security/authentication) topic on the Web Security Academy.


## Before you start

 Obtain lists of potential usernames and passwords. For the example below, you can use the following lists:


-  [Usernames](https://portswigger.net/web-security/authentication/auth-lab-usernames)
-  [Passwords](https://portswigger.net/web-security/authentication/auth-lab-passwords)

 In practice, we recommend sorting the list in order of how likely you think the username or password is to be correct.


## Steps

 You can follow along with the process below using the [Username enumeration via subtly different responses](https://portswigger.net/web-security/authentication/password-based/lab-username-enumeration-via-subtly-different-responses) lab from our Web Security Academy.


1.
 Send the request for submitting the login form to Burp Intruder.

1.
 Go to **Intruder** and select **Cluster bomb attack** from the attack type drop-down menu.

1.

 In the request, highlight the username value and click **Add §** to mark it as a payload position. Do the same for the password.


![Setting the payload positions for brute-forcing a login](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/brute-forcing-logins/positions.png)

1.
 In the **Payloads** side panel, select position `1` from the **Payload position** drop-down list.

1.

 Under **Payload configuration**, paste the list of usernames.


![Setting the payloads for brute-forcing a login](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/brute-forcing-logins/payloads-one.png)

1.
 Select position `2` from the **Payload position** drop-down list, and paste the list of passwords.

1.
 Click ** Start attack**. The attack starts running in the new dialog. Intruder sends a request for every possible combination of the provided usernames and passwords.

1.

 When the attack is finished, study the responses to look for any behavior that may indicate a valid login. For example, look for any anomalous error messages, response times, or status codes. In the example below, one of the requests has received a 302 response.


![Viewing the results of a login brute-force attempt](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/brute-forcing-logins/results.png)

1.
 To investigate the contents of a response in detail, right-click and select **Send to Comparer (response)**. Do the same for the original response.

1.
 Go to the **Comparer** tab. Select the two responses and click **Words** or **Bytes** to compare the responses. Any differences are highlighted.


#### Related pages

- [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder)
- [Predefined payload lists](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/payload-lists)
- [Grep - match](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/settings#grep-match)
- [Grep - extract](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/settings#grep-extract)
- [Burp Comparer](https://portswigger.net/burp/documentation/desktop/tools/comparer)
- Academy: [Vulnerabilities in password-based login mechanisms](https://portswigger.net/web-security/authentication/password-based)
