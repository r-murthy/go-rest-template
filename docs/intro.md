## Case

Is Go just a fluff?

An additional setup to the populor Java Springboot choice for building microservice

In companies' tech world, with the increasing ease of running and managing services on the cloud, it has become convienient to seperate the stateful concerns into micoservices. This is a common practice. More and more companies have been breaking down large monoliths into tiny microservices. In our techstack, Java (springboot) is probably the best setup to develop scalable and robust microservices.

## Potential areas of improvements

- Longer compile time of java springboot makes it harder to debug and run microservices applications.

- DB connection speed is a slow

- Vast number and keywords and third party library allow developers to write complex code making it harder for the other developers to follow.

- Perception of OOPs differ from person to person leaving always that room for different opinions. Eg. Its common pattern we hear how some interfaces are seen redundant.

- Bulky IDEs to work with when developing microservices. Eg. A simple editor doesnt quite suffice develop a small simple microservice in Javaspring boot.

## Plight of the current choice

In the world of microservices, it can be frequently observed that we do not end up using most of the benefits of OOPs concepts. Thus in Java springboot setup, plenty irrelavant code shenanigans are brought in. Java hence has a hard time putting up a fight as the best choice.

Most OOPs concepts are obselete when it comes to developing simple CRED services.
There are few other alternatives to pick from the techstack, however it is devoid of an alternative that is proven to bring more just in building microservices.

### Further benefits of Go over Java

- Fast compile time
  	- Extreamly fast build time, saving plenty of development time. Go’s fast compile times are a major productivity win compared to languages like Java and C++ which are famous for sluggish compilation speed. [Here are some comparisions](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go.html)

- It has built-in support for HTTP/2
  	- No need to use a framework to build services. It has built-in support for HTTP/2 and brings flexibility, scalability, and simplicity to web development in a way that Java can’t.

- Simple yet powerful and scalable
  	- The syntax of Go is very simple and comprises only about 25 keywords, hence learning the language is quite simple. It avoids programmers to write complex code, and hence easier to maintain for newcommers.

- Memory efficient, thus exteamly low latency
	- Go eschews features, cleverness and flexibility in favor of ruthless pragmatism and simplicity. Enforced coding style, strong typing, garbage collection, simple concurrency primitives, native binaries and fast compile times make Go an excellent choice when you’re concerned about ensuring that a codebase remains maintainable at scale
  	- There are plenty of comparisions available on internet on why Go is great choice in building microserives.
    It is a simple yet powerful language.
    [Here is an informative comparision between Java and Go](https://blog.boot.dev/golang/golang-vs-java-go/#:~:text=Golang%20doesn't%20possess%20the,footprint%20and%20faster%20start%20time.)
    Few excerpts from the above comparisions,

    > Garbage collectors are famous for weighing down performance, however, Go’s compiler is newer and optimized, so even with the added feature, it doesn’t add nearly as much drag as it did to Java in the tests.

    > Multithreading doesn’t inherently eat up memory, but Java’s concurrency is dense and bulky compared to Go’s, and it takes an excessive amount of memory to craft and destroy Java threads. A server can only handle a small amount of Java threads, whereas it can handle thousands of threads in Go.

    > In fact, one of Go’s best features is its multithreading abilities. Its cutting-edge ‘goroutines’ provide a smooth application, and even allows you to check your code in binary to ensure its safety.

    > In the end, a comparison of Go’s and Java’s memory usage, when sitting idle, showed Java consumed upwards of 160 MB of memory, while Go only used 0.86 MB – an insane difference.

- Increasing telent pool
  	- Last but certainly not the least is the 'hip' factor Go brings along. Young developers have mostly been opting Go as their preffered language to develop. It makes it easier to make the talent pool future proof (its already hard enough to find developers to work in J2EE applications).

In the end, there is nothing that Go can that Java cannot. And there are probably plently that Java can that Go cannot. However when it comes to developing microservices, Go is simple and performant choice.

### Downside of Go

Only downside of Go is that its not as old as Java and about 13 years old. Although it has great penentration in the industry, it hasnt the acceptance of the likes of Java, yet.

#### [For those who already are coding in go](https://go-proverbs.github.io/)
