package show

import (
	"strings"
)

// Banner returns show's banner string.
func Banner() string {
	return strings.TrimSpace(bannerString)
}

// bannerString defines the banner string.
const bannerString = `
------------------------------------------------------------
██████╗ ██████╗  ██████╗ ███╗   ██╗██████╗ ██████╗ ██╗   ██╗
██╔══██╗██╔══██╗██╔═══██╗████╗  ██║██╔══██╗██╔══██╗██║   ██║
██████╔╝██████╔╝██║   ██║██╔██╗ ██║██║  ██║██████╔╝██║   ██║
██╔═══╝ ██╔══██╗██║   ██║██║╚██╗██║██║  ██║██╔══██╗██║   ██║
██║     ██║  ██║╚██████╔╝██║ ╚████║██████╔╝██║  ██║╚██████╔╝
╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝ ╚═════╝
------------------------------------------------------------
 CLI for Office of Scientific and Technical Information API
             https://www.osti.gov/api/v1/docs
------------------------------------------------------------
               Welcome to the prOndru v18.0!
                          ---
        Ondra's terminal gateway to the fantastic
             world of physical discoveries.
               ~ Wednesday, 20 Jan 2021

       Copyright (c) 2021 Tommy Chu & Ondrej Skrna
             Licensed under the MIT license
------------------------------------------------------------
Discover over 70 years of research results
from 3 million scientific records
provided by OSTI.
`
