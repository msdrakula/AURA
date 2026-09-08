> Source: https://portswigger.net/burp/documentation/desktop/tools/intruder/uses/fuzzing

ProfessionalCommunity Edition

# Fuzzing for vulnerabilities

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 You can use Burp Intruder to identify input-based vulnerabilities by analyzing the attack results for error messages and other exceptions.


#### Related pages

 For more information on input-based vulnerabilities, see:


- [SQL injection](https://portswigger.net/web-security/sql-injection).
- [Cross-site scripting](https://portswigger.net/web-security/cross-site-scripting).
- [Directory traversal](https://portswigger.net/web-security/file-path-traversal).

 Given the size and complexity of today's applications, manually fuzzing for vulnerabilities is a time-consuming process. You can automate the process with Burp Intruder.


## Step 1: Set the payload positions

 Set payload positions at the values of all request parameters.


## Step 2: Set the payload type

 Select the simple list payload type, then add a list of attack strings under **Payload configuration**.


 You can use your own list of attack strings, or one of Burp's predefined payload lists of common fuzz strings if you're using Burp Suite Professional. For more information, see [Predefined payload lists](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/payload-lists).


![Setting payloads for fuzzing](https://portswigger.net/burp/documentation/desktop/images/intruder-gettingstarted-8-940.png)

## Step 3: Set the match grep

 Configure the **Grep - Match** settings to flag responses that contain various common error message strings. For more information, see [Burp Intruder attack settings](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/settings#grep-match).


 You can use your own expressions, or use the default items.


![Fuzzing grep match](https://portswigger.net/burp/documentation/desktop/images/intruder-fuzzing-grep-match-940.png)

## Step 4: Analyze the results

 Sort the attack results by the match grep expressions to identify any results with the common error message strings.


![Intruder grep match fuzzing](https://portswigger.net/burp/documentation/desktop/images/intruder-grep-match.png)

#### Note

 To test a large number of requests with the same payloads and match grep configuration, you can save or copy the tab's configuration. For more information, see [Configuring Burp Intruder attacks](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack).
