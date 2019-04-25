package core

const (
	//AreYouNode ask a connection if it is a node
	AreYouNode MessageType = 1 //DELETE

	//IAmNode say that you are a node
	IAmNode MessageType = 3 //

	//IAmReady say you are ready
	IAmReady MessageType = 5

	//MyConfig send my config for check
	MyConfig MessageType = 7

	//GiveMeNodes ask for other connected nodes
	GiveMeNodes MessageType = 9

	//GiveYouNodes give other connected nodes
	GiveYouNodes MessageType = 11

	//GiveMeChain ask for the full chain starting from index
	GiveMeChain MessageType = 13

	//GiveYouChain sends the chain starting from index
	GiveYouChain MessageType = 15

	//MyStake communicate your intention to validate and your stake
	MyStake MessageType = 19

	//GiveMeValidators ask fot other validators connected
	GiveMeValidators MessageType = 21

	//GiveYouValidators send other validators connected
	GiveYouValidators MessageType = 23

	//ErrorCommitted error in the last block
	ErrorCommitted MessageType = 25

	//ErrorFound error in the full chain
	ErrorFound MessageType = 27

	//BanThisValidator ban a validator after suspiscius behavior
	BanThisValidator MessageType = 29

	//MyNextBlock propose a new block
	MyNextBlock MessageType = 31
)

type BCMessage interface {
	Message
	GenerateBCMessage() []byte
}

func (m *message) GenerateBCMessage() []byte {
	var b []byte
	switch m.mType {
	case IAmNode:
		{
			//TODO: write send to write generate
			b = make([]byte, 1)
			b[0] = (byte)(m.mType)
			return b
		}
	}
	return b
}
