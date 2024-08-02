# f5-awaf-attack-signature-enforcer

A tool written in *Go* to help manage and enforce attack signatures on a *BIG-IP Advanced WAF System*. 

It can be used to:

1. List all WAF policies on the system ( *-action list-waf-policies* ). 
2. List all attack signatures for a WAF policy and their respective status ( *-action list-attack-signatures* ). Optionally, a filter for the status can be specified. 
3. Print the *signatures enforcement readiness summary* for a specifc WAF policy or for all WAF policies on the system ( *-action print-enforcement-summary*).
4. Enforce all signatures which are *ready to be enforced* for a specific WAF policy ( *-action enforce-ready-signatures* ).

## Usage 

### Environment Variables

|  Variable Name  | Mandatory |          Description            |
|-----------------|-----------|---------------------------------|
| `BIGIP_ADDRESS` |    Yes    | BIGIP IP address or hostname.   |
| `BIGIP_USERNAME`|    Yes    | BIGIP admin username.           |
| `BIGIP_PASSWORD`|    Yes    | BIGIP admin password.           |

### Command-line options

|    Option   | Mandatory |        Default       |         Description            |
|-------------|-----------|----------------------|--------------------------------|
| `-action`   |    Yes    | *list-waf-policies*  | Specify the action which will be performed. Allowed values are: 'list-waf-policies','list-attack-signatures','print-enforcement-summary','enforce-ready-signatures'. |
| `-policy`   |    No     | Specify the WAF policy in whch the *action* will be applied. Mandatory for the actions *list-attack-signatures* and *enforce-ready-signatures*. Optional for the action *print-enforcement-summary* |
| `-sigstatus`|    No     | Specify a *status* filter when listing attack signatures. This option is optional and the allowed values are: 'all','ready to be enforced','not enforced (has suggestions)','not enforced','enforced (has suggestions)','enforced'. |