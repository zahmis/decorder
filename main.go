package main

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// UUIDの文字列表現
	originalUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("uuid 生成失敗", err)
		return
	}

	// UUIDを16バイトのバイナリ表現として取得
	uuidBytes := originalUUID[:]

	// Base64エンコード
	encoded := base64.RawURLEncoding.EncodeToString(uuidBytes)

	// Base64デコード
	decoded, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("デコード失敗", err)
		return
	}

	// バイナリ表現を文字列表現に変換
	decodedUUID, err := uuid.FromBytes(decoded)
	if err != nil {
		fmt.Println("文字列置換失敗", err)
		return
	}

	// 結果
	fmt.Println("オリジナル:", originalUUID)
	fmt.Println("エンコード:", encoded)
	fmt.Println("デコード:", decodedUUID.String())
}
