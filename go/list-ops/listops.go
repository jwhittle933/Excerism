package listops

// import "fmt"

// IntList ...
type IntList []int

type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

// Foldr iterates from len(il)-1 to 0
func (il IntList) Foldr(fn binFunc, init int) int {
	if len(il) == 0 {
		return init
	}

	for i := len(il); i > 0; i-- {
		// accumulator as second argument
		init = fn(il[i-1], init)
	}
	return init
}

// Foldl iterates from 0 to len(il)-1
func (il IntList) Foldl(fn binFunc, init int) int {
	if len(il) == 0 {
		return init
	}

	for i := 0; i < len(il); i++ {
		// accumulator as first argument
		init = fn(init, il[i])
	}

	return init
}

// Filter ...
func (il IntList) Filter(fn predFunc) IntList {
	newIl := IntList{}
	for _, x := range il {
		if fn(x) {
			newIl = append(newIl, x)
		}
	}

	return newIl
}

// Length ...
func (il IntList) Length() int {
	return len(il)
}

// Map ...
func (il IntList) Map(fn unaryFunc) IntList {
	newIl := make(IntList, len(il))
	for i, val := range il {
		newIl[i] = fn(val)
	}
	return newIl
}

// Reverse ...
func (il IntList) Reverse() IntList {
	newIl := make(IntList, len(il))

	for i, val := range il {
		newIl[len(il)-(i+1)] = val
	}

	return newIl
}

// Append ...
func (il IntList) Append(ilist IntList) IntList {

	l := len(il)+len(ilist)
	newIl := make(IntList, l, l)

	for i := 0; i < len(il); i++ {
		newIl[i] = il[i]
	}

	for j := 0; j < len(ilist); j++ {
		newIl[j+len(il)] = ilist[j]
	}

	return newIl
}

// Concat ..
func (il IntList) Concat(ils []IntList) IntList {
	for _, list := range ils {
		for _, x := range list {
			il = append(il, x)
		}
	}

	return il
}
