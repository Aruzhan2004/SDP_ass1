package main

import (
	"fmt"
)

type FilterStrategy interface {
	Process(image Image) Image
}

type Image struct {
	data string
}

type SepiaFilter struct{}

func (sf *SepiaFilter) Process(image Image) Image {
	fmt.Println("Apply SEPIA filter to image")
	return image
}

type BlackAndWhiteFilter struct{}

func (bwf *BlackAndWhiteFilter) Process(image Image) Image {
	fmt.Println("Apply BLACKANDWHITE filter to image")
	return image
}

type DistortionFilter struct{}

func (df *DistortionFilter) Process(image Image) Image {
	fmt.Println("Apply DISTORTION filter to image")
	return image
}

type Filter struct {
	filterStrategy FilterStrategy
}

func (f *Filter) ApplyFilter(image Image) Image {
	return f.filterStrategy.Process(image)
}

func main() {
	filter := Filter{}
	image := Image{"YourImageDataHere"}

	filter.filterStrategy = &SepiaFilter{}
	filteredImage := filter.ApplyFilter(image)
	fmt.Println("Filtered Image (Sepia):", filteredImage)

	filter.filterStrategy = &BlackAndWhiteFilter{}
	filteredImage = filter.ApplyFilter(image)
	fmt.Println("Filtered Image (Black and White):", filteredImage)

	filter.filterStrategy = &DistortionFilter{}
	filteredImage = filter.ApplyFilter(image)
	fmt.Println("Filtered Image (Distortion):", filteredImage)
}
