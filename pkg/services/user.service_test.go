package services

import (
	"test_go_project/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFindByUserId(t *testing.T) {
	service := NewUserService(testDB)
	input := models.NewUser("test FindByUserId")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.FindByUserId(testCtx, input.UserId)
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, docs[0])
	t.Log("[TestFindOne] Complete. result:", docs[0])
}
func TestFindByName(t *testing.T) {
	service := NewUserService(testDB)
	input := models.NewUser("test FindByName")
	input.Name = "test Name"
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.FindByName(testCtx, input.Name)
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, docs[0])
	t.Log("[TestFindOne] Complete. result:", docs[0])
}
func TestCreate_User(t *testing.T) {
	service := NewUserService(testDB)
	input := models.NewUser("test Create")
	doc, err := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, doc.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, input, doc)
	t.Log("[TestCreate] Complete. result:", doc)
}
func TestFindById_User(t *testing.T) {
	service := NewUserService(testDB)
	input := models.NewUser("test FindById")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	doc, err := service.FindById(testCtx, expectDoc.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, doc)
	t.Log("[TestFindById] Complete. result:", doc)
}
func TestFindOne_User(t *testing.T) {
	service := NewUserService(testDB)
	input := models.NewUser("test FindOne")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	doc, err := service.FindOne(testCtx, bson.M{"_id": expectDoc.Id})
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, doc)
	t.Log("[TestFindOne] Complete. result:", doc)
}
func TestFind_User(t *testing.T) {
	service := NewUserService(testDB)
	input := models.NewUser("test Find")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.Find(testCtx, bson.M{"createdAt": input.CreatedAt})
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, docs[0])
	t.Log("[TestFindOne] Complete. result:", docs)
}
func TestFindAll_User(t *testing.T) {
	service := NewUserService(testDB)
	input := models.NewUser("test FindAll")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.FindAll(testCtx)
	t.Log("[TestFindAll] docs length =", len(docs))
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(docs), 1)
	t.Log("[TestFindAll] Complete. result:", docs)
}
func TestUpdateById_User(t *testing.T) {
	service := NewUserService(testDB)
	createInput := models.NewUser("test UpdateById::Before")
	createdDoc, _ := service.Create(testCtx, createInput)
	defer service.DeleteById(testCtx, createdDoc.Id.Hex())
	input := &models.User{
		UserId: "test UpdateById::After",
		Name:   "test UpdateById::After",
	}
	expectedDoc := createdDoc
	expectedDoc.UserId = input.UserId
	expectedDoc.Name = input.Name
	doc, err := service.UpdateById(testCtx, createdDoc.Id.Hex(), input)
	assert.NoError(t, err)
	assert.Equal(t, expectedDoc, doc)
	t.Log("[TestUpdateById] Complete. result:", doc)
}
func TestDeleteById_User(t *testing.T) {
	service := NewUserService(testDB)
	input := models.NewUser("test DeleteById")
	doc, _ := service.Create(testCtx, input)
	err := service.DeleteById(testCtx, doc.Id.Hex())
	assert.NoError(t, err)
	founded, _ := service.FindById(testCtx, doc.Id.Hex())
	assert.Nil(t, founded)
	t.Log("[TestDeleteById] Complete.")
}
