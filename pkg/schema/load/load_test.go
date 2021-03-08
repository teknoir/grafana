package load

import (
	"fmt"
	"os"
	"testing"

	"cuelang.org/go/cue"
	"github.com/grafana/grafana/pkg/schema"
	"github.com/stretchr/testify/require"
)

func TestLoadBaseDashboard(t *testing.T) {
	currentpath, _ := os.Getwd()
	loadpaths := &BaseLoadPaths{
		BaseCueFS:       currentpath,
		DistPluginCueFS: currentpath,
		InstanceCueFS:   currentpath,
	}
	t.Run("Test lookup dashboardFamily with success", func(t *testing.T) {
		mockBuildDashboardFamily()
		t.Cleanup(resetBuildDashboardFamily)
		loadpaths.packageName = "grafanaschema"
		_, err := BaseDashboard(*loadpaths)
		require.EqualError(t, err, "dashboardFamily found but build go object failed")
	})
	t.Run("Test create BuildDashboardFamily with success", func(t *testing.T) {
		loadpaths.packageName = "grafanaschema"
		_, err := BaseDashboard(*loadpaths)
		require.EqualError(t, err, "dashboardFamily found but build go object failed")
	})
}

func mockBuildDashboardFamily() {
	buildFamilyFunc = func(fam *schema.Family, famval cue.Value) (*schema.Family, error) {
		return fam, fmt.Errorf("dashboardFamily found but build go object failed")
	}
}

func resetBuildDashboardFamily() {
	buildFamilyFunc = BuildDashboardFamily
}
