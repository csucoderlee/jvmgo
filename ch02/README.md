# 类路径

```
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("hell world");
    }
}
```

最简单的输出hello world的代码，加载 - 验证 - 解析 - 准备 - 初始化 - 使用 - 卸载 这7个步骤是类生命周期，
类的生命周期的第一步就是要加载这个HelloWorld类。
加载 HelloWorld类之前，首页要加载它的超类 Object。
调用main()方法之前，虚拟机要准备好参数数组，所以要加载java.lang.String和java.lang.String[]类。
把字符串打印出来，需要加载java.lang.System类，等等。
第二节的内容，就是讨论Java虚拟机从哪里来寻找这些类。


java 虚拟机规范中并没有规定虚拟机应该从哪里寻找类，因此不同的虚拟机可以采用不同的方法。
而Oracle的Java虚拟机实现是根据类路径class path来搜索类的，按照搜索顺序可以分为三个部分：
启动类路径 bootstrap classpath
扩展类路径 extension classpath
用户类路径 user classpath

启动类路径默认对应jre/lib目录，可以通过-Xbootclasspath选项修改启动类路径
扩展类路径默认对应jre/lib/ext目录
自己实现的类以及第三方类库则位于用户类路径，用户类路径的默认值是当前目录"."，可以设置CLASSPATH环境变量来修改用户类路径，
不推荐使用这种方式来修改用户类路径，建议可以使用java -classpath 或者 java -cp 来指定用户类路径，此命令优先级高于CLASSPATH环境变量
所以，即便设置了CLASSPATH环境变量，也可以使用java -cp来覆盖，达到修改用户类路径的目的