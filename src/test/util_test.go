package test

import (
	"testing"
)
import util "util"
import "github.com/stretchr/testify/assert"

func TestAddElement(t *testing.T)  {
	assert := assert.New(t)
	util.Size = 10
	util.AddElement("orgId", "12345")
	assert.Equal(1, len(util.GetAllElement()), "map doesn't contain added element")
	util.AddElement("OrgId", "54321")
	assert.Equal(2, len(util.GetAllElement()), "map doesn't contain added element")
}

func TestDeleteElement(t *testing.T)  {
	assert := assert.New(t)
	util.Size = 10
	util.DeleteElement("orgId")
	assert.Equal(1, len(util.GetAllElement()), "map doesn't contain added element")
}

func TestUpdateElement(t *testing.T)  {
	assert := assert.New(t)
	util.Size = 10
	util.AddElement("orgId", "12345")
	util.UpdateElement("orgId","54321")
	assert.Equal(2, len(util.GetAllElement()), "map doesn't contain added element")
	assert.Equal("54321", util.GetElement("orgId"), "key doesn't exist in the map")
}

func TestGetAllElement(t *testing.T)  {
	assert := assert.New(t)
	util.Size = 10
	util.AddElement("orgId", "12345")
	util.AddElement("orgId2", "12345")
	util.AddElement("orgId3", "12345")
	util.AddElement("orgId4", "12345")
	assert.Equal(5, len(util.GetAllElement()), "map doesn't contain added element")
}

func TestAddElementError(t *testing.T)  {
	assert := assert.New(t)
	util.Size = 5
	util.AddElement("orgId", "12345")
	util.AddElement("orgId2", "12345")
	util.AddElement("orgId3", "12345")
	util.AddElement("orgId4", "12345")
	util.AddElement("orgId5", "12345")
	error := util.AddElement("orgId6", "12345")
	assert.Equal("You have reached the max size of the cache, please remove unwanted keys & try again", error.Error())
}