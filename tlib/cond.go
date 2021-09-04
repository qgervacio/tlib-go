// Copyright (c) 2021. Quirino Gervacio

package tlib

// Tiff is your good 'ol elvis operator
func Tiff(cond bool, tiffThen, tiffElse interface{}) interface{} {
	if cond {
		return tiffThen
	}
	return tiffElse
}
