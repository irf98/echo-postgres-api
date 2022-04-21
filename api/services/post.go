package services

import (
	"awesomeProject/api/models"
	"awesomeProject/db"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

func CreatePost(post *models.Post, user models.User, b string, t string) error {
	post.ID = xid.New().String()
	post.Author = user
	post.Body = b
	post.Tag = t
	post.UpVotes, post.DownVotes, post.Reputation = 0, 0, 0

	res := db.DB().Create(post)

	return res.Error
}

func GetPostById(post *models.Post, id string) error {
	res := db.DB().Raw("SELECT * FROM posts WHERE id = ?;", id).Scan(post)

	return res.Error
}

func GetUserPosts(posts *[]models.Post, id string) error {
	res := db.DB().Where("fk_user = ?", id).Find(posts)

	return res.Error
}

func GetPostsByTag(posts *[]models.Post, t string) error {
	res := db.DB().Where("tag = ?", t).Find(posts)

	return res.Error
}

func UpVotePost(post *models.Post, id string) error {
	res := db.DB().Raw("UPDATE posts SET up_votes = ?, reputation = ? WHERE id = ?;", gorm.Expr("up_votes + 1"), gorm.Expr("up_votes - down_votes"), id).Scan(post)

	return res.Error
}

func DownVotePost(post *models.Post, id string) error {
	res := db.DB().Raw("UPDATE posts SET down_votes = ?, reputation = ? WHERE id = ?;", gorm.Expr("down_votes + 1"), gorm.Expr("up_votes - down_votes"), id).Scan(post)

	return res.Error
}

func DeletePost(post *models.Post, id string) error {
	res := db.DB().Raw("DELETE FROM posts WHERE id = ?;", id).Scan(post)

	return res.Error
}
