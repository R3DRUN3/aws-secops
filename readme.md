# AWS SECOPS
[![Vuln scan for aws-secops using Snyk](https://github.com/R3DRUN3/aws-secops/actions/workflows/snyk-scan.yml/badge.svg)](https://github.com/R3DRUN3/aws-secops/actions/workflows/snyk-scan.yml)

## ABSTRACT

Make use of aws Security Hub service to launch a fast vulnerability assessment against your aws asset.

## Requirements

`aws account` `golang version >= 1.16`

## Instructions

Clone this folder and run the code:

```console
git clone  \
&& cd \
&& go run main.go
```

You can also build the docker image yourself or download it from docker hub.

To run as a docker container you need to set some mandatory env vars:

```docker
docker run -it --rm \
--env AWS_REGION=AWS_REGION_HERE \
--env AWS_ACCESS_KEY_ID=AWS_ACCESS_KEY_HERE \
--env AWS_SECRET_ACCESS_KEY=AWS_ACCESS_SECRET_HERE \
r3drun3/aws-secops:v1.0.0
```

Output Sample:

```console
      ___           ___           ___                   ___           ___           ___           ___           ___           ___
     /\  \         /\__\         /\  \                 /\  \         /\  \         /\  \         /\  \         /\  \         /\  \
    /::\  \       /:/ _/_       /::\  \               /::\  \       /::\  \       /::\  \       /::\  \       /::\  \       /::\  \
   /:/\:\  \     /:/ /\__\     /:/\ \  \             /:/\ \  \     /:/\:\  \     /:/\:\  \     /:/\:\  \     /:/\:\  \     /:/\ \  \
  /::\~\:\  \   /:/ /:/ _/_   _\:\~\ \  \           _\:\~\ \  \   /::\~\:\  \   /:/  \:\  \   /:/  \:\  \   /::\~\:\  \   _\:\~\ \  \
 /:/\:\ \:\__\ /:/_/:/ /\__\ /\ \:\ \ \__\         /\ \:\ \ \__\ /:/\:\ \:\__\ /:/__/ \:\__\ /:/__/ \:\__\ /:/\:\ \:\__\ /\ \:\ \ \__\
 \/__\:\/:/  / \:\/:/ /:/  / \:\ \:\ \/__/         \:\ \:\ \/__/ \:\~\:\ \/__/ \:\  \  \/__/ \:\  \ /:/  / \/__\:\/:/  / \:\ \:\ \/__/
      \::/  /   \::/_/:/  /   \:\ \:\__\            \:\ \:\__\    \:\ \:\__\    \:\  \        \:\  /:/  /       \::/  /   \:\ \:\__\
      /:/  /     \:\/:/  /     \:\/:/  /             \:\/:/  /     \:\ \/__/     \:\  \        \:\/:/  /         \/__/     \:\/:/  /
     /:/  /       \::/  /       \::/  /               \::/  /       \:\__\        \:\__\        \::/  /                     \::/  /
     \/__/         \/__/         \/__/                 \/__/         \/__/         \/__/         \/__/                       \/__/


GRAVITY:  LOW
[1.1 Avoid the use of the root user]
The root user has unrestricted access to all resources in the AWS account. It is highly recommended that the use of this user be avoided.For directions on how to fix this issue, consult the AWS Security Hub CIS documentation. ===> https://docs.aws.amazon.com/console/securityhub/standards-cis-1.1/remediation

GRAVITY:  MEDIUM
[2.5 Ensure AWS Config is enabled]
AWS Config is a web service that performs configuration management of supported AWS resources within your account and delivers log files to you. The recorded information includes the configuration item (AWS resource), relationships between configuration items (AWS resources), and any configuration changes between resources. It is recommended to enable AWS Config in all regions.For directions on how to fix this issue, consult the AWS Security Hub CIS documentation. ===> https://docs.aws.amazon.com/console/securityhub/standards-cis-2.5/remediation

......................CENSORED...........................
......................CENSORED...........................
......................CENSORED...........................
......................CENSORED...........................
......................CENSORED...........................


VULNERABILITY RECAP: FOUND 15 LOW, 2 MEDIUM AND 0 HIGH

```