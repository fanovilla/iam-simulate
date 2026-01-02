# AWS Service Control Policy Baseline

For AWS Service Control Policies (SCPs), there isn't a single universal "minimum set" since it depends on your organization's needs, but here are the commonly recommended baseline services that should typically remain accessible:

## Core Management Services

- **AWS Organizations** - Required to manage SCPs themselves
- **IAM** - Identity and access management (though often restricted to specific actions)
- **CloudTrail** - Audit logging and compliance
- **CloudWatch** - Monitoring and alerting
- **AWS Support** - Access to AWS support resources

## Security & Compliance

- **AWS Config** - Configuration compliance tracking
- **GuardDuty** - Threat detection
- **Security Hub** - Centralized security findings
- **KMS** - Encryption key management

## Billing & Cost Management

- **AWS Billing Console** - Cost visibility
- **Cost Explorer** - Cost analysis
- **Budgets** - Cost controls

## Common Operational Services

- **EC2** - Compute instances
- **S3** - Object storage
- **VPC** - Networking
- **CloudFormation** - Infrastructure as code
- **Systems Manager** - Operational management

## Deny-List Strategy

The typical approach is to use a deny-list strategy rather than an allow-list - meaning you explicitly deny problematic services/actions while allowing everything else by default. Common denials include:

- Leaving AWS Organizations
- Disabling CloudTrail logging
- Disabling required security services
- Access to specific restricted regions
- Root user actions