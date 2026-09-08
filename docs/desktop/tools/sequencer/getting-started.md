> Source: https://portswigger.net/burp/documentation/desktop/tools/sequencer/getting-started

ProfessionalCommunity Edition

# Getting started with Burp Sequencer

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 In this tutorial, you'll use Burp Sequencer to analyze the quality of randomness in an application's session tokens.


#### Note

 Burp Sequencer may have unexpected results in some applications. Until you are fully familiar with its functionality and settings, only use Burp Sequencer against non-production systems.


1.
 Open Burp's browser and access a deliberately vulnerable test website, such as `https://ginandjuice.shop/`.

1.
 Go to **Proxy > HTTP history** and find an entry with a response that issues a session token, for example in a `Set-Cookie` header. To quickly find issued cookies, you can sort the **Cookies** column in the history.

1.
 Right-click the entry and click **Send to Sequencer**.

1.
 Go to the **Sequencer** tab. The entry you just sent to Sequencer is automatically selected in the **Select live capture request** panel.

1.
 Select a cookie in the **Token location within response** panel.

1.
 Click **Start live capture**.

1.
 When Burp has captured a few hundred tokens, click **Pause**.

1.
 To run randomness tests on the tokens, click **Analyze now**.


 The analysis results are displayed in the **Live capture** window. They show a summary of the quality of randomness in the sample.


#### Related pages

-  [Obtaining a token sample](https://portswigger.net/burp/documentation/desktop/tools/sequencer/sample).

-  [Burp Sequencer live capture](https://portswigger.net/burp/documentation/desktop/tools/sequencer/sample/live-capture).

-  [Burp Sequencer results](https://portswigger.net/burp/documentation/desktop/tools/sequencer/results).

-  [Penetration testing workflow](https://portswigger.net/burp/documentation/desktop/testing-workflow).
