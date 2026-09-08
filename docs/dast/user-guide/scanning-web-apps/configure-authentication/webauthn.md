> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/webauthn

DAST

# WebAuthn passkeys in recorded logins

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can configure a recorded login to authenticate using WebAuthn. This includes passkeys, biometrics such as fingerprint or facial recognition, device PINs such as Windows Hello, and hardware security keys such as YubiKey.

 Support is limited to applications that implement the WebAuthn standard. Non-standard implementations or flows tied to specific hardware may not be supported.

 If your login uses a passkey, you must capture it using the [Login Recorder for Burp Suite extension](https://chrome.google.com/webstore/detail/burp-suite-navigation-rec/anpapjclbjicacakeoggghfldppbkepg?hl=en-GB) before recording your login sequence. Existing passkeys in your application or identity provider cannot be used unless they were captured and exported using the extension.

 To capture a passkey, do one of the following:

- Capture a new passkey using the Login Recorder for Burp Suite extension.
- Import a passkey that was previously captured and exported from the extension.

#### Note

 Do not delete any passkeys used in recorded logins from your application or identity provider. Burp uses them to authenticate during scans, simulating the WebAuthn authentication process automatically.


## Capturing a new passkey

 Use this option if you don't already have a passkey captured using the Login Recorder for Burp Suite extension.

1. Open the Login Recorder for Burp Suite extension.
1. Enable the **My login uses a passkey** toggle.
1. Click **Add passkey**.
1. Click **Capture new passkey**. This enables passkey capture in your browser.
1. Go to your application or identity provider and create a new passkey. The extension captures it automatically.
1. Return to the extension and click **Stop capturing**.

 You can now record your login sequence using this passkey. For more information, see [Recording login sequences](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/using-recorded-logins).

#### Note

 Passkeys don't persist between sessions in the extension. Export any passkeys you want to reuse in future recorded login sequences before closing the browser.


## Importing a passkey

 Use this option if you've previously captured and exported a passkey using the Login Recorder for Burp Suite extension.

1. Open the Login Recorder for Burp Suite extension.
1. Enable the **My login uses a passkey** toggle.
1. Click **Add passkey**.
1. Click **Import passkey**.
1. Upload the passkey file.

 You can now record your login sequence using this passkey. For more information, see [Recording login sequences](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/using-recorded-logins).

#### Related pages

-  [Using recorded logins](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/recorded-logins).

-  [Managing steps in a recorded login](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/managing-recorded-login-steps).

-  [Best practice for recording login sequences](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/recorded-login-best-practice).
