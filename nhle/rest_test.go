package nhle_test

import (
	"github.com/StephenGriese/hello-api/nhle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUsers(t *testing.T) {
	ps := nhle.NewPlayerService()
	players, err := ps.Players()
	assert.NoError(t, err)
	assert.True(t, len(players) > 0)
}
