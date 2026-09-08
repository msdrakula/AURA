> Source: https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/external-database-requirements

DAST

# System requirements for the external database (Kubernetes)

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 For Kubernetes, you need to [set up an external database](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/setup-external-database). The external database stores the accumulated data from your scans. The required size depends on the number of scans that you perform and the number of issues found.


 The following table indicates the quantity of data that you are likely to accumulate:

| **
 Number of scans
 **  | **
 Data storage
 **
|
 1,000
  |
 500 MB

|
 10,000
  |
 5 GB

|
 100,000
  |
 50 GB


## Supported database versions

 We support the following external databases. We expect newer versions from these vendors to be supported, but these are the latest that we have actively tested.


 We try to support older versions as long as they are supported by the vendor.

| **Type**  | **Latest version tested**
|

MariaDB |

11.2.3
|

Microsoft SQL server |

2022 (16.0.1000.6)
|

MySQL |

8.3.0
|

Oracle |

23c (23.0.0.0.0)
|

PostgreSQL |

17.4

#### Note

 AWS Aurora databases are not supported.


## Database hardware requirements

 The size of your database instance depends on the number of concurrent scans that you want to run, and the type of database.


 For a database instance such as AWS RDS, we recommend a minimum of 2 vCPUs and 8GB RAM, for example a db.t3.large instance.


 If you need help with the system requirements, please [email our support team](mailto:support@portswigger.net).


#### Related pages

-  [Setting up the external database](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/setup-external-database).

-  [Planning the deployment process](https://portswigger.net/burp/documentation/dast/setup).

-  [Architecture overview - Standard instance](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/architecture-overview).

-  [Architecture overview - Kubernetes instance](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/architecture-overview).


**Next step - **Setting up the external database  [CONTINUE](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes/setup-external-database)
