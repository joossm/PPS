# struct:
```
using System;

struct s {
    public int foo;
    public string bar;
}

class Program
{
    static void Main() {
        s x = new s(); 
        x.foo = 3;
        x.bar = "bar";
        s y = x;
        y.bar = "foo";
        Console.WriteLine (x.bar);
    }
}
```
The code creates a struct named `s`, which contains two fields: `foo` and `bar`. It then creates two instances of `s`, `x` and `y`. It assigns the value `3` to the field `foo` in `x`, and the value `"bar"` to the field `bar` in `x`. Finally, it assigns `y` to `x`, and then changes the value of the `bar` field in `y` to `"foo"`.

The result of the code is that the value of `x.bar` is `"foo"`.

# class:
```
using System;

struct s {
    public int foo;
    public string bar;
}

class Program
{
    static void Main() {
        s x = new s(); 
        x.foo = 3;
        x.bar = "bar";
        s y = x;
        y.bar = "foo";
        Console.WriteLine (x.bar);
    }
}
The code above creates a new struct `s`, with two fields `foo` and `bar`. It then creates a new instance of `s` called `x`, and sets the value of `x.foo` to 3 and `x.bar` to `"bar"`. It then creates a new instance of `s` called `y`, and sets it equal to `x`. Finally, it changes the value of `y.bar` to `"foo"`.

The output of this code is `"foo"`.
```
# comparison
| Class                                                                               | Structure                                                                                                                              |
|-------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| Classes are of reference types\.                                                    | Structs are of value types\.                                                                                                           |
| All the reference types are allocated on heap memory\.                              | All the value types are allocated on stack memory\.                                                                                    |
| Allocation of large reference type is cheaper than allocation of large value type\. | Allocation and de\-allocation is cheaper in value type as compare to reference type\.                                                  |
| Class has limitless features\.                                                      | Struct has limited features\.                                                                                                          |
| Class is generally used in large programs\.                                         | Struct are used in small programs\.                                                                                                    |
| Classes can contain constructor or destructor\.                                     | Structure does not contain parameter less constructor or destructor, but can contain Parameterized constructor or static constructor\. |
| Classes used new keyword for creating instances\.                                   | Struct can create an instance, with or without new keyword\.                                                                           |
| A Class can inherit from another class\.                                            | A Struct is not allowed to inherit from another struct or class\.                                                                      |
| The data member of a class can be protected\.                                       | The data member of struct can’t be protected\.                                                                                         |

https://www.geeksforgeeks.org/difference-between-class-and-structure-in-c-sharp/

# nice explanation: 
https://www.youtube.com/watch?v=HlzAtIHFRk0