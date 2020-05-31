package qx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTxSign(t *testing.T) {
	k := "c39fb9103419af8be42385f3d6390b4c0c8f2cb67cf24dd43a059c4045d1a409"
	tx := "0100000001255fea249c9747f7f4a8c432ca6f6bbed20db023fa9101288cad1a4e8056a5f600000000ffffffff0100943577000000001976a914c50b62be2f7c23cf0b9d904fa9984efbdb75859888ac0000000000000000a2b54c5e0100"
	net := "testnet"
	rs, _ := TxSign(k, tx, net)
	fmt.Println(rs)
	assert.Equal(t, rs, "0100000001255fea249c9747f7f4a8c432ca6f6bbed20db023fa9101288cad1a4e8056a5f600000000ffffffff0100943577000000001976a914c50b62be2f7c23cf0b9d904fa9984efbdb75859888ac0000000000000000a2b54c5e016b483045022100ae3a535c09d005c0ceca3029cbf28cc45791f9710f401ee4ad4925e5163fbe0302202ed3256c2cbec121d8c1fd0a1bded5ca8e4e44f9de9d42ca421b55c3ccdf5ccf012102b3e7c21a906433171cad38589335002c34a6928e19b7798224077c30f03e835e")
}

func TestTxEncode(t *testing.T) {
	inputs := make(map[string]uint32)
	outputs := make(map[string]uint64)
	inputs["25517e3b3759365e80a164a3d4d2db2462c5d6888e4bd874c5fbfbb6fb130b41"] = 0
	outputs["Tmeyuj8ZBaQC8F47wNKxDmYAWUFti3XMrLb"] = 2083509771
	outputs["TmfTUZcZNrtvuyqfZym5LJ2sT2MN3p5WES8"] = 100000000
	timestamp, _ := time.Parse("2016-01-02 15:04:05", "2019-13-14 00:00:00")
	rs, _ := TxEncode(1, 0, &timestamp, inputs, outputs)

	fmt.Println(rs)
	assert.Equal(t, rs, "0100000001410b13fbb6fbfbc574d84b8e88d6c56224dbd2d4a364a1805e3659373b7e512500000000ffffffff020bd62f7c000000001976a914afda839fa515ffdbcbc8630b60909c64cfd73f7a88ac00e1f505000000001976a914b51127b89f9b704e7cfbc69286f0de2e00e7196988ac000000000000000000096e880100")
}

func TestNewEntropy(t *testing.T) {
	s, _ := NewEntropy(32)
	fmt.Printf("%s\n", s)
	assert.Equal(t, len(s), 64)

}

func TestEcNew(t *testing.T) {
	s, _ := EcNew("secp256k1", "7686a4df8171ebf04ede968167d0593fd4fbd8ee9feb07d453e768e06cc5e51d")
	assert.Equal(t, s, "dbae6e0b3174330ad24be8d952307e95106eb8d573defdc1f393ef2abf2e7b9c")
}

func TestEcPrivateKeyToEcPublicKey(t *testing.T) {
	s, _ := EcPrivateKeyToEcPublicKey(false, "dbae6e0b3174330ad24be8d952307e95106eb8d573defdc1f393ef2abf2e7b9c")
	assert.Equal(t, s, "02addd806e8813f85fad05b97541915eb3a1f27528d3156f2ef8166823d6722b58")
}

func TestEcPubKeyToAddress(t *testing.T) {
	s, _ := EcPubKeyToAddress("testnet", "02addd806e8813f85fad05b97541915eb3a1f27528d3156f2ef8166823d6722b58")
	assert.Equal(t, s, "TmgMiXziDuFiyLc159zagcCnmVxhReojytr")
}

func TestCreateAddress(t *testing.T) {
	s, _ := NewEntropy(32)
	k, _ := EcNew("secp256k1", s)
	p, _ := EcPrivateKeyToEcPublicKey(false, k)
	a, _ := EcPubKeyToAddress("testnet", p)
	fmt.Printf("%s\n%s\n%s\n%s\n", s, k, p, a)
	assert.Contains(t, a, "Tm")
}

func TestCreateMixParamsAddressPublicKeyHash(t *testing.T) {
	times := 0
	for {
		if times > 20000 {
			break
		}
		s, _ := NewEntropy(32)
		k, _ := EcNew("secp256k1", s)
		p, _ := EcPrivateKeyToEcPublicKey(false, k)
		a, _ := EcPubKeyToAddress("mixnet", p)
		//fmt.Printf("%s\n%s\n%s\n%s\n", s, k, p, a)
		if !assert.Contains(t, a, "Xm") {
			break
		}
		times++
	}
}

func TestCreateMixParamsSciptToHashAddress(t *testing.T) {
	times := 0
	for {
		if times > 20000 {
			break
		}
		s, _ := NewEntropy(32)
		k, _ := EcNew("secp256k1", s)
		p, _ := EcPrivateKeyToEcPublicKey(false, k)
		a, _ := EcScriptKeyToAddress("mixnet", p)
		//fmt.Printf("%s\n%s\n%s\n%s\n", s, k, p, a)
		if !assert.Contains(t, a, "Xd") {
			break
		}
		times++
	}

}

func TestCompactToTarget(t *testing.T) {
	CompactToTarget("471859199")
	// output :
	// 0x000000001fffff00000000000000000000000000000000000000000000000000
}

func TestCompactToHashrate(t *testing.T) {
	CompactToHashrate("471859199", 1)
	// output :
	// 34.35975475 GH/s
}

func TestHashrateToCompact(t *testing.T) {
	// 34.35975475 GH/s
	hashrate := 34.35975475 * 1000 * 1000 * 1000
	HashrateToCompact(fmt.Sprintf("%d", uint64(hashrate)))
	// output :
	// 471859199
}

func TestTargetToCompact(t *testing.T) {
	TargetToCompact("0x000000001fffff00000000000000000000000000000000000000000000000000")
	// output :
	// 471859199
}
