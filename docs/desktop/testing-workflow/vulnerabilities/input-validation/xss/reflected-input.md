> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/reflected-input

ProfessionalCommunity Edition

# Identifying reflected input with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Reflected input is when data is copied from a request and echoed into the application's immediate response. This is a prerequisite for a range of vulnerabilities, including reflected cross-site scripting (XSS). You can use Burp to test for reflected input:


- Professional Use Burp Scanner to automatically flag reflected input.
- Use Burp Repeater to manually identify reflected input.

## Steps

 You can follow the processes below using the lab [Reflected XSS into HTML context with nothing encoded](https://portswigger.net/web-security/cross-site-scripting/reflected/lab-html-context-nothing-encoded).


### Scanning for reflected inputs

 If you're using Burp Suite Professional, you can use Burp Scanner to test for reflected input:


1. Identify a request that you want to investigate.
1. In **Proxy > HTTP history**, right-click the request and select **Do active scan**. Burp Scanner audits the request.
1.

Review the **Issues** tab on the **Dashboard** to identify any reflected input that Burp Scanner flags.

![Identify reflected input that Burp Scanner flags](https://portswigger.net/burp/documentation/desktop/images/tutorials/reflected-input-1.png)

### Identifying reflected input manually

1. In **Proxy > HTTP history**, right-click the message you want to investigate and select **Send to
 Repeater**.
1. Replace the value of a parameter with an easily identifiable unique string. For example, `xsstest`.
1. Click **Send**.
1.

Review the response. Use the text editor search function to identify whether the unique string is
 reflected in the
 response.

![Use text editor search function](https://portswigger.net/burp/documentation/desktop/images/tutorials/reflected-input-2.png)

1. Repeat Steps 2 to 4 to test each parameter in the request.

## Next steps

 Once you have identified reflected input, you can investigate the request further to test for vulnerabilities, such as:


- Cross-site scripting
- Client-side template injection
- Server-side template injection
- Input transformations

#### Related pages

- [Burp Scanner](https://portswigger.net/burp/documentation/scanner)
- [Burp Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater)
