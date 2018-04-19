package goavro

const EMPTY64 uint64 = 0xc15d213aa4d7a795
type Rabin struct {
    table []uint64
}

func NewRabin() *Rabin{
    return &Rabin{
        table: makeTable(),
    }
}

func (r *Rabin) Fingerprint(buf []byte) []byte {
    fp := r.fingerprint64(buf)
    result := make([]byte, 8)
    for i := 0; i < 8; i++ {
        result[i] = byte(fp)
        fp >>= 8
    }
    return result;
}
func (r *Rabin) fingerprint64(buf []byte) uint64 {
    result := EMPTY64
    for i := 0; i < len(buf); i++ {
        result = (result >> 8) ^ r.table[(result ^ uint64(buf[i])) & 0xff]
    }
    return result
}

func makeTable() []uint64{
    table := make([]uint64, 256)
    for i := 0; i < 256; i++ {
        fp := uint64(i)
        for j := 0; j < 8; j++ {
            fp = (fp >> 1) ^ (EMPTY64 & -(fp & 1))
        }
        table[i] = fp
    }
    return table
}
