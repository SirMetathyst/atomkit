// Package atomkit ...
// Generated by the atom tool.  DO NOT EDIT!
// Source: atom_LocalRotation2
package atomkit

import (
	"github.com/SirMetathyst/atom"

)

// LocalRotation2Key ...
const LocalRotation2Key uint = 3546324404

// LocalRotation2Data ...
type LocalRotation2Data struct {
	X float32
	Y float32	
}

// LocalRotation2Component ...
type LocalRotation2Component struct {
	ctx atom.CTX
	data map[atom.EntityID]LocalRotation2Data
}

// NewLocalRotation2Component ...
func NewLocalRotation2Component() *LocalRotation2Component {
	return &LocalRotation2Component{
		data: make(map[atom.EntityID]LocalRotation2Data),
	}
}

// SetContext ...
func (c *LocalRotation2Component) SetContext(ctx atom.CTX) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

func init() {
	x := NewLocalRotation2Component()
	ctx := atom.Default().RegisterComponent(LocalRotation2Key, x)
	x.SetContext(ctx)
}

// DeleteEntity ...
func (c *LocalRotation2Component) DeleteEntity(id atom.EntityID) {
	delete(c.data, id)
}

// HasEntity ...
func (c *LocalRotation2Component) HasEntity(id atom.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// SetLocalRotation2 ...
func (c *LocalRotation2Component) SetLocalRotation2(id atom.EntityID, localrotation2 LocalRotation2Data) {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = localrotation2
			c.ctx.ComponentUpdated(LocalRotation2Key, id)
		} else {
			c.data[id] = localrotation2
			c.ctx.ComponentAdded(LocalRotation2Key, id)
		}
	}
}

// LocalRotation2 ...
func (c *LocalRotation2Component) LocalRotation2(id atom.EntityID) LocalRotation2Data {
	return c.data[id]
}

// DeleteLocalRotation2 ...
func (c *LocalRotation2Component) DeleteLocalRotation2(id atom.EntityID) {
	delete(c.data, id)
	c.ctx.ComponentDeleted(LocalRotation2Key, id)
}

// SetLocalRotation2X ...
func SetLocalRotation2X(e *atom.EntityManager, id atom.EntityID, localrotation2 LocalRotation2Data) {
	v, _ := e.Component(LocalRotation2Key)
	c := v.(*LocalRotation2Component)
	c.SetLocalRotation2(id, localrotation2)
}

// SetLocalRotation2 ...
func SetLocalRotation2(id atom.EntityID, localrotation2 LocalRotation2Data) {
	SetLocalRotation2X(atom.Default(), id, localrotation2)
}

// LocalRotation2X ...
func LocalRotation2X(e *atom.EntityManager, id atom.EntityID) LocalRotation2Data {
	v, _ := e.Component(LocalRotation2Key)
	c := v.(*LocalRotation2Component)
	return c.LocalRotation2(id)
}

// LocalRotation2 ...
func LocalRotation2(id atom.EntityID) LocalRotation2Data {
	return LocalRotation2X(atom.Default(), id)
}

// DeleteLocalRotation2X ...
func DeleteLocalRotation2X(e *atom.EntityManager, id atom.EntityID) {
	v, _ := e.Component(LocalRotation2Key)
	c := v.(*LocalRotation2Component)
	c.DeleteLocalRotation2(id)
}

// DeleteLocalRotation2 ...
func DeleteLocalRotation2(id atom.EntityID) {
	DeleteLocalRotation2X(atom.Default(), id)
}

// HasLocalRotation2X ...
func HasLocalRotation2X(e *atom.EntityManager, id atom.EntityID) bool {
	v, _ := e.Component(LocalRotation2Key)
	return v.HasEntity(id)
}

// HasLocalRotation2 ...
func HasLocalRotation2(id atom.EntityID) bool {
	return HasLocalRotation2X(atom.Default(), id)
}