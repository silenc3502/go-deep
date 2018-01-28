package deep

import "math"

type Mode int

const (
	ModeDefault    Mode = 0 // Use default output layer activations
	ModeMultiClass Mode = 1 // Softmax output layer for multiclass classification
	ModeRegression Mode = 2 // Linear output layer
	ModeBinary     Mode = 3 // Sigmoid output layer for binary classification
	ModeMultiLabel Mode = 4 // Sigmoid output layer with multiclass CE (no softmax)
)

func OutputActivation(c Mode) ActivationType {
	switch c {
	case ModeMultiClass:
		return ActivationSoftmax
	case ModeRegression:
		return ActivationLinear
	case ModeBinary, ModeMultiLabel:
		return ActivationSigmoid
	}
	return ActivationNone
}

func GetActivation(act ActivationType) Differentiable {
	switch act {
	case ActivationSigmoid:
		return Sigmoid{}
	case ActivationTanh:
		return Tanh{}
	case ActivationReLU:
		return ReLU{}
	case ActivationLinear:
		return Linear{}
	case ActivationSoftmax:
		return Linear{}
	}
	return Linear{}
}

type ActivationType int

const (
	ActivationNone    ActivationType = 0
	ActivationSigmoid ActivationType = 1
	ActivationTanh    ActivationType = 2
	ActivationReLU    ActivationType = 3
	ActivationLinear  ActivationType = 4
	ActivationSoftmax ActivationType = 5
)

type Differentiable interface {
	F(float64) float64
	Df(float64) float64
}

type Sigmoid struct{}

func (a Sigmoid) F(x float64) float64  { return Logistic(x, 1) }
func (a Sigmoid) Df(y float64) float64 { return y * (1 - y) }

func Logistic(x, a float64) float64 {
	return 1 / (1 + math.Exp(-a*x))
}

type Tanh struct{}

func (a Tanh) F(x float64) float64  { return (1 - math.Exp(-2*x)) / (1 + math.Exp(-2*x)) }
func (a Tanh) Df(y float64) float64 { return 1 - math.Pow(y, 2) }

type ReLU struct{}

func (a ReLU) F(x float64) float64 { return math.Max(x, 0) }
func (a ReLU) Df(y float64) float64 {
	if y > 0 {
		return 1
	}
	return 0
}

type Linear struct{}

func (a Linear) F(x float64) float64  { return x }
func (a Linear) Df(x float64) float64 { return 1 }
