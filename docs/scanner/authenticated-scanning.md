> Source: https://portswigger.net/burp/documentation/scanner/authenticated-scanning

DASTProfessional

# Authenticated scanning

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 When crawling a target application, Burp Scanner attempts to cover as much of the application's attack surface as possible. Authenticated scanning enables Burp to crawl privileged content that requires a login to access, such as user dashboards and admin panels.


 Burp Scanner can authenticate with target applications in two ways:


- **Login credentials** are simple username and password pairs. They are intended for sites that use a single-step login mechanism.
- **Recorded login sequences** are user-defined sequences of instructions. They are intended for sites that use complex login mechanisms such as Single Sign-On. You can record login sequences manually, or use Burp AI to record them automatically.

 You can only use one authentication method per scan. If you enter both login credentials and a recorded login sequence, Burp Scanner ignores the provided login credentials.


#### In this section

-

[Login credentials](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/login-credentials)
-

[Identifying login and registration forms](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/identifying-login-forms)
-

[Recorded login sequences](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/recorded-login-sequences)

  -

[Best practice for recording login sequences](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/recorded-login-best-practice)
  -

[Recording login sequences](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/using-recorded-logins)
  -

[Troubleshooting recorded login sequences](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/troubleshooting-recorded-logins)
