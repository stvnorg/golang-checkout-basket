package main

import (
  "github.com/gin-gonic/gin"
  "net/http/httptest"
  "testing"
)

func Test_GET_CustomHTTP404(t *testing.T) {
  w := httptest.NewRecorder()
  c, _ := gin.CreateTestContext(w)

  GET_CustomHTTP404(c)
  if w.Code != 404 {
    t.Fail()
  }
}

func Test_POST_CreateBasket(t *testing.T) {
  w := httptest.NewRecorder()
  c, _ := gin.CreateTestContext(w)

  POST_CreateBasket(c)
  if w.Code != 201 {
    t.Fail()
  }
}

func Test_POST_AddProductToBasket(t *testing.T) {
  w := httptest.NewRecorder()
  c, _ := gin.CreateTestContext(w)
  c.Params = []gin.Param{gin.Param{Key: "code", Value: "v"}}

  POST_AddProductToBasket(c)
  if w.Code != 200 {
    t.Fail()
  }
}

func Test_GET_TotalAmountInBasket(t *testing.T) {
  w := httptest.NewRecorder()
  c, _ := gin.CreateTestContext(w)

  GET_TotalAmountInBasket(c)
  if w.Code != 200 {
    t.Fail()
  }
}

func Test_DELETE_DeleteBasket(t *testing.T) {
  w := httptest.NewRecorder()
  c, _ := gin.CreateTestContext(w)

  DELETE_DeleteBasket(c)
  if w.Code != 200 {
    t.Fail()
  }
}
