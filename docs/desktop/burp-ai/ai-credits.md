> Source: https://portswigger.net/burp/documentation/desktop/burp-ai/ai-credits

Professional

# AI credits

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 AI credits give you access to Burp Suite's AI-powered features and to Burp AT. Whenever you use an AI-powered tool or an extension that interacts with an AI model, or run a task with Burp AT, credits are deducted from your balance.


## Buying AI credits

 You can buy AI credits from [My Account](https://portswigger.net/users/youraccount/licenses) on the PortSwigger site. You must have a Burp Suite Professional license associated with your account to buy AI credits.


 To check your AI credit balance, click the  AI icon in the bottom-right corner of Burp. You can also view your balance in My Account.


 Unused credits expire 12 months after purchase.


#### Note

 AI credits are assigned to an individual user. They cannot be shared or pooled across multiple users.


 If you have bought AI credits but can't see them in Burp, see [Troubleshooting AI connectivity](https://portswigger.net/burp/documentation/desktop/burp-ai/ai-connectivity-troubleshooting).


## Using AI credits

 The amount of credits required by a feature depends on how many AI requests the feature needs to make, and how complex those requests are. For example, Explainer generally uses credits at a low rate, as it only needs to send and receive snippets of text. However, Explore Issue requires Burp to send larger requests to the AI, and usually uses credits at a higher rate. Likewise, Burp AT uses credits as it works through a task, so longer or more complex tasks generally use more credits.


 If you run out of AI credits, AI features pause until you buy more. For AI scan enhancements, you can choose whether a scan should stop running or carry on without AI features if you run out of credits part way through the scan.


#### Note

 Your credit balance can sometimes go below zero. Burp checks your balance before it sends a request to the AI service, then deducts credits after the request completes. If you make several requests at the same time, they may pass the balance check before any deductions happen.


## Extensions

 The number of credits needed for an AI-powered extension depends on the extension's implementation. You can track credit use for individual extensions. To see how many credits an extension has used in the current Burp session, go to **Extensions > Installed** and check the **AI credits used this session** column.


#### More information

-

[AI features (Burp AI)](https://portswigger.net/burp/documentation/desktop/burp-ai)
-

[AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust)
