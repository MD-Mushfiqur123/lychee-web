//go:build !cgo

package mlx

import (
	"errors"
)

type Array struct{}

func (a *Array) Dtype() Dtype { return 0 }
func (a *Array) Shape() []int32 { return nil }
func (a *Array) DataFloat32() []float32 { return nil }

type Dtype int

const (
	DtypeFloat32 Dtype = iota
	DtypeBFloat16
)

type SafetensorsFile struct{}

func (s *SafetensorsFile) Free() {}
func (s *SafetensorsFile) GetMetadata(key string) string { return "" }
func (s *SafetensorsFile) Get(name string) *Array { return nil }

func LoadSafetensorsNative(path string) (*SafetensorsFile, error) {
	return nil, errors.New("MLX requires CGO")
}

func AsType(arr *Array, dtype Dtype) *Array { return arr }
func Contiguous(arr *Array) *Array { return arr }
func Eval(arrs ...*Array) {}
func NewArray(data []float32, shape []int32) *Array { return nil }
func ClipScalar(arr *Array, min, max float32) *Array { return arr }
func Transpose(arr *Array, axes []int32) *Array { return arr }
func Slice(arr *Array, start, end []int32) *Array { return arr }

func GPUIsAvailable() bool { return false }
func SetDefaultDeviceGPU() {}
func EnableCompile() {}
func Collect(v any) []*Array { return nil }
func ToBFloat16(arr *Array) *Array { return arr }
func MetalGetActiveMemory() uint64 { return 0 }
func MetalResetPeakMemory() {}
func MetalStartCapture(path string) {}
func MetalIsAvailable() bool { return false }
func Pin(arr *Array) {}
func Sweep() {}
func Load(path string) map[string]*Array { return nil }
