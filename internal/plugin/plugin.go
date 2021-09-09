/*
Package plugin export interfaces to be implemented by application plugins.

The Pluggable interface must be implemented by all plugins.

Other interfaces are optionals and guarded by type assertions. If a plugin type
implements a specific interface that plugin will gain specified functionalities.

plugin/manager is responsible for leveraging plugins and their functionalities.

This is not a third-party plugin system where you can load plugins from
different sources.
The goal is to allow an easy integration for different functionalities while
retaining control on code and execution behaviour. Using interfaces allow
for plugin modularity and provide a standard integration point.
devid is aimed to be a tool to easen and help security posture. Loading
arbitrary unverified code from unknown source is not appealing at all in this
context.
For the time being, all plugins will need to be added manually to the "plugin
registry" in plugin/manager package.
*/
package plugin
