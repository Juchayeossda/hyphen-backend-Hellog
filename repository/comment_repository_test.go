package repository

import (
	"context"
	"fmt"
	"hyphen-backend-hellog/initializer"
	"testing"
)

func TestXxx(t *testing.T) {
	component := initializer.NewCompnent("../.env")
	db := initializer.NewDatabse(component)

	userRepo := NewUserRepository(db)
	// postRepo := NewPostRepository(db)
	// commentRepo := NewCommentRepository(db)
	// fmt.Println(commentRepo)

	u, err := userRepo.UpdateByID(context.Background(), 1, 22, "fdsfsfsf", "fdsfsfsfs")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(u)

	err = userRepo.DeleteByID(context.Background(), 1)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	// // user, err := userRepo.Create(context.Background(), 1, "JunBeomHan", "exampleImage0")
	// // if err != nil {
	// // 	fmt.Println("Error: ", err)
	// // }

	// user, err := userRepo.SelectByID(context.Background(), 1)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// // post, err := postRepo.Create(context.Background(), "what is golang?", "good", "exampleimage0", false, user)
	// // if err != nil {
	// // 	fmt.Println("Error: ", err)
	// // }

	// post, err := postRepo.SelectByID(context.Background(), 1)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// // _, err = commentRepo.Create(context.Background(), "content1", post, user, nil)
	// // if err != nil {
	// // 	fmt.Println("Error: ", err)
	// // }

	// comment, err := commentRepo.SelectByID(context.Background(), 1)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// // fmt.Println(comment)

	// _, err = commentRepo.Create(context.Background(), "hi", post, user, comment)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// _, err = commentRepo.Create(context.Background(), "two", post, user, comment)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// comments, err := commentRepo.SelectChilds(context.Background(), comment)
	// fmt.Println(comments)

	// commentss, err := commentRepo.SelectParentsByPost(context.Background(), post)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("\n\n\n", commentss)

	// // fmt.Println(comment)

}
