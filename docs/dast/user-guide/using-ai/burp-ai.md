> Source: https://portswigger.net/burp/documentation/dast/user-guide/using-ai/burp-ai

DAST

# Burp AI

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

 Burp Suite DAST includes AI-powered features that help you to triage issues faster and to simplify scan setup.

## Burp AI features in DAST

 Burp Suite DAST includes:

### AI-enhanced scanning

 AI-enhanced scanning automatically investigates issues found during a scan to see if they can be reproduced and exploited. Burp AI uses evidence from Burp Scanner to plan and execute targeted validation steps. It then displays its results on the issue detail page, including an impact analysis and manual reproduction steps, to help you prioritize and resolve issues faster.

#### Related pages

 For more information on AI-enhanced scanning, see [Configuring AI-enhanced scanning](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/configuring-ai-enhanced-scanning).


### AI-powered recorded logins

 Configuring authentication for web apps can be time-consuming and error-prone. Burp Suite DAST can use AI to generate recorded login sequences automatically, saving time and eliminating the possibility of human error.

#### Related pages

 For more information on AI-generated recorded login sequences, see [Using recorded logins](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/recorded-logins#recording-login-sequences-using-ai).


## Enabling Burp AI features in your DAST instance

 Burp AI features only appear in Burp Suite DAST if they have been enabled for your instance. You must have the **Enable/Disable Burp AI for the installation** permission to enable Burp AI.

 To enable Burp AI in DAST:

1. From the settings menu, select **Burp AI**.
1. Activate the **Enable Burp AI** toggle.
1. Click **Save**.

## Security and privacy

 We've designed Burp Suite DAST's AI features with security, privacy, and transparency in mind:

- AI features only run when you choose, giving you full control over when and where they execute.
- AI request data is processed securely through our trusted AI infrastructure. It is never stored by our AI providers.
- Burp AI complies with ISO 27001 standards and implements robust encryption, ensuring data is protected in transit and at rest.

#### Related pages

- For more information on trust and safety in Burp AI, see the [AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust).
