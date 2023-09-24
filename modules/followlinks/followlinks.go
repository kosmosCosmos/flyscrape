// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package followlinks

import (
	"github.com/philippta/flyscrape"
)

func init() {
	flyscrape.RegisterModule(new(Module))
}

type Module struct{}

func (m *Module) OnResponse(resp *flyscrape.Response) {
	for _, link := range flyscrape.ParseLinks(string(resp.Body), resp.Request.URL) {
		resp.Visit(link)
	}
}

var _ flyscrape.OnResponse = (*Module)(nil)
