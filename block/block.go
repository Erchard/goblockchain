package block

import "../transaction"

type Block struct {
	BlHash    string                    `json:"bl_hash"`
	Height    uint32                    `json:"height"`
	Previous  string                    `json:"prev"`
	Timestamp uint32                    `json:"timestamp"`
	Nonce     string                    `json:"nonce"`
	TxList    []transaction.Transaction `json:"tx_list"`
}

type BlockRaw struct {
	BlHash    []byte
	BlData    []byte
	PublicKey []byte
	Signature []byte
}
