> Source: https://portswigger.net/burp/documentation/desktop/burp-ai/ai-connectivity-troubleshooting

Professional

# Troubleshooting AI connectivity

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Burp AI and Burp AT connect to PortSwigger's AI platform. To use them, your network must allow outbound HTTPS traffic to `ai.portswigger.net` on port 443.


 If your organization uses a firewall or content filtering tool, make sure `ai.portswigger.net` is allowlisted. You may need to ask a network administrator to do this.


## Common connectivity problems

 If you can't see your AI credits or use AI features in Burp, the cause is likely to be one of the following:


-

You're offline.
-

Your network requires an upstream proxy.
-

An intercepting proxy is intercepting and re-signing traffic with its own certificate. Burp doesn't trust self-signed certificates by default, so it blocks the connection.

 This page provides a step-by-step guide to help you troubleshoot common connectivity problems.


## Step 1: Check your internet connection

 Make sure your computer is online and able to access PortSwigger's AI services. To do this:


1.

Open a browser and visit `https://portswigger.net` to confirm general internet access.
1.

Try visiting `https://ai.portswigger.net`. You should see a `404 Not Found` message - this is expected and confirms that your network can reach Burp's AI platform.

 If either step fails, check your network settings or contact your administrator.


## Step 2: Configure an upstream proxy

 Some networks require an upstream proxy to access the internet. If this applies to your environment:


1.

In Burp, click **Settings**. The **Settings** dialog opens.
1.

Go to **Network > Connections**.
1.

Under **Upstream proxy servers**, click **Add**. The **Add upstream proxy rule** dialog opens.
1.

Enter the required proxy details. For more information, see [Connections settings - Upstream proxy servers](https://portswigger.net/burp/documentation/desktop/settings/network/connections#upstream-proxy-servers).
1.

Click **OK**.
1.

The new upstream proxy rule is added to the table.

## Step 3: Identify and resolve an intercepting proxy

 Some networks use intercepting proxies, such as ZScaler, that decrypt HTTPS traffic and re-sign certificates. Burp doesn't trust these certificates by default, which can block access to AI services.


 To check if an intercepting proxy is affecting your connection:


1.

In Burp's browser, go to `https://ai.portswigger.net`.
1.

In Burp, go to **Settings > Network > TLS**.
1.

Under **Server TLS certificates**, find the entry for `ai.portswigger.net`.
1.

Check the certificate issuer:

  -

If the issuer is a well-known CA provider such as Amazon, it's unlikely that an intercepting proxy is interfering with the connection.
  -

If the certificate is issued by an intercepting proxy (such as ZScaler), or your company's security system, then your traffic is being intercepted.

 If your traffic is being intercepted by a proxy, Burp might not trust the proxy's certificate. This can prevent access to AI features.


 To fix this, add the proxy's certificate authority (CA) to Burp's trusted certificates:


1.

Find the path to the proxy's certificate. If you're unsure, check your system settings or ask a network administrator.
1.

In Burp, go to **Settings > Network > TLS**.
1.

Under **Custom CA certificates**, click **Add**.
1.

Select the certificate file.
1.

Restart Burp, then check whether you can access AI features.
