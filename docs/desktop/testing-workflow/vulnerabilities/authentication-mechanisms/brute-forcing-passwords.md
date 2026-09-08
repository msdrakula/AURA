> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/brute-forcing-passwords

ProfessionalCommunity Edition

# Brute-forcing passwords with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 4 Minutes

 Burp Suite provides a number of features that can help you brute-force the password of a given user, gaining access to their account and additional attack surface. For example, you can:


-
 Use a list of common passwords. This is commonly known as a dictionary attack. For details on how to do this, see [Running a dictionary attack](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/brute-forcing-passwords#running-a-dictionary-attack).

-
 Try every permutation of a character set. For details on how to do this, see [Running an exhaustive brute-force attack](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/brute-forcing-passwords#running-an-exhaustive-brute-force-attack).


#### Note

 The examples below are simplified to demonstrate how to use the relevant features of Burp Suite. To run these attacks on real websites, you usually need to also bypass defenses such as rate limiting. For some ideas on how to do this, see the [Authentication](https://portswigger.net/web-security/authentication) topic on the Web Security Academy.


## Before you start

 Identify one or more valid usernames for the target website. For example, you can potentially [enumerate a list of usernames](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/enumerating-usernames) using Burp. For the examples below, you can assume that the username `wiener` is valid.


 For details on how to brute-force both the username and password in a single attack, see [Brute-forcing a login with Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/brute-forcing-logins).


## Running a dictionary attack

 One approach for brute-forcing passwords is to use a list of potential passwords, usually collated from previous data breaches. This is far more efficient than an exhaustive brute-force attack, but relies on the user's password being present in your list, which may not always be the case.


 You can follow along with the process below using the [User role controlled by request parameter](https://portswigger.net/web-security/access-control/lab-user-role-controlled-by-request-parameter) lab from our Web Security Academy.


1.
 Send the request for submitting the login form to Burp Intruder.

1.
 Go to **Intruder**. Make sure that **Sniper attack** is selected.

1.

 Highlight the password value and click **Add §** to mark it as a payload position. Make sure that you're using a valid username. If you're following along with the lab, set the username to `wiener`.


![Configuring payload positions for brute-forcing a password](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/brute-forcing-passwords/positions.png)

1.

 In the **Payloads** side panel, under **Payload configuration**, add a list of passwords that you want to test. Ideally, sort the list in order of how likely you think the password is to be correct. This could be based on prior knowledge of the user in question or just how common the password is in general.


  -
 If you're using Burp Suite Professional, you can open the **Add from list** dropdown menu and select the **Passwords** list.

  -
 If you're using Burp Suite Community Edition, manually add a list of potential passwords.


![Adding a list of payloads for brute-forcing a password with a dictionary attack](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/brute-forcing-passwords/payloads-list.png)

1.
 Click ** Start attack**. The attack starts running in the new dialog. Intruder sends a request for each password in the list.

1.

 When the attack is finished, study the responses to look for any behavior that may indicate a valid password. For example, look for any anomalous error messages, response times, or status codes. In the example below, one of the requests has received a 302 response.


![Viewing the results of brute-forcing a password with a dictionary attack.](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/brute-forcing-passwords/results.png)

1.
 To investigate the contents of a response in detail, right-click and select **Send to Comparer (response)**. Do the same for the original response.

1.
 Go to the **Comparer** tab. Select the two responses and click **Words** or **Bytes** to compare the responses. Any differences are highlighted.


## Running an exhaustive brute-force attack

 Another approach is to attempt every possible permutation of a character set. This enables you to brute-force passwords that don't necessarily appear in a wordlist. However, for longer passwords and larger character sets, this type of attack is often impractical due to the number of requests needed. For example, an alphabetical password with five characters has over 11 million possible combinations. It's often better to try running a dictionary attack first.


1.
 Send the request for submitting the login form to Burp Intruder.

1.

 In the request, highlight the password value and click **Add §** to mark it as a payload position. Make sure that you're using a valid username.


![Configuring payload positions for an exhaustive password brute-force attack](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/brute-forcing-passwords/positions.png)

1.
 In the **Payloads** side panel, change the **Payload type** to **Brute forcer**.

1.

 Under **Payload configuration**, enter the full character set and set the minimum and maximum password length that you want to test. If you're able to create your own account on the site, you can potentially get clues about the password requirements to help you determine the appropriate values.


![Configuring payloads for the brute-forcer payload type](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/brute-forcing-passwords/payloads-brute-forcer.png)

1.
 Click ** Start attack**. The attack starts running in the new dialog. Intruder sends a request for every possible password based on your settings.

1.
 When the attack is finished, study the responses to look for any behavior that may indicate a valid password. For example, look for any anomalous error messages, response times, or status codes. In the example below, one of the requests has received a 302 response.

1.
 To investigate the contents of a response in detail, right-click and select **Send to Comparer (response)**. Do the same for the original response.

1.
 Go to the **Comparer** tab. Select the two responses and click **Words** or **Bytes** to compare the responses. Any differences are highlighted.


#### Related pages

- [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder)
- [Predefined payload lists](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/payload-lists)
- [Burp Comparer](https://portswigger.net/burp/documentation/desktop/tools/comparer)
- Academy: [Vulnerabilities in password-based login mechanisms](https://portswigger.net/web-security/authentication/password-based)
