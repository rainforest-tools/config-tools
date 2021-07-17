package tests

import (
	"encoding/json"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/rainforest-tools/config-tools/models"
)

func TestJob(t *testing.T) {
	raw := models.Job{Name: "test", Partition: "v100-32g", GPU: 2, Modules: []string{"opt", "gcc", "cuda/10"}, Command: "python --version"}
	bytes, _ := json.Marshal(raw)
	var job models.Job
	json.Unmarshal(bytes, &job)
	assert.Equal(t, job, raw)
}
