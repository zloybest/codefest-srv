package controllers

import (
  "github.com/labstack/echo"
  "codefest/db"
  "codefest/model"
  "gopkg.in/mgo.v2/bson"
)

func Likes(c echo.Context) error {
  return c.JSON(OK, db.Lectures(nil))
}

func Like(c echo.Context) error {
  likes := new([]model.Like)
  c.Bind(likes)
  copyLikes := *likes
  var lectures []*model.Lecture
  for i := range copyLikes {
    like := copyLikes[i]
    lecture := db.Lecture(bson.M{"lectureid": like.LectureId})
    if lecture.LectureId == "" {
      lecture = &model.Lecture{
        LectureId: like.LectureId,
      }
      db.LectureInsert(lecture)
    }
    var likeIndex, dislikeIndex int = -1, -1
    for l, val := range lecture.Likes {
      if val == like.Uid {
        likeIndex = l
      }
    }

    for l, val := range lecture.Dislikes {
      if val == like.Uid {
        dislikeIndex = l
      }
    }

    if like.Type == "like" {
      if likeIndex == -1 && dislikeIndex == -1 {
        lecture.Likes = append(lecture.Likes, like.Uid)
      }
      if likeIndex != -1 {
        lecture.Likes = append(lecture.Likes[:likeIndex], lecture.Likes[likeIndex+1:]...)
      }
      if likeIndex == -1 && dislikeIndex != -1 {
        lecture.Dislikes = append(lecture.Dislikes[:dislikeIndex], lecture.Dislikes[dislikeIndex+1:]...)
        lecture.Likes = append(lecture.Likes, like.Uid)
      }
    }

    if like.Type == "dislike" {
      if likeIndex != -1 {
        lecture.Likes = append(lecture.Likes[:likeIndex], lecture.Likes[likeIndex+1:]...)
        lecture.Dislikes = append(lecture.Dislikes, like.Uid)
      }
      if likeIndex == -1 && dislikeIndex != -1 {
        lecture.Dislikes = append(lecture.Dislikes[:dislikeIndex], lecture.Dislikes[dislikeIndex+1:]...)
      }
      if likeIndex == -1 && dislikeIndex == -1 {
        lecture.Dislikes = append(lecture.Dislikes, like.Uid)
      }
    }

    db.LectureUpdate(lecture.LectureId, lecture)

    lectures = append(lectures, lecture)
  }
  return c.JSON(OK, lectures)
}
