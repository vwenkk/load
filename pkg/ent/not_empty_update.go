// Code generated by ent, DO NOT EDIT.

package ent

// set field if value is not empty. e.g. string does not equal to ""
func (gr *GroupUpdate) SetNotEmptyStatus(value uint8) *GroupUpdate {
	if value != 0 {
		return gr.SetStatus(value)
	}
	return gr
}

// set field if value is not empty. e.g. string does not equal to ""
func (gr *GroupUpdateOne) SetNotEmptyStatus(value uint8) *GroupUpdateOne {
	if value != 0 {
		return gr.SetStatus(value)
	}
	return gr
}

// set field if value is not empty. e.g. string does not equal to ""
func (gr *GroupUpdate) SetNotEmptyName(value string) *GroupUpdate {
	if value != "" {
		return gr.SetName(value)
	}
	return gr
}

// set field if value is not empty. e.g. string does not equal to ""
func (gr *GroupUpdateOne) SetNotEmptyName(value string) *GroupUpdateOne {
	if value != "" {
		return gr.SetName(value)
	}
	return gr
}

// set field if value is not empty. e.g. string does not equal to ""
func (gr *GroupUpdate) SetNotEmptyRemark(value string) *GroupUpdate {
	if value != "" {
		return gr.SetRemark(value)
	}
	return gr
}

// set field if value is not empty. e.g. string does not equal to ""
func (gr *GroupUpdateOne) SetNotEmptyRemark(value string) *GroupUpdateOne {
	if value != "" {
		return gr.SetRemark(value)
	}
	return gr
}
