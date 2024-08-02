package main

import (
	"fmt"
	"os"
)

func main() {

	var err error
	var args Args

	err = GetSensitiveArgs(&args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = GetArgs(&args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	awaf := NewAWAFSystem(args.device, args.username, args.password)

	err = awaf.GetPolicies()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch args.action {

	case "list-waf-policies":

		awaf.ListPolicies()
		os.Exit(0)

	case "list-attack-signatures":

		err = awaf.LoadAttackSignaruresByPolicyName(args.policy)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = awaf.ListAttackSignaturesByPolicyName(args.policy, args.sigstatus)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Exit(0)

	case "print-enforcement-summary":

		if args.policy == "" {

			err = awaf.LoadAttackSignaruresAllPolicies()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			awaf.PrintSignaturesEnforcementReadinessSummaryAllPolicies()

		} else {

			err = awaf.LoadAttackSignaruresByPolicyName(args.policy)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = awaf.PrintSignaturesEnforcementReadinessSummaryByPolicyName(args.policy)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		}

	case "enforce-ready-signatures":

		err = awaf.LoadAttackSignaruresByPolicyName(args.policy)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = awaf.EnforceSignaturesReadyToBeEnforced(args.policy)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

}
