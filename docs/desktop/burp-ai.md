> Source: https://portswigger.net/burp/documentation/desktop/burp-ai

Professional

# AI features (Burp AI)

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Burp AI is a set of AI-powered features in Burp Suite Professional that helps you uncover vulnerabilities more efficiently, understand complex web technologies, and streamline authentication setup.


 Burp AI features prioritize privacy and security, and none of them run unless you explicitly activate them. Using them requires AI credits. For more information, see [AI credits](https://portswigger.net/burp/documentation/desktop/burp-ai/ai-credits).


## Burp AI features

 Burp AI includes the following features:


### Burp AI in Repeater

 Burp AI is built into Repeater, enabling you to run custom prompts against any tab. This flexible workflow gives you full control over what Burp AI examines, making it easy to tailor each task to your needs. For example, you can analyze a suspicious request, test for a specific vulnerability, or ask for suggestions on what to try next when you're unsure how to proceed.


#### More information

-

[Using Burp AI in Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater/http-messages/burp-ai-in-repeater)
-

[Prompting best practices with Burp AI](https://portswigger.net/burp/documentation/desktop/burp-ai/prompting)

### Explore Issue

 Explore Issue autonomously investigates vulnerabilities identified by Burp Scanner, saving you time and effort. It follows up on issues like a human pentester would - attempting exploits, identifying additional attack vectors, and summarizing findings so you can validate and demonstrate impact more efficiently.


#### More information

 For more information on Explore Issue, see [Exploring Issues with AI](https://portswigger.net/burp/documentation/desktop/running-scans/explore-issue-with-ai).


### Explainer

 Explainer enables you to quickly understand unfamiliar technologies without leaving Burp Suite. Highlight any part of a Repeater message and click a button to get an AI-generated explanation. Explainer provides instant insights into headers, cookies, JavaScript functions, and more, to help you quickly identify potential security implications without disrupting your workflow.


#### More information

 For more information on Explainer, see [Generating AI-powered explanations](https://portswigger.net/burp/documentation/desktop/tools/repeater/http-messages/ai-explainer).


### Broken access control false positive reduction

 False positives in automated security testing can waste valuable time. Burp enhances Broken Access Control scan checks by intelligently filtering out false positives before they're reported, helping to free up your time to focus on real threats.


#### More information

 For more information on BAC false positive reduction, see [Configure AI scan enhancements](https://portswigger.net/burp/documentation/desktop/running-scans/webapp-scans/full-crawl-and-audit#step-2-configure-ai-scan-enhancements-optional).


### AI-powered recorded logins

 Configuring authentication for web apps can be time-consuming and error-prone. Burp can use AI to generate recorded login sequences automatically, saving time and eliminating the possibility of human error.


#### More information

 For more information on AI-generated recorded login sequences, see [Adding recorded login sequences](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-app-logins/adding-recorded-logins#generating-recorded-login-sequences-using-ai).


### AI-powered extensions

 The Montoya API enables you to add advanced AI features into your Burp Suite extensions. Your extensions can send prompts to an AI model, allowing for real-time input analysis and intelligent responses. There's no need for complex setup, such as managing API keys, as all AI interactions are handled within Burp Suite's secure AI infrastructure.


#### More information

 For more information on creating AI extensions, see [Creating AI extensions](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/creating/creating-ai-extensions).


### AI-powered custom actions in Burp Repeater

 Burp Repeater supports custom actions enhanced with AI, enabling real-time, context-aware analysis of HTTP messages.


 All AI interactions are handled within Burp Suite's secure AI infrastructure, so there's no need for complex setup, such as managing API keys.


#### More information

 For more information on creating AI custom actions, see [Developing AI features in custom actions](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/writing-custom-actions/developing-ai-features).


## When to use Burp AI

 Use Burp AI when you want help with a specific task in a tool you're already working in, such as analyzing a request in Repeater, explaining an unfamiliar technology, or following up a single Scanner finding with Explore Issue. Each feature runs only when you invoke it, so you stay in control of what it examines and when.


#### Note

 Looking for Burp AT, which brings agentic AI to human-led pentesting? See [Burp AT](https://portswigger.net/burp/documentation/desktop/burp-at).


#### Related pages

-

[AI credits](https://portswigger.net/burp/documentation/desktop/burp-ai/ai-credits)
-

[AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust)
-

[Troubleshooting AI connectivity](https://portswigger.net/burp/documentation/desktop/burp-ai/ai-connectivity-troubleshooting)
