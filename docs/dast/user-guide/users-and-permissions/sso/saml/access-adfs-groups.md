> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-adfs-groups

DAST

# Enabling Burp Suite DAST to access your ADFS groups

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you're not using SCIM, you can match local Burp Suite DAST groups to your ADFS groups. This enables you to manage these groups locally.


 To make sure that the group membership of your users is in a format that Burp Suite DAST can recognize, you have the following options:


-  [Create a central claim issuance policy](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-adfs-groups#create-a-central-claim-issuance-policy) that handles all of your groups in the same way.

-
 Configure [claim rules for each group](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-adfs-groups#create-claim-rules-for-each-group-individually) that you want to expose to Burp Suite DAST individually.


#### Note

 You can also use a combination of both approaches. In this case, the groups available to Burp Suite DAST would be the union of the groups covered by the claim issuance policy and any additional groups that have their own claim rules.


## Create a central claim issuance policy

 To expose all of your users' groups to Burp Suite DAST, configure a central claim issuance policy. This allows you to manage the claim rules for all of your groups in one place. It also removes the need to configure claim rules each time you add a new group.


 The downside to this approach is that your groups must keep their existing group names. For example, if your group is called `BSEE_View_Scans` in Active Directory, you need to use this exact name for the corresponding user group in Burp Suite DAST. For more information, see [Configuring user permissions for SSO](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).


1.
 Open the ADFS Management tool and go to the list of relying party trusts.

1.
 Right-click on the entry you created for Burp Suite DAST and select **Edit claim issuance policy**.

1.
 Use the wizard to configure the following rules:


### Rule 1

- **Template:** Send LDAP attributes as claims
- **Name:** `Send UPN as nameId`
- **Rule:** `User-Principle-Name` => `Name ID`

### Rule 2

- **Template:** Send LDAP attributes as claims
- **Name:** `AccountName`
- **Rule:** `SAM-Account-Name` = `Windows Account Name`

### Rule 3

- **Template:** Send claims using a custom rule
- **Name:** `nameDN`
- **Rule:** `c:[Type == "http://schemas.microsoft.com/ws/2008/06/identity/claims/windowsaccountname", Issuer == "AD AUTHORITY"] => add(store = "Active Directory", types = ("http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameDN"), query = ";distinguishedName;{0}", param = c.Value);`

### Rule 4

- **Template:** Send claims using a custom rule
- **Name:** `Group`
- **Rule:** `c1:[Type == "http://schemas.microsoft.com/ws/2008/06/identity/claims/windowsaccountname", Issuer == "AD AUTHORITY"] && c2:[Type == "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameDN"] => add(store = "Active Directory", types = ("http://schemas.xmlsoap.org/claims/Group"), query = "(member:1.2.840.113556.1.4.1941:={1});samaccountname;{0}", param = c1.Value, param = c2.Value);`

### Rule 5

- **Template:** Pass through or filter an incoming claim
- **Name:** `IssuedGroup`
- **Rule:** Group - Pass through all claim values

 All the groups that the user belongs to are sent with every claim to Burp Suite DAST. If you add new groups, these rules automatically apply to them as well.


## Create claim rules for each group individually

 You can create claim rules on a group-by-group basis. This gives you more granular control over the groups and related information that each claim exposes to Burp Suite DAST.


 You can output the group with a different name than the one used in Active Directory. For example, if your group is called `BSEE_View_Scans`, you can output this with a more user-friendly name, such as "Scan viewers". You can then use this name for the corresponding group in Burp Suite DAST. For more information, see [Configuring user permissions for SSO](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).


1.
 Open the ADFS Management tool and go to the list of relying party trusts.

1.
 Right-click on the entry you created for Burp Suite DAST and select **Edit claim issuance policy**.

1.
 From the **Claim rule template** drop-down list, select **Send Group Membership as Claim** and click **Next**.

1.
 Enter a name for the claim rule.

1.
 To configure a claim rule, select **User's group** and select the group.

1.
 From the **Outgoing claim type** drop-down list, select **Group**.

1.
 In the **Outgoing claim value** field, enter a new name that you want to use for this group when sending a claim.

1.
 Repeat this process for each group that you want to expose to Burp Suite DAST.


 If you add new groups in the future, you will need to repeat this process for each of them.


## Adding your groups to Burp Suite DAST

 To grant permissions, set the **Identity provider group** value on each group that you create in Burp Suite DAST. For more information, see [Creating local groups for SAML or LDAP](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).
