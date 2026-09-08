> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/analyzing/opaque-data/parts-of-token

ProfessionalCommunity Edition

# Identifying which parts of a token impact the response with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Sometimes an application may only inspect certain parts of a token. You can use the Character frobber and Bit flipper payload types in Burp Intruder to modify the value of each character or bit position of a token in turn. This enables you to identify which parts of the token impact the response you receive.


 For example, if you modify the value of a character in a session token and your request is still processed in your session, it is likely the character is not used to track your session.


## Steps

 To identify which parts of a token impact the response:


1. Identify a message that includes a token that you want to investigate further, such as a session token.
1. Highlight the token, then right-click the message and click **Send to Intruder**.
1.

Go to **Intruder**. Notice that the token is automatically marked as a payload position.
1. In the **Payloads** side panel, change the **Payload type** to **Character
 frobber**.
1.

Add an extract grep for sections of interest in the response, such as a verbose error message:

  1. Under **Grep - Extract**, click **Add**. The **Define extract grep
 item** dialog opens.
  1. Highlight the section in the response that you want to extract, such as an error message.
  1. Click **OK**. The item is added to the list. When you start the attack, Intruder extracts the text at this location in each response and displays it in the results table.

1. Click ** Start attack**. Burp Intruder modifies the encrypted value of each character in the token in turn.
1. When the attack is finished, study the responses to look for any behavior that may indicate the modification has impacted the token validity or changed the response. For example, look for any anomalous error messages or status codes.

#### Note

 If tokens contain data encrypted with a CBC cipher, use the **Bit flipper** payload type instead of the **Character frobber**. This modifies bits in the preceding cipher block.


#### Related pages

- [Payload types - Character frobber](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/payload-types#character-frobber)
- [Payload types - Bit flipper](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/payload-types#bit-flipper)
- [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder)
- [Grep - extract](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/settings#grep-extract)
