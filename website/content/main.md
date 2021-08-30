---
weight: 1
title: "Main"
draft: false
---

## Installation

### With Go

```bash
$ go get github.com/chutommy/prondru
```

### From source

```bash
$ git clone https://github.com/chutommy/prondru.git
```

The binary files are located in bin directory:

#### Linux and MacOS

```bash
$ ./prondru/bin/prondru
```

#### Windows

```shell
> prondru\bin\prondru.exe
```

## Usage

### Help

Run `prondru` with `--help` or `-h` flag:

```bash
$ prondru --help
Usage of prondru:
  -ansi
    	ansi colors enabled? (default true)
  -banner string
    	banner.txt file (default "banner.txt")
  -show-banner
    	print the banner? (default true)
```

### Run

```bash
$ prondru
--------------------------------------------------------------
  ██████╗ ██████╗  ██████╗ ███╗   ██╗██████╗ ██████╗ ██╗   ██╗
  ██╔══██╗██╔══██╗██╔═══██╗████╗  ██║██╔══██╗██╔══██╗██║   ██║
  ██████╔╝██████╔╝██║   ██║██╔██╗ ██║██║  ██║██████╔╝██║   ██║
  ██╔═══╝ ██╔══██╗██║   ██║██║╚██╗██║██║  ██║██╔══██╗██║   ██║
  ██║     ██║  ██║╚██████╔╝██║ ╚████║██████╔╝██║  ██║╚██████╔╝
  ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝ ╚═════╝
--------------------------------------------------------------
   CLI for Office of Scientific and Technical Information API
               https://www.osti.gov/api/v1/docs
--------------------------------------------------------------
                 Welcome to the prOndru v18.0!
                            ---
          Ondra's terminal gateway to the fantastic
               world of physical discoveries.
                 ~ Wednesday, 20 Jan 2021
  
         Copyright (c) 2021 Tommy Chu & Ondrej Skrna
               Licensed under the MIT license
--------------------------------------------------------------
Discover over 70 years of research results
from 3 million scientific records
provided by OSTI.
++++++++++++++++
✗ Keyword: █
```

### Discover

Type a keyword (at least 4 characters):

```bash
✔ Keyword: Newton
```

Use arrows to navigate:

```bash
✔ Keyword: Newton
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select article: 
    Optimization and supervised machine learning methods for fitting numerical physics models without derivatives
    Nonlinear convergence in contact mechanics: Immersed boundary finite volume
    Oldroyd's model and the foundation of modern rheology of yield stress fluids
  ▸ Levenberg–Marquardt multi-classification using hinge loss function
    Hydroxide melt induced corrosion of Ni at elevated temperatures under steam electrolysis conditions
    Recent progress on neoclassical impurity transport in stellarators with implications for a stellarator reactor
    Structural stability and artificial buckling modes in topology optimization
    Heat transfer and pseudo phase transition for low-Reynolds, mixed-convection channel flow in the supercritical thermodynamic regime
    de Sitter instanton from Euclidean dynamical triangulations
↓   Flight directions of songbirds are unaffected by the topography of Lake Erie’s southern coastline during fall migration
```

Use enter to discover articles:

```bash
✔ Keyword: Newton█
✔ Information transfer with a gravitating bath

================
Information transfer with a gravitating bath [OSTI: 1782949]
Authors: Geng, Hao [University of Washington], Karch, Andreas [The University of Texas at Austin, University of Washington],
Perez-Pardavila, Carlos [The University of Texas at Austin], Raju, Suvrat [Tata Institute of Fundamental Research], Randall,
Lisa [Harvard University], Riojas, Marcos [The University of Texas at Austin], Shashi, Sanjit [The University of Texas at Austin]
Publisher: Stichting SciPost - Netherlands
Subjects: 
Date: 2021-05-14T00:00:00Z (https://doi.org/10.21468/SciPostPhys.10.5.103)
----------------
Research

Late-time dominance of entanglement islands plays a critical role in addressing
the information paradox for black holes in AdS coupled to an asymptotic
non-gravitational bath. A natural question is how this observation can be
extended to gravitational systems. To gain insight into this question, we
explore how this story is modified within the context of Karch-Randall
braneworlds when we allow the asymptotic bath to couple to dynamical gravity.
We find that because of the inability to separate degrees of freedom by spatial
location when defining the radiation region, the entanglement entropy of
radiation emitted into the bath is a time-independent constant, consistent with
recent work on black hole information in asymptotically flat space. If we
instead consider an entanglement entropy between two sectors of a specific
division of the Hilbert space, we then find non-trivial time-dependence, with
the Page time a monotonically decreasing function of the brane angle---provided
both branes are below a particular angle. However, the properties of the entropy
depend discontinuously on this angle, which is the first example of such
discontinuous behavior for an AdS brane in AdS space.</p>

Press enter to continue...
```

## License

The project is under the MIT open source software license.
