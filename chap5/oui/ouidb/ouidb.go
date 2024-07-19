package ouidb

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"slices"
	"strconv"
)

// 200000 以内倍增的质数列表.
var primes = [...]uint32{7, 17, 37, 79, 163, 331, 673, 1361, 2729, 5471, 10949, 21911, 43853, 87719, 175447}

// nextPrime 在 primes 里寻找下一个质数.
// 用质数做桶数量, 可以最大程度减少哈希冲突的概率, 使哈希后的数据分布的更加均匀
func nextPrime(n uint32) uint32 {
	i, _ := slices.BinarySearch(primes[:], n)
	return primes[i]
}

// 文件头
const fileHeader = "OUI\x00"

// 此文件中编码数值时用到的字节序
var Order binary.ByteOrder = binary.LittleEndian

type OUI struct {
	Id      uint32
	Company string
}

// Write 写出 data 到文件哈希表 w
func Write(w io.Writer, data []OUI) error {
	// 桶个数. 是项目个数的 1.25倍(10/8)
	var nBucket = nextPrime(uint32(len(data) * 10 / 8))
	// 桶数据的内存表现形式
	var buckets = make([][]*OUI, nBucket)
	// 把数据逐条加入桶
	for _, oui := range data {
		// 直接使用 Id 做 hash
		b := &buckets[oui.Id%nBucket]
		*b = append(*b, &oui)
	}
	// 桶索引. bucketIndex[i] 为 桶 i 的数据起始位置偏移.
	var bucketIndex = make([]int, nBucket)
	// 文件中的桶数据
	var bucketsData bytes.Buffer
	var buf [4]byte
	for i, b := range buckets {
		// 记录桶 i 的数据偏移.
		bucketIndex[i] = bucketsData.Len()
		bucketLen := len(b)
		if bucketLen > 0xFF {
			return fmt.Errorf("too many buckets %v", bucketLen)
		}

		// 桶数据的格式为
		// 1字节列表长度n 应列表项0 应列表项1 ... 应列表项n

		bucketsData.WriteByte(byte(bucketLen)) // 写列表长度
		for _, oui := range b {
			// 列表项格式为
			// 4字节Id 1字节Company长度n 长度为n的UTF-8序列
			Order.PutUint32(buf[:], oui.Id)
			bucketsData.Write(buf[:]) // 写 Id
			companyLen := len(oui.Company)
			if companyLen > 0xFF {
				return errors.New("company too long")
			}
			bucketsData.WriteByte(byte(companyLen)) // 写 Company 的长度
			bucketsData.WriteString(oui.Company)    // 写 Company 数据
		}
	}

	// 写文件头
	if _, err := io.WriteString(w, fileHeader); err != nil {
		return err
	}
	// 写桶个数
	Order.PutUint32(buf[:], nBucket)
	if _, err := w.Write(buf[:]); err != nil {
		return err
	}
	// 写桶索引
	dataOffset := uint32(len(fileHeader)) + 4 + 4*nBucket // bucketsData 在文件数据中的偏移
	for _, i := range bucketIndex {
		// 修正偏移
		Order.PutUint32(buf[:], uint32(i)+dataOffset)
		if _, err := w.Write(buf[:]); err != nil {
			return err
		}
	}
	// 写桶数据
	if _, err := io.Copy(w, &bucketsData); err != nil {
		return err
	}
	return nil
}

// Generate 从 https://standards-oui.ieee.org/ 读取数据并生成数据库文件到 w
func Generate(w io.Writer) error {
	resp, err := http.Get("https://standards-oui.ieee.org/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data []OUI
	lineRegexp := regexp.MustCompile(`([0-9A-F]{6})\s+\(base 16\)\s+(.+)`)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return err
		}
		m := lineRegexp.FindSubmatch(scanner.Bytes())
		if m == nil {
			continue
		}
		id, err := strconv.ParseUint(string(m[1]), 16, 32)
		if err != nil {
			return fmt.Errorf("invalid id %v: %w", string(m[1]), err)
		}
		if id > 0xFFFFFF {
			return fmt.Errorf("id %v is too large", id)
		}
		data = append(data, OUI{Id: uint32(id), Company: string(m[2])})
	}
	return Write(w, data)
}

type DB struct {
	r       io.ReadSeeker
	nBucket uint32
}

// NewDB 创建一个 *DB
// r 必须为 Write 写出的文件哈希表.
func NewDB(r io.ReadSeeker) (db *DB, err error) {
	var buf [4]byte
	// 读文件头
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	if string(buf[:]) != fileHeader {
		err = fmt.Errorf("invalid file header %v", buf[:])
		return
	}
	// 读桶个数
	if _, err = io.ReadFull(r, buf[:]); err != nil {
		return
	}
	nBucket := Order.Uint32(buf[:])
	return &DB{r, nBucket}, nil
}

// Lookup 从数据中查询指定 OUI 对应的 company
// 如果找到, 返回非空 company 和 nil error
func (db *DB) Lookup(oui string) (company string, err error) {
	// 把 oui 解析为 uint32
	id, err := strconv.ParseUint(oui, 16, 32)
	if err != nil {
		err = fmt.Errorf("invalid id %v: %w", oui, err)
		return
	}
	if id > 0xFFFFFF {
		err = fmt.Errorf("id %v is too large", id)
		return
	}
	// 跳到索引
	// 8 为文件头和桶个数所占的字节数
	// 4 为每个索引所占用的字节数
	if _, err = db.r.Seek(8+(int64(id)%int64(db.nBucket))*4, io.SeekStart); err != nil {
		return
	}
	var buf [4]byte
	// 读索引
	if _, err = io.ReadFull(db.r, buf[:]); err != nil {
		return
	}
	offset := Order.Uint32(buf[:])
	// 跳到桶数据
	if _, err = db.r.Seek(int64(offset), io.SeekStart); err != nil {
		return
	}
	// 读桶列表长度
	if _, err = io.ReadFull(db.r, buf[:1]); err != nil {
		return
	}
	// 依次读列表项
	for range buf[0] {
		// 读 Id
		if _, err = io.ReadFull(db.r, buf[:]); err != nil {
			return
		}
		// 读 Company 长度
		var lenBuf [1]byte
		if _, err = io.ReadFull(db.r, lenBuf[:]); err != nil {
			return
		}
		if id == uint64(Order.Uint32(buf[:])) {
			// Id 匹配
			companyBuf := make([]byte, lenBuf[0])
			// 读 Company
			if _, err = io.ReadFull(db.r, companyBuf); err != nil {
				return
			}
			company = string(companyBuf)
			return
		}
		// Id 不匹配, 跳过 Company
		if _, err = db.r.Seek(int64(lenBuf[0]), io.SeekCurrent); err != nil {
			return
		}
	}
	return
}
