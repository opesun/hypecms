// Package skeleton implements a minimalistic but idiomatic plugin for HypeCMS.
package skeleton

import (
	"github.com/opesun/hypecms/api/context"
	"github.com/opesun/jsonp"
	"github.com/opesun/routep"
	"launchpad.net/mgo/bson"
)

// Create a type only to spare ourselves from typing map[string]interface{} every time.
type m map[string]interface{}

// mod.GetHook accesses certain functions dynamically trough this.
var Hooks = map[string]func(*context.Uni){
	"Front":     Front,
	"Back":      Back,
	"Install":   Install,
	"Uninstall": Uninstall,
	"Test":      Test,
	"AD":        AD,
}

// main.runFrontHooks invokes this trough mod.GetHook.
func Front(uni *context.Uni) {
	if _, err := routep.Comp("/skeleton", uni.P); err == nil {
		uni.Dat["_hijacked"] = true		// This is important, this stops the main front loop from executing any further modules.
		uni.Put("Hello, this is the skeleton module here.")		// This is just a basic output to allow you to see your freshly written module.
	}
	// You can insert code here which will decide wich view to call.
}

// main.runBackHooks invokes this trough mod.GetHook.
func Back(uni *context.Uni) {
	action := uni.Dat["_action"].(string)
	switch action {
	// You can dispatch your background operations here.
	}
}

// main.runDebug invokes this trough mod.GetHook.
func Test(uni *context.Uni) {
	res := map[string]interface{}{}
	res["Front"] = jsonp.HasVal(uni.Opt, "Hooks.Front", "skeleton")
	_, ok := jsonp.Get(uni.Opt, "Modules.Skeleton")
	res["Modules"] = ok
	uni.Dat["_cont"] = res
}

// admin.AD invokes this trough mod.GetHook.
func AD(uni *context.Uni) {
	uni.Dat["_points"] = []string{"skeleton/index"}
	// You can dispatch your different admin views here, based on url structure.
	
}

// admin.Install invokes this trough mod.GetHook.
func Install(uni *context.Uni) {
	id := uni.Dat["_option_id"].(bson.ObjectId)
	skeleton_options := m{
		"example": "any value",
	}
	uni.Db.C("options").Update(m{"_id": id}, m{"$addToSet": m{"Hooks.Front": "skeleton"}, "$set": m{"Modules.skeleton": skeleton_options}})
}

// Admin Install invokes this trough mod.GetHook.
func Uninstall(uni *context.Uni) {
	id := uni.Dat["_option_id"].(bson.ObjectId)
	uni.Db.C("options").Update(m{"_id": id}, m{"$pull": m{"Hooks.Front": "skeleton"}, "$unset": m{"Modules.skeleton": 1}})
}
