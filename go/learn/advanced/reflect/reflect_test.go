package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

/*
什麼時候需要反射？
1. 不能預先確定參數的類型，可能是沒有約定好，也可能是傳入的參數類型有很多，且不能統一表示
2. 函數需要根據輸入的參數類型來動態的執行不同的行為
常用場景：結構體json標籤序列化 JobId int `job_id`，
適配器函數作為統一處理接口 bridge=func(call interface{},args ...interface{})
*/

/*
1.反射可以在運行時動態獲取變量的各種信息，比如變量的類型(type)，類別(kind)
2.如果是結構體變量，還可以獲取到結構體本身的信息（包誇結構體的字段，方法）
3.通過反射，可以修改結構體的值，可以調用結構體關聯的方法
*/

/* 重要函數＆概念
1. reflect.TypeOf(變量名)，獲取變量的類型，返回reflect.Type類型
2. reflect.ValueOf(變量名)，獲取變量的值，返回reflect.Value類型（是一個結構體類型）
3. 變量，interface{}和reflect.Value是可以互相轉換的，這點在實際開發中經常使用
*/

type Student struct {
	Name string
}

func TestReflect(t *testing.T) {

	var student1 = Student{Name: "Raymond"}
	reflect1(student1)
}

// 專門用於做反射
func reflect1(i interface{}) {
	// 1.如何將interface{}轉成reflect.Value
	rVal := reflect.ValueOf(i)

	// 2.如何將reflect.Value -> interface{}
	iVal := rVal.Interface()

	// 3.如何將interface{}轉換成原來的變量類型，使用類型斷言
	v := iVal.(Student)

	fmt.Println(v)
}

// 變量，interface{}，reflect.Value 之間轉換的示意圖
/*
	(*´∀`)~♥	i n t e r f a c e { }
		  ⬀	   ⬃		⬁			⬂
	傳遞參數  類型斷言  v.Interface()  reflect.ValueOf()函數
		⬀	⬃				⬁		⬃
		變 量				reflect.Value
*/

//專門用於做反射[基本數據類型]
func reflect2(i interface{}) {
	// 通過反射獲取傳入的變量的 type , kind ,值
	// 1. 先獲取到 reflect.Type
	rType := reflect.TypeOf(i)
	// 這樣輸出看到的是轉換過的類型
	fmt.Println("rType= ", rType) //int
	// 這樣輸出才看得到的rType類型
	fmt.Printf("rType= %T \n", rType) //*reflect.rtype

	// 2.獲取到 reflect.Value
	rVal := reflect.ValueOf(i)
	fmt.Printf("rVal=%v ,rVal type=%T\n", rVal, rVal) // 100,reflect.Value
	// sum:= 10 +rVal // 不能加，因為rVal是reflect.Value類型，不是int
	// 要調用方法
	fmt.Println(100 + rVal.Int())
	// fmt.Println(100 + rVal.Float()) //會panic

	// 3.獲取變量對應的Kind
	// (1) rVal.Kind() ==>
	fmt.Printf("rVal.Kind=%v\n", rVal.Kind())
	// (2) rType.Kind() ==>
	fmt.Printf("rType.Kind=%v\n", rType.Kind())

	// 將rVal轉成interface{}
	iv := rVal.Interface()
	// 將interface{} 通過斷言轉成需要的類型
	fmt.Println(iv.(int))
	// 要小心，斷言是運行時才會進行，如果iv進來不是float，下面這樣會panic
	// fmt.Println(iv.(float32))
}
func TestReflect2(t *testing.T) {
	// 演示對反射的基本操作
	var a int = 100
	reflect2(a)
}

// 專門用於做反射[對結構體的反射]
func reflectStruct(i interface{}) {
	rVal := reflect.ValueOf(i)
	iv := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iv, iv)
	struct1, ok := iv.(Struct1)
	if ok {
		fmt.Printf("struct1.Name= %v\n", struct1.Name)
	}
}

type Struct1 struct {
	Name string
	Age  int
}

func TestReflectStruct(t *testing.T) {
	s1 := Struct1{
		Name: "Raymond",
		Age:  25,
	}
	reflectStruct(s1)
}

/* 反射的注意事項和細節說明 */
// 1. reflect.Value.Kind 獲取變量的"類別"，返回的是一個常量const int,string,bool..
// 2. Type 是類型，Kind是類別，Type和Kind可能是相同的，也可能是不同的
// 比如： var num int = 10，num的Type是int，Kind也是int
// 		 var stu Student ，stu的Type是 包名(package).Student，Kind是struct
// 3. 將interface{} 通過斷言轉成需要的類型時，或調用reflect,Value._const_()時，要保證有辦法轉換，不然會panic
// 比如： rVal.Interface().(float32)斷言 ， rVal.Float()取值，如果rVal傳近來是int就會panic

// 4. 通過反射來修改變量，要注意當使用SetXXX方法來設置時，需要通過對應的指針類型來完成，
// 這樣才能改變傳入的變量的值，同時需要使用到reflect.Value.Elem()方法

// 通過反射修改num int的值
func reflectModify(i interface{}) {
	// 1.獲取reflect.Value
	rVal := reflect.ValueOf(i)
	// 2.看看 rVal的Kind是
	fmt.Printf("rVal kind=%v", rVal.Kind())
	// 3. 改變值
	// rVal.SetInt(100) //會panic
	rVal.Elem().SetInt(100)
	/*
		rVal.Elem() 用於獲取指針變量，行為類似：
		var num = 10
		var b *int = &num
		*b = 100
	*/
}
func TestReflectModify(t *testing.T) {
	var num int = 10
	reflectModify(&num)
	fmt.Println("num:\n", num)
}

/*反射練習： 給定一個變量 var v float64 = 1.2 ，請使用反射來得到它的reflect.Value，
然後獲取對應的Type，Kind和值，並將reflect.Value轉換成interface{}，再將interface{}轉換成float64*/

func reflectFloat64(b interface{}) {
	rVal := reflect.ValueOf(b)
	rType := reflect.TypeOf(b)
	rVal_kind := rVal.Kind()
	fmt.Printf("rVal:%v ,rval_type:%v,rval_kind:%v\n", rVal, rType, rVal_kind)
	iV := rVal.Interface()
	fmt.Println(iV.(float64))
}

func TestReflectFloat64(t *testing.T) {
	var v float64 = 1.2
	reflectFloat64(v)
}

// 注意使用SetXXX方法時Elem()的使用
func TestSetString(t *testing.T) {
	var str string = "aaaa" //ok
	// fs := reflect.ValueOf(str) //ok fs -> string
	// fs.SetString("rrrrr") //error
	// 上面這樣寫會報錯，要像下面
	fmt.Printf("%v\n", str)
	// 要取地址
	fs := reflect.ValueOf(&str)
	// 然後使用Elem()
	fs.Elem().SetString("rrrrr")
	fmt.Printf("%v\n", str)
}

/*
	反射最佳實踐
*/
// 1.使用反射遍歷結構體的字段，調用結構體的方法，並獲取結構體標籤的值

//定義了一個Monster結構體
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

//方法，顯示s的值
func (m Monster) Print() {
	fmt.Println("-----start-----")
	fmt.Println(m)
	fmt.Println("------end------")
}

//方法，返回兩個數的和
func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//方法，返回平方
func (m Monster) Square(n1 int) int {
	return n1 * n1
}

func testStruct(a interface{}) {
	// 獲取reflect.Type類型
	rTyp := reflect.TypeOf(a)
	// 獲取reflect.Value類型
	rVal := reflect.ValueOf(a)
	// 獲取到a對應的類別
	kd := rVal.Kind()
	// 如果傳入的不是struct，就退出
	if kd != reflect.Struct {
		fmt.Println("except struct")
		return
	}
	// 獲取到該結構體有幾個字段
	num_field := rVal.NumField()
	fmt.Printf("struct has %d fields \n", num_field)
	// 遍歷結構體的所有字段
	for i := 0; i < num_field; i++ {
		fmt.Printf("Field %d: 值為=%v\n", i, rVal.Field(i)) // 可以看一下field的值，但是不能拿來運算
		//***重點：獲取到struct標籤，注意需要通過reflect.Type來獲取tag標籤的值，不能用Value***
		tagVal := rTyp.Field(i).Tag.Get("json")
		// 如果該字段有tag就顯示，沒有就不顯示
		if tagVal != "" {
			fmt.Printf("Field %d: tag=%v\n", i, tagVal)
		}
	}
	// 獲取到該結構體有多少個方法
	numOfMethod := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//var params []reflect.Value
	// 獲取到第二個方法(從0開始)，然後調用它
	rVal.Method(1).Call(nil)
	/*
		要注意Print()，GetSum()，Square()，三個方法
		Index的順序不是按照寫的順序，而是按照函數名稱ASCII code的大小
		看開頭大小排序應該是G<P<S，所以方法的順序為
		0:GetSum(),1:Print(),2:Square()
	*/

	//調用結構體的第1個方法Method(0)
	var params []reflect.Value //宣告了一個 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rVal.Method(0).Call(params) //傳入的參數是 []reflect.Value
	fmt.Println("res=", res[0].Int())  //返回結果，返回的結果也是 []reflect.Value
}

func TestReflectBestApply(t *testing.T) {
	//創建一個Monster實體
	var a Monster = Monster{
		Name:  "Raymond",
		Age:   25,
		Score: 77.7,
	}
	//將Monster實體傳遞給testStruct函數
	testStruct(a)
}

// 2. 要進行賦值操作時，可以這樣幹
func testStruct2(a interface{}) {
	rType := reflect.TypeOf(a)
	rVal := reflect.ValueOf(a)
	kd := rVal.Kind()
	fmt.Printf("kind: %v,%T,elem_kind:%v\n", kd, kd, rVal.Elem().Kind())
	// Kind要是指針，且是指向的是一個struct
	if kd != reflect.Ptr && rVal.Elem().Kind() == reflect.Struct {
		fmt.Println("except struct")
		return
	}

	num_fields := rVal.Elem().NumField()
	fmt.Printf("struct has %d fields \n", num_fields)

	for i := 0; i < num_fields; i++ {
		fmt.Printf("%d %v\n", i, rVal.Elem().Field(i).Kind())
	}
	/*要用reflect.Type*/
	tag := rType.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	numOfMethod := rVal.NumMethod() // 不能加Elem()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	// 遍歷顯示所有的方法
	for i := 0; i < numOfMethod; i++ {
		method := rType.Method(i)
		fmt.Printf("method name:%s ,type:%s, exported:%t\n", method.Name, method.Type, method.IsExported())
	}
	// 呼叫Set()方法改變值
	var params []reflect.Value
	var score float32 = 99.9
	params = append(params, reflect.ValueOf("GGGGG"), reflect.ValueOf(99), reflect.ValueOf(score), reflect.ValueOf("M"))

	rVal.Method(2).Call(params)
	// 不能寫 rVal.Elem().Method(2).Call(params)，這樣會抓不到Monster的指針方法(Set方法)，
	// 因為近來的是結構體指針，加了Elem()就轉成非指針的結構體，自然就抓不到結構體的指針方法了
	/*
		userValue := userPtrValue.Elem()                    //Elem() 指針Value轉為非指針Value
		fmt.Println(userValue.Kind(), userPtrValue.Kind())  //struct ptr
		userPtrValue3 := userValue.Addr()                   //Addr() 非指針Value轉為指針Value
		fmt.Println(userValue.Kind(), userPtrValue3.Kind()) //struct ptr
	*/
	// 但是如果是調用SetXXX方法，還是要記得用Elem()
	// 由於rVal對應的原始對象是指針，需要通過Elem()返回指針指向的對象
	/*	強調一下，要想修改原始數據的值，給ValueOf傳的必須是指針，而指針Value不能調用Set和FieldByName方法，所以得先通過Elem()轉為非指針Value */
	rVal.Elem().Field(0).SetString("RRRRRRRR")

}

//方法，接收4個值給m賦值
func (m *Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func TestReflectBestApply2(t *testing.T) {
	//創建一個Monster實體
	var a Monster = Monster{
		Name:  "Raymond",
		Age:   25,
		Score: 77.7,
	}
	//將Monster實體 指針 傳遞給testStruct2函數
	testStruct2(&a)
	fmt.Println(a)
}

// 3. 定義兩個函數test1和test2，定義一個適配器函數用作為統一處理的接口
func TestReflectBestApply3(t *testing.T) {
	call1 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}
	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}
	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}
	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}

// 4. 使用反射操作任意的結構體類型
type user struct {
	UserId int
	Name   string
}

func TestReflectBestApply4(t *testing.T) {
	var (
		model *user
		sv    reflect.Value
	)
	model = &user{}
	sv = reflect.ValueOf(model) //ptr
	t.Log("reflect.ValueOf", sv.Kind().String())

	sv = sv.Elem() //相當於取到真正的reflect.Value
	t.Log("reflect.ValueOf.Elem", sv.Kind().String())
	sv.FieldByName("UserId").SetInt(1234567)
	sv.FieldByName("Name").SetString("Raymond")
	t.Log("model:", model)
}

// 5. 使用反射[創建]並操作結構體
func TestReflectBestApply5(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)                  //獲取類型 *user
	t.Log("reflect.TypeOf", st.Kind().String()) // ptr

	st = st.Elem()                                   //st指向真正的類型
	t.Log("reflect.TypeOf.Elem", st.Kind().String()) //struct

	elem = reflect.New(st)                                 //New返回一個reflect.Value類型值，該值持有一個指向類型為type的新申請的零值的指針(指向user)
	t.Log("reflect.New", elem.Kind().String())             // ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) //struct

	// model就是創建的user結構體的變量(實體)
	model = elem.Interface().(*user)           // 轉成空接口，再用類型斷言，讓model和elem都指向一個user
	elem = elem.Elem()                         //取得elem指向的值
	elem.FieldByName("UserId").SetInt(1234567) //賦值
	elem.FieldByName("Name").SetString("Raymond")
	t.Log("model:", model)

}

/*...........一些其他的使用..............*/
//使用switch來斷言
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)

	}
}
func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
}

// struct可以帶tag 跟JAVA的 annotation tag類似
type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Raymond", 25}
	// 按名字獲取成員
	t.Logf("Name: value(%[1]v),Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {

		t.Error("Failed to get 'Name' field.")

	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Update Age", e)
}

// 通用程序: 對Employee跟Customer 改Name跟Age
func fillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		// Elem() 獲取指針指向的值
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("The first param type should be a pointer to the struct")
		}
	}
	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		// 檢查有無這個field
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		// 若StructField裡面的type跟 map裡面的type一致 才能寫入
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}

	}
	return nil
}
func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Raymond", "Age": 25}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

/*
reflect.DeepEqual
寫測試的時候，常常需要比較兩個變量是否相等
但是golang不支援直接比較map,slice,...等等的
這時候就可以用到DeepEqual (但是testify的assert.Equal比較常用)
*/
func TestSliceEqual(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	// fmt.Println(slice1==slice2)
	fmt.Println(reflect.DeepEqual(slice1, slice2))
}
func TestStructEqual(t *testing.T) {
	struct1 := Employee{Name: "Raymond"}
	struct2 := Employee{Name: "Raymond"}
	fmt.Println(reflect.DeepEqual(struct1, struct2))
	// 不用反射可以這樣
	fmt.Println(struct1 == struct2)
	fmt.Println(&struct1 == &struct2) //記憶體位置不一樣

	// 但是如果結構含有slice map之類的，就會不能比較
	type Employee2 struct {
		Employee
		Slice []int
	}
	struct3 := Employee2{Employee: struct1, Slice: []int{1, 2, 3}}
	struct4 := Employee2{Employee: struct2, Slice: []int{1, 2, 3}}
	// fmt.Println(struct3 == struct4) //會不能比較

	fmt.Println(reflect.DeepEqual(struct3, struct4))
}

/* 反射三大定律 Rob Pike在 The Laws of Reflection提出的*/
/*反射第一定律: Reflection goes from interface value to reflection object
interface{}變量可以透過reflect.TypeOf獲取變量類型，reflect.ValueOf獲取變量的值
如果我們知道一個變量的類型和值，那就代表我們知道了這個變量的全部信息
*/

type Coder struct {
	Name string
}

func (c *Coder) String() string {
	return c.Name
}

func TestReflectLaw1(t *testing.T) {
	coder := &Coder{Name: "Raymond"}
	typ := reflect.TypeOf(coder)
	val := reflect.ValueOf(coder)
	typeOfStringer := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	fmt.Println("kind of coder:", typ.Kind())
	fmt.Println("type of coder", typ)
	fmt.Println("value of coder ", val)

	// fmt.Println("field 0 of coder ", typ.Field(0))
	fmt.Println("implements stringer:", typ.Implements(typeOfStringer))
}

/* 反射第二定律: Reflection goes from reflection object to interface value
就是通過interface()方法，從反射對象獲取interface{}變量
*/
func TestReflectLaw2(t *testing.T) {
	coder := &Coder{Name: "Raymond"}
	val := reflect.ValueOf(coder)
	c, ok := val.Interface().(*Coder)
	if ok {
		fmt.Println(c.Name)
	} else {
		panic("type assert to *Coder err")
	}
}

/*
簡單總結一下，第一第二定律描述了通過ValueOf方法可以從Interface獲取反射對象，
反射對象調用Interface()獲取Interface變量
Interface變量可以用個類型轉換再得到原來的變量
 -----				  ---------------	--TypeOf/ValueOf-->  -----------------
|value| <-Typecast-> |interface value|						|reflection object|
 -----				  ---------------	  <--Interface()--   -----------------
*/

/*反射第三定律: To modify a reflection object, the value must be settable
要更新一個reflect.Value，那麼他的值必須是可被更新的(可被尋址)，否則就會panic
*/

func TestReflectLaw3v1(t *testing.T) {
	// 下面這樣會panic
	i := 1
	v := reflect.ValueOf(i)
	v.SetInt(10)
	fmt.Println(i)
}

/*
	通過獲取變量的指針來更新值
	相當於
		i:=1
		v:= &i
		*v = 10
*/
func TestReflectLaw3v2(t *testing.T) {
	i := 1
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)
}
