/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"context"
	"github.com/spf13/cobra"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	
)

var (
	chainID 				string
	from 					string
	delegatorAdd 			string
	amount 					uint64
	validatorName 			string
	identity 				string
	website 				string
	details 				string
	commissionRate 			float32
	maxCommissionRate		float32
	maxCommissionChangeRate	float32
	sideChainID 			uint64
	sideConsAddr 			string
	sideFeeAddr 			string
	home 					string
)

var infuraURL = "https://testnet.dexit.network";

// startCmd represents the start command
var chainidcmd = &cobra.Command{
	Use:   "chainid",
	Short: "this command allow validator to stake DXT",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chainid called");
		client, err := ethclient.DialContext(context.Background(),infuraURL);
		if err != nil {
			log.Fatalf(err)
		}
		cont, err :=  CD.NewCDTransactor(common.HexToAddress("0x83318095186B28f8195d29c1902c1288381406F8"), client)
		if err != nil {
			log.Fatalf(err)
		}
		methodname, err := cont.getchainID(100)
		if err != nil {
			log.Fatalf(err)
		}
		fmt.Println("Method Name:",methodname)
	},
}

var stakeCmd = &cobra.Command{
	Use:   "staking",
	Short: "this command allow validator to stake DXT",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Staking method called");
	},
}

var createValidator = &cobra.Command{
	Use:   "dexit-create-validator",
	Short: "Create dexit validator",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createValidator is called");
	},
}

var editValidator = &cobra.Command{
	Use:   "dexit-edit-validator",
	Short: "edit dexit validator",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("editValidator is called");
	},
}

func init() {
	rootCmd.AddCommand(stakeCmd)
	stakeCmd.AddCommand(createValidator)
	stakeCmd.AddCommand(editValidator)
	rootCmd.AddCommand(chainidcmd)
	chainidFlags()
	createValidatorFlags();
	editValidatorFlags();



	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createValidatorFlags(){
		//Chain ID flag
		createValidator.Flags().StringVarP(&chainID,"chain-id","c", "", "the chain id of dexit chain.");
		createValidator.MarkFlagRequired("chain-id");
	
		//from flag
		createValidator.Flags().StringVarP(&from,"from","f", "", "address of private key with which to sign this tx, also be used as the validator operator address.");
		createValidator.MarkFlagRequired("from")
	
		//Address delegator flag
		createValidator.Flags().StringVarP(&delegatorAdd,"address-delegator","d", "", "optional, bech32 address of the self-delegator. if not provided, --from address will be used as self-delegator.");
	
		//Amount flag
		createValidator.Flags().Uint64VarP(&amount,"amount","", 0, "self-delegation amount, it has 8 decimal places.");
		createValidator.MarkFlagRequired("amount")
	
		//Moniker flag
		createValidator.Flags().StringVarP(&validatorName,"moniker","m", "", "validator name");
		createValidator.MarkFlagRequired("moniker")
	
		//Identity flag
		createValidator.Flags().StringVarP(&identity,"identity","i", "", "optional identity signature (ex. UPort or Keybase)");
		
		//Website flag
		createValidator.Flags().StringVarP(&website,"website","w", "", "optional website");
		
		//Details flag
		createValidator.Flags().StringVarP(&details	,"details","", "", "optional details");
		
		//Commission rate flag
		createValidator.Flags().Float32VarP(&commissionRate,"commission-rate","", 0, "The initial commission rate percentage, it has 8 decimal places.");
		createValidator.MarkFlagRequired("commission-rate")
		
		//Max Commission rate flag
		createValidator.Flags().Float32VarP(&maxCommissionRate,"commission-max-rate","", 0, "The maximum commission rate percentage, it has 8 decimal places. You can not update this rate.");
		createValidator.MarkFlagRequired("commission-max-rate")
		
		//Max Commission Change rate
		createValidator.Flags().Float32VarP(&maxCommissionChangeRate,"commission-max-change-rate","", 0, "	The maximum commission change rate percentage (per day). You can not update this rate.");
		createValidator.MarkFlagRequired("commission-max-change-rate")
		
		//side Chain_ID  flag
		createValidator.Flags().Uint64VarP(&sideChainID,"side-chain-id","", 0, "chain-id of the side chain the validator belongs to");
		createValidator.MarkFlagRequired("side-chain-id")
		
		//side cons Address flag
		createValidator.Flags().StringVarP(&sideConsAddr,"side-cons-addr","", "", "consensus address of the validator on side chain, please use hex format prefixed with 0x.");
		createValidator.MarkFlagRequired("side-cons-addr")
	
		//side fee Address flag
		createValidator.Flags().StringVarP(&sideFeeAddr,"side-fee-addr","", "", "address that validator collects fee rewards on side chain, please use hex format prefixed with 0x.");
		createValidator.MarkFlagRequired("side-fee-addr")
	
		//Home flag
		createValidator.Flags().StringVarP(&home,"home","", "", "home directory of bnbcli data and config, default to “~/.dxtcli”");
}

func editValidatorFlags(){
	//Chain ID flag
	editValidator.Flags().StringVarP(&chainID,"chain-id","c", "", "the chain id of dexit chain.");
	editValidator.MarkFlagRequired("chain-id");

	//from flag
	editValidator.Flags().StringVarP(&from,"from","f", "", "address of private key with which to sign this tx, also be used as the validator operator address.");
	editValidator.MarkFlagRequired("from")

	//Moniker flag
	editValidator.Flags().StringVarP(&validatorName,"moniker","m", "", "validator name");
	editValidator.MarkFlagRequired("moniker")

	//Identity flag
	editValidator.Flags().StringVarP(&identity,"identity","i", "", "optional identity signature (ex. UPort or Keybase)");
	
	//Website flag
	editValidator.Flags().StringVarP(&website,"website","w", "", "optional website");
	
	//Details flag
	editValidator.Flags().StringVarP(&details	,"details","", "", "optional details");
	
	//Commission rate flag
	editValidator.Flags().Float32VarP(&commissionRate,"commission-rate","", 0, "The initial commission rate percentage, it has 8 decimal places.");
	editValidator.MarkFlagRequired("commission-rate")
	
	//side Chain_ID  flag
	editValidator.Flags().Uint64VarP(&sideChainID,"side-chain-id","", 0, "chain-id of the side chain the validator belongs to");
	editValidator.MarkFlagRequired("side-chain-id")

	//side fee Address flag
	editValidator.Flags().StringVarP(&sideFeeAddr,"side-fee-addr","", "", "address that validator collects fee rewards on side chain, please use hex format prefixed with 0x.");
	editValidator.MarkFlagRequired("side-fee-addr")

}

func chainidFlags () {
	chainidcmd.Flags().StringVarP(&chainID,"chain-id","c", "", "the chain id of dexit chain.");
	chainidcmd.MarkFlagRequired("chain-id");
}