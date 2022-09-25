package services

import (
	"test_go_project/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFindByText_Message(t *testing.T) {
	service := NewMessageService(testDB)
	input := models.NewMessage("testId")
	input.Text = "test FindByMessage"
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.FindByText(testCtx, input.Text)
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, docs[0])
	t.Log("[TestFindByText] Complete. result:", docs[0])
}
func TestFindByMessageId_Message(t *testing.T) {
	service := NewMessageService(testDB)
	input := models.NewMessage("testId0")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.FindByMessageId(testCtx, input.MessageId)
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, docs[0])
	t.Log("[TestFindByMessageId] Complete. result:", docs[0])
}
func TestFindByUserId_Message(t *testing.T) {
	service := NewMessageService(testDB)
	input := models.NewMessage("testId1")
	input.UserId = "test UserId"
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.FindByUserId(testCtx, input.UserId)
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, docs[0])
	t.Log("[TestFindByUserId] Complete. result:", docs[0])
}

func TestCreate_Message(t *testing.T) {
	service := NewMessageService(testDB)
	input := models.NewMessage("testId2")
	doc, err := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, doc.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, input, doc)
	t.Log("[TestCreate] Complete. result:", doc)
}
func TestFindById_Message(t *testing.T) {
	service := NewMessageService(testDB)
	input := models.NewMessage("testId3")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	doc, err := service.FindById(testCtx, expectDoc.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, doc)
	t.Log("[TestFindById] Complete. result:", doc)
}
func TestFindOne_Message(t *testing.T) {
	service := NewMessageService(testDB)
	input := models.NewMessage("testId4")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	doc, err := service.FindOne(testCtx, bson.M{"_id": expectDoc.Id})
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, doc)
	t.Log("[TestFindOne] Complete. result:", doc)
}
func TestFind_Message(t *testing.T) {
	service := NewMessageService(testDB)
	input := models.NewMessage("testId5")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.Find(testCtx, bson.M{"createdAt": input.CreateAt})
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, docs[0])
	t.Log("[TestFindOne] Complete. result:", docs)
}
func TestFindAll_Message(t *testing.T) {
	service := NewMessageService(testDB)
	input := models.NewMessage("testId6")
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.FindAll(testCtx)
	t.Log("[TestFindAll] docs length =", len(docs))
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(docs), 1)
	t.Log("[TestFindAll] Complete. result:", docs)
}
func TestUpdateById_Message(t *testing.T) {
	service := NewMessageService(testDB)
	createInput := models.NewMessage("testId7")
	createdDoc, _ := service.Create(testCtx, createInput)
	defer service.DeleteById(testCtx, createdDoc.Id.Hex())
	input := &models.Message{
		Text: "test UpdateById::After",
	}
	expectedDoc := createdDoc
	expectedDoc.Text = input.Text
	doc, err := service.UpdateById(testCtx, createdDoc.Id.Hex(), input)
	assert.NoError(t, err)
	assert.Equal(t, expectedDoc, doc)
	t.Log("[TestUpdateById] Complete. result:", doc)
}
func TestDeleteById_Message(t *testing.T) {
	service := NewMessageService(testDB)
	input := models.NewMessage("testId8")
	doc, _ := service.Create(testCtx, input)
	err := service.DeleteById(testCtx, doc.Id.Hex())
	assert.NoError(t, err)
	founded, _ := service.FindById(testCtx, doc.Id.Hex())
	assert.Nil(t, founded)
	t.Log("[TestDeleteById] Complete.")
}
