package repository

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"math/rand"
	"time"

	"github.com/guregu/dynamo"
	"github.com/ia17011/hollein/model"
)

type Data struct {
	Table dynamo.Table
}

// NOTE: for test
func RandomString(n int) string {
	var letter = []rune("abcdfghijkmwxyzABCPQRSTWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func (d *Data) Save(userName string, contributionCount int) {
	HashedUserName := md5.Sum([]byte(userName))

	// TODO: どんどんデータを増やしていく
	w := model.Data{
		UserID: hex.EncodeToString(HashedUserName[:]),
		CreatedAt: time.Now().Unix(),
		Name: userName,
		GitHubTodaysContributionCount: contributionCount,
	}
	err := d.Table.Put(w).Run()
	if err != nil {
		log.Fatalf("%v", err)
	}
}