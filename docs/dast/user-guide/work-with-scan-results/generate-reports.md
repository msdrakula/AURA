> Source: https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/generate-reports

DAST

# Downloading reports

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This section describes how to generate **Standard** and **Compliance** reports in HTML or PDF format. You can send scan summary reports automatically, by email.


 For more information about the contents of these reports, refer to [Reports](https://portswigger.net/burp/documentation/dast/user-guide/reference/reports).


 You can also export scan data as an XML file, for import into other tools or reporting frameworks.


## Download a standard report

 To download a standard report in HTML or PDF format:


1.
 Open the **Scans** tab and select a scan.

1.
 Select the **Reporting** tab.

1.
 From the **Report type** drop-down menu, select **Summary** or **Detailed** in HTML or PDF format.

1.
 From the **Include severities** drop-down menu, select the severity levels you want to include in the report.

1.
 To include false positives in the report, select **Include false positive issues**.

1.
 Click **Download**.


## Download a compliance report

 To download a compliance report in HTML or PDF format:


1.
 Open the **Scans** tab and select a scan.

1.
 Select the **Reporting** tab.

1.
 From the **Report type** drop-down menu, select **OWASP Top 10: 2025** or **PCI DSS v4.0.1** in HTML or PDF format.

1.
 Click **Download**.


#### Note

 Burp Suite DAST's compliance reports do not guarantee compliance or non-compliance with any specific security standard.


## Send scan summary reports automatically

 You can configure Burp Suite DAST to automatically send scan summary reports.


#### Note

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 For self-hosted instances, you need to configure a connection to an SMTP server first. For more information, see [Configuring your SMTP server](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server).


1.
 Open the **Sites** tab and select a site.

1.
 In the **Details** tab, click **Edit**.

1.
 In **Scan settings**, select the **Scan notifications** tab.

1.
 In the **Send scan summary reports by email** section, enter an email address.

1.
 To send the report to more than one email address, click  and enter another address.

1.
 To remove an email address, click the trash icon .

1.
 Click **Save**.


## Export issue data

 To export issue data in XML format:

1.
 Open the **Scans** tab and select a scan.

1.
 Select the **Reporting** tab.

1.
 From the **Report type** drop-down menu, select **Export Issue Data**.

1.
 From the **Include severities** drop-down menu, select the severity levels you want to include in the report.

1.
 To include false positives in the report, select **Include false positive issues**.

1.
 To encode requests and responses in Base64, select **Base64-encode requests and responses**.

1.
 Click **Download**.


#### Note

 The XML file uses an internal DTD. If you're an author of interoperability code, we recommend that you review a sample report to obtain the current DTD. These XML elements are specific to the scan results:


-
 The `serialNumber` element contains a long integer that uniquely identifies the individual issue instance. If you export issues several times from the same instance of Burp, you can use the serial number to identify incrementally new issues.

-
 The `type` element contains an integer that uniquely identifies the issue type (such as SQL injection, or XSS). This value is stable across different instances of Burp. See the list of [scan issue types](https://portswigger.net/kb/issues) for a list of all numeric type identifiers.

-
 The `name` element contains the descriptive name for the issue type. See the list of [scan issue types](https://portswigger.net/kb/issues) for a list of all issue names.

-
 The `path` element contains the URL for the issue (excluding query string).

-
 The `location` element includes both the URL and a description of the entry point for the attack, where relevant. For example, a specific URL parameter or request header.

-
 The `request` and `response` elements have a `base64` attribute, which contains a Boolean value to indicate whether the messages are Base64-encoded.


#### Related pages

-  [Reports](https://portswigger.net/burp/documentation/dast/user-guide/reference/reports).
