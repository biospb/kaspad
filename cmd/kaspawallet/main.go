package main

import "github.com/pkg/errors"

func main() {
	subCmd, config := parseCommandLine()

	var err error
	switch subCmd {
	case createSubCmd:
		err = create(config.(*createConfig))
	case balanceSubCmd:
		err = balance(config.(*balanceConfig))
	case sendSubCmd:
		err = send(config.(*sendConfig))
	case createUnsignedTransactionSubCmd:
		err = createUnsignedTransaction(config.(*createUnsignedTransactionConfig))
	case signSubCmd:
		err = sign(config.(*signConfig))
	case broadcastSubCmd:
		err = broadcast(config.(*broadcastConfig))
	case showAddressSubCmd:
		err = showAddress(config.(*showAddressConfig))
	case dumpUnencryptedDataSubCmd:
		err = dumpUnencryptedData(config.(*dumpUnencryptedDataConfig))
	default:
		err = errors.Errorf("Unknown sub-command '%s'\n", subCmd)
	}

	if err != nil {
		printErrorAndExit(err)
	}
}
