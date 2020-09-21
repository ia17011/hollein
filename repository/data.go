package repository

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"time"

	"github.com/guregu/dynamo"
	"github.com/ia17011/hollein/model"
)

type Data struct {
	Table dynamo.Table
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