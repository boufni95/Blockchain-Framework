package blockchain

//IsBlockValid : chech if block is valid
func IsBlockValid(newBlock, oldBlock Block) (bool, error) {
	if oldBlock.Index+1 != newBlock.Index {
		return false, nil
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false, nil
	}
	chash, err := newBlock.CalculateHash()
	if err != nil {
		return false, err
	}

	if chash != newBlock.Hash {
		return false, nil
	}

	return true, nil
}
