package variable

import (
	"fmt"

	"github.com/tomnz/glowpher/internal/color"
)

type Param interface{}

type ParamVariable interface {
	Param
	WireVariable(variables map[string]Variable) error
}

type ParamInt interface {
	Param
	Value() int
}

type ParamIntLiteral int

func (i ParamIntLiteral) Value() int {
	return int(i)
}

type ParamIntVariable struct {
	varName       string
	variable      NumberVariable
	multiply, add float32
}

func NewParamIntVariable(varName string, multiply, add float32) *ParamIntVariable {
	return &ParamIntVariable{
		varName:  varName,
		multiply: multiply,
		add:      add,
	}
}

func (i *ParamIntVariable) WireVariable(variables map[string]Variable) error {
	vari, found := variables[i.varName]
	if !found {
		return fmt.Errorf("unknown variable %q", i.varName)
	}
	numVari, ok := vari.(NumberVariable)
	if !ok {
		return fmt.Errorf("variable %q cannot be connected to an integer param", i.varName)
	}

	i.variable = numVari
	return nil
}

func (i *ParamIntVariable) Value() int {
	return int(i.variable.Value()*i.multiply + i.add)
}

type ParamFloat interface {
	Param
	Value() float32
}

type ParamFloatLiteral float32

func (f ParamFloatLiteral) Value() float32 {
	return float32(f)
}

type ParamFloatVariable struct {
	varName       string
	variable      NumberVariable
	multiply, add float32
}

func NewParamFloatVariable(varName string, multiply, add float32) *ParamFloatVariable {
	return &ParamFloatVariable{
		varName:  varName,
		multiply: multiply,
		add:      add,
	}
}
func (f *ParamFloatVariable) WireVariable(variables map[string]Variable) error {
	vari, found := variables[f.varName]
	if !found {
		return fmt.Errorf("unknown variable %q", f.varName)
	}
	numVari, ok := vari.(NumberVariable)
	if !ok {
		return fmt.Errorf("variable %q cannot be connected to a float param", f.varName)
	}

	f.variable = numVari
	return nil
}

func (f *ParamFloatVariable) Value() float32 {
	return f.variable.Value()*f.multiply + f.add
}

type ParamColor interface {
	Param
	Value() color.Color
}

type ParamColorLiteral color.Color

func (c ParamColorLiteral) Value() color.Color {
	return color.Color(c)
}