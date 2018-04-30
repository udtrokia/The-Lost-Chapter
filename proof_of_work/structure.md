# structure

```go

func NewBlock (data string, prevBlockhash []byte) *Block {}

type Block struct {
	TimeStamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

---

func NewBlockchain() *Blockchain {}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {}

---

func NewProofOfWork(b *Block) *ProofOfWork {}

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {}
func (pow *ProofOfWork) Run() (int, []byte) {}
func (pow *ProofOfWork) Validate() bool {}

```
