package expsql

// A TxFn is a function that will be called with an initialized `Transaction` object.
type TxFn func(Tx) error
