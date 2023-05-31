package main

import "fmt"

//Implementation Tech_task.txt
// Структура Product, id - уникальный номер продукта, name - наименование
type Product struct {
	ID   string
	Name string
}

// Класс ProductsImpl для реализации 4 функций
type ProductsImpl struct {
	products []Product
}

// Добавляет новый продукт
// Возвращает true, если продукта с таким id еще не было
// Возвращает false, если был такой id, вставка отменяется
func (p *ProductsImpl) AddProduct(product Product) bool {
	for _, prod := range p.products {
		if prod.ID == product.ID {
			return false
		}
	}
	p.products = append(p.products, product)
	return true
}

// Удаляет продукт
// Возвращает true, если продукт с таким id был
// Возвращает false, если id не было (удаление не происходит)
func (p *ProductsImpl) DeleteProduct(product Product) bool {
	for i, prod := range p.products {
		if prod.ID == product.ID {
			p.products = append(p.products[:i], p.products[i+1:]...)
			return true
		}
	}
	return false
}

// Получает имя (name) продукта
// Возвращает name продукта у которого идентификатор равен (=) id
// Если продукта нет, возвращает пустую строку ""
func (p *ProductsImpl) GetName(id string) string {
	for _, prod := range p.products {
		if prod.ID == id {
			return prod.Name
		}
	}
	return ""
}

// Возвращает массив (список) идентификаторов (id),
// у которых наименование равно (=) name
// Если таких нет, возвращается пустой массив (список)
func (p *ProductsImpl) FindByName(name string) []string {
	var result []string
	for _, prod := range p.products {
		if prod.Name == name {
			result = append(result, prod.ID)
		}
	}
	return result
}

func main() {
	// Создаем объект структуры ProductsImpl
	productsImpl := ProductsImpl{}

	// Создаем объекты продуктов
	product1 := Product{"1", "Kiwi"}
	product2 := Product{"2", "Orange"}
	product3 := Product{"3", "Grapes"}

	// Добавляем продукты
	fmt.Println(productsImpl.AddProduct(product1)) // true
	fmt.Println(productsImpl.AddProduct(product2)) // true
	fmt.Println(productsImpl.AddProduct(product3)) // true

	// Удаляем продукт
	fmt.Println(productsImpl.DeleteProduct(product2)) // true
	fmt.Println(productsImpl.DeleteProduct(product2)) // false, так как продукта с id=2 больше нет

	// Получаем имя продукта по идентификатору
	fmt.Println(productsImpl.GetName("1")) // Kiwi
	fmt.Println(productsImpl.GetName("2")) // "", так как продукта с id=2 больше нет

	// Ищем идентификаторы продуктов по имени
	fmt.Println(productsImpl.FindByName("Kiwi"))   // ["1"]
	fmt.Println(productsImpl.FindByName("Orange")) // []
	fmt.Println(productsImpl.FindByName("Grapes")) // ["3"]
}
