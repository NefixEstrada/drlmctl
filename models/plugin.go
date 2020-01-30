// SPDX-License-Identifier: AGPL-3.0-only

package models

// Plugin is an individual DRLM plugin. It's read from the plugin.json
type Plugin struct {
	Repo        string `json:"-"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	License     string `json:"license,omitempty"`
}
