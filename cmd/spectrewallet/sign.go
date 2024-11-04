package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/cmd/spectrewallet/daemon/server"
	"github.com/spectre-project/spectred/cmd/spectrewallet/keys"
	"github.com/spectre-project/spectred/cmd/spectrewallet/libspectrewallet"
)

func sign(conf *signConfig) error {
	if conf.Transaction == "" && conf.TransactionFile == "" {
		return errors.Errorf("Either --transaction or --transaction-file is required")
	}
	if conf.Transaction != "" && conf.TransactionFile != "" {
		return errors.Errorf("Both --transaction and --transaction-file cannot be passed at the same time")
	}

	keysFile, err := keys.ReadKeysFile(conf.NetParams(), conf.KeysFile)
	if err != nil {
		return err
	}

	if len(conf.Password) == 0 {
		conf.Password = keys.GetPassword("Password:")
	}
	privateKeys, err := keysFile.DecryptMnemonics(conf.Password)
	if err != nil {
		return err
	}

	transactionsHex := conf.Transaction
	if conf.TransactionFile != "" {
		transactionHexBytes, err := ioutil.ReadFile(conf.TransactionFile)
		if err != nil {
			return errors.Wrapf(err, "Could not read hex from %s", conf.TransactionFile)
		}
		transactionsHex = strings.TrimSpace(string(transactionHexBytes))
	}
	partiallySignedTransactions, err := server.DecodeTransactionsFromHex(transactionsHex)
	if err != nil {
		return err
	}

	updatedPartiallySignedTransactions := make([][]byte, len(partiallySignedTransactions))
	for i, partiallySignedTransaction := range partiallySignedTransactions {
		updatedPartiallySignedTransactions[i], err =
			libspectrewallet.Sign(conf.NetParams(), privateKeys, partiallySignedTransaction, keysFile.ECDSA)
		if err != nil {
			return err
		}
	}

	areAllTransactionsFullySigned := true
	for _, updatedPartiallySignedTransaction := range updatedPartiallySignedTransactions {
		// This is somewhat redundant to check all transactions, but we do that just-in-case
		isFullySigned, err := libspectrewallet.IsTransactionFullySigned(updatedPartiallySignedTransaction)
		if err != nil {
			return err
		}
		if !isFullySigned {
			areAllTransactionsFullySigned = false
		}
	}

	if areAllTransactionsFullySigned {
		fmt.Fprintln(os.Stderr, "The transaction is signed and ready to broadcast")
	} else {
		fmt.Fprintln(os.Stderr, "Successfully signed transaction")
	}

	fmt.Println(server.EncodeTransactionsToHex(updatedPartiallySignedTransactions))
	return nil
}
