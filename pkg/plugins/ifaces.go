package plugins

// Manager is the plugin manager service interface.
type Manager interface {
	// Renderer gets the renderer plugin.
	Renderer() *RendererPlugin
	// GetDataSource gets a data source plugin with a certain ID.
	GetDataSource(id string) *DataSourcePlugin
	// GetDataPlugin gets a data plugin with a certain ID.
	GetDataPlugin(id string) DataPlugin
	// GetPlugin gets a plugin with a certain ID.
	GetPlugin(id string) *PluginBase
	// NumDataSources gets the number of data sources.
	NumDataSources() int
}
