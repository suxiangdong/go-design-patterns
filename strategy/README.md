# 策略模式

### 介绍
策略模式是对变化的封装，多个类只有在算法或行为上稍有不同的场景均可使用策略模式。
即：如果对一个功能，有多种算法实现，那么可以将这些算法，分别封装为一个类，通过context上下文进行管理，以达到算法随时切换并屏蔽算法实现规则的目的。

### 应用场景
以支付系统为例，如果存在银行卡支付、支付宝支付、微信支付等多种方式，那么可以分别将其封装为支付类，同时实现支付接口。然后交由上下文类做统一管理及
切换。
参考示例代码。

### 对比简单工厂模式

1、简单工厂模式是创建型模式，侧重于封装类的创建。策略模式是行为型模式，侧重于封装不同算法类之间的变化。
2、简单工厂模式只需对外暴露创建方法及对象接口即可，调用者无需关心对象的创建。策略模式需要对外暴露所有类，调用者自己管理需要使用的类对象。
3、新增算法类，对于简单工厂模式创建算法类并修改工厂创建对象函数。对于策略模式仅需要增加算法类即可。
4、策略模式可随时自由切换类对象，但调用者需要知道每种算法的差异，当算法类增加到一定程度时，对调用者来说是一场灾难。

实际使用过程中，策略模式结合简单工厂模式应该是个不错的选择，就是在context上下文类中实现简单工厂模式。如此一来，调用者再也不需要关心类的创建。
在工厂创建类的过程中，应用反射会写更高质量的代码。