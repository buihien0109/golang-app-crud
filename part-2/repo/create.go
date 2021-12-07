package repo

import (
	"fiber-postgres/model"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

/*
Tạo 1 số dữ liệu mẫu trong database
*/
func MockData() (err error) {
	// Khởi tạo transaction khi chúng ta muốn insert dữ liệu trong nhiều bảng
	transaction, err := DB.Begin()
	if err != nil {
		return err
	}

	// Insert 1 số bản ghi user
	for i := 0; i < 10; i++ {
		userId := NewID()
		_, err = transaction.Model(&model.User{
			Id:        userId,
			FullName:  gofakeit.Animal(),
			Email:     gofakeit.Email(),
			Phone:     gofakeit.Phone(),
			CreatedAt: gofakeit.Date(),
		}).Insert()

		if err != nil {
			transaction.Rollback()
			return err
		}

		// Với mỗi bản ghi user, insert 1 số bản ghi post liên quan đến user đó
		for j := 0; j < 5+rand.Intn(5); j++ {
			_, err = transaction.Model(&model.Post{
				Id:        NewID(),
				Title:     gofakeit.Quote(),
				Content:   gofakeit.LoremIpsumSentence(200),
				AuthorId:  userId,
				CreatedAt: gofakeit.Date(),
			}).Insert()

			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	return transaction.Commit()
}
