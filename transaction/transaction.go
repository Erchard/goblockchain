package transaction

type Transaction struct {
	TxHash    []byte
	Block     int32
	Sender    []byte
	Receiver  []byte
	Amount    int64
	Fee       int64
	Timestamp int32
	PublicKey []byte
	Signature []byte
}

type TransactionJSON struct {
	TxHash    string
	Block     int32
	Sender    string
	Receiver  string
	Amount    int64
	Fee       int64
	Timestamp int32
	PublicKey string
	Signature string
}

type TransactionRaw struct {
	TxHash    []byte
	TxData    []byte
	PublicKey []byte
	Signature []byte
}
