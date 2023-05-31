#Implementation Tech_task.txt

class Product:
    def __init__(self, id, name):
        self.ID = id
        self.Name = name


class ProductsImpl:
    def __init__(self):
        self.products = []  # Список для хранения продуктов

    def AddProduct(self, product):
        """
        Добавляет новый продукт.

        Возвращает True, если продукта с таким id еще не было.
        Возвращает False, если был такой id, вставка отменяется.
        """
        for prod in self.products:
            if prod.ID == product.ID:
                return False
        self.products.append(product)
        return True

    def DeleteProduct(self, product):
        """
        Удаляет продукт.

        Возвращает True, если продукт с таким id был.
        Возвращает False, если id не было (удаление не происходит).
        """
        for i, prod in enumerate(self.products):
            if prod.ID == product.ID:
                self.products.pop(i)
                return True
        return False

    def GetName(self, id):
        """
        Получает имя (name) продукта.

        Возвращает name продукта у которого идентификатор равен (=) id.
        Если продукта нет, возвращает пустую строку "".
        """
        for prod in self.products:
            if prod.ID == id:
                return prod.Name
        return ""

    def FindByName(self, name):
        """
        Возвращает массив (список) идентификаторов (id),
        у которых наименование равно (=) name.
        Если таких нет, возвращается пустой массив (список).
        """
        result = []
        for prod in self.products:
            if prod.Name == name:
                result.append(prod.ID)
        return result


# Создаем объект класса ProductsImpl
products_impl = ProductsImpl()

# Создаем объекты продуктов
product1 = Product("1", "Kiwi")
product2 = Product("2", "Orange")
product3 = Product("3", "Grapes")

# Добавляем продукты
print(products_impl.AddProduct(product1))  # True
print(products_impl.AddProduct(product2))  # True
print(products_impl.AddProduct(product3))  # True

# Удаляем продукт
print(products_impl.DeleteProduct(product2))  # True
print(products_impl.DeleteProduct(product2))  # False, так как продукта с id=2 больше нет

# Получаем имя продукта по идентификатору
print(products_impl.GetName("1"))  # Kiwi
print(products_impl.GetName("2"))  # "", так как продукта с id=2 больше нет

# Ищем идентификаторы продуктов по имени
print(products_impl.FindByName("Kiwi"))   # ['1']
print(products_impl.FindByName("Orange")) # []
print(products_impl.FindByName("Grapes")) # ['3']
