package crc16

import (
	"fmt"
	"testing"
)

var testData = []byte("123456789")

func testChecksum(t *testing.T, params Params) {
	table := MakeTable(params)
	if table == nil {
		t.Errorf("Failed to create %q computer", params.Name)
	}

	crc := Checksum(testData, table)

	if crc != table.params.Check {
		t.Errorf("Invalid %q sample calculation, expected: %X, actual: %X", table.params.Name, table.params.Check, crc)
	}
}

func TestChecksum_ARC(t *testing.T) {
	testChecksum(t, ARC)
}

func TestChecksum_AUGCCIT(t *testing.T) {
	testChecksum(t, AUGCCITT)
}

func TestChecksum_BUYPASS(t *testing.T) {
	testChecksum(t, BUYPASS)
}

func TestChecksum_CCITTFALSE(t *testing.T) {
	testChecksum(t, CCITTFALSE)
}

func TestChecksum_CDMA2000(t *testing.T) {
	testChecksum(t, CDMA2000)
}

func TestChecksum_DDS110(t *testing.T) {
	testChecksum(t, DDS110)
}

func TestChecksum_DECTR(t *testing.T) {
	testChecksum(t, DECTR)
}

func TestChecksum_DECTX(t *testing.T) {
	testChecksum(t, DECTX)
}

func TestChecksum_DNP(t *testing.T) {
	testChecksum(t, DNP)
}

func TestChecksum_EN13757(t *testing.T) {
	testChecksum(t, EN13757)
}

func TestChecksum_GENIBUS(t *testing.T) {
	testChecksum(t, GENIBUS)
}

func TestChecksum_MAXIM(t *testing.T) {
	testChecksum(t, MAXIM)
}

func TestChecksum_MCRF4XX(t *testing.T) {
	testChecksum(t, MCRF4XX)
}

func TestChecksum_RIELLO(t *testing.T) {
	testChecksum(t, RIELLO)
}

func TestChecksum_T10DIF(t *testing.T) {
	testChecksum(t, T10DIF)
}

func TestChecksum_TELEDISK(t *testing.T) {
	testChecksum(t, TELEDISK)
}

func TestChecksum_TMS37157(t *testing.T) {
	testChecksum(t, TMS37157)
}

func TestChecksum_USB(t *testing.T) {
	testChecksum(t, USB)
}

func TestChecksum_CRCA(t *testing.T) {
	testChecksum(t, CRCA)
}

func TestChecksum_KERMIT(t *testing.T) {
	testChecksum(t, KERMIT)
}

func TestChecksum_MODBUS(t *testing.T) {
	testChecksum(t, MODBUS)
}

func TestChecksum_X25(t *testing.T) {
	testChecksum(t, X25)
}

func TestChecksum_XMODEM(t *testing.T) {
	testChecksum(t, XMODEM)
}

func TestHash(t *testing.T) {
	tbl := MakeTable(XMODEM)
	h := New(tbl)

	fmt.Fprint(h, "standard")
	fmt.Fprint(h, " library hash interface")
	sum1 := h.Sum16()
	h.Reset()
	fmt.Fprint(h, "standard library hash interface")
	sum2 := h.Sum16()

	if sum1 != sum2 {
		t.Errorf("CRC16 checksums for chunked input %x should be equal %x", sum1, sum2)
	}

	var crc uint16 = 0xe698
	if sum1 != crc {
		t.Errorf("CRC16 for input should equal %x but was %x", crc, sum1)
	}

	if h.Size() != 2 {
		t.Errorf("CRC16 checksum should have a size of 2 byte. But was %d", h.Size())
	}
	buf := make([]byte, 0, 10)
	buf = h.Sum(buf)
	expected := []byte{0xe6, 0x98}
	if len(buf) != 2 || buf[0] != expected[0] || buf[1] != expected[1] {
		t.Errorf("CRC16 checksum should append %v to slice, but was %v", expected, buf)
	}

	// for the 100% test coverage
	if h.BlockSize() != 1 {
		t.Errorf("Block size should return 1")
	}
}
