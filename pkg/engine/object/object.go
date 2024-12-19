package object

// представляет собой элементарный объект для обработки
type Object interface {
	// возращение имени типа в виде строки
	TypeName() string
	// возвращение строковое представление типа
	String() string
	// выполнение бинарной операции
	BinaryOp(int, Object) (Object, error)
	// выполнение унарной операции
	UnaryOp(int) (Object, error)
	// может ли данный объект быть вызван как функция
	CanCall() bool
	// вызов объекта как функции с аргументами
	Call(...Object) (Object, error)
	// call method of object
	MethodCall(string, ...Object) (Object, error)
	// может ли данный объект итерироваться
	CanIterate() bool
	// возвращает итератор для объекта
	Iterate() Iterator
	// получение значения из объекта
	GetValue() any
	// получение значения объекта по его индексу внутри
	IndexGet(Object) (Object, error)
	// установка значения объекта по индексу
	IndexSet(Object, Object) error
}

// Impl базовая реализация объекта
type Impl struct{}

func (o *Impl) TypeName() string {
	panic(ErrNotImplemented)
}

func (o *Impl) String() string {
	panic(ErrNotImplemented)
}

func (o *Impl) BinaryOp(_ int, _ Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (o *Impl) UnaryOp(_ int) (Object, error) {
	return nil, ErrNotImplemented
}

func (o *Impl) CanCall() bool {
	return false
}

func (o *Impl) Call(_ ...Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (o *Impl) CanIterate() bool {
	return false
}

func (o *Impl) Iterate() Iterator {
	return nil
}

func (o *Impl) GetValue() any {
	return nil
}

func (o *Impl) IndexGet(Object) (Object, error) {
	return nil, ErrInvalidOp
}

func (o *Impl) IndexSet(Object, Object) error {
	return ErrInvalidOp
}

func (o *Impl) MethodCall(string, ...Object) (Object, error) {
	return nil, ErrUnknownMethod
}
