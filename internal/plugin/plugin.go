/*
Package plugin export interfaces to be implemented by application plugins.

The Pluggable interface must be implemented by all plugins.

Other interfaces are optionals and guarded by type assertions. If a plugin type
implements a specific interface that plugin will gain specified functionalities.

plugin/manager is responsible for leveraging plugins and their functionalities.
*/
package plugin
