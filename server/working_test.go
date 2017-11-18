package server

import (
	"testing"

	"github.com/contribsys/faktory/client"
	"github.com/stretchr/testify/assert"
)

func TestAcknowledge(t *testing.T) {
	t.Parallel()

	set := &fakeSet{}

	job, err := acknowledge("", &fakeSet{})
	assert.NoError(t, err)
	assert.Nil(t, job)

	wid := "123876"
	job = client.NewJob("sometest", 1, 2, 3)

	err = reserve(wid, job, set)
	assert.NoError(t, err)

	jid := job.Jid
	job, err = acknowledge(jid, set)
	assert.NoError(t, err)
	assert.Equal(t, jid, job.Jid)

	job, err = acknowledge(jid, set)
	assert.NoError(t, err)
	assert.Nil(t, job)
}

type fakeSet struct {
}

func (fs *fakeSet) AddElement(ts string, jid string, data []byte) error {
	return nil
}

func (fs *fakeSet) RemoveElement(ts string, jid string) error {
	return nil
}
