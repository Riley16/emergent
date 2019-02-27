// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leabra

import "reflect"

// leabra.Synapse holds state for the synaptic connection between neurons
type Synapse struct {
	Wt     float32 `desc:"synaptic weight value -- sigmoid contrast-enhanced"`
	LWt    float32 `desc:"linear (underlying) weight value -- learns according to the lrate specified in the connection spec -- this is converted into the effective weight value, Wt, via sigmoidal contrast enhancement (see WtSigPars)"`
	DWt    float32 `desc:"change in synaptic weight, from learning"`
	Norm   float32 `desc:"dwt normalization factor -- reset to max of abs value of dwt, decays slowly down over time -- serves as an estimate of variance in weight changes over time"`
	Moment float32 `desc:"momentum -- time-integrated dwt changes, to accumulate a consistent direction of weight change and cancel out dithering contradictory changes"`
}

var SynapseVars = []string{"Wt", "LWt", "DWt", "Norm", "Moment"}

var SynapseVarsMap map[string]int

func init() {
	SynapseVarsMap = make(map[string]int, len(SynapseVars))
	for i, v := range SynapseVars {
		SynapseVarsMap[v] = i
	}
}

func (sy *Synapse) VarNames() []string {
	return SynapseVars
}

func (sy *Synapse) VarByName(varNm string) (float32, bool) {
	i, ok := SynapseVarsMap[varNm]
	if !ok {
		return 0, false
	}
	// todo: would be ideal to avoid having to use reflect here..
	v := reflect.ValueOf(sy)
	return v.Elem().Field(i).Interface().(float32), true
}

func (sy *Synapse) SetVarByName(varNm string, val float64) bool {
	i, ok := SynapseVarsMap[varNm]
	if !ok {
		return false
	}
	// todo: would be ideal to avoid having to use reflect here..
	v := reflect.ValueOf(sy)
	v.Elem().Field(i).SetFloat(val)
	return true
}
