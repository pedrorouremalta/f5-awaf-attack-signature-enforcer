# f5-awaf-attack-signature-enforcer

A tool written in *Go* to help manage and enforce attack signatures on a *BIG-IP Advanced WAF System*. 

It can be used to:

1. List all WAF policies on the system ( *-action list-waf-policies* ). 
2. List all attack signatures for a WAF policy and their respective status ( *-action list-attack-signatures* ). 
3. Print the *signatures enforcement readiness summary* for a specifc WAF policy or for all WAF policies on the system ( *-action print-enforcement-summary*).
4. Enforce all signatures which are *ready to be enforced* for a specific WAF policy ( *-action enforce-ready-signatures* ).

## Usage 

### Environment Variables

|  Variable Name  | Mandatory |          Description            |
|-----------------|-----------|---------------------------------|
| `BIGIP_ADDRESS` |    Yes    | BIGIP IP address or hostname.   |
| `BIGIP_USERNAME`|    Yes    | BIGIP admin username.           |
| `BIGIP_PASSWORD`|    Yes    | BIGIP admin password.           |

