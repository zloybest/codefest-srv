package db

import (
  "codefest/model"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "os"
)

var _DB_HOST = "localhost"
var _DB_USER = ""
var _DB_PASSWORD = ""
var _DB_NAME = "test"
var _db *mgo.Database = nil

func getDB() (*mgo.Database, error) {

  if _db != nil {
    return _db, nil
  }

  _DB_HOST = os.Getenv("CODEFEST_DB_HOST")
  if _DB_HOST == "" {
    _DB_HOST = "localhost"
  }

  _DB_NAME = os.Getenv("CODEFEST_DB_NAME")
  if _DB_NAME == "" {
    _DB_NAME = "test"
  }

  _DB_USER = os.Getenv("CODEFEST_DB_USER")
  _DB_PASSWORD = os.Getenv("CODEFEST_DB_PASSWORD")

  var authString = ""
  if _DB_USER != "" {
    authString = _DB_USER
    if _DB_PASSWORD != "" {
      authString = authString + ":" + _DB_PASSWORD
    }
    authString = authString + "@"
  }
  session, err := mgo.Dial("mongodb://" + authString + _DB_HOST)
  if err != nil {
    return nil, err
  }

  _db = session.DB(_DB_NAME)
  return _db, nil
}

func col(name string) *mgo.Collection {
  hDB, err := getDB()
  if err != nil {
    panic(err)
    return nil
  }
  return hDB.C(name)
}

func CLectures() *mgo.Collection {
  return col("lectures")
}

func Lectures(query interface{}) *[]model.Lecture {
  result := &[]model.Lecture{}
  CLectures().Find(query).All(result)
  return result
}

func Lecture(query interface{}) *model.Lecture {
  result := &model.Lecture{}
  CLectures().Find(query).One(result)
  return result
}

func LectureInsert(lecture *model.Lecture) error {
  return CLectures().Insert(lecture)
}

func LectureUpdate(lectureId string, lecture *model.Lecture) error {
  return CLectures().Update(bson.M{"lectureid": lectureId}, lecture)
}