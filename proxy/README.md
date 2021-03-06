# 代理模式

### 介绍
为其他对象提供代理，以控制实际对象访问。比如远程代理、虚拟代理、安全代理、智能指引。  

远程代理：提供一个代理对象，以达到访问远程对象的目的。  
虚拟代理：代理开销比较大的对象创建。比如浏览器中图片的加载，图片加载是比较慢的，但是并不影响页面的渲染，图片会在之后缓慢加载到页面中。  
这里就是应用了代理模式，用于性能优化。  
安全代理：通过代理对象，控制真实对象的权限访问。  
智能指引：就是在真实对象方法调用之前，做一些其他事情，示例代码就是这种场景：在支付之前做参数检测。  

### 装饰模式 VS 代理模式
这里可能会有人发现，对于智能指引场景的代理模式与装饰模式基本一致，都是对其他类的功能增强。从实现方面来说，确实是一致的，UML类图都是类似的。
但是，从控制权来说就有很大的区别了：装饰模式是将控制权交予调用者的，而代理模式的控制权在开发者手里。说白了就是，对于装饰模式来说，调用者可以
选择如何去装饰，在支付类的支付方法调用之前，可以装饰日志的记录，也可以装饰参数检测，还能控制顺序。对于代理模式就不行了，控制权在开发者手里，
调用者只能调用增强后的功能！  
实际开发中大家可以根据场景以及控制权的归属来确定应用哪种模式。