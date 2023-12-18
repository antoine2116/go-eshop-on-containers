package type_mapper

// https://github.com/mehdihadeli/go-ecommerce-microservices/blob/main/internal/pkg/reflection/type_mappper/unsafe_types.go

import "unsafe"

//go:linkname typelinks2 reflect.typelinks
func typelinks2() (sections []unsafe.Pointer, offset [][]int32)

//go:linkname resolveTypeOff reflect.resolveTypeOff
func resolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

type emptyInterface struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}
