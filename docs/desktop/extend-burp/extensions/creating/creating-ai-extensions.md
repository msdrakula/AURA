> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/creating/creating-ai-extensions

Professional

# Creating AI extensions

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 The Montoya API enables you to add advanced AI features into your Burp Suite extensions. Your extensions can now send prompts to a Large Language Model (LLM), allowing real-time input analysis and intelligent responses.


 For example, you could use the API to build extensions that:


-

Automatically evaluate HTTP messages for potential vulnerabilities.
-

Create detailed security reports.
-

Develop extensions that dynamically interact with testers, suggesting payloads or workflows.
-

Generate explanations or training material for complex issues.

 The Montoya API integrates directly with Burp, requiring no additional setup or external configuration to send prompts. All AI interactions are securely managed within PortSwigger's AI platform.


## AI credits

 Ai credits give you access to Burp Suite's AI-powered features. When your extension interacts with an LLM, it deducts AI credits from the user's balance. The cost varies depending on the number and complexity of AI requests needed.


 For tips on writing extensions that manage AI credit usage efficiently, see [Optimize AI requests for efficiency and security](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/creating/creating-ai-extensions/best-practices#optimize-ai-requests-for-efficiency-and-security).


## Examples

 To explore practical examples, check out our [extension repository](https://github.com/PortSwigger/burp-extensions-montoya-api-examples). It contains ready-made extensions that demonstrate how to integrate AI features using the Montoya API. This includes a new AI-enabled version of our popular Hackvertor extension, created by Gareth Heyes from PortSwigger Research.


#### More information

-

[Developing AI features in extensions](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/creating/creating-ai-extensions/developing-ai-features)
-

[Best practices for writing AI extensions](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/creating/creating-ai-extensions/best-practices)
-

[AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust)
