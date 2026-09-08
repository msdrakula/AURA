> Source: https://portswigger.net/burp/documentation/desktop/tools/decoder

ProfessionalCommunity Edition

# Burp Decoder

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Burp Decoder enables you to transform data using common encoding and decoding formats. You can use Decoder to:


- Manually decode data.
- Automatically identify and decode recognizable encoding formats, such as URL-encoding.
- Transform raw data into various encoded and hashed formats.

 Decoder enables you to apply layers of transformations to the same data. This enables you to unpack or apply complex encoding schemes. For example, to generate modified data in the correct format for an attack, you could:


1. Apply URL-decoding, then HTML-decoding.
1. Edit the decoded data.
1. Reapply the HTML-encoding, then the URL-encoding.

## Carrying out transformations

 You can send data to Burp Decoder from the message editor in various Burp tools, such as HTTP history. To carry out a data transformation using Burp Decoder:


1. Locate the data that you want to analyze.
1. Right-click the data in the message editor and select **Send to Decoder**.
1. Go to the **Decoder** tab. The data is in the top panel.
1. Select the [operation](https://portswigger.net/burp/documentation/desktop/tools/decoder#operations) you want to perform on the data from the controls beside the data panel. For example, **Encode as** or **Smart decode**.

 You can view the data in either **Text** or **Hex** form.


#### Note

 To send a portion of a message from Burp, select the relevant section before you send it to Decoder.


 You can also type or paste data directly into the Decoder editor panel.


 When you carry out a transformation, a new editor panel opens with the transformed data. You can then apply further transformations as required. For each transformation, the following applies:


- The transformation applies to the whole data set. To apply the transformation to only a portion of the data, select the relevant section before you choose an operation.
- The data is color-coded to indicate the type of encoding or decoding that is applied.
- Any parts of the data that aren't transformed are copied into the new panel in their raw form.

## Operations

- **Decode as** - Apply a decoding function to the data.
- **Encode as** - Apply an encoding function to the data.
- **Hash** - Apply a hash function to the data. The available functions depend upon the capability of your Java platform.
- **Smart decode** - Burp looks for encoded data, and applies layers of decoding until there aren't any further recognizable data formats. This is often useful as an automated first decoding step.

 The following decode and encode functions are available:


-
 URL.

-
 HTML.

-
 Base64.

-
 ASCII hex.

-
 Hex.

-
 Octal.

-
 Binary.

-
 GZIP.


#### Note

 The smart decode operation uses heuristic techniques to recognize common encoding formats, and can therefore make mistakes. You can quickly identify and fix incorrect transformations, as each layer of decoding is shown in a separate panel, and the type of decoding is indicated using color-coding.
