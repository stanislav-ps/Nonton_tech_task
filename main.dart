//Implementation Tech_task.txt

class Product {
  late String id;
  late String name;

  Product(this.id, this.name);
}

class ProductsImpl {
  List<Product> products = [];

  // Добавляет новый продукт
  // Возвращает true, если продукта с таким id еще не было
  // Возвращает false, если был такой id, вставка отменяется
  bool addProduct(Product product) {
    for (var p in products) {
      if (p.id == product.id) {
        return false;
      }
    }
    products.add(product);
    return true;
  }

  // Удаляет продукт
  // Возвращает true, если продукт с таким id был
  // Возвращает false, если id не было (удаление не происходит)
  bool deleteProduct(Product product) {
    for (var p in products) {
      if (p.id == product.id) {
        products.remove(p);
        return true;
      }
    }
    return false;
  }

  // Получает имя (name) продукта
  // Возвращает name продукта у которого идентификатор равен (=) id
  // Если продукта нет, возвращает пустую строку ""
  String getName(String id) {
    for (var p in products) {
      if (p.id == id) {
        return p.name;
      }
    }
    return '';
  }

  // Возвращает массив (список) идентификаторов (id),
  // у которых наименование равно (=) name
  // Если таких нет, возвращается пустой массив (список)
  List<String> findByName(String name) {
    List<String> result = [];
    for (var p in products) {
      if (p.name == name) {
        result.add(p.id);
      }
    }
    return result;
  }
}

void main() {
  // Создаем объект класса ProductsImpl
  var productsImpl = ProductsImpl();

  // Создаем объекты продуктов
  var product1 = Product('1', 'Kiwi');
  var product2 = Product('2', 'Orange');
  var product3 = Product('3', 'Grapes');

  // Добавляем продукты
  print(productsImpl.addProduct(product1));  // true
  print(productsImpl.addProduct(product2));  // true
  print(productsImpl.addProduct(product3));  // true

  // Удаляем продукт
  print(productsImpl.deleteProduct(product2));  // true
  print(productsImpl.deleteProduct(product2));  // false, так как продукта с id=2 больше нет

  // Получаем имя продукта по идентификатору
  print(productsImpl.getName('1'));  // Kiwi
  print(productsImpl.getName('2'));  // "", так как продукта с id=2 больше нет

  // Ищем идентификаторы продуктов по имени
  print(productsImpl.findByName('Kiwi'));   // ['1']
  print(productsImpl.findByName('Orange')); // []
  print(productsImpl.findByName('Grapes')); // ['3']
}
