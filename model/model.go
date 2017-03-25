package model

type Lecture struct {
  LectureId string `json:"lectureId" form:"lectureId" query:"lectureId"`
  Likes []string `json:"likes" form:"likes" query:"likes"`
  Dislikes []string `json:"dislikes" form:"dislikes" query:"dislikes"`
}

type Like struct {
  LectureId string `json:"lectureId"`
  Uid       string `json:"uid"`
  Type      string `json:"type"`
}