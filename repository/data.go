package repository

import (
	"log"
	"time"
	"math/rand"

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

// TODO: contributionCount:int -> data:obj
func (d *Data) Save(contributionCount int) {
	w := model.Data{
		ID: rand.Intn(10000),
		UserID: RandomString(rand.Intn(100)),
		Name: RandomString(rand.Intn(100)),
		GitHubTodaysContributionCount: contributionCount,
		CreatedAt: time.Now(),
	}
	err := d.Table.Put(w).Run()
	if err != nil {
		log.Fatalf("%v", err)
	}
}