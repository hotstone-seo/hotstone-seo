package dependency

// Autogenerated by Typical-Go. DO NOT EDIT.

import "github.com/hotstone-seo/hotstone-server/typical"

func init() {
	typical.Context.TestTargets.Append("./app")
	typical.Context.TestTargets.Append("./app/config")
	typical.Context.TestTargets.Append("./app/repository")
	typical.Context.TestTargets.Append("./app/service")
}
