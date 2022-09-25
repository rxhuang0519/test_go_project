package services

import (
	"context"
	"test_go_project/pkg/models"
	"test_go_project/pkg/repository"
	"test_go_project/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var testDB *mongo.Database

var testCtx = context.Background()

/* Only for testing */
func newBaseService(db *mongo.Database) *Service[models.Base] {
	return &Service[models.Base]{collection: db.Collection("base")}
}

func TestMain(m *testing.M) {
	testDB = utils.SetupTestDB("", "test")
	defer repository.Disconnect(testDB.Client())
	m.Run()

}
func TestCreate(t *testing.T) {
	service := newBaseService(testDB)
	input := models.NewBase()
	doc, err := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, doc.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, input, doc)
	t.Log("[TestCreate] Complete. result:", doc)
}
func TestFindById(t *testing.T) {
	service := newBaseService(testDB)
	input := models.NewBase()
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	doc, err := service.FindById(testCtx, expectDoc.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, doc)
	t.Log("[TestFindById] Complete. result:", doc)
}
func TestFindOne(t *testing.T) {
	service := newBaseService(testDB)
	// service := &Service[models.Base]{collection: db.Collection("base")}
	input := models.NewBase()
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	doc, err := service.FindOne(testCtx, bson.M{"_id": expectDoc.Id})
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, doc)
	t.Log("[TestFindOne] Complete. result:", doc)
}
func TestFind(t *testing.T) {
	service := newBaseService(testDB)
	input := models.NewBase()
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.Find(testCtx, bson.M{"createdAt": input.CreateAt})
	assert.NoError(t, err)
	assert.Equal(t, expectDoc, docs[0])
	t.Log("[TestFindOne] Complete. result:", docs)
}
func TestFindAll(t *testing.T) {
	service := newBaseService(testDB)
	input := models.NewBase()
	expectDoc, _ := service.Create(testCtx, input)
	defer service.DeleteById(testCtx, expectDoc.Id.Hex())
	docs, err := service.FindAll(testCtx)
	t.Log("[TestFindAll] docs length =", len(docs))
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(docs), 1)
	t.Log("[TestFindAll] Complete. result:", docs)
}
func TestUpdateById(t *testing.T) {
	service := newBaseService(testDB)
	createInput := models.NewBase()
	createdDoc, _ := service.Create(testCtx, createInput)
	defer service.DeleteById(testCtx, createdDoc.Id.Hex())
	input := &models.Base{
		CreateAt: time.Now().Truncate(time.Millisecond).AddDate(0, 0, 1).UTC(),
		UpdateAt: time.Now().Truncate(time.Millisecond).AddDate(0, 0, 1).UTC(),
	}
	expectedDoc := createdDoc
	expectedDoc.CreateAt = input.CreateAt
	expectedDoc.UpdateAt = input.UpdateAt
	doc, err := service.UpdateById(testCtx, createdDoc.Id.Hex(), input)
	assert.NoError(t, err)
	assert.Equal(t, expectedDoc, doc)
	t.Log("[TestUpdateById] Complete. result:", doc)
}
func TestDeleteById(t *testing.T) {
	service := newBaseService(testDB)
	input := models.NewBase()
	doc, _ := service.Create(testCtx, input)
	err := service.DeleteById(testCtx, doc.Id.Hex())
	assert.NoError(t, err)
	founded, _ := service.FindById(testCtx, doc.Id.Hex())
	assert.Nil(t, founded)
	t.Log("[TestDeleteById] Complete.")
}
