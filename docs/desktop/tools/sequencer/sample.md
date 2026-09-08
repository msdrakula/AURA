> Source: https://portswigger.net/burp/documentation/desktop/tools/sequencer/sample

ProfessionalCommunity Edition

# Obtaining a token sample

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You need to obtain a sample of tokens to run a Burp Sequencer analysis. You can do this in two ways:


- Perform an automatic live capture of tokens directly from the target.
- Manually load a sample of tokens that you have already acquired.

#### Note

 Sequencer lets you perform an analysis with a sample of only 100 tokens, but this should not be considered reliable for any serious purpose. A sample of 5,000 tokens is usually sufficient, although this may depend on the sample's characteristics. Sequencer supports a maximum sample size of 20,000 tokens, for compliance with the FIPS standards.


## Configuring a live capture of tokens

 To automatically capture tokens from the target response:


1. Locate a request that returns a token that you want to analyze.
1. Right-click the request and choose **Send to Sequencer** from the context menu. You can send a request from anywhere in Burp.
1. Select the token in the **Token location within response** panel.
1. Click **Start live capture**.

 A new results window opens, in which Sequencer repeatedly issues the request and extracts the relevant token from the application's responses. To learn about the live capture process, see [Burp Sequencer live capture](https://portswigger.net/burp/documentation/desktop/tools/sequencer/sample/live-capture).


### Selecting a token location

 You can specify the token you want to analyze in the **Token location within response** panel. The following options are available:


-  **Cookie** - Select a cookie from the drop-down list. This is the most common method to pass session tokens to clients.

-  **Form field** - Select an HTML form field value from the drop-down list. You can use this method to transmit anti-CSRF tokens and other per-page tokens to clients.

-  **Custom location** - Specify a custom location within the response that contains the token you want to analyze. This uses the [response extraction rule dialog](https://portswigger.net/burp/documentation/desktop/settings/response-extraction). This is useful when the token doesn't appear in a usual form field.


#### Related pages

- For more information on the live capture process, see [Burp Sequencer live capture](https://portswigger.net/burp/documentation/desktop/tools/sequencer/sample/live-capture).
- You can control how Sequencer makes HTTP requests and harvests tokens during a live capture. For more information, see [Live capture
 settings](https://portswigger.net/burp/documentation/desktop/settings/tools/sequencer#live-capture).

## Manually loading tokens

 You can load Sequencer with a sample of tokens that you have already obtained, for example from an earlier live capture, or an Intruder attack:


1. Go to **Sequencer > Manual load**.
1.

 Insert the tokens:


  - To paste tokens from the clipboard, click **Paste**.
  - To load tokens from a file, click **Load**.

1. Use the details of the shortest and longest lengths in the display field to make sure that the sample has loaded correctly.
1. Click **Analyze now**. The analysis begins in a new results window.

#### Note

 Make sure that the tokens are in a simple newline-delimited text format.
