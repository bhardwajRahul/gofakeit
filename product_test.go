package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleProduct() {
	Seed(11)
	product := Product()
	fmt.Println(product.Name)
	fmt.Println(product.Description)
	fmt.Println(product.Categories)
	fmt.Println(product.Price)
	fmt.Println(product.Features)
	fmt.Println(product.Color)
	fmt.Println(product.Material)
	fmt.Println(product.UPC)
	fmt.Println(product.Audience)
	fmt.Println(product.Dimension)
	fmt.Println(product.UseCase)
	fmt.Println(product.Benefit)
	fmt.Println(product.Suffix)

	// Output: Wave Precision Lamp
	// This upset product is crafted from wood and includes wireless, making it perfect for personal grooming and delivering comfort for travelers and professionals.
	// [cosmetics outdoor gear]
	// 73.93
	// [touchscreen ultra-lightweight gps-enabled biometric]
	// maroon
	// silver
	// 041026894059
	// [gamers musicians]
	// heavy
	// learning
	// efficiency
	// dash
}

func ExampleFaker_Product() {
	f := New(11)
	product := f.Product()
	fmt.Println(product.Name)
	fmt.Println(product.Description)
	fmt.Println(product.Categories)
	fmt.Println(product.Price)
	fmt.Println(product.Features)
	fmt.Println(product.Color)
	fmt.Println(product.Material)
	fmt.Println(product.UPC)
	fmt.Println(product.Audience)
	fmt.Println(product.Dimension)
	fmt.Println(product.UseCase)
	fmt.Println(product.Benefit)
	fmt.Println(product.Suffix)

	// Output: Wave Precision Lamp
	// This upset product is crafted from wood and includes wireless, making it perfect for personal grooming and delivering comfort for travelers and professionals.
	// [cosmetics outdoor gear]
	// 73.93
	// [touchscreen ultra-lightweight gps-enabled biometric]
	// maroon
	// silver
	// 041026894059
	// [gamers musicians]
	// heavy
	// learning
	// efficiency
	// dash
}

func TestProduct(t *testing.T) {
	for i := 0; i < 1000; i++ {
		product := Product()
		if product.Name == "" {
			t.Error("Name is empty")
		}

		if product.Description == "" {
			t.Error("Description is empty")
		}

		if len(product.Categories) == 0 {
			t.Error("Categories is empty")
		}

		if product.Price == 0 {
			t.Error("Price is empty")
		}

		if len(product.Features) == 0 {
			t.Error("Features is empty")
		}

		if product.Color == "" {
			t.Error("Color is empty")
		}

		if product.Material == "" {
			t.Error("Material is empty")
		}

		if product.UPC == "" {
			t.Error("UPC is empty")
		}

		if len(product.Audience) == 0 {
			t.Error("Audience is empty")
		}

		if product.Dimension == "" {
			t.Error("Dimension is empty")
		}

		if len(product.UseCase) == 0 {
			t.Error("UseCase is empty")
		}

		if len(product.Benefit) == 0 {
			t.Error("Benefit is empty")
		}

		if product.Suffix == "" {
			t.Error("Suffix is empty")
		}
	}
}

func BenchmarkProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Product()
	}
}

func ExampleProductName() {
	Seed(11)
	fmt.Println(ProductName())

	// Output: Green Glass Hair Dryer
}

func ExampleFaker_ProductName() {
	f := New(11)
	fmt.Println(f.ProductName())

	// Output: Green Glass Hair Dryer
}

func BenchmarkProductName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductName()
	}
}

func ExampleProductDescription() {
	Seed(11)
	fmt.Println(ProductDescription())

	// Output: This product, ideal for seniors and families, features puzzled gold and incorporates gps-enabled to ensure robust construction during remote work.
}

func ExampleFaker_ProductDescription() {
	f := New(11)
	fmt.Println(f.ProductDescription())

	// Output: This product, ideal for seniors and families, features puzzled gold and incorporates gps-enabled to ensure robust construction during remote work.
}

// Runs 10,000 tests to ensure description doesnt have any { or } in it
func TestProductDescriptionReplacement(t *testing.T) {
	for i := 0; i < 100000; i++ {
		desc := ProductDescription()
		if strings.ContainsAny(desc, "{") || strings.ContainsAny(desc, "}") {
			t.Error("Description contains { or }")
		}
	}
}

func BenchmarkProductDescription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductDescription()
	}
}

func ExampleProductCategory() {
	Seed(11)
	fmt.Println(ProductCategory())

	// Output: pet food
}

func ExampleFaker_ProductCategory() {
	f := New(11)
	fmt.Println(f.ProductCategory())

	// Output: pet food
}

func BenchmarkProductCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductCategory()
	}
}

func ExampleProductFeature() {
	Seed(11)
	fmt.Println(ProductFeature())

	// Output: fast-charging
}

func ExampleFaker_ProductFeature() {
	f := New(11)
	fmt.Println(f.ProductFeature())

	// Output: fast-charging
}

func BenchmarkProductFeature(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductFeature()
	}
}

func ExampleProductMaterial() {
	Seed(11)
	fmt.Println(ProductMaterial())

	// Output: quartz
}

func ExampleFaker_ProductMaterial() {
	f := New(11)
	fmt.Println(f.ProductMaterial())

	// Output: quartz
}

func BenchmarkProductMaterial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductMaterial()
	}
}

func ExampleProductUPC() {
	Seed(11)
	fmt.Println(ProductUPC())

	// Output: 088125275989
}

func ExampleFaker_ProductUPC() {
	f := New(11)
	fmt.Println(f.ProductUPC())

	// Output: 088125275989
}

func BenchmarkProductUPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductUPC()
	}
}

func ExampleProductAudience() {
	Seed(11)
	fmt.Println(ProductAudience())

	// Output: [DIY enthusiasts students]
}

func ExampleFaker_ProductAudience() {
	f := New(11)
	fmt.Println(f.ProductAudience())

	// Output: [DIY enthusiasts students]
}

func BenchmarkProductAudience(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductAudience()
	}
}

func ExampleProductDimension() {
	Seed(11)
	fmt.Println(ProductDimension())

	// Output: standard
}

func ExampleFaker_ProductDimension() {
	f := New(11)
	fmt.Println(f.ProductDimension())

	// Output: standard
}

func BenchmarkProductDimension(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductDimension()
	}
}

func ExampleProductUseCase() {
	Seed(11)
	fmt.Println(ProductUseCase())

	// Output: remote work
}

func ExampleFaker_ProductUseCase() {
	f := New(11)
	fmt.Println(f.ProductUseCase())

	// Output: remote work
}

func BenchmarkProductUseCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductUseCase()
	}
}

func ExampleProductBenefit() {
	Seed(11)
	fmt.Println(ProductBenefit())

	// Output: minimal maintenance
}

func ExampleFaker_ProductBenefit() {
	f := New(11)
	fmt.Println(f.ProductBenefit())

	// Output: minimal maintenance
}

func BenchmarkProductBenefit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductBenefit()
	}
}

func ExampleProductSuffix() {
	Seed(11)
	fmt.Println(ProductSuffix())

	// Output: turbo
}

func ExampleFaker_ProductSuffix() {
	f := New(11)
	fmt.Println(f.ProductSuffix())

	// Output: turbo
}

func BenchmarkProductSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductSuffix()
	}
}

func ExampleProductISBN() {
	Seed(11)
	fmt.Println(ProductISBN(&ISBNOptions{Version: "13", Separator: "-"}))

	// Output: 978-8-8125-2759-5
}

func ExampleFaker_ProductISBN() {
	f := New(11)
	fmt.Println(f.ProductISBN(&ISBNOptions{Version: "13", Separator: "-"}))

	// Output: 978-8-8125-2759-5
}

func BenchmarkProductISBN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductISBN(&ISBNOptions{Version: "13", Separator: "-"})
	}
}

func ExampleProductISBN_isbn10() {
	Seed(11)
	fmt.Println(ProductISBN(&ISBNOptions{Version: "10", Separator: "-"}))

	// Output: 8-8125-275-7
}

func ExampleFaker_ProductISBN_isbn10() {
	f := New(11)
	fmt.Println(f.ProductISBN(&ISBNOptions{Version: "10", Separator: "-"}))

	// Output: 8-8125-275-7
}
