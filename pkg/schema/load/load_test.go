package load

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadBaseDashboard(t *testing.T) {
	currentpath, _ := os.Getwd()
	loadpaths := &BaseLoadPaths{
		BaseCueFS:       currentpath,
		DistPluginCueFS: currentpath,
		InstanceCueFS:   currentpath,
	}

	t.Run("Test basics of load base dashboard", func(t *testing.T) {
		_, err := BaseDashboard(*loadpaths)
		require.NoError(t, err)
	})
}
