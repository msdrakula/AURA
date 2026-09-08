> Source: https://portswigger.net/burp/documentation/dast/user-guide/working-with-scans/false-positives-best-practice

DAST

# Best practices for managing false positives

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 All automated scanners can generate false positives, which are issues flagged as vulnerabilities that don't pose actual risks.

 Managing false positives enables you to reduce noise, avoid wasting resources, and focus on genuine threats. This section outlines some best practices for identifying and addressing false positives, so that you can integrate them into your workflows.

 Use the **Dashboard** for your scan to review the issues found by Burp Scanner. We recommend the following workflow for managing false positives.

1.

**Review the confidence ratings**

  - The **Issues** tab shows you the confidence rating for each issue.
  - Issues rated as `Certain` are highly likely to be valid vulnerabilities.
  - Issues rated as `Tentative` or `Firm` may require further investigation.
  - Prioritize a deeper analysis for any rating below `Certain`.

1.

**Understand the issue type**

  - In the **Issues** tab, click on an issue and go to the **Advisory** tab.
  - Expand the headings to see detailed information about the issue and possible steps for remediation.

1.

**Understand the issue context**

  - Cross-check flagged issues against your application's architecture and configurations.
  - Use historical scan data to identify recurring patterns that might indicate false positives.

1.

**Use manual testing to validate issues**

  - For issues flagged with lower confidence ratings, manually validate them to confirm whether they pose an actual risk.
  - You can use tools in Burp Suite Professional such as **Repeater** or **Intruder** to replicate the vulnerability and assess its impact.

1.

**Record your decision in Burp Suite DAST**

  - When you have confirmed that an issue is not a genuine vulnerability, set its status to **False positive**. For more information, see [Managing issues in Burp Suite DAST](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues).
  - The status stays with that issue in every later scan of the site, so you do not need to mark it again.
  - The status applies to that issue only. The same vulnerability at another location, or on another site, is a separate issue with its own status.
  - To give the same status to several issues at once, including issues on different sites, use bulk editing. For more information, see [Updating issues in bulk](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues#updating-issues-in-bulk).

1.

**Collaborate with development teams**

  - Work with your development team to determine whether flagged issues align with the intended functionality of the application.
  - Generate reports in Burp Suite DAST to share with your development team, streamlining this process.

1.

**Fine-tune your scanning configurations**

  - Regularly review and update your scanning configurations to minimize noise.
  - Modify scan settings to exclude known false positives or non-applicable checks for your application environment.
  - Explore the scan accuracy options within custom scan configurations. For more information, see [Audit settings](https://portswigger.net/burp/documentation/scanner/scan-configurations/audit-settings).

1.

**Document and refine your process**

  - Create internal documentation to record common false positives and describe how they are handled.
  - Regularly review and refine your process for handling false positives, based on feedback from your team and your evolving business needs.

#### Related pages

- [Managing issues in Burp Suite DAST](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues)
- [Issue statuses](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-statuses)
