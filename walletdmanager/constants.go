// Copyright (c) 2018, The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in TRTL
	DefaultTransferFee float64 = 2

	logWalletdCurrentSessionFilename     = "QwertyZero-service-session.log"
	logWalletdAllSessionsFilename        = "QwertyZero-service.log"
	logTurtleCoindCurrentSessionFilename = "QwertyZeroCoind-session.log"
	logTurtleCoindAllSessionsFilename    = "QwertyZeroCoind.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "QwertyZero-service"
	turtlecoindCommandName               = "QwertyZeroCoind"
)
