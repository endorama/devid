package plugin

// PluggableInstantiator is a proxy type for the init function for a plugin.
type PluggableInstantiator func() Pluggable
