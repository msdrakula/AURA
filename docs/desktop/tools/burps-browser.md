> Source: https://portswigger.net/burp/documentation/desktop/tools/burps-browser

ProfessionalCommunity Edition

# Burp's browser

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Burp Suite comes with its own browser, which is ready to use for a variety of manual and automated testing purposes.


## Manual testing with Burp's browser

 Burp's browser is preconfigured to work with the full functionality of Burp Suite right out of the box. All of the necessary proxy listener settings are automatically adjusted for you. This means you can launch Burp for the first time and immediately start testing, even using HTTPS, without performing any additional configuration.


 To launch Burp's browser, go to the **Proxy > Intercept** tab and click **Open browser**. You can then visit and interact with websites just like you would with any other browser. All in-scope traffic is automatically proxied through Burp. This means that as you browse your target website, you can take advantage of Burp Suite's manual testing features. For example, you can intercept and modify requests using [Burp Proxy](https://portswigger.net/burp/documentation/desktop/tools/proxy) and study the complete HTTP history from the corresponding tabs. You can then send these requests to other tools, such as [Burp Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater) and [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder), to perform additional testing of interesting items that you encounter.


 While you browse, Burp's default [live tasks](https://portswigger.net/burp/documentation/desktop/running-scans/live-tasks) will also passively crawl and audit the locations that you visit. This will automatically populate the [site map](https://portswigger.net/burp/documentation/desktop/tools/target/site-map) and [report any potential security issues](https://portswigger.net/burp/documentation/desktop/running-scans/results/audit-log) as they are identified.


#### Read more
[Penetration testing workflow](https://portswigger.net/burp/documentation/desktop/testing-workflow)

 If you prefer, you can still use an external browser for testing. In this case, you just need to perform some [additional configuration steps](https://portswigger.net/burp/documentation/desktop/external-browser-config).


## Scanning websites with Burp's browser

 Burp's browser offers a convenient way to perform manual testing with minimal setup. However, it's even more powerful when integrated into your automated testing workflow through [browser-powered scanning with Burp Scanner](https://portswigger.net/burp/documentation/scanner/browser-powered-scanning).


## Health check for Burp's browser

 If you are experiencing any issues with Burp's browser, you can use the **Health check for Burp's browser** tool to help diagnose the problem. You can access this tool from Burp's **Help** menu. The health check runs a series of tests to check whether the browser is working correctly and provides feedback on any issues that arise.
