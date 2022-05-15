/*
Copyright © 2022 xiexianbin
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import beego "github.com/beego/beego/v2/server/web"

type BaseController struct {
	beego.Controller
}

// ServeJSON sends a json response with encoding charset.
func (c *BaseController) ServeJSON(encoding ...bool) error {
	var (
		hasEncoding = len(encoding) > 0 && encoding[0]
	)

	return c.Ctx.Output.JSON(c.Data["json"], true, hasEncoding)
}
