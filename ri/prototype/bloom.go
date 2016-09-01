package prototype

import (
	"github.com/reusee/mmh3"
	"hash/fnv"

	"io"
	"fmt"
)

const (
	MinimalBloomFilterSize int = 16
	DefaultBloomFilterSize int = 512 * 5
	BloomFilterDoubleHashI0 int = 123
	BloomFilterDoubleHashI1 int = 456
)

func Hash(key string) []int {

	m := []int{0,0,0,0}

	h0 := fnv.New64()
	io.WriteString(h0,key)
	hash0 := h0.Sum(nil)
	
	h1 := mmh3.New128()
	io.WriteString(h1,key)
	hash1 := h1.Sum(nil)

	m[0] += int(hash0[0])
	m[0] += int(hash0[1])
	m[0] += int(hash0[2])
	m[0] += int(hash0[3])
	m[0] += int(hash0[4])
	m[0] += int(hash0[5])
	m[0] += int(hash0[6])
	m[0] += int(hash0[7])

	m[3] += int(hash1[0])
	m[3] += int(hash1[1])
	m[3] += int(hash1[2])
	m[3] += int(hash1[3])
	m[3] += int(hash1[4])
	m[3] += int(hash1[5])
	m[3] += int(hash1[6])
	m[3] += int(hash1[7])	
	m[3] += int(hash1[8])
	m[3] += int(hash1[9])
	m[3] += int(hash1[10])
	m[3] += int(hash1[11])
	m[3] += int(hash1[12])
	m[3] += int(hash1[13])
	m[3] += int(hash1[14])
	m[3] += int(hash1[15])	

	m[1] = m[0] + (BloomFilterDoubleHashI0 * m[3]) + (BloomFilterDoubleHashI0 * BloomFilterDoubleHashI0)
	m[2] = m[0] + (BloomFilterDoubleHashI1 * m[3]) + (BloomFilterDoubleHashI1 * BloomFilterDoubleHashI1)

	return m
}

type BloomFilter struct {
	bits []int
	keys int
}

func (f *BloomFilter) Raw() (int,[]int) {
	bits := make([]int,len(f.bits))
	copy(bits,f.bits)
	return f.keys,bits
} 

func (f *BloomFilter) Stats() (int,float64,float64) {
	if f.keys == 0 {
		return f.keys,0.0,1.0
	}

	zeros := 0
	for i := 0; i < len(f.bits); i++ {
		if f.bits[i] == 0 {
			zeros ++
		}
	}

	ratio := (float64(len(f.bits) - zeros) / 4) / float64(f.keys)
	sparse := (float64(zeros) / float64(len(f.bits)))	

	return f.keys,ratio,sparse
}

func (f *BloomFilter) String() string {
	_,ratio,sparse := f.Stats()
	return fmt.Sprintf("Bloom Filter -- size %dbytes, %d key(s), ratio %.2f, sparseness %.2f",
											len(f.bits) * 4 + 4,f.keys,ratio,sparse)
}

func (f *BloomFilter) Len() int {
	return len(f.bits)
} 

func (f *BloomFilter) Append(keys ...string) *BloomFilter {
	if len(keys) == 0 {
		return f
	}
	
	bits := make([]int,len(f.bits))
	copy(bits,f.bits)

	dups := make(map[string]bool,0)

	for _,key := range keys {
		if _,exists := dups[key]; exists {
			continue
		}

		m := Hash(key)
		for _,q := range m {
			bits[q % len(f.bits)] ++
		}

		dups[key] = true
	}
	return &BloomFilter{bits,f.keys + len(dups)}
}
	
func (f *BloomFilter) IsMember(keys ...string) bool {
	if len(keys) == 0 || f.keys == 0 {
		return false
	}

	dups := make(map[string]bool,0)
	for _,key := range keys {
		if _,exists := dups[key]; exists {
			continue
		}
		m := Hash(key)

		c := 0
		for _,q := range m {
			if f.bits[q % len(f.bits)] > 0 {
				c ++
			}
		}
		if c != 4 {
			return false
		}
		dups[key] = true
	}
	return true
}

func NewBloomFilter(size int) *BloomFilter {
	if size < MinimalBloomFilterSize {
		size = MinimalBloomFilterSize
	}	
	bits := make([]int,size)
	return &BloomFilter{bits,0}
}

func SetBloomFilter(bits []int,size int) *BloomFilter {
	bits2 := make([]int,len(bits))
	copy(bits2,bits)
	return &BloomFilter{bits2,size}
}
	

func (f *BloomFilter) Print() string {

	blocks := make([]string,0)
	blocks = append(blocks,f.String() + "\n\n")	
	l0 := ""
	l01 := ""
	l1 := ""
	columns := 0
	for i := 0; i < len(f.bits); i++ {
			if f.bits[i] > 0 {
				l0 += fmt.Sprintf("%03d|",i)
				l01 += fmt.Sprintf("---+")
				l1 += fmt.Sprintf("%03d ",f.bits[i])
				columns++
			}

			if columns >= 20 {
				columns = 0
				blocks = append(blocks,l0 + "\n" + l01 + "\n" +  l1 + "\n\n")
				l0 = ""
				l01 = ""
				l1 = ""
			}
	}
	if columns > 0 {
		blocks = append(blocks,l0 + "\n" + l01 + "\n" + l1 + "\n\n")
	}

	out := ""
	for _,b := range blocks {
		out += b
	}

	return out 
}





