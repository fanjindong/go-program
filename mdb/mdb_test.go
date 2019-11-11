package mdb_test

import (
	"fmt"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMdb(t *testing.T) {
	uid := "1234560"
	uidStr := fmt.Sprintf("%024s", uid)
	uidObj, err := primitive.ObjectIDFromHex(uidStr)
	t.Log(uid, uidStr, uidObj, err)
	t.Log(uidObj.Hex())
	result := strings.TrimPrefix(uidObj.Hex(), "0")
	t.Log(result)
}
