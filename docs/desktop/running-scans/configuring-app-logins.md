> Source: https://portswigger.net/burp/documentation/desktop/running-scans/configuring-app-logins

Professional

# Configuring application logins

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Adding application logins to a scan enables Burp Scanner to discover content that is only accessible to authenticated users.


#### Note

Burp Scanner uses application logins when it crawls an application. You cannot specify application logins for **Audit selected items** scans, because these scans do not crawl the target.

 There are two types of application login you can add in Burp Suite:


- Username and password pairs are intended for sites that use a basic, single-step login mechanism.
- Recorded login sequences are intended for sites that use more complex login mechanisms, such as Single Sign-On. You can record login sequences manually or use AI to generate them automatically.

 You can only use one of the available login mechanisms per scan. If you specify both simple login credentials and a recorded login sequence, Burp Scanner uses the recorded login when scanning.


#### Related pages

- [Adding usernames and passwords](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-app-logins/adding-usernames-passwords).
- [Adding recorded login sequences](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-app-logins/adding-recorded-logins).
- [Managing application logins using the configuration library](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-app-logins/application-login-config-library).
