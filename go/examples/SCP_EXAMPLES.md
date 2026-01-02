# SCP Example Files - Source Documentation

This document maps each SCP example file to its source in the AWS Organizations documentation.

## General Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_general.html

| Filename | Description |
|----------|-------------|
| `scp-allow-all.json` | FullAWSAccess - Default allow all policy |
| `scp-deny-by-region.json` | Deny access to AWS based on the requested AWS Region |
| `scp-prevent-iam-role-changes.json` | Prevent IAM users and roles from making certain changes |
| `scp-prevent-iam-role-changes-with-admin-exception.json` | Prevent IAM users and roles from making specified changes, with an exception for a specified admin role |
| `scp-require-mfa-stop-ec2.json` | Require MFA to stop an Amazon EC2 instance |
| `scp-block-root-user-access.json` | Block service access for the root user |
| `scp-prevent-leave-organization.json` | Prevent member accounts from leaving the organization |

## Amazon Bedrock Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_bedrock.html

| Filename | Description |
|----------|-------------|
| `scp-bedrock-deny-specific-models.json` | Deny access to specific Amazon Bedrock models |
| `scp-bedrock-restrict-to-permitted-models.json` | Restrict access to specific Amazon Bedrock models (deny all except permitted) |
| `scp-bedrock-restrict-api-keys.json` | Restrict creation and use of Amazon Bedrock API keys |

## Amazon CloudWatch Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_cloudwatch.html

| Filename | Description |
|----------|-------------|
| `scp-cloudwatch-prevent-disable-and-modify.json` | Prevent users from disabling CloudWatch or altering its configuration |

## AWS Config Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_config.html

| Filename | Description |
|----------|-------------|
| `scp-config-prevent-disable-and-modify.json` | Prevent users from disabling AWS Config or changing its rules |

## Amazon EC2 Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_ec2.html

| Filename | Description |
|----------|-------------|
| `scp-ec2-require-instance-type.json` | Require Amazon EC2 instances to use a specific type |
| `scp-ec2-require-imdsv2.json` | Prevent launching EC2 instances without IMDSv2 |

## Amazon GuardDuty Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_guardduty.html

| Filename | Description |
|----------|-------------|
| `scp-guardduty-prevent-disable-and-modify.json` | Prevent users from disabling GuardDuty or modifying its configuration |

## AWS Resource Access Manager (RAM) Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_ram.html

| Filename | Description |
|----------|-------------|
| `scp-ram-prevent-external-sharing.json` | Prevent external sharing |
| `scp-ram-restrict-to-specific-accounts.json` | Restrict resource sharing to specific account IDs |
| `scp-ram-prevent-sharing-with-orgs.json` | Prevent sharing with organizations or organizational units (OUs) |

## Amazon S3 Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_s3.html

| Filename | Description |
|----------|-------------|
| `scp-s3-prevent-unencrypted-uploads-basic.json` | Prevent Amazon S3 unencrypted object uploads (Basic) |
| `scp-s3-prevent-unencrypted-uploads-with-type.json` | Prevent Amazon S3 unencrypted object uploads (With Encryption Type Enforcement) |

## Resource Tagging Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_tagging.html

| Filename | Description |
|----------|-------------|
| `scp-tagging-require-tags-on-resources.json` | Require a tag on specified created resources |
| `scp-tagging-prevent-tag-modification.json` | Prevent tags from being modified except by authorized principals |

## Amazon VPC Examples

Source: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_vpc.html

| Filename | Description |
|----------|-------------|
| `scp-vpc-prevent-delete-flow-logs.json` | Prevent users from deleting Amazon VPC flow logs |
| `scp-vpc-prevent-internet-access.json` | Prevent any VPC that doesn't already have internet access from getting it |

## Summary

- **Total Examples**: 25 SCP policy files
- **Main Documentation**: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples.html
- **Last Updated**: 2025-12-15

## Notes

- All examples are taken directly from AWS official documentation
- These are example policies and should be customized for your specific use cases
- Some policies contain placeholder values (e.g., role names, account IDs) that need to be replaced
- Test policies in a non-production environment before applying to production accounts
